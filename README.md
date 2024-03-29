# github.com/speakeasy-sdks/test-collab

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/test-collab
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"github.com/speakeasy-sdks/test-collab/models/components"
	"log"
)

func main() {
	s := testcollab.New(
		testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
		AuthID:    "713189b8-5555-4b08-83ba-75d08780aebd",
		RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
		Subjects:  components.AuthContinueRequestSubjects{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AuthContinueResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Auth](docs/sdks/auth/README.md)

* [AuthContinueRequest](docs/sdks/auth/README.md#authcontinuerequest) - AuthContinue /v1/server/auth/continue
* [AuthExchangeRequest](docs/sdks/auth/README.md#authexchangerequest) - ServerAuthExchange /v1/server/auth/exchange
* [AuthFinishRequest](docs/sdks/auth/README.md#authfinishrequest) - AuthFinish /v1/server/auth/finish
* [AuthStartRequest](docs/sdks/auth/README.md#authstartrequest) - AuthStart /v1/server/auth/start

### [Device](docs/sdks/device/README.md)

* [DeviceUnregisterRequest](docs/sdks/device/README.md#deviceunregisterrequest) - ServerDeviceUnregister /v1/server/device/unregister

### [Push](docs/sdks/push/README.md)

* [PushTokenDeleteRequest](docs/sdks/push/README.md#pushtokendeleterequest) - ServerPushTokenDelete /v1/server/push/delete
* [UserInfoRequest](docs/sdks/push/README.md#userinforequest) - ServerUserInfo /v1/server/user/info
* [UserMobileActiveRequest](docs/sdks/push/README.md#usermobileactiverequest) - ServerUserMobileActive /v1/server/user/mobileactive
* [UserRemoveRequest](docs/sdks/push/README.md#userremoverequest) - ServerUserRemove /v1/server/user/remove
* [UserUnbindRequest](docs/sdks/push/README.md#userunbindrequest) - ServerUserUnbind /v1/server/user/unbind
<!-- End Available Resources and Operations [operations] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 400,500            | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"github.com/speakeasy-sdks/test-collab/models/components"
	"github.com/speakeasy-sdks/test-collab/models/sdkerrors"
	"log"
)

func main() {
	s := testcollab.New(
		testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
		AuthID:    "713189b8-5555-4b08-83ba-75d08780aebd",
		RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
		Subjects:  components.AuthContinueRequestSubjects{},
	})
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://oapi.uat.prove-auth.proveapis.com/` | None |

#### Example

```go
package main

import (
	"context"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"github.com/speakeasy-sdks/test-collab/models/components"
	"log"
)

func main() {
	s := testcollab.New(
		testcollab.WithServerIndex(0),
		testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
		AuthID:    "713189b8-5555-4b08-83ba-75d08780aebd",
		RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
		Subjects:  components.AuthContinueRequestSubjects{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AuthContinueResponse != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"github.com/speakeasy-sdks/test-collab/models/components"
	"log"
)

func main() {
	s := testcollab.New(
		testcollab.WithServerURL("https://oapi.uat.prove-auth.proveapis.com/"),
		testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
		AuthID:    "713189b8-5555-4b08-83ba-75d08780aebd",
		RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
		Subjects:  components.AuthContinueRequestSubjects{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AuthContinueResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name    | Type    | Scheme  |
| ------- | ------- | ------- |
| `Token` | apiKey  | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	testcollab "github.com/speakeasy-sdks/test-collab"
	"github.com/speakeasy-sdks/test-collab/models/components"
	"log"
)

func main() {
	s := testcollab.New(
		testcollab.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Auth.AuthContinueRequest(ctx, &components.AuthContinueRequest{
		AuthID:    "713189b8-5555-4b08-83ba-75d08780aebd",
		RequestID: testcollab.String("eba12f3a-5555-47bc-b85d-21c0cbc4b973"),
		Subjects:  components.AuthContinueRequestSubjects{},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AuthContinueResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
