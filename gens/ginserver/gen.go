package genginserver

import (
	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/common"
)

// Gen 执行GinServer生成程序
func Gen() {
	log.Infof("start gen_gin_server, name=%s, port=%d", genutils.Values["AppName"], genutils.Values["AppPort"])
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
		"Name":     genutils.Values["AppName"],
		"Port":     genutils.Values["AppPort"],
		"RootPath": genutils.Values["RootPath"],
	}

	// mkdirs
	genutils.InitDirs(dirs)
	// copy files
	genutils.CopyFiles(cpFiles)
	// generate files
	genutils.GenFiles(files, model)
}
