package client_test

import (
	"context"
	"sync"
	"testing"

	"github.com/ffbb/api/go/client"
)

var (
	setupOnce         sync.Once
	directusClient    *client.ClientWithResponses
	meilisearchClient *client.ClientWithResponses
	setupErr          error
)

func testClients(t *testing.T) (*client.ClientWithResponses, *client.ClientWithResponses) {
	t.Helper()
	if testing.Short() {
		t.Skip("skipping e2e test in short mode")
	}

	setupOnce.Do(func() {
		directusClient, meilisearchClient, setupErr = client.NewAuthenticatedClient(context.Background())
	})
	if setupErr != nil {
		t.Fatalf("NewAuthenticatedClient: %v", setupErr)
	}
	return directusClient, meilisearchClient
}

func TestGetConfiguration(t *testing.T) {
	directus, _ := testClients(t)

	resp, err := directus.GetConfigurationWithResponse(context.Background())
	if err != nil {
		t.Fatalf("GetConfigurationWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		t.Fatal("expected non-nil JSON200")
	}
	if resp.JSON200.Data.KeyDh == "" {
		t.Fatal("expected non-empty key_dh")
	}
	if resp.JSON200.Data.KeyMs == "" {
		t.Fatal("expected non-empty key_ms")
	}
}

func TestListCompetitions(t *testing.T) {
	directus, _ := testClients(t)

	limit := 5
	resp, err := directus.ListCompetitionsWithResponse(context.Background(), &client.ListCompetitionsParams{
		Limit: &limit,
	})
	if err != nil {
		t.Fatalf("ListCompetitionsWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil || len(resp.JSON200.Data) == 0 {
		t.Fatal("expected non-empty competition data")
	}
}

func TestGetCompetition(t *testing.T) {
	directus, _ := testClients(t)

	// First, get a valid competition ID.
	limit := 1
	list, err := directus.ListCompetitionsWithResponse(context.Background(), &client.ListCompetitionsParams{
		Limit: &limit,
	})
	if err != nil {
		t.Fatalf("ListCompetitionsWithResponse: %v", err)
	}
	if list.JSON200 == nil || len(list.JSON200.Data) == 0 {
		t.Fatal("need at least one competition")
	}
	id := list.JSON200.Data[0].Id

	resp, err := directus.GetCompetitionWithResponse(context.Background(), id, &client.GetCompetitionParams{})
	if err != nil {
		t.Fatalf("GetCompetitionWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		t.Fatal("expected non-nil JSON200")
	}
	if resp.JSON200.Data.Id != id {
		t.Fatalf("expected id %s, got %s", id, resp.JSON200.Data.Id)
	}
}

func TestGetOrganisme(t *testing.T) {
	directus, meilisearch := testClients(t)

	// Search for an organisme ID via Meilisearch.
	q := "paris"
	msLimit := 1
	search, err := meilisearch.MultiSearchWithResponse(context.Background(), client.MultiSearchJSONRequestBody{
		Queries: []client.SearchQuery{
			{
				IndexUid: client.FfbbserverOrganismes,
				Q:        &q,
				Limit:    &msLimit,
			},
		},
	})
	if err != nil {
		t.Fatalf("MultiSearchWithResponse: %v", err)
	}
	if search.JSON200 == nil || len(search.JSON200.Results) == 0 {
		t.Fatal("need at least one search result")
	}
	hits := search.JSON200.Results[0].Hits
	if hits == nil || len(*hits) == 0 {
		t.Fatal("need at least one hit")
	}

	// Use the typed union accessor to get the organisme hit.
	orgHit, err := (*hits)[0].AsOrganismesHit()
	if err != nil {
		t.Fatalf("AsOrganismesHit: %v", err)
	}
	if orgHit.Id == nil || *orgHit.Id == "" {
		t.Fatal("organisme hit missing id")
	}
	orgID := *orgHit.Id

	resp, err := directus.GetOrganismeWithResponse(context.Background(), orgID, &client.GetOrganismeParams{})
	if err != nil {
		t.Fatalf("GetOrganismeWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		t.Fatal("expected non-nil JSON200")
	}
	if resp.JSON200.Data.Id != orgID {
		t.Fatalf("expected id %s, got %s", orgID, resp.JSON200.Data.Id)
	}
}

func TestGetPoule(t *testing.T) {
	directus, _ := testClients(t)

	// Get a competition with its poules to find a valid poule ID.
	limit := 1
	fields := client.Fields{"id", "poules"}
	list, err := directus.ListCompetitionsWithResponse(context.Background(), &client.ListCompetitionsParams{
		Limit:  &limit,
		Fields: &fields,
	})
	if err != nil {
		t.Fatalf("ListCompetitionsWithResponse: %v", err)
	}
	if list.JSON200 == nil || len(list.JSON200.Data) == 0 {
		t.Fatal("need at least one competition")
	}
	comp := list.JSON200.Data[0]
	if comp.Poules == nil || len(*comp.Poules) == 0 {
		t.Skip("competition has no poules")
	}

	// Poules are returned as string IDs when not expanded.
	pouleID, err := (*comp.Poules)[0].AsCompetitionPoules0()
	if err != nil {
		t.Fatalf("AsCompetitionPoules0: %v", err)
	}

	resp, err := directus.GetPouleWithResponse(context.Background(), pouleID, &client.GetPouleParams{})
	if err != nil {
		t.Fatalf("GetPouleWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil {
		t.Fatal("expected non-nil JSON200")
	}
	if resp.JSON200.Data.Id != pouleID {
		t.Fatalf("expected id %s, got %s", pouleID, resp.JSON200.Data.Id)
	}
}

func TestListSaisons(t *testing.T) {
	directus, _ := testClients(t)

	resp, err := directus.ListSaisonsWithResponse(context.Background(), &client.ListSaisonsParams{})
	if err != nil {
		t.Fatalf("ListSaisonsWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil || len(resp.JSON200.Data) == 0 {
		t.Fatal("expected non-empty saison data")
	}
}

func TestGetLives(t *testing.T) {
	directus, _ := testClients(t)

	resp, err := directus.GetLivesWithResponse(context.Background())
	if err != nil {
		t.Fatalf("GetLivesWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
}

func TestMultiSearch(t *testing.T) {
	_, meilisearch := testClients(t)

	q := "paris"
	limit := 5
	resp, err := meilisearch.MultiSearchWithResponse(context.Background(), client.MultiSearchJSONRequestBody{
		Queries: []client.SearchQuery{
			{
				IndexUid: client.FfbbserverOrganismes,
				Q:        &q,
				Limit:    &limit,
			},
		},
	})
	if err != nil {
		t.Fatalf("MultiSearchWithResponse: %v", err)
	}
	if resp.StatusCode() != 200 {
		t.Fatalf("expected status 200, got %d", resp.StatusCode())
	}
	if resp.JSON200 == nil || len(resp.JSON200.Results) == 0 {
		t.Fatal("expected non-empty search results")
	}
}
