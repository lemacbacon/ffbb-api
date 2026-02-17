# Client Go FFBB API

Client Go généré automatiquement depuis la spécification OpenAPI.

## Installation

```bash
go get github.com/ffbb/api/go
```

## Utilisation

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ffbb/api/go/client"
)

func main() {
	ctx := context.Background()

	// Récupère automatiquement les tokens et crée les clients
	directus, meilisearch, err := client.NewAuthenticatedClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Lister les compétitions
	resp, err := directus.ListCompetitionsWithResponse(ctx, &client.ListCompetitionsParams{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.JSON200.Data)

	_ = meilisearch
}
```

## Développement

### Regénérer le client

```bash
make generate-go
```

### Vérifier que le code généré est à jour

```bash
make verify-go
```

## Structure

```
go/
  go.mod
  tools.go              # dépendance oapi-codegen
  oapi-codegen.yaml     # configuration de génération
  client/
    client.gen.go       # généré — ne pas modifier manuellement
    auth.go             # helper d'authentification
```
