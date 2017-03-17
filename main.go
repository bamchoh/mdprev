package main

import (
	"fmt"
	"os"

	static "github.com/Code-Hex/echo-static"
	assetfs "github.com/elazarl/go-bindata-assetfs"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:generate cp node_modules/marked/lib/marked.js _assets/
//go:generate cp node_modules/jquery/dist/jquery.min.js _assets/
//go:generate go-bindata data/... _assets/... node_modules/marked/marked.min.js

var (
	// SavePath ...
	SavePath string
)

// NewAssets ...
func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}

func initSavePath(savePath string) error {
	SavePath = savePath
	return os.MkdirAll(SavePath, 0777)
}

func main() {
	c := parseArgs()
	err := initSavePath(c.storePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(static.ServeRoot("/_assets", NewAssets("_assets")))
	e.Use(static.ServeRoot("/node_modules", NewAssets("node_modules")))
	e.GET("/", getRoot)
	e.GET("/api/markdown", getMarkdown)
	e.POST("/api/markdown", postMarkdown)
	e.Logger.Fatal(e.Start(":" + c.port))
}
