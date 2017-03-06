package main

import (
	static "github.com/Code-Hex/echo-static"
	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:generate go-bindata data/... _assets/... node_modules/marked/marked.min.js

// NewAssets ...
func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}

func main() {
	c := parseArgs()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(static.ServeRoot("/_assets", NewAssets("_assets")))
	e.Use(static.ServeRoot("/node_modules", NewAssets("node_modules")))
	e.GET("/", getRoot)
	e.GET("/api/markdown", getMarkdown)
	e.POST("/api/markdown", postMarkdown)
	e.Logger.Fatal(e.Start(":" + c.port))
}
