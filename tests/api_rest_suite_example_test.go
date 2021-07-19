package tests

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleApiRestSuite struct {
	suite.Suite
}

func TestExampleApiRestSuite(t *testing.T) {
	suite.Run(t, new(ExampleApiRestSuite))
}

func (suite *ExampleApiRestSuite) SetupSuite() {

}

func (suite *ExampleApiRestSuite) SetupTest() {

}

func (suite *ExampleApiRestSuite) TestGet() {

}
