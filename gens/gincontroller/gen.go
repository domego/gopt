package gengincontroller

import (
	"io/ioutil"
	"os"
	"strings"

	sh "github.com/codeskyblue/go-sh"
	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
	yaml "gopkg.in/yaml.v1"
)

type GenControllerGroup struct {
	Name        string      `yaml:"Name"`
	PackageName string      `yaml:"PackageName"`
	PackagePath string      `yaml:"PackagePath"`
	Desc        string      `yaml:"Desc"`
	RoutePrefix string      `yaml:"RoutePrefix"`
	AllPost     bool        `yaml:"AllPost"`
	Routes      interface{} `yaml:"Routes"`
}

// Gen 执行GenGinController生成程序
func Gen(apiFile string) {
	log.Infof("start gen_gin_controller, use %s", apiFile)
	groups := []*GenControllerGroup{}
	var bs []byte
	var err error
	if bs, err = ioutil.ReadFile(apiFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &groups); err != nil {
		log.Fatalf("%s", err)
	}

	for _, group := range groups {
		genRouteGroup(group)
	}
}

func genRouteGroup(group *GenControllerGroup) {
	var outputPath string
	if group.PackageName == "" {
		group.PackageName = strings.ToLower(group.Name) + "controller"
	}
	if group.PackagePath == "" {
		group.PackagePath = "controller/" + strings.ToLower(group.Name)
	}
	genutils.MkDirIfNotExists(group.PackagePath)
	genutils.GenFileWithTargetPath("controller/gen_controller.go.tmpl", group.PackagePath+"/gen_controller.go", group)
	outputPath = group.PackagePath + "/controller.go"
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		genutils.GenFileWithTargetPath("controller/controller.go.tmpl", outputPath, group)
	} else {
		log.Infof("%s is already exist.", outputPath)
	}

	genutils.GenFileWithTargetPath("controller/gen_route.go.tmpl", group.PackagePath+"/gen_route.go", group)
	outputPath = group.PackagePath + "/route.go"
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		genutils.GenFileWithTargetPath("controller/route.go.tmpl", outputPath, group)
	} else {
		log.Infof("%s is already exist.", outputPath)
	}
	sh.Command("gofmt", "-w", ".", sh.Dir(group.PackagePath)).Run()
}
