package mymodels

type FormatResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func MyResponse(msg string, data ...interface{}) FormatResponse {
	var mydata interface{}
	if len(data) != 0 {
		mydata = data[0]
	}
	return FormatResponse{
		Message: msg,
		Data:    mydata,
	}
}
