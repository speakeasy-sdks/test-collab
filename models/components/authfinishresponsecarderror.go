// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponseCardError struct {
	// Code is an internal error code that describes the problem category of the request.
	Code *int64 `json:"code,omitempty"`
	// Message is an error message describing the problem with the request.
	Message string `json:"message"`
	// PUID is the unique ID associated with an AirKey card that was extracted from the AirKey message.
	Puid *string `json:"puid,omitempty"`
}

func (o *AuthFinishResponseCardError) GetCode() *int64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *AuthFinishResponseCardError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *AuthFinishResponseCardError) GetPuid() *string {
	if o == nil {
		return nil
	}
	return o.Puid
}