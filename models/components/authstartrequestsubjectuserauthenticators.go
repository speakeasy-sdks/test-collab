// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthStartRequestSubjectUserAuthenticators struct {
	Passive *AuthStartRequestSubjectUserAuthenticatorPassive `json:"passive,omitempty"`
	Present *SubjectUserAuthenticatorPresent                 `json:"present,omitempty"`
}

func (o *AuthStartRequestSubjectUserAuthenticators) GetPassive() *AuthStartRequestSubjectUserAuthenticatorPassive {
	if o == nil {
		return nil
	}
	return o.Passive
}

func (o *AuthStartRequestSubjectUserAuthenticators) GetPresent() *SubjectUserAuthenticatorPresent {
	if o == nil {
		return nil
	}
	return o.Present
}
