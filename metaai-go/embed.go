package metaai

import "embed"

// SmartStudioDistFS holds the SMART Studio React production build (smart-studio/dist). In
// dev this only contains a dev.txt marker; `make build-prod` (which runs
// `make build-jenta`) populates it with the real Vite build so a single Go
// binary serves the SPA + the API on the same origin (no CORS in prod).
//
//go:embed smart-studio/dist
var SmartStudioDistFS embed.FS
