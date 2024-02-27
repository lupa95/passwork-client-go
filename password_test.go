package passwork

func (suite *PassworkTestSuite) TestPassword() {
	suite.Run("Add", func() {
		request := PasswordRequest{
			Name:            "provider-test-entry",
			VaultId:         suite.VaultId,
			Login:           "provider-test-user",
			CryptedPassword: "cHJvdmlkZXItdGVzdC1wYXNzd29yZA==",
			Description:     "provider-test-description",
			Url:             "https://login.com",
			Color:           1,
			Tags:            []string{"test", "foo", "bar"},
		}

		result, err := suite.client.AddPassword(request)

		if suite.NoError(err) {
			suite.Equal(nil, err, "AddPassword() should not return an error.")
			suite.Equal("success", result.Status, "AddPassword() should return success.")
			suite.Equal("provider-test-entry", result.Data.Name, "Result name should be the same as request name.")
			suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
			suite.Equal("provider-test-user", result.Data.Login, "Result Login should be the same as request Login.")
			suite.Equal("cHJvdmlkZXItdGVzdC1wYXNzd29yZA==", result.Data.CryptedPassword, "Result CryptedPassword should be the same as request CryptedPassword.")
			suite.Equal("provider-test-description", result.Data.Description, "Result Description should be the same as request Description.")
			suite.Equal("https://login.com", result.Data.Url, "Result Url should be the same as request Url.")
			suite.Equal(1, result.Data.Color, "Result Color should be the same as request Color.")
			suite.Equal(3, len(result.Data.Tags), "Number of result Tags should be the same as number of request Tags.")
		}
		suite.PasswordId = result.Data.Id
	})

	suite.Run("Edit", func() {
		request := PasswordRequest{
			Name:            "provider-test-entry-changed",
			VaultId:         suite.VaultId,
			Login:           "provider-test-user-changed",
			CryptedPassword: "cHJvdmlkZXItdGVzdC1wYXNzd29yZC1jaGFuZ2Vk",
			Description:     "provider-test-description-changed",
			Url:             "https://login-changed.com",
			Color:           2,
			Tags:            []string{"changed", "bar"},
		}

		result, err := suite.client.EditPassword(suite.PasswordId, request)

		suite.Equal(nil, err, "EditPassword() should not return an error.")
		suite.Equal(suite.PasswordId, result.Data.Id, "Result Password ID should be the same as previously created Password ID.")
		suite.Equal("success", result.Status, "EditPassword() should return success.")
		suite.Equal("provider-test-entry-changed", result.Data.Name, "Result name should be the same as request name.")
		suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
		suite.Equal("provider-test-user-changed", result.Data.Login, "Result Login should be the same as request Login.")
		suite.Equal("cHJvdmlkZXItdGVzdC1wYXNzd29yZC1jaGFuZ2Vk", result.Data.CryptedPassword, "Result CryptedPassword should be the same as request CryptedPassword.")
		suite.Equal("provider-test-description-changed", result.Data.Description, "Result Description should be the same as request Description.")
		suite.Equal("https://login-changed.com", result.Data.Url, "Result Url should be the same as request Url.")
		suite.Equal(2, result.Data.Color, "Result Color should be the same as request Color.")
		suite.Equal(2, len(result.Data.Tags), "Number of result Tags should be the same as number of request Tags.")

		suite.PasswordName = result.Data.Name
	})

	suite.Run("Search", func() {
		request := PasswordSearchRequest{
			Query:   suite.PasswordName,
			VaultId: suite.VaultId,
		}

		result, err := suite.client.SearchPassword(request)

		suite.Equal(nil, err, "SearchPassword() should not return an error.")
		suite.Equal("success", result.Status, "SearchPassword() should return success.")
		suite.Equal("provider-test-entry-changed", result.Data[0].Name, "Result name should be the same as request name.")
		suite.Equal(suite.VaultId, result.Data[0].VaultId, "Result VaultId should be the same as request VaultId.")
		suite.Equal("provider-test-user-changed", result.Data[0].Login, "Result Login should be the same as request Login.")
		suite.Equal(suite.PasswordId, result.Data[0].Id, "Result Password ID should be the same as previously created Password ID.")
		suite.Equal("provider-test-description-changed", result.Data[0].Description, "Result Description should be the same as request Description.")
		suite.Equal("https://login-changed.com", result.Data[0].Url, "Result Url should be the same as request Url.")
		suite.Equal(2, result.Data[0].Color, "Result Color should be the same as request Color.")
		suite.Equal(2, len(result.Data[0].Tags), "Number of result Tags should be the same as number of request Tags.")
	})

	suite.Run("Get", func() {
		result, err := suite.client.GetPassword(suite.PasswordId)

		suite.Equal(nil, err, "GetPassword() should not return an error.")
		suite.Equal(suite.PasswordId, result.Data.Id, "Result Password ID should be the same as previously created Password ID.")
		suite.Equal("success", result.Status, "GetPassword() should return success.")
		suite.Equal("provider-test-entry-changed", result.Data.Name, "Result name should be the same as request name.")
		suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
		suite.Equal("provider-test-user-changed", result.Data.Login, "Result Login should be the same as request Login.")
		suite.Equal("cHJvdmlkZXItdGVzdC1wYXNzd29yZC1jaGFuZ2Vk", result.Data.CryptedPassword, "Result CryptedPassword should be the same as request CryptedPassword.")
		suite.Equal("provider-test-description-changed", result.Data.Description, "Result Description should be the same as request Description.")
		suite.Equal("https://login-changed.com", result.Data.Url, "Result Url should be the same as request Url.")
		suite.Equal(2, result.Data.Color, "Result Color should be the same as request Color.")
		suite.Equal(2, len(result.Data.Tags), "Number of result Tags should be the same as number of request Tags.")
	})

	suite.Run("Delete", func() {
		result, err := suite.client.DeletePassword(suite.PasswordId)

		suite.Equal(nil, err, "DeletePassword() should not return an error.")
		suite.Equal("success", result.Status, "DeletePassword() should return success.")
	})
}
