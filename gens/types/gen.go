package gentypes

import (
	"fmt"
	"io/ioutil"

	sh "github.com/codeskyblue/go-sh"
	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
	yaml "gopkg.in/yaml.v1"
)

type GenTypeGroup struct {
	Package     string      `yaml:"Package"`
	PackagePath string      `yaml:"Path"`
	Types       interface{} `yaml:"Types"`
}

func Gen(typesFile string) {
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
	if bs, err = genutils.Asset("gen_types.go.tmpl"); err != nil {
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

		sh.Command("gofmt", "-w", outputPath).Run()
	}
}
