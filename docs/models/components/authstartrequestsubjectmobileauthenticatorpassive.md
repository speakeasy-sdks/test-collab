# AuthStartRequestSubjectMobileAuthenticatorPassive


## Fields

| Field                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                          | Example                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `CacheResult`                                                                                                                                                                                                                                                                                                                        | **bool*                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                   | If true, bind will persist between authenticators. If false, bind will be reattempted on every auth flow.                                                                                                                                                                                                                            | true                                                                                                                                                                                                                                                                                                                                 |
| `LocalDomain`                                                                                                                                                                                                                                                                                                                        | **bool*                                                                                                                                                                                                                                                                                                                              | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                   | If true, will allow authentication against a local domain for testing FIDO.                                                                                                                                                                                                                                                          | true                                                                                                                                                                                                                                                                                                                                 |
| `MaxAge`                                                                                                                                                                                                                                                                                                                             | **int64*                                                                                                                                                                                                                                                                                                                             | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                   | If not set or set to 0, the age of the bind will be ignored during authentication. If set to -1, then it will<br/>force a rebind of the mobile number to the device. If set greater than 0, it specifies the maximum age<br/>(in seconds) of the bind that is allowed or else a rebind occurs during authentication. Ex: 7776000 is 90 days. | 7776000                                                                                                                                                                                                                                                                                                                              |
| `TestMode`                                                                                                                                                                                                                                                                                                                           | **string*                                                                                                                                                                                                                                                                                                                            | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                   | N/A                                                                                                                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                                                                                      |