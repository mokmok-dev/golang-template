//go:build tools
// +build tools

package tools

import (
	_ "github.com/cosmtrek/air"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/k0kubun/sqldef/cmd/psqldef"
	_ "golang.org/x/tools/cmd/goimports"
	_ "mvdan.cc/gofumpt"
)
