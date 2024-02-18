package passwork

type FolderResponse struct {
	Status string
	Data   FolderResponseData
	Code   string
}

type FolderSearchResponse struct {
	Status string
	Data   []FolderResponseData
}

type FolderResponseData struct {
	VaultId         string
	Name            string
	Id              string
	ParentId        string
	Path            PathData
	FoldersAmount   int
	PasswordsAmount int
	Access          int
}

type FolderAddRequest struct {
	VaultId  string `json:"vaultId"`
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}

type FolderSearchRequest struct {
	Query   string `json:"query"`
	VaultId string `json:"vaultId"`
}
