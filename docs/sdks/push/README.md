# Push
(*Push*)

### Available Operations

* [PushTokenDeleteRequest](#pushtokendeleterequest) - ServerPushTokenDelete /v1/server/push/delete
* [UserInfoRequest](#userinforequest) - ServerUserInfo /v1/server/user/info
* [UserMobileActiveRequest](#usermobileactiverequest) - ServerUserMobileActive /v1/server/user/mobileactive
* [UserRemoveRequest](#userremoverequest) - ServerUserRemove /v1/server/user/remove
* [UserUnbindRequest](#userunbindrequest) - ServerUserUnbind /v1/server/user/unbind

## PushTokenDeleteRequest

Delete push tokens for a list of devices so they will no longer receive push notifications.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/push/delete

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
    res, err := s.Push.PushTokenDeleteRequest(ctx, &components.PushTokenDeleteRequest{
        Devices: []string{
            "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
            "fba12f3a-5555-47bc-b85d-21c0cbc4b974",
        },
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.PushTokenDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [components.PushTokenDeleteRequest](../../models/components/pushtokendeleterequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PushTokenDeleteRequestResponse](../../models/operations/pushtokendeleterequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UserInfoRequest

Return device records associated with a user ID.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/user/info

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
    res, err := s.Push.UserInfoRequest(ctx, &components.UserInfoRequest{
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        UserID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [components.UserInfoRequest](../../models/components/userinforequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.UserInfoRequestResponse](../../models/operations/userinforequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UserMobileActiveRequest

Set user ID as active on their associated device.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/user/mobileactive

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
    res, err := s.Push.UserMobileActiveRequest(ctx, &components.UserMobileActiveRequest{
        AuthID: "713189b8-5555-4b08-83ba-75d08780aebd",
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        UserID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserMobileActiveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.UserMobileActiveRequest](../../models/components/usermobileactiverequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UserMobileActiveRequestResponse](../../models/operations/usermobileactiverequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UserRemoveRequest

Deactivate user ID and associated devices.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/user/remove

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
    res, err := s.Push.UserRemoveRequest(ctx, &components.UserRemoveRequest{
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        UserID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserRemoveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.UserRemoveRequest](../../models/components/userremoverequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UserRemoveRequestResponse](../../models/operations/userremoverequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UserUnbindRequest

Deactivate user ID on either a specified device or all associated devices. Device(s) will remain active.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/user/unbind

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
    res, err := s.Push.UserUnbindRequest(ctx, &components.UserUnbindRequest{
        DeviceID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
        UserID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UserUnbindResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [components.UserUnbindRequest](../../models/components/userunbindrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UserUnbindRequestResponse](../../models/operations/userunbindrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
