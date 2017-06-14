package gentypes

import (
	"github.com/lenbo-ma/ginpt/gens/common"
	"testing"
)

func TestGenTypes(t *testing.T) {
	genutils.Asset = Asset
	GenTypes("types.yaml")
}
