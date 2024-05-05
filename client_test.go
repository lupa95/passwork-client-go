package passwork

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PassworkTestSuite struct {
	suite.Suite
	ApiKey       string
	Host         string
	VaultId      string
	VaultName    string
	FolderId     string
	FolderName   string
	PasswordId   string
	PasswordName string
	client       *Client
}

func (suite *PassworkTestSuite) SetupSuite() {
	suite.ApiKey = os.Getenv("PASSWORK_API_KEY")
	suite.Host = os.Getenv("PASSWORK_HOST")
	suite.VaultId = os.Getenv("PASSWORK_VAULT_ID")

	suite.client = NewClient(suite.Host, suite.ApiKey, time.Second*30)
	err := suite.client.Login()
	if err != nil {
		suite.Fail("Could not login to Passwork!, Aborting test suite.")
	}
}

func TestPassworkTestSuite(t *testing.T) {
	suite.Run(t, new(PassworkTestSuite))
}
