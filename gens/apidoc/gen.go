package genapidoc

import (
	"io/ioutil"

	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
	yaml "gopkg.in/yaml.v1"
)

// Gen 执行GenAPIDoc生成程序
func Gen(apiFile string) {
	log.Infof("start gen_api_doc, api_file=%s", apiFile)
	groups := []*genutils.GenControllerGroup{}
	var bs []byte
	var err error
	if bs, err = ioutil.ReadFile(apiFile); err != nil {
		log.Fatalf("%s", err)
	}
	if err = yaml.Unmarshal(bs, &groups); err != nil {
		log.Fatalf("%s", err)
	}

	model := map[string]interface{}{
		"Name":     genutils.Values["AppName"],
		"Port":     genutils.Values["AppPort"],
		"RootPath": genutils.Values["RootPath"],
		"Groups":   groups,
	}
	genutils.MkDirIfNotExists("doc")
	genutils.GenFile("doc/api_doc.md.tmpl", model)
}
