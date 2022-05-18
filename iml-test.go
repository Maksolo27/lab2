package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestPrefixToInfix(c *C) {
	examples := map[string]string{
		"+ 5 * - 4 2 3": "3 2 4 - * 5 +",
		"+ 2 3":         "3 2 +",
	}

	for prefix, infix := range examples {
		res, err := PrefixToPostfix(prefix)
		if err != nil {
			c.Assert(err, ErrorMatches, infix)
		} else {
			c.Assert(res, Equals, infix)
		}
	}
}

func ExamplePrefixToPostfix() {
	res, err := PrefixToPostfix("+ 2 3")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}
}
