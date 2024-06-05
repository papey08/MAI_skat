package httpserver

type refreshRequest struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type logoutRequest struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
