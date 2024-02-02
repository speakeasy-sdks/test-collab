// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponse struct {
	// AppNamespace describes the Prove Auth environment.
	AppNamespace string `json:"appNamespace"`
	// AuthID is the unique ID of the auth flow.
	AuthID   string                      `json:"authId"`
	Delivery *AuthFinishResponseDelivery `json:"delivery,omitempty"`
	Subjects AuthFinishResponseSubjects  `json:"subjects"`
}

func (o *AuthFinishResponse) GetAppNamespace() string {
	if o == nil {
		return ""
	}
	return o.AppNamespace
}

func (o *AuthFinishResponse) GetAuthID() string {
	if o == nil {
		return ""
	}
	return o.AuthID
}

func (o *AuthFinishResponse) GetDelivery() *AuthFinishResponseDelivery {
	if o == nil {
		return nil
	}
	return o.Delivery
}

func (o *AuthFinishResponse) GetSubjects() AuthFinishResponseSubjects {
	if o == nil {
		return AuthFinishResponseSubjects{}
	}
	return o.Subjects
}
