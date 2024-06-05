package leads

type updateLeadRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Status      string `json:"status"`
	Responsible uint64 `json:"responsible"`
}
