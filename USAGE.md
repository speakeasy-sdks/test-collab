<!-- Start SDK Example Usage [usage] -->
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