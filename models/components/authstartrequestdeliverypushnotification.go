// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthStartRequestDeliveryPushNotification struct {
	// Body is the body of the push message.
	Body string `json:"body"`
	// ConfirmButton is the text of the push message confirm button.
	ConfirmBtn *string `json:"confirmBtn,omitempty"`
	// DenyButton is the text of the push message deny button.
	DenyBtn *string `json:"denyBtn,omitempty"`
	// OriginatingDevice is the type of device where where the push flow was initiated.
	OriginatingDevice *string `json:"originatingDevice,omitempty"`
	// OriginatingIp is the IP address of the web client where the push flow was initiated.
	OriginatingIP *string `json:"originatingIp,omitempty"`
	// Title is the title of the push message.
	Title string `json:"title"`
}

func (o *AuthStartRequestDeliveryPushNotification) GetBody() string {
	if o == nil {
		return ""
	}
	return o.Body
}

func (o *AuthStartRequestDeliveryPushNotification) GetConfirmBtn() *string {
	if o == nil {
		return nil
	}
	return o.ConfirmBtn
}

func (o *AuthStartRequestDeliveryPushNotification) GetDenyBtn() *string {
	if o == nil {
		return nil
	}
	return o.DenyBtn
}

func (o *AuthStartRequestDeliveryPushNotification) GetOriginatingDevice() *string {
	if o == nil {
		return nil
	}
	return o.OriginatingDevice
}

func (o *AuthStartRequestDeliveryPushNotification) GetOriginatingIP() *string {
	if o == nil {
		return nil
	}
	return o.OriginatingIP
}

func (o *AuthStartRequestDeliveryPushNotification) GetTitle() string {
	if o == nil {
		return ""
	}
	return o.Title
}
