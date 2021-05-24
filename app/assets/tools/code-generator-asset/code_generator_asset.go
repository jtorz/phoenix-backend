// Package codegenasset uses go:embed to load the templates for code generation.
package codegenasset

import _ "embed"

//go:embed model.tpl
var ModelTPL string
