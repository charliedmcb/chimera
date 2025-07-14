# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Local Development
- `make web` - Run the web server locally (starts on port 80)
- `make codegen` - Generate Go code files from data sources 
- `make build` - Build the binary and Docker image

### Deployment
- `make login` - Authenticate with Azure
- `make publish` - Push Docker image to Azure Container Registry
- `make release` - Deploy to AKS cluster
- `make deploy` - Full deployment pipeline (login + build + publish + release)
- `make cleanup` - Remove deployment from cluster

## Project Architecture

This is a Go web application for generating Netrunner deck lists in the "Chimera" format. The application serves both a web interface and plain-text API endpoints.

### Core Components

**Main Web Server** (`deck-builder/main.go`)
- HTTP server listening on port 80
- Serves HTML pages for deck generation with embedded CSS/JS
- Provides plain-text API endpoints at `/api/plain-text-runner` and `/api/plain-text-corp`
- Static file serving for CSS, favicon, and pre-generated HTML lists

**Deck Building Logic** (`deck-builder/deckbuilder/deckbuilder.go`)
- `MakeCorpDeck()` - Generates 49-card corp decks with specific constraints:
  - Minimum 20 agenda points
  - Exactly 9 economy cards
  - Exactly 3 "End the Run" (ETR) ice cards
  - Hardcoded list of ETR ice types
- `MakeRunnerDeck()` - Generates 40-card runner decks with constraints:
  - Exactly 8 economy cards
  - Always includes Peacock, Aurora, and Creeper breakers

**Data Management**
- Card data stored in JSON files (`data/corp-cards.json`, `data/runner-cards.json`)
- Generated Go files with embedded card data (`deck-builder/generateddata/`)
- Tag-based card filtering system (`data/tags/` with econ and banned card lists)

**Code Generation** (`hack/codegen/`)
- Converts JSON card data to Go source files
- Generates static HTML files for ban lists and economy card lists
- Run via `make codegen` when card data changes

### Deployment Architecture

The application is deployed on Azure Kubernetes Service (AKS):
- Docker containerized Go binary
- Uses Azure Container Registry for image storage
- LoadBalancer service exposes the application
- Domain routing through Azure DNS (playchimera.net)

### Module Structure
- Module name: `netrunner-erng` (Go 1.19)
- No external Go dependencies - uses only standard library
- Self-contained with embedded static assets

## Branch Naming Convention

When creating pull request branches, use the format: `claude/<description-of-change>`
- `<description-of-change>` should be a minimal descriptor of the change being worked on
- Examples: `claude/custom-seed`, `claude/shareable-urls`, `claude/card-sorting`
- This helps clearly identify Claude-created branches vs. other branches

## Version Management

### Patch Command
When the user says "patch", it means to bump the patch version by one in both locations:
- `Makefile` - Update the `VERSION ?= vX.X.X-alpha` line
- `deck-builder/static/homepage.html` - Update the version display in the paragraph

Example: v0.0.48-alpha → v0.0.49-alpha