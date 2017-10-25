package generrors

import (
	"fmt"
	"io/ioutil"

	sh "github.com/codeskyblue/go-sh"
	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
	yaml "gopkg.in/yaml.v1"
)

type GenErrorGroup struct {
	Package     string      `yaml:"Package"`
	PackagePath string      `yaml:"Path"`
	Errors      interface{} `yaml:"Errors"`
}

func Gen(errorsFile string) {
	log.Infof("start gen_errors, use %s", errorsFile)
	groups := []*GenErrorGroup{}
	var bs []byte
	var err error
	if bs, err = ioutil.ReadFile(errorsFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &groups); err != nil {
		log.Fatalf("%s", err)
	}

	// gen_errors.go
	if bs, err = genutils.Asset("gen_errors.go.tmpl"); err != nil {
		log.Fatalf("%s", err)
	}
	for _, g := range groups {
		outputPath := "gen_errors.go"
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

	// config/error_msg.yaml
	if bs, err = genutils.Asset("config/error_msg.yaml.tmpl"); err != nil {
		log.Fatalf("%s", err)
	}

	outputPath := "config/error_msg.yaml"
	genutils.TemplateExecute(bs, outputPath, map[string]interface{}{
		"Groups": groups,
	})

}
