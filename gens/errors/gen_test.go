package generrors

import (
	"github.com/domego/gopt/gens/common"
	"testing"
)

func TestGenErrors(t *testing.T) {
	genutils.Asset = Asset
	Gen("errors.yaml")
}
