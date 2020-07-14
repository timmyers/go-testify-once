package b_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/timmyers/go-testify-once/pkg/b"
	"github.com/timmyers/go-testify-once/pkg/test"
)

type BSuite struct {
	suite.Suite
}

func (suite *BSuite) SetupTest() {
	test.Init()
}

func (suite *BSuite) TestDoB() {
	b.DoB()
	assert := assert.New(suite.T())
	assert.Equal(1, test.Count)
}

func TestSuiteB(t *testing.T) {
	suite.Run(t, new(BSuite))
}
