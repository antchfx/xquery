package xmlquery

import (
	"github.com/antchfx/xpath"
	"github.com/antchfx/xpath/test"
	"testing"
)

func TestXQuery(t *testing.T) {
	create := func(doc string) (xpath.NodeNavigator, error) {
		return CreateXPathNavigator(loadXML(doc)), nil
	}

	xquerytest.TestAll(t, create, xquerytest.EnableAll)
}
