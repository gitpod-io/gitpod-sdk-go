// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/gitpod-io/gitpod-sdk-go/internal/apijson"
	"github.com/gitpod-io/gitpod-sdk-go/shared"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	// The status code, which should be an enum value of
	// [google.rpc.Code][google.rpc.Code].
	Code shared.ErrorCode `json:"code"`
	// A developer-facing error message, which should be in English. Any user-facing
	// error message should be localized and sent in the
	// [google.rpc.Status.details][google.rpc.Status.details] field, or localized by
	// the client.
	Message     string                 `json:"message"`
	ExtraFields map[string]interface{} `json:"-,extras"`
	JSON        errorJSON              `json:"-"`
	StatusCode  int
	Request     *http.Request
	Response    *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorJSON) RawJSON() string {
	return r.raw
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.RawJSON())
}

func (r *Error) DumpRequest(body bool) []byte {
	if r.Request.GetBody != nil {
		r.Request.Body, _ = r.Request.GetBody()
	}
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}
