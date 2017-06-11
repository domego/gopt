package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/lenbo-ma/gokits/log"
	yaml "gopkg.in/yaml.v1"
)

type GenTypeGroup struct {
	Package     string      `yaml:"Package"`
	PackagePath string      `yaml:"Path"`
	Types       interface{} `yaml:"Types"`
}

func genTypes(name, desc string) {
	log.Infof("start gen_types, use %s", typesFile)
	groups := []*GenTypeGroup{}
	var bs []byte
	var err error
	if bs, err = ioutil.ReadFile(typesFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &groups); err != nil {
		log.Fatalf("%s", err)
	}
	if bs, err = Asset("gen_types.go.tmpl"); err != nil {
		log.Fatalf("%s", err)
	}
	for _, g := range groups {
		outputPath := "gen_types.go"
		if g.PackagePath == "" {
			g.PackagePath = "."
		}
		if g.PackagePath != "." && g.PackagePath != "./" {
			mkDirIfNotExists(g.PackagePath)
		}
		outputPath = fmt.Sprintf("%s/%s", g.PackagePath, outputPath)
		templateExecute(bs, outputPath, g)
	}
}

func genGinServer(name, desc string) {
	log.Infof("start gen_gin_server, name=%s, port=%d", appName, appPort)
	dirs := []string{
		"config",
	}
	cpFiles := []string{
		"Makefile",
	}
	files := []string{
		"config/config.yaml",
		"app.go.tmpl",
		"config.go.tmpl",
		"main.go.tmpl",
		"NAME",
		"README.md",
		"router.go.tmpl",
	}
	model := map[string]interface{}{
		"Name":     appName,
		"Port":     appPort,
		"RootPath": rootPath,
	}

	// mkdirs
	for _, dir := range dirs {
		if err := mkDirIfNotExists(dir); err != nil {
			log.Fatalf("%s", err)
		}
	}
	// copy files
	for _, n := range cpFiles {
		bs, err := Asset(n)
		if err != nil {
			log.Fatalf("%s", err)
		}
		n = updateFileName(n)
		var mode os.FileMode = 0644
		if strings.Contains(n, "bin/") || strings.HasSuffix(n, ".sh") {
			mode = 0755
		}
		if err = ioutil.WriteFile(n, bs, mode); err != nil {
			log.Fatalf("%s", err)
		}
	}
	// generate files
	for _, n := range files {
		log.Infof("start gen file %s", n)
		var (
			bs  []byte
			err error
		)
		if bs, err = Asset(n); err != nil {
			log.Fatalf("%s", err)
		}
		n = updateFileName(n)
		templateExecute(bs, n, model)
	}
}
