// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthStartRequestSubjectMobileAuthenticatorOTP struct {
	// If AllowOtpRetry is set to true, the end user will have up to 3 OTP entry attempts in a Prove Auth flow. If
	// false, the end user will only get 1 OTP attempt. Once the maximum number of attempts is reached, the auth flow
	// will return an error.
	AllowOtpRetry *bool `json:"allowOtpRetry,omitempty"`
	// MessageText is the message that will be sent to end user with the one-time passcode (OTP) represented by pound
	// signs. The number of pound signs should be equal to the OTP length
	MessageText string  `json:"messageText"`
	TestMode    *string `json:"testMode,omitempty"`
}

func (o *AuthStartRequestSubjectMobileAuthenticatorOTP) GetAllowOtpRetry() *bool {
	if o == nil {
		return nil
	}
	return o.AllowOtpRetry
}

func (o *AuthStartRequestSubjectMobileAuthenticatorOTP) GetMessageText() string {
	if o == nil {
		return ""
	}
	return o.MessageText
}

func (o *AuthStartRequestSubjectMobileAuthenticatorOTP) GetTestMode() *string {
	if o == nil {
		return nil
	}
	return o.TestMode
}
