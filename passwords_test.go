package passwork

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	apiKey   string = os.Getenv("PASSWORK_API_KEY")
	host     string = os.Getenv("PASSWORK_HOST")
	vaultId  string = os.Getenv("PASSWORK_VAULT_ID")
	folderId string = os.Getenv("PASSWORK_FOLDER_ID")
	pwId     string
	pwName   string
)

func TestAddPassword(t *testing.T) {
	c := NewClient(host, apiKey)
	err := c.Login()
	if err != nil {
		log.Fatal("Could not login to Passwork!, Aborting test.")
		return
	}

	request := PasswordRequest{
		Name:            "provider-test-entry",
		VaultId:         vaultId,
		FolderId:        folderId,
		Login:           "provider-test-user",
		CryptedPassword: "cHJvdmlkZXItdGVzdC1wYXNzd29yZA==",
		Description:     "provider-test-description",
		Url:             "https://login.com",
		Color:           1,
		Tags:            []string{"test", "foo", "bar"},
	}

	// Create password
	result, err := c.AddPassword(request)

	assert.Equal(t, nil, err, "AddPassword() should not return an error.")
	assert.Equal(t, "success", result.Status, "AddPassword() should return success.")
	assert.Equal(t, "provider-test-entry", result.Data.Name, "Result name should be the same as request name.")
	assert.Equal(t, vaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
	assert.Equal(t, folderId, result.Data.FolderId, "Result FolderId should be the same as request FolderId.")
	assert.Equal(t, "provider-test-user", result.Data.Login, "Result Login should be the same as request Login.")
	assert.Equal(t, "cHJvdmlkZXItdGVzdC1wYXNzd29yZA==", result.Data.CryptedPassword, "Result CryptedPassword should be the same as request CryptedPassword.")
	assert.Equal(t, "provider-test-description", result.Data.Description, "Result Description should be the same as request Description.")
	assert.Equal(t, "https://login.com", result.Data.Url, "Result Url should be the same as request Url.")
	assert.Equal(t, 1, result.Data.Color, "Result Color should be the same as request Color.")
	assert.Equal(t, 3, len(result.Data.Tags), "Number of result Tags should be the same as number of request Tags.")

	// Set pwId for subsequent test cases
	pwId = result.Data.Id
}

func TestEditPassword(t *testing.T) {
	c := NewClient(host, apiKey)
	err := c.Login()
	if err != nil {
		log.Fatal("Could not login to Passwork!, Aborting test.")
		return
	}

	request := PasswordRequest{
		Name:            "provider-test-entry-changed",
		VaultId:         vaultId,
		FolderId:        folderId,
		Login:           "provider-test-user-changed",
		CryptedPassword: "cHJvdmlkZXItdGVzdC1wYXNzd29yZC1jaGFuZ2Vk",
		Description:     "provider-test-description-changed",
		Url:             "https://login-changed.com",
		Color:           2,
		Tags:            []string{"changed", "bar"},
	}

	// Edit password
	result, err := c.EditPassword(pwId, request)

	assert.Equal(t, nil, err, "EditPassword() should not return an error.")
	assert.Equal(t, pwId, result.Data.Id, "Result Password ID should be the same as previously created Password ID.")
	assert.Equal(t, "success", result.Status, "EditPassword() should return success.")
	assert.Equal(t, "provider-test-entry-changed", result.Data.Name, "Result name should be the same as request name.")
	assert.Equal(t, vaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
	assert.Equal(t, folderId, result.Data.FolderId, "Result FolderId should be the same as request FolderId.")
	assert.Equal(t, "provider-test-user-changed", result.Data.Login, "Result Login should be the same as request Login.")
	assert.Equal(t, "cHJvdmlkZXItdGVzdC1wYXNzd29yZC1jaGFuZ2Vk", result.Data.CryptedPassword, "Result CryptedPassword should be the same as request CryptedPassword.")
	assert.Equal(t, "provider-test-description-changed", result.Data.Description, "Result Description should be the same as request Description.")
	assert.Equal(t, "https://login-changed.com", result.Data.Url, "Result Url should be the same as request Url.")
	assert.Equal(t, 2, result.Data.Color, "Result Color should be the same as request Color.")
	assert.Equal(t, 2, len(result.Data.Tags), "Number of result Tags should be the same as number of request Tags.")

	// Set pwName for subsequent test cases
	pwName = result.Data.Name
}

func TestSearchPassword(t *testing.T) {
	c := NewClient(host, apiKey)
	err := c.Login()
	if err != nil {
		log.Fatal("Could not login to Passwork!, Aborting test.")
		return
	}

	request := PasswordSearchRequest{
		Query:   pwName,
		VaultId: vaultId,
	}

	result, err := c.SearchPassword(request)

	assert.Equal(t, nil, err, "SearchPassword() should not return an error.")
	assert.Equal(t, "success", result.Status, "SearchPassword() should return success.")
	assert.Equal(t, "provider-test-entry-changed", result.Data[0].Name, "Result name should be the same as request name.")
	assert.Equal(t, vaultId, result.Data[0].VaultId, "Result VaultId should be the same as request VaultId.")
	assert.Equal(t, folderId, result.Data[0].FolderId, "Result FolderId should be the same as request FolderId.")
	assert.Equal(t, "provider-test-user-changed", result.Data[0].Login, "Result Login should be the same as request Login.")
	assert.Equal(t, pwId, result.Data[0].Id, "Result Password ID should be the same as previously created Password ID.")
	assert.Equal(t, "provider-test-description-changed", result.Data[0].Description, "Result Description should be the same as request Description.")
	assert.Equal(t, "https://login-changed.com", result.Data[0].Url, "Result Url should be the same as request Url.")
	assert.Equal(t, 2, result.Data[0].Color, "Result Color should be the same as request Color.")
	assert.Equal(t, 2, len(result.Data[0].Tags), "Number of result Tags should be the same as number of request Tags.")
}

func TestGetPassword(t *testing.T) {
	c := NewClient(host, apiKey)
	err := c.Login()
	if err != nil {
		log.Fatal("Could not login to Passwork!, Aborting test.")
		return
	}

	result, err := c.GetPassword(pwId)

	assert.Equal(t, nil, err, "GetPassword() should not return an error.")
	assert.Equal(t, pwId, result.Data.Id, "Result Password ID should be the same as previously created Password ID.")
	assert.Equal(t, "success", result.Status, "GetPassword() should return success.")
	assert.Equal(t, "provider-test-entry-changed", result.Data.Name, "Result name should be the same as request name.")
	assert.Equal(t, vaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
	assert.Equal(t, folderId, result.Data.FolderId, "Result FolderId should be the same as request FolderId.")
	assert.Equal(t, "provider-test-user-changed", result.Data.Login, "Result Login should be the same as request Login.")
	assert.Equal(t, "cHJvdmlkZXItdGVzdC1wYXNzd29yZC1jaGFuZ2Vk", result.Data.CryptedPassword, "Result CryptedPassword should be the same as request CryptedPassword.")
	assert.Equal(t, "provider-test-description-changed", result.Data.Description, "Result Description should be the same as request Description.")
	assert.Equal(t, "https://login-changed.com", result.Data.Url, "Result Url should be the same as request Url.")
	assert.Equal(t, 2, result.Data.Color, "Result Color should be the same as request Color.")
	assert.Equal(t, 2, len(result.Data.Tags), "Number of result Tags should be the same as number of request Tags.")
}

func TestDeletePassword(t *testing.T) {
	c := NewClient(host, apiKey)
	err := c.Login()
	if err != nil {
		log.Fatal("Could not login to Passwork!, Aborting test.")
		return
	}

	result, err := c.DeletePassword(pwId)

	assert.Equal(t, nil, err, "DeletePassword() should not return an error.")
	assert.Equal(t, "success", result.Status, "DeletePassword() should return success.")
}
