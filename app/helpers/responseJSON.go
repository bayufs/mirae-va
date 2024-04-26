package helpers

type Response struct {
	Status      uint        `json:"status"`
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
	Meta        interface{} `json:"meta"`
}

func CreateResponse(status uint, message string, description string, additionalData ...interface{}) Response {
	var responseData interface{}
	if len(additionalData) > 0 {
		responseData = additionalData[0]
	} else {
		responseData = map[string]interface{}{}
	}
	var responseMeta interface{}
	if len(additionalData) > 1 {
		responseMeta = additionalData[1]
	} else {
		responseMeta = map[string]interface{}{}
	}

	return Response{
		status,
		message,
		description,
		responseData,
		responseMeta,
	}
}
