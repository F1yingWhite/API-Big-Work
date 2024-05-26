package utils

type errorResponse struct {
	ErrorStr string `json:"error"`
}

func ErrorResponse(err error) errorResponse {
	if err == nil {
		return errorResponse{}
	}
	return errorResponse{
		ErrorStr: err.Error(),
	}
}

type response struct {
	Data any `json:"data,omitempty"`
}

func Response(obj any) response {
	if obj == nil {
		return response{}
	}
	return response{
		Data: obj,
	}
}
