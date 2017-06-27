package main

import (
	"github.com/domego/gopt/gens/common"
	"github.com/domego/gopt/gens/gincontroller"
	"github.com/domego/gopt/gens/ginserver"
	"github.com/domego/gopt/gens/orm"
	"github.com/domego/gopt/gens/types"
)

func genTypes(name, desc string) {
	gentypes.Gen(typesFile)
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
	gengincontroller.Gen(apiFile)
}
