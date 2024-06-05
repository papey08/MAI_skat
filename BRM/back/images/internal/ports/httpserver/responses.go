package httpserver

type idResponse struct {
	Id  *uint64 `json:"data"`
	Err *string `json:"err"`
}

func errorResponse(err error) idResponse {
	if err == nil {
		return idResponse{}
	}
	errStr := err.Error()
	return idResponse{
		Err: &errStr,
	}
}
