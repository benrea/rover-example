//go:build tools
// +build tools

package main

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/99designs/gqlgen/plugin/federation"
	_ "github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

func Tools() {
	// This function does nothing. It's just there to provide a declaration to reference the imported package(s).
}
