package handlers

type Response struct {
	OK      bool        `json:"ok"`
	Error   string      `json:"error,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func (h *Handler) ok(data interface{}) *Response {
	return &Response{
		OK:      true,
		Details: data,
	}
}

func (h *Handler) error(err error) *Response {
	return &Response{
		OK:    false,
		Error: err.Error(),
	}
}
