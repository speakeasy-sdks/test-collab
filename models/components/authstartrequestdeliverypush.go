// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthStartRequestDeliveryPush struct {
	Notification AuthStartRequestDeliveryPushNotification `json:"notification"`
}

func (o *AuthStartRequestDeliveryPush) GetNotification() AuthStartRequestDeliveryPushNotification {
	if o == nil {
		return AuthStartRequestDeliveryPushNotification{}
	}
	return o.Notification
}