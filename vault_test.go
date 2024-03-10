package passwork

import "fmt"

func (suite *PassworkTestSuite) TestVault() {
	suite.Run("Add", func() {
		vaultName := "provider-test-vault"
		request := VaultAddRequest{
			Name:         vaultName,
			IsPrivate:    true,
			PasswordHash: "test123",
			Salt:         "test123",
			MpCrypted:    "test123",
		}

		result, err := suite.client.AddVault(request)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "AddVault() should return success.")
			suite.Equal("vaultCreated", result.Code, "AddVault() should return vaultUpdated.")
		}

		suite.VaultId = result.Data
	})

	suite.Run("Edit", func() {
		newVaultName := "provider-test-vault-renamed"

		request := VaultEditRequest{
			Name: newVaultName,
		}

		result, err := suite.client.EditVault(suite.VaultId, request)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "EditVault() should return success.")
			suite.Equal("vaultUpdated", result.Code, "EditVault() should return vaultUpdated.")
			suite.Equal(suite.VaultId, result.Data, "Result VaultId should be the same as request VaultId.")
		}

		suite.VaultName = newVaultName
	})

	suite.Run("Get", func() {
		result, err := suite.client.GetVault(suite.VaultId)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "EditVault() should return success.")
			suite.Equal(suite.VaultId, result.Data.Id, "Result VaultId should be the same as request VaultId.")
			suite.Equal(suite.VaultName, result.Data.Name, "Result VaultName should be the same as suite VaultName.")
		}
	})

	suite.Run("Delete", func() {
		result, err := suite.client.DeleteVault(suite.VaultId)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "DeleteVault() should return success.")
			suite.Equal(fmt.Sprintf("Vault %s deleted", suite.VaultId), result.Data, "Result data should be vaultDeleted.")
		}
	})
}
