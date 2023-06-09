package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, data any) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
func FormatValidationError(err error) []string {
	var errors []string
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			errors = append(errors, e.Error())
		}
	} else {
		errors = append(errors, err.Error())
	}
	return errors
	//var errors []string
	//for _, e := range err.(validator.ValidationErrors) {
	//	errors = append(errors, e.Error())
	//}
	//return errors
}
