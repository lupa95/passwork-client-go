package passwork

func (suite *PassworkTestSuite) TestFolder() {
	suite.Run("Add", func() {
		folderName := "provider-test-folder"
		request := FolderRequest{
			Name:    folderName,
			VaultId: suite.VaultId,
		}

		result, err := suite.client.AddFolder(request)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "AddFolder() should return success.")
			suite.Equal("folderCreated", result.Code, "AddFolder() should return folderCreated.")
			suite.Equal(folderName, result.Data.Name, "Result name should be the same as request name.")
			suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
			suite.Equal("", result.Data.ParentId, "Result parentId should be empty (top level folder).")
			suite.Equal(0, result.Data.FoldersAmount, "Result CryptedPassword should be the same as request CryptedPassword.")
			suite.Equal(0, result.Data.FoldersAmount, "Result Description should be the same as request Description.")
			suite.Equal(0, result.Data.ShortcutsAmount, "Result Url should be the same as request Url.")
			suite.Equal(0, result.Data.Access, "Result Color should be the same as request Color.")
		}

		suite.FolderId = result.Data.Id
	})

	suite.Run("Edit", func() {
		newFolderName := "provider-test-folder-renamed"

		request := FolderRequest{
			Name: newFolderName,
		}

		result, err := suite.client.EditFolder(suite.FolderId, request)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "EditFolder() should return success.")
			suite.Equal("folderRenamed", result.Code, "EditFolder() should return folderRenamed.")
			suite.Equal(newFolderName, result.Data.Name, "Result name should be the same as request name.")
			suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as request VaultId.")
			suite.Equal("", result.Data.ParentId, "Result parentId should be empty (top level folder).")
			suite.Equal(0, result.Data.FoldersAmount, "Result CryptedPassword should be the same as request CryptedPassword.")
			suite.Equal(0, result.Data.FoldersAmount, "Result Description should be the same as request Description.")
			suite.Equal(0, result.Data.ShortcutsAmount, "Result Url should be the same as request Url.")
			suite.Equal(0, result.Data.Access, "Result Color should be the same as request Color.")
		}

		suite.FolderName = result.Data.Name
	})

	suite.Run("Search", func() {
		request := FolderSearchRequest{
			VaultId: suite.VaultId,
			Query:   suite.FolderName,
		}

		result, err := suite.client.SearchFolder(request)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "EditFolder() should return success.")
			suite.Equal(suite.FolderName, result.Data[0].Name, "Result name should be the same as request name.")
			suite.Equal(suite.VaultId, result.Data[0].VaultId, "Result VaultId should be the same as request VaultId.")
			suite.Equal("", result.Data[0].ParentId, "Result parentId should be empty (top level folder).")
			suite.Equal(suite.FolderId, result.Data[0].Id, "Result FolderId should be same as previously created folder.")
		}
	})

	suite.Run("Get", func() {
		result, err := suite.client.GetFolder(suite.FolderId)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "EditFolder() should return success.")
			suite.Equal(suite.FolderName, result.Data.Name, "Result FolderName should be the same as suite FolderName.")
			suite.Equal(suite.VaultId, result.Data.VaultId, "Result VaultId should be the same as suite VaultId.")
			suite.Equal("", result.Data.ParentId, "Result parentId should be empty (top level folder).")
			suite.Equal(suite.FolderId, result.Data.Id, "Result Id should be the same as suite FolderId.")
		}
	})

	suite.Run("Delete", func() {
		result, err := suite.client.DeleteFolder(suite.FolderId)

		if suite.NoError(err) {
			suite.Equal("success", result.Status, "DeleteFolder() should return success.")
			suite.Equal("folderDeleted", result.Data, "Result data should be folderDeleted.")
		}
	})
}
