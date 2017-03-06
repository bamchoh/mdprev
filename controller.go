package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/schollz/jsonstore"
)

const (
	storeFile = "mk.json.gz"
)

type markdown struct {
	String string `json:"md"`
}

func getRoot(c echo.Context) error {
	html, err := Asset("data/index.html")
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, string(html))
}

func storeJSON(mk markdown) (err error) {
	ks := new(jsonstore.JSONStore)
	err = ks.Set("mk:1", mk)
	if err != nil {
		return
	}

	err = jsonstore.Save(ks, storeFile)
	return
}

func getMarkdown(c echo.Context) (err error) {
	var ks2 *jsonstore.JSONStore
	ks2, err = jsonstore.Open(storeFile)
	if err != nil {
		return
	}

	var mk markdown
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
	var m markdown
	dec.Decode(&m)

	err := storeJSON(m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, `{"message":"success"}`)
}
