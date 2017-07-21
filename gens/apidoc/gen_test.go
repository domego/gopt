package genapidoc

import (
	"github.com/lenbo-ma/ginpt/gens/common"
	"testing"
)

func TestGen(t *testing.T) {
	genutils.Asset = Asset
	Gen("api.yaml")
}
