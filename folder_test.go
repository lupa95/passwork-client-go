package passwork

func (suite *PassworkTestSuite) TestAddFolder() string {

	request := FolderRequest{
		Name:    "provider-test-folder",
		VaultId: suite.VaultId,
	}

	// Create folder
	result, err := suite.client.AddFolder(request)

	if suite.NoError(err) {
		suite.Equal("success", result.Status, "AddFolder() should return success.")
		suite.Equal("folderCreated", result.Code, "AddFolder() should return folderCreated.")
		suite.Equal("provider-test-folder", result.Data.Name, "Result name should be the same as request name.")
		suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
		suite.Equal("", result.Data.ParentId, "Result parentId should be empty (top level folder).")
		suite.Equal(0, result.Data.FoldersAmount, "Result CryptedPassword should be the same as request CryptedPassword.")
		suite.Equal(0, result.Data.FoldersAmount, "Result Description should be the same as request Description.")
		suite.Equal(0, result.Data.ShortcutsAmount, "Result Url should be the same as request Url.")
		suite.Equal(0, result.Data.Access, "Result Color should be the same as request Color.")

	}

	return result.Data.Id
}
