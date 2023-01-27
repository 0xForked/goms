package service_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type bffServiceTestSuite struct {
	suite.Suite
}

func (suite *bffServiceTestSuite) SetupSuite() {

}

func TestBFFService(t *testing.T) {
	suite.Run(t, new(bffServiceTestSuite))
}
