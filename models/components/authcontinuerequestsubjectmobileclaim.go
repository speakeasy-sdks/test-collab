// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthContinueRequestSubjectMobileClaim struct {
	// MobileNumber to bind to the device. Mobile number may be prefixed with a plus sign, but does not need to be for
	// US. A plus sign is required for international.
	MobileNumber string `json:"mobileNumber"`
}

func (o *AuthContinueRequestSubjectMobileClaim) GetMobileNumber() string {
	if o == nil {
		return ""
	}
	return o.MobileNumber
}
