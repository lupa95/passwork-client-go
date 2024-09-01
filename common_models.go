package passwork

type PathData struct {
	Order int
	Name  string
	Type  string // Allowed: vault, folder, inbox
	Id    string
}

type DeleteResponse struct {
	Status string
	Code   string
	Data   string
}
