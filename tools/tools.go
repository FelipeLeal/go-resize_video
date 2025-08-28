//go:build tools

// This package imports tool dependencies so they are tracked by go mod.
// This ensures that everyone on the team, and the CI/CD, uses the same
// version of the tools.
// See: https://www.jba.dev/posts/tools-as-dependencies/

package tools

import (
	_ "github.com/air-verse/air"
	_ "github.com/swaggo/swag/cmd/swag"
)