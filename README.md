# FFBB API

Spécification [OpenAPI 3.0.3](https://spec.openapis.org/oas/v3.0.3) de l'API de la Fédération Française de Basketball (FFBB), accompagnée de clients générés automatiquement.

## Services

L'API combine deux services :

| Service | Base URL | Auth |
|---|---|---|
| **Directus** (données basketball) | `https://api.ffbb.app` | Bearer `key_dh` |
| **Meilisearch** (recherche full-text) | `https://meilisearch-prod.ffbb.app` | Bearer `key_ms` |

Les tokens sont récupérés automatiquement via `GET /items/configuration` (sans authentification).

## Clients

| Langage | Dossier | Documentation |
|---|---|---|
| Go | [`go/`](go/) | [README](go/README.md) |

## Développement

### Valider la spécification

```bash
make lint
```

### Regénérer les clients

```bash
make generate-go
```

## Contribuer

Les contributions sont les bienvenues ! Consultez [CONTRIBUTING.md](CONTRIBUTING.md) pour les détails.

## Licence

[MIT](LICENSE)
