package main

import (
	"encoding/json"
	"net/http"

	static "github.com/Code-Hex/echo-static"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/schollz/jsonstore"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//go:generate go-bindata data/... _assets/... node_modules/marked/marked.min.js

var ks *jsonstore.JSONStore

func getRoot(c echo.Context) error {
	html, err := Asset("data/index.html")
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, string(html))
}

// Markdown ...
type Markdown struct {
	String string `json:"md"`
}

// NewAssets ...
func NewAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}

func storeJSON(mk Markdown) (err error) {
	err = ks.Set("mk:1", mk)
	if err != nil {
		return
	}

	err = jsonstore.Save(ks, "mk.json.gz")
	return
}

func getMarkdown(c echo.Context) (err error) {
	var ks2 *jsonstore.JSONStore
	ks2, err = jsonstore.Open("mk.json.gz")
	if err != nil {
		return
	}

	var mk Markdown
	err = ks2.Get("mk:1", &mk)
	if err != nil {
		return
	}

	mkRaw, err := json.Marshal(mk)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, string(mkRaw))
}

func postMarkdown(c echo.Context) error {
	req := c.Request()
	dec := json.NewDecoder(req.Body)
	var m Markdown
	dec.Decode(&m)

	err := storeJSON(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, `{"message":"success"}`)
}

func main() {
	ks = new(jsonstore.JSONStore)
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(static.ServeRoot("/_assets", NewAssets("_assets")))
	e.Use(static.ServeRoot("/node_modules", NewAssets("node_modules")))
	e.GET("/", getRoot)
	e.GET("/api/markdown", getMarkdown)
	e.POST("/api/markdown", postMarkdown)
	e.Logger.Fatal(e.Start(":3000"))
}
