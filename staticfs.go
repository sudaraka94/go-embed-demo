package main

import (
	"embed"
	"io/fs"
	"path"
)

type StaticFS struct {
	content embed.FS
}

func (c StaticFS) Open(name string) (fs.File, error) {
	return c.content.Open(path.Join("static", name))
}
