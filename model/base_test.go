package model

import (
	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"time"
)



type BaseTestSuite struct {
	suite.Suite
	b *base
	fakeID, fakeETag string
	pastDate, pasterDate, futureDate, nowDate time.Time
	deleted bool
}

func (suite *BaseTestSuite) SetupTest() {
	//fmt.Println("here...")
	suite.b = newBase()
	suite.fakeID = shared.NewUUID()
	suite.pastDate = time.Now().UTC().AddDate(0, -1, 0)
	suite.pasterDate = time.Now().UTC().AddDate(0, -2, 0)
	suite.futureDate = time.Now().UTC().AddDate(0, 1, 0)
	suite.nowDate = time.Now().UTC()
	suite.fakeETag = shared.NewUUID()
	suite.deleted = false
}

func (suite *BaseTestSuite) TestNewBase() {
	assert.NotNil(suite.T(), suite.b.ID(), "Missing Id from Base")
	assert.Nil(suite.T(), shared.ParseUUID(suite.b.ID()), "Unable to parse ID() into valid guid format")
	assert.NotNil(suite.T(), suite.b.DateCreatedUTC(), "Unable to set the create date")
	assert.NotNil(suite.T(), suite.b.DateUpdatedUTC(), "Unable to set the update date")
}

func (suite *BaseTestSuite) TestModelInitinializeWithFakeEtag() {
	err := suite.b.LoadBase(suite.fakeID, suite.nowDate, suite.nowDate, suite.deleted, "fakeETag")
	assert.Error(suite.T(), err)
}

func (suite *BaseTestSuite) TestModelHanedlesInvalidGuidErrorProperly() {
	//model := &base{}
	err := suite.b.LoadBase("fakeID", time.Now().UTC(), time.Now().UTC(), suite.deleted, "etag")
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errInvalidGuid, err.Error())
}
func (suite *BaseTestSuite) TestLoadExistingModel() {
	b, err := LoadExistingBase(suite.fakeID, suite.nowDate, suite.nowDate, suite.deleted)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), b)
	assert.Equal(suite.T(), suite.fakeID, b.ID())
	assert.Equal(suite.T(), suite.nowDate, suite.b.DateCreatedUTC())
	assert.Equal(suite.T(), suite.nowDate, suite.b.DateUpdatedUTC())
	assert.Equal(suite.T(), suite.deleted, suite.b.Deleted())
	//assert.Equal(suite.T(), suite.fakeETag, suite.b.ETag())
}

func (suite *BaseTestSuite) TestLoadModelWithFutureAndPastDateValidation() {
	_, err := LoadExistingBase(suite.fakeID, suite.nowDate, suite.futureDate, suite.deleted)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errFutureUpdated, err.Error())

	_, err = LoadExistingBase(suite.fakeID, suite.pastDate, suite.pasterDate, suite.deleted)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errBeforeCreated, err.Error())
}
