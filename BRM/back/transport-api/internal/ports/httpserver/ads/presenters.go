package ads

type addAdRequest struct {
	Title    string `json:"title"`
	Text     string `json:"text"`
	Industry string `json:"industry"`
	Price    uint   `json:"price"`
	ImageURL string `json:"image_url"`
}

type updateAdRequest struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	Industry    string `json:"industry"`
	Price       uint   `json:"price"`
	Responsible uint64 `json:"responsible"`
	ImageURL    string `json:"image_url"`
}
