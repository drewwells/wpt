package wpt

import (
	"io/ioutil"
	"testing"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&MySuite{})

type MySuite struct {
	bytes []byte
	locs  []Location
}

func (s *MySuite) SetUpTest(c *C) {

	var err error

	s.bytes, err = ioutil.ReadFile("./test/location.json")

	if err != nil {
		c.Log(err.Error())
	}

	s.locs, err = processLoc(s.bytes)
	if err != nil {
		c.Log(err.Error())
	}
}

func (s *MySuite) TestLocations(c *C) {
	loc0 := s.locs[0]

	c.Assert(loc0.Name, Equals, "wpt-001")
	c.Assert(loc0.Browser, Equals, "Chrome")
	c.Assert(loc0.Label, Equals, "uswest - Agent 001")
	c.Assert(loc0.Total, Equals, 1)
	c.Assert(loc0.Testing, Equals, 1)
	c.Assert(loc0.Busy, Equals, false)
	c.Assert(s.locs[4].Name, Equals, "wpt-002")
	c.Assert(s.locs[4].Busy, Equals, true)
}
