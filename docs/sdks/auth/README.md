# Auth
(*Auth*)

### Available Operations

* [AuthContinueRequest](#authcontinuerequest) - AuthContinue /v1/server/auth/continue
* [AuthExchangeRequest](#authexchangerequest) - ServerAuthExchange /v1/server/auth/exchange
* [AuthFinishRequest](#authfinishrequest) - AuthFinish /v1/server/auth/finish
* [AuthStartRequest](#authstartrequest) - AuthStart /v1/server/auth/start

## AuthContinueRequest

If the initial Prove Auth authenticators fail, the Force Bind authenticator can be used which permits using another
authentication method to verify a mobile number (like Prove Instant Link™). Once the mobile number is verified, send
this AuthContinue request.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/auth/continue

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-collab/models/components"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"context"
	"log"
)

func main() {
    s := testcollab.New(
        testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
        AuthID: "713189b8-5555-4b08-83ba-75d08780aebd",
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        Subjects: components.AuthContinueRequestSubjects{
            Mobile: &components.AuthContinueRequestSubjectMobile{
                Claim: &components.AuthContinueRequestSubjectMobileClaim{
                    MobileNumber: "12065550100",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthContinueResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.AuthContinueRequest](../../models/components/authcontinuerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.AuthContinueRequestResponse](../../models/operations/authcontinuerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AuthExchangeRequest

Exchange AuthID for AuthToken.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-collab/models/components"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"context"
	"log"
)

func main() {
    s := testcollab.New(
        testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Auth.AuthExchangeRequest(ctx, &components.AuthExchangeRequest{
        AuthID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthExchangeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [components.AuthExchangeRequest](../../models/components/authexchangerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.AuthExchangeRequestResponse](../../models/operations/authexchangerequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AuthFinishRequest

To complete the auth flow and get the authentication result, send this AuthFinish request.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/auth/finish

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-collab/models/components"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"context"
	"log"
)

func main() {
    s := testcollab.New(
        testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Auth.AuthFinishRequest(ctx, &components.AuthFinishRequest{
        AuthID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthFinishResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.AuthFinishRequest](../../models/components/authfinishrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.AuthFinishRequestResponse](../../models/operations/authfinishrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## AuthStartRequest

To start an auth flow, send this AuthStart request to control how Prove Auth will authenticate the end users and
their devices. The expected authenticators should be included in the body parameters. There are many supported
scenarios to use Prove Auth so each of the request types are explained in the "Server Integration Guide".

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/auth/start

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/test-collab/models/components"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"context"
	"log"
)

func main() {
    s := testcollab.New(
        testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Auth.AuthStartRequest(ctx, &components.AuthStartRequest{
        CallbackURL: testcollab.String("https://example.com/webhook"),
        Delivery: &components.AuthStartRequestDelivery{
            Push: components.AuthStartRequestDeliveryPush{
                Notification: components.AuthStartRequestDeliveryPushNotification{
                    Body: "Do you want to allow login from this device?",
                    ConfirmBtn: testcollab.String("Confirm"),
                    DenyBtn: testcollab.String("Deny"),
                    OriginatingDevice: testcollab.String("iPhone12"),
                    OriginatingIP: testcollab.String("198.51.100.10"),
                    Title: "Confirm Login",
                },
            },
        },
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        Subjects: components.AuthStartRequestSubjects{
            Device: &components.AuthStartRequestSubjectDevice{
                Authenticators: &components.AuthStartRequestSubjectDeviceAuthenticators{
                    Passive: &components.AuthStartRequestSubjectDeviceAuthenticatorPassive{},
                },
            },
            Mobile: &components.AuthStartRequestSubjectMobile{
                Authenticators: &components.AuthStartRequestSubjectMobileAuthenticators{
                    Instant: &components.AuthStartRequestSubjectMobileAuthenticatorInstant{
                        Consent: &components.AuthStartRequestInstantAuthConsent{
                            CollectedTimestamp: testcollab.String("2022-02-17T00:00:00Z"),
                            Description: testcollab.String("Customer Application Name"),
                            Status: "optedIn",
                            TransactionID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
                        },
                    },
                    Otp: &components.AuthStartRequestSubjectMobileAuthenticatorOTP{
                        AllowOtpRetry: testcollab.Bool(true),
                        MessageText: "Your pin is: ####",
                    },
                    Passive: &components.AuthStartRequestSubjectMobileAuthenticatorPassive{
                        CacheResult: testcollab.Bool(true),
                        LocalDomain: testcollab.Bool(true),
                        MaxAge: testcollab.Int64(7776000),
                    },
                },
                Claim: &components.AuthStartRequestSubjectMobileClaim{
                    MobileNumber: testcollab.String("12065550100"),
                },
            },
            User: &components.AuthStartRequestSubjectUser{
                Authenticators: components.AuthStartRequestSubjectUserAuthenticators{
                    Passive: &components.AuthStartRequestSubjectUserAuthenticatorPassive{},
                    Present: &components.SubjectUserAuthenticatorPresent{},
                },
                Claim: components.AuthStartRequestSubjectUserClaim{
                    UserID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthStartResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [components.AuthStartRequest](../../models/components/authstartrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.AuthStartRequestResponse](../../models/operations/authstartrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
