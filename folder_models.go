package passwork

type FolderResponse struct {
	Status string
	Code   string // folderCreated, folderRenamed
	Data   FolderResponseData
}

type FolderSearchResponse struct {
	Status string
	Code   string
	Data   []FolderResponseData
}

type FolderDeleteResponse struct {
	Status string
	Code   string
	Data   string // folderDeleted
}

type FolderResponseData struct {
	VaultId         string
	Name            string
	Id              string
	ParentId        string
	Path            []PathData
	FoldersAmount   int
	PasswordsAmount int
	ShortcutsAmount int
	Access          int
}

type FolderRequest struct {
	VaultId  string `json:"vaultId,omitempty"`
	Name     string `json:"name"`
	ParentId string `json:"parentId,omitempty"`
}

type FolderSearchRequest struct {
	Query   string `json:"query"`
	VaultId string `json:"vaultId"`
}
