// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthStartRequestSubjectMobileAuthenticatorInstant struct {
	Consent *AuthStartRequestInstantAuthConsent `json:"consent,omitempty"`
	// MobileAuthTestMode simulates responses. Acceptable values are: off (default), pass, failStart, failFinish, and
	// failService.
	TestMode *string `json:"testMode,omitempty"`
}

func (o *AuthStartRequestSubjectMobileAuthenticatorInstant) GetConsent() *AuthStartRequestInstantAuthConsent {
	if o == nil {
		return nil
	}
	return o.Consent
}

func (o *AuthStartRequestSubjectMobileAuthenticatorInstant) GetTestMode() *string {
	if o == nil {
		return nil
	}
	return o.TestMode
}
