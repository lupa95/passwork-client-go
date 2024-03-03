package passwork

type VaultResponse struct {
	Status string
	Data   VaultResponseData
}

type VaultResponseData struct {
	Id                   string
	Name                 string
	VaultPasswordCrypted string
	Access               string
	Scope                string
}

type VaultAddRequest struct {
	Name            string `json:"name"`
	PasswordCrypted string `json:"passwordCrypted,omitempty"`
	PasswordHash    string `json:"passwordHash"`
	Salt            string `json:"salt"`
	IsPrivate       bool   `json:"isPrivate,omitempty"`
	MpCrypted       string `json:"mpCrypted"`
}

type VaultEditRequest struct {
	Name string `json:"name"`
}

type VaultEditResponse struct {
	Status string // success
	Code   string // always vaultUpdated
	Data   string // Id of the Vault
}

type VaultDeleteResponse struct {
	Status string // success
	Data   string // Vault <VaultId> deleted
}
