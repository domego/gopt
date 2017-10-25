package main

import (
	"github.com/domego/gokits/log"
	"github.com/domego/gopt/gens/apidoc"
	"github.com/domego/gopt/gens/common"
	"github.com/domego/gopt/gens/errors"
	"github.com/domego/gopt/gens/gincontroller"
	"github.com/domego/gopt/gens/ginserver"
	"github.com/domego/gopt/gens/jsapi"
	"github.com/domego/gopt/gens/orm"
	"github.com/domego/gopt/gens/types"
)

func genTypes(name, desc string) {
	gentypes.Gen(typesFile)
}

func genErrors(name, desc string) {
	generrors.Gen(errorsFile)
}

func genGinServer(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	genginserver.Gen()
}

func genORM(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	genorm.Gen(ormFile)
}

func genGinController(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	log.Debugf("RootPath: %s", rootPath)
	gengincontroller.Gen(apiFile)
}

func genJavascriptAPI(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	log.Debugf("RootPath: %s", rootPath)
	genjsapi.Gen(apiFile, template)
}

func genAPIDoc(name, desc string) {
	genutils.SetValues(map[string]interface{}{
		"AppName":  appName,
		"AppPort":  appPort,
		"RootPath": rootPath,
	})
	log.Debugf("RootPath: %s", rootPath)
	genapidoc.Gen(apiFile)
}
