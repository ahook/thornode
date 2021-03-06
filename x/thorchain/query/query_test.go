package query

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestPackage(t *testing.T) { TestingT(t) }

type QuerySuite struct{}

var _ = Suite(&QuerySuite{})

func (s QuerySuite) TestQuery(c *C) {
	c.Check(QueryTx.Endpoint("foo", "bar"), Equals, "/foo/tx/{bar}")
	c.Check(QueryTx.Path("foo", "bar"), Equals, "custom/foo/tx/bar")
}
