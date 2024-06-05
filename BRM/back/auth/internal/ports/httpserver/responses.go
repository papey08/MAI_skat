package httpserver

func errorResponse(err error) tokensResponse {
	if err == nil {
		return tokensResponse{}
	}
	errStr := err.Error()
	return tokensResponse{
		Data: nil,
		Err:  &errStr,
	}
}

type tokensData struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type tokensResponse struct {
	Data *tokensData `json:"data"`
	Err  *string     `json:"error"`
}

type logoutResponse struct {
	Data any     `json:"data"`
	Err  *string `json:"error"`
}
