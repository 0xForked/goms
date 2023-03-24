package rest_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type bffRegistrarRESTHandlerTestSuite struct {
	suite.Suite
}

func (suite *bffRegistrarRESTHandlerTestSuite) SetupSuite() {

}

func TestBFFRegistrarRESTHandler(t *testing.T) {
	suite.Run(t, new(bffRegistrarRESTHandlerTestSuite))
}
