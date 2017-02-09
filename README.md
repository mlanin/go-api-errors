# go-apierr
Set of handy HTTP errors for JSON web services on go.

# Install

```bash
go get github.com/mlanin/go-apierr
```

# About

Prepared set of errors for API usage. Can be transformed into JSON like:

```json
{
    "error": {
      "id": "not_found",
      "message": "Requested object not found."
    }
}
```

Also contain corresponding HTTP status code.

# Use

```go
import "github.com/mlanin/go-apierr"

fail := *apierr.BadRequest
fail.
  // Add more info for the user in meta attribute.
  AddMeta(struct {
    Fail string `json:"fail"`
  }{
    Fail: "JWT is malformed.",
  }).
  // panic(&err) helper
  Send()
```

# More

Errors have helpers for logs management.

```go
import "github.com/mlanin/go-apierr"

fail := *apierr.BadRequest.
  // Enable reporting for the error.
  Report().
  // We want to add trace.
  WithTrace().
  // Add error context.
  AddContext(fmt.Sprintf("Request with malformed JWT"))
```

Example of handling:
```go
import "github.com/mlanin/go-apierr"

defer func() {
  if err := recover(); err != nil {
    var fail apierr.APIError = convertToAPIError(err)

    if fail.WantsToBeReported() {
      log("[APIError] %+v [%+v]", err, fail.Context)
      if fail.WantsToShowTrace() {
        log("[Trace] %s", debug.Stack())
      }
    }

    showJSON(fail.HTTPCode, fail)
  }
}()
```
