# Device
(*Device*)

### Available Operations

* [DeviceUnregisterRequest](#deviceunregisterrequest) - ServerDeviceUnregister /v1/server/device/unregister

## DeviceUnregisterRequest

Mark a device as inactive so it can no longer be used in an auth flow.

Production URL: https://oapi.prove-auth.proveapis.com/v1/server/device/unregister

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
    res, err := s.Device.DeviceUnregisterRequest(ctx, &components.DeviceUnregisterRequest{
        DeviceID: "eba12f3a-5555-47bc-b85d-21c0cbc4b973",
        RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeviceUnregisterResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [components.DeviceUnregisterRequest](../../models/components/deviceunregisterrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeviceUnregisterRequestResponse](../../models/operations/deviceunregisterrequestresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
