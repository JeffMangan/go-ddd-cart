package model

import (
	"fmt"
	"github.com/JeffMangan/go-ddd-cart/shared"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite
	u, uu                                   *User
	bytes                                   []byte
	firstName, lastName, displayName, email string
	params                                  map[string]interface{}
}

func (suite *UserTestSuite) SetupTest() {
	//fmt.Println("here....")
	suite.firstName = "fName"
	suite.lastName = "lName"
	suite.displayName = "DisplayNamex"
	suite.email = "Emailx@somethingheretest.com"

	params := make(map[string]interface{})
	params["FirstName"] = suite.firstName
	params["LastName"] = suite.lastName
	params["DisplayName"] = suite.displayName
	params["Email"] = suite.email
	suite.params = params
	suite.u, _ = NewUser(suite.params)

	suite.bytes = []byte(fmt.Sprintf(`{
		"ID": "%s",
		"DateCreatedUTC": "%s",
		"DateUpdatedUTC": "%s",
		"Email": "%s",
		"FirstName": "%s",
		"LastName": "%s",
		"DisplayName": "%s",
		"Deleted": "%t"
	}`, suite.u.ID(),
		suite.u.DateCreatedUTCx.UTC().Format("2006-01-02 15:04:05"),
		suite.u.DateUpdatedUTCx.UTC().Format("2006-01-02 15:04:05"),
		suite.u.Emailx,
		suite.u.FirstNamex,
		suite.u.LastNamex,
		suite.u.DisplayNamex,
		suite.u.Deletedx,
		//suite.u.tracerID,
	))
}

func (suite *UserTestSuite) TestBaseFieldsLoaded() {
	assert.EqualValues(suite.T(), shared.ParseUUID, suite.u.ID(), nil)
	assert.NotEmpty(suite.T(), suite.uu.DateCreatedUTC())
	assert.NotEmpty(suite.T(), suite.uu.DateUpdatedUTC())
}

func (suite *UserTestSuite) TestUserFieldsPopulated() {
	assert.NotNil(suite.T(), suite.u.LastName(), "Expected to not be empty")
	assert.NotNil(suite.T(), suite.u.FirstName(), "Expected to not be empty")
	assert.NotNil(suite.T(), suite.u.DisplayName(), "Expected to not be empty")
	assert.NotNil(suite.T(), suite.u.Email(), "Expected to not be empty")
}

func (suite *UserTestSuite) TestGettersAndSetters() {
	//suite.uu, _ = NewUser(suite.params)
	assert.Equal(suite.T(), suite.uu.Email(), suite.u.Email())
	assert.Equal(suite.T(), suite.uu.FirstName(), suite.u.FirstName())
	assert.Equal(suite.T(), suite.uu.LastName(), suite.u.LastName())
	assert.Equal(suite.T(), suite.uu.DisplayName(), suite.u.DisplayName())
	assert.NotNil(suite.T(), suite.u.LastName(), "Expected to not be empty")
	assert.NotNil(suite.T(), suite.uu.FirstName(), "Expected to not be empty")
	assert.NotNil(suite.T(), suite.uu.DisplayName(), "Expected to not be empty")
	assert.NotNil(suite.T(), suite.uu.Email(), "Expected to not be empty")
}

func (suite *UserTestSuite) TestNewUser() {
	assert.Equal(suite.T(), "*model.User", fmt.Sprintf("%T", suite.uu))
}

func (suite *UserTestSuite) TestDisplayName() {
	panic("blah")
	//fmt.Println(suite.u.SetDisplayName("jasldfkjalsfjaslfdjsaljfsalkfjaslfjdslfj"))
	assert.NotNil(suite.T(), suite.u.SetDisplayName("x"), "Expected an err but was empty")
	assert.NotNil(suite.T(), suite.u.SetDisplayName(")"), "Expected an err but was empty")
	assert.NotNil(suite.T(), suite.u.SetDisplayName(""), "Expected an err but was empty")
	assert.NotNil(suite.T(), suite.u.SetDisplayName("jasldfkjalsfjaslfdjsaljfsalkfjaslfjdslfj"), "Expected an err but was empty")
}

func (suite *UserTestSuite) TestEmail() {
	//uu := NewUser(shared.NewUUID())

	assert.Nil(suite.T(), suite.uu.SetEmail(suite.u.Email()))
	assert.Equal(suite.T(), suite.uu.Email(), suite.uu.Email())
}

func (suite *UserTestSuite) TestEmailValidationFails() {
	//uu := NewUser(shared.NewUUID())
	assert.Error(suite.T(), suite.uu.SetEmail("Emailx.com"))
	assert.Error(suite.T(), suite.uu.SetEmail("Emailx@com"))
}
