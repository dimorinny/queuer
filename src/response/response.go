package response

func GenerateErrorResponse(code int, errorString string) map[string]interface{} {
	return map[string]interface{}{
		"Status": statusError,
		"Code":   code,
		"Error":  errorString,
	}
}

func GenerateSuccessResponse(response interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Status":   statusOk,
		"Code":     codeOk,
		"Response": response,
	}
}
