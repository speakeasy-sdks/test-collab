// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponseError struct {
	// Code is an internal error code that describes the problem category of the request.
	Code *int64 `json:"code,omitempty"`
	// Message is an error message describing the problem with the request.
	Message string `json:"message"`
}

func (o *AuthFinishResponseError) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *AuthFinishResponseError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}
