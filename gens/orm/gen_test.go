package genorm

import (
	"github.com/lenbo-ma/ginpt/gens/common"
	"testing"
)

func TestGen(t *testing.T) {
	genutils.Asset = Asset
	genutils.SetValues(map[string]interface{}{
		"AppName":  "example",
		"RootPath": "github.com/lenbo-ma/ginpt/gens/orm",
	})
	Gen("db.yaml")
}
