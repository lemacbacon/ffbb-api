# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This repository contains the **OpenAPI 3.0.3 specification** for the FFBB (Fédération Française de Basketball) API. It is a specification-only repo — there is no application code, build system, or tests. The single file `openapi.yaml` (~2,200 lines) documents the REST API.

## Validating the Spec

```bash
# Install a validator if needed
npx @redocly/cli lint openapi.yaml
```

## Architecture

The spec describes two integrated services:

1. **Directus API** (`api.ffbb.app`) — CMS backend for basketball data (competitions, teams, matches, clubs, seasons)
2. **Meilisearch** (`meilisearch-prod.ffbb.app`) — Full-text search across multiple indices

### Authentication Flow

1. Call `GET /items/configuration` (no auth) to retrieve tokens
2. Use `key_dh` (Bearer token) for Directus endpoints
3. Use `key_ms` (Bearer token) for Meilisearch endpoints

### Key Endpoints & Entities

| Endpoint Pattern | Entity | Notes |
|---|---|---|
| `/items/ffbbserver_competitions` | Competition | Supports deep nested queries (phases → poules → rencontres) |
| `/items/ffbbserver_poules/{id}` | Poule | Includes matches and team rankings |
| `/items/ffbbserver_saisons` | Saison | Basketball seasons |
| `/items/ffbbserver_organismes/{id}` | Organisme | Clubs, leagues, committees |
| `/json/lives.json` | Live | Real-time match scores with quarter breakdowns |
| `/multi-search` (Meilisearch) | Multi-index search | Indices: organismes, rencontres, terrains, salles, tournois, competitions, pratiques |

### Directus Query Patterns

All Directus endpoints use common parameters: `fields[]` (field selection), `filter` (JSON filtering), `limit`, and `deep` (for nested relation limits).

## Conventions

- Documentation is written in **French**
- All schema names and field names use the FFBB domain vocabulary (rencontre = match, poule = pool, organisme = club/organization)
- Directus collection names are prefixed with `ffbbserver_` or `ffbbnational_`
