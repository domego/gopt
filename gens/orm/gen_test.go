package genorm

import (
	"github.com/domego/gopt/gens/common"
	"testing"
)

func TestGen(t *testing.T) {
	genutils.Asset = Asset
	genutils.SetValues(map[string]interface{}{
		"RootPath": "github.com/domego/gopt/gens/orm",
	})
	Gen("db.yaml")
}
