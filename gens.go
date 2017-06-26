package main

import (
	"github.com/lenbo-ma/ginpt/gens/common"
	"github.com/lenbo-ma/ginpt/gens/ginserver"
	"github.com/lenbo-ma/ginpt/gens/orm"
	"github.com/lenbo-ma/ginpt/gens/types"
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
