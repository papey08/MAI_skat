package companies

type updateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    string `json:"industry"`
	OwnerId     uint64 `json:"owner_id"`
}
