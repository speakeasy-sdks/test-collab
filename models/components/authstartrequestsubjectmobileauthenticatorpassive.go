// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthStartRequestSubjectMobileAuthenticatorPassive struct {
	// If true, bind will persist between authenticators. If false, bind will be reattempted on every auth flow.
	CacheResult *bool `json:"cacheResult,omitempty"`
	// If true, will allow authentication against a local domain for testing FIDO.
	LocalDomain *bool `json:"localDomain,omitempty"`
	// If not set or set to 0, the age of the bind will be ignored during authentication. If set to -1, then it will
	// force a rebind of the mobile number to the device. If set greater than 0, it specifies the maximum age
	// (in seconds) of the bind that is allowed or else a rebind occurs during authentication. Ex: 7776000 is 90 days.
	MaxAge   *int64  `json:"maxAge,omitempty"`
	TestMode *string `json:"testMode,omitempty"`
}

func (o *AuthStartRequestSubjectMobileAuthenticatorPassive) GetCacheResult() *bool {
	if o == nil {
		return nil
	}
	return o.CacheResult
}

func (o *AuthStartRequestSubjectMobileAuthenticatorPassive) GetLocalDomain() *bool {
	if o == nil {
		return nil
	}
	return o.LocalDomain
}

func (o *AuthStartRequestSubjectMobileAuthenticatorPassive) GetMaxAge() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxAge
}

func (o *AuthStartRequestSubjectMobileAuthenticatorPassive) GetTestMode() *string {
	if o == nil {
		return nil
	}
	return o.TestMode
}
