package client

import (
	"context"
	"fmt"
	"net/http"
)

const (
	directusBaseURL    = "https://api.ffbb.app"
	meilisearchBaseURL = "https://meilisearch-prod.ffbb.app"
	defaultUserAgent   = "okhttp/4.12.0"
)

// NewAuthenticatedClient appelle GET /items/configuration (sans auth)
// pour récupérer les tokens, puis retourne deux clients pré-configurés :
//   - Client Directus (api.ffbb.app) avec Bearer key_dh
//   - Client Meilisearch (meilisearch-prod.ffbb.app) avec Bearer key_ms
func NewAuthenticatedClient(ctx context.Context, opts ...ClientOption) (directus *ClientWithResponses, meilisearch *ClientWithResponses, err error) {
	uaEditor := func(_ context.Context, req *http.Request) error {
		req.Header.Set("User-Agent", defaultUserAgent)
		return nil
	}

	tmp, err := NewClientWithResponses(directusBaseURL, append(opts, WithRequestEditorFn(uaEditor))...)
	if err != nil {
		return nil, nil, fmt.Errorf("creating temporary client: %w", err)
	}

	resp, err := tmp.GetConfigurationWithResponse(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("fetching configuration: %w", err)
	}
	if resp.JSON200 == nil {
		return nil, nil, fmt.Errorf("unexpected status %s from /items/configuration", resp.Status())
	}

	cfg := resp.JSON200.Data

	authEditor := func(token string) RequestEditorFn {
		return func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("User-Agent", defaultUserAgent)
			return nil
		}
	}

	directus, err = NewClientWithResponses(directusBaseURL, append(opts, WithRequestEditorFn(authEditor(cfg.KeyDh)))...)
	if err != nil {
		return nil, nil, fmt.Errorf("creating directus client: %w", err)
	}

	meilisearch, err = NewClientWithResponses(meilisearchBaseURL, append(opts, WithRequestEditorFn(authEditor(cfg.KeyMs)))...)
	if err != nil {
		return nil, nil, fmt.Errorf("creating meilisearch client: %w", err)
	}

	return directus, meilisearch, nil
}
