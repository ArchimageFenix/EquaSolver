// File: assets.go
// Purpose: embed the Web templates and static presentation assets in the binary.
// Receives: template and static files included at build time by go:embed.
// Previous stage: project Web resources at build time.
// Next stage: present.go uses templates; server.go serves static assets.
// Restrictions: no HTTP routing, domain logic, calculations or presentation formatting.

package web

import "embed"

// webFiles contains the Web interface resources required at runtime.
// Keeping them embedded preserves EquaSolver as a self-contained binary.
//
//go:embed templates/*.html static/css/*.css
var webFiles embed.FS
