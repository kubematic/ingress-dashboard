package static

import (
	"embed"
	"io/fs"
)

//go:embed assets/templates/**
var Templates embed.FS

//go:embed assets/static/**
var Files embed.FS

func Static() fs.FS {
	f, err := fs.Sub(Files, "assets/static")
	if err != nil {
		panic(err) // this should never happen
	}
	return f
}
