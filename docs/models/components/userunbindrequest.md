# UserUnbindRequest


## Fields

| Field                                                                                                                                                                    | Type                                                                                                                                                                     | Required                                                                                                                                                                 | Description                                                                                                                                                              | Example                                                                                                                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `DeviceID`                                                                                                                                                               | **string*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                       | DeviceID is the UUID of the device to disassociate from the user ID. If no device is specified, all devices<br/>associated with the user ID will be unbound.             | eba12f3a-5555-47bc-b85d-21c0cbc4b973                                                                                                                                     |
| `RequestID`                                                                                                                                                              | **string*                                                                                                                                                                | :heavy_minus_sign:                                                                                                                                                       | RequestID is a UUID generated on the customer side to be associated with the unique request for tracking.<br/>Acceptable characters are: alphanumeric with symbols '-._+=/'. | eba12f3a-5555-47bc-b85d-21c0cbc4b973                                                                                                                                     |
| `UserID`                                                                                                                                                                 | *string*                                                                                                                                                                 | :heavy_check_mark:                                                                                                                                                       | UserID is a user identifier generated on the customer side to remove association with a device.<br/>Acceptable characters are: alphanumeric with symbols '-._+=/:'.      | eba12f3a-5555-47bc-b85d-21c0cbc4b973                                                                                                                                     |