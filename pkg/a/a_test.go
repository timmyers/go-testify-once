package a_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/timmyers/go-testify-once/pkg/a"
	"github.com/timmyers/go-testify-once/pkg/test"
)

type ASuite struct {
	suite.Suite
}

func (suite *ASuite) SetupTest() {
	test.Init()
}

func (suite *ASuite) TestDoA() {
	a.DoA()
	assert := assert.New(suite.T())
	assert.Equal(1, test.Count)
}

func TestSuiteA(t *testing.T) {
	suite.Run(t, new(ASuite))
}
