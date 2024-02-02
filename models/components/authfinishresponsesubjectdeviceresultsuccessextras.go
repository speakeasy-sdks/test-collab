// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponseSubjectDeviceResultSuccessExtras struct {
	// DeviceCapabilities contains the available device capabilities (webauthn, face, fingerprint).
	DeviceCapabilities []string `json:"deviceCapabilities,omitempty"`
	// DeviceAuthType contains description of the type of client device used for auth: 'web' or 'mobile'.
	DeviceType *string `json:"deviceType,omitempty"`
	// SupportedUvLevel indicates the level of user verification supported on the device.
	SupportedUvLevel []int64                                `json:"supportedUvLevel,omitempty"`
	Trust            *SubjectDeviceResultSuccessExtrasTrust `json:"trust,omitempty"`
}

func (o *AuthFinishResponseSubjectDeviceResultSuccessExtras) GetDeviceCapabilities() []string {
	if o == nil {
		return nil
	}
	return o.DeviceCapabilities
}

func (o *AuthFinishResponseSubjectDeviceResultSuccessExtras) GetDeviceType() *string {
	if o == nil {
		return nil
	}
	return o.DeviceType
}

func (o *AuthFinishResponseSubjectDeviceResultSuccessExtras) GetSupportedUvLevel() []int64 {
	if o == nil {
		return nil
	}
	return o.SupportedUvLevel
}

func (o *AuthFinishResponseSubjectDeviceResultSuccessExtras) GetTrust() *SubjectDeviceResultSuccessExtrasTrust {
	if o == nil {
		return nil
	}
	return o.Trust
}
