package client

type ClientResponse struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
	Err        string            `json:"err"`
}

/*
func (r *Response) SetOutputBody(body interface{}) *error {
	b, SystemError := json.Marshal(body)
	if SystemError == nil {
		r.Body = string(b)
		return nil
	} else {
		return &SystemError
	}
}
*/
func NewClientResponse() *ClientResponse {
	Output := &ClientResponse{}
	Output.Headers = make(map[string]string)
	Output.Headers["Access-Control-Allow-Origin"] = "*"
	return Output
}

