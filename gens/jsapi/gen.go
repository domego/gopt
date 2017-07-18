package genjsapi

import (
	"io/ioutil"

	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
	yaml "gopkg.in/yaml.v1"
)

// Gen 执行GenGinController生成程序
func Gen(apiFile, template string) {
	log.Infof("start gen_js_api, api_file=%s, template=%s", apiFile, template)
	if template == "vue" {
		template = "webapp/api.vue_resource.js.tmpl"
	} else {
		log.Fatalf("template is invalid")
	}
	groups := []*genutils.GenControllerGroup{}
	var bs []byte
	var err error
	if bs, err = ioutil.ReadFile(apiFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &groups); err != nil {
		log.Fatalf("%s", err)
	}

	genutils.MkDirIfNotExists("webapp/src")
	genutils.GenFileWithTargetPath(template, "webapp/src/api.js", map[string]interface{}{"Groups": groups})
	//	sh.Command("gofmt", "-w", "router.go").Run()
}
