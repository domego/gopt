package main

import (
	"fmt"
	"io/ioutil"

	"github.com/lenbo-ma/ginpt/gens/common"
	"github.com/lenbo-ma/ginpt/gens/ginserver"
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
			genutils.MkDirIfNotExists(g.PackagePath)
		}
		outputPath = fmt.Sprintf("%s/%s", g.PackagePath, outputPath)
		genutils.TemplateExecute(bs, outputPath, g)
	}
}

func genGinServer(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	genginserver.Gen()
}
