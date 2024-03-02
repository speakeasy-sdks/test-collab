// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package testcollab

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/test-collab/internal/hooks"
	"github.com/speakeasy-sdks/test-collab/internal/utils"
	"github.com/speakeasy-sdks/test-collab/models/components"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	"https://oapi.uat.prove-auth.proveapis.com/",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Proveapi - Prove Auth API: This specification describes the Prove Auth API.
//
// Swagger 2.0 Spec - generated by [go-swagger](https://github.com/go-swagger/go-swagger)
type Proveapi struct {
	Auth   *Auth
	Device *Device
	Push   *Push

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Proveapi)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Proveapi) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Proveapi) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Proveapi) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Proveapi) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(token string) SDKOption {
	return func(sdk *Proveapi) {
		security := components.Security{Token: &token}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (components.Security, error)) SDKOption {
	return func(sdk *Proveapi) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Proveapi) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Proveapi {
	sdk := &Proveapi{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.37.3",
			SDKVersion:        "0.2.4",
			GenVersion:        "2.277.0",
			UserAgent:         "speakeasy-sdk/go 0.2.4 2.277.0 1.37.3 github.com/speakeasy-sdks/test-collab",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.DefaultClient = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.DefaultClient)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.Device = newDevice(sdk.sdkConfiguration)

	sdk.Push = newPush(sdk.sdkConfiguration)

	return sdk
}
