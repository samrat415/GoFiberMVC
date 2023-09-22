package Resources

type ApiResponse struct {
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

func (r *ApiResponse) SetProperties(success bool, data interface{}, message string) *ApiResponse {
	if !r.Success {
		r.Success = success
	}
	if r.Data == nil {
		r.Data = data
	}
	if r.Message == "" {
		r.Message = message
	}
	if r.StatusCode == 0 {
		r.StatusCode = 200
	}
	return r
}
