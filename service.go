// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gitpod

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/gitpod-go/internal/apijson"
	"github.com/stainless-sdks/gitpod-go/internal/param"
	"github.com/stainless-sdks/gitpod-go/internal/requestconfig"
	"github.com/stainless-sdks/gitpod-go/option"
)

// ServiceService contains methods and other services that help with interacting
// with the gitpod API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceService] method instead.
type ServiceService struct {
	Options []option.RequestOption
}

// NewServiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServiceService(opts ...option.RequestOption) (r *ServiceService) {
	r = &ServiceService{}
	r.Options = opts
	return
}

// UpdateService
func (r *ServiceService) Update(ctx context.Context, params ServiceUpdateParams, opts ...option.RequestOption) (res *ServiceUpdateResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/UpdateService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// DeleteService deletes a service. This call does not block until the service is
// deleted. If the service is not stopped it will be stopped before deletion.
func (r *ServiceService) Delete(ctx context.Context, params ServiceDeleteParams, opts ...option.RequestOption) (res *ServiceDeleteResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/DeleteService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StartService starts a service. This call does not block until the service is
// started. This call will not error if the service is already running or has been
// started.
func (r *ServiceService) Start(ctx context.Context, params ServiceStartParams, opts ...option.RequestOption) (res *ServiceStartResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StartService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// StopService stops a service. This call does not block until the service is
// stopped. This call will not error if the service is already stopped or has been
// stopped.
func (r *ServiceService) Stop(ctx context.Context, params ServiceStopParams, opts ...option.RequestOption) (res *ServiceStopResponse, err error) {
	if params.ConnectProtocolVersion.Present {
		opts = append(opts, option.WithHeader("Connect-Protocol-Version", fmt.Sprintf("%s", params.ConnectProtocolVersion)))
	}
	if params.ConnectTimeoutMs.Present {
		opts = append(opts, option.WithHeader("Connect-Timeout-Ms", fmt.Sprintf("%s", params.ConnectTimeoutMs)))
	}
	opts = append(r.Options[:], opts...)
	path := "gitpod.v1.EnvironmentAutomationService/StopService"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ServiceUpdateResponse = interface{}

type ServiceDeleteResponse = interface{}

type ServiceStartResponse = interface{}

type ServiceStopResponse = interface{}

type ServiceUpdateParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceUpdateParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                    `json:"id" format:"uuid"`
	Metadata               param.Field[ServiceUpdateParamsMetadata]               `json:"metadata"`
	// Changing the spec of a service is a complex operation. The spec of a service can
	// only be updated if the service is in a stopped state. If the service is running,
	// it must be stopped first.
	Spec param.Field[ServiceUpdateParamsSpec] `json:"spec"`
	// Service status updates are only expected from the executing environment. As a
	// client of this API you are not expected to provide this field. Updating this
	// field requires the `environmentservice:update_status` permission.
	Status param.Field[ServiceUpdateParamsStatus] `json:"status"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceUpdateParamsConnectProtocolVersion float64

const (
	ServiceUpdateParamsConnectProtocolVersion1 ServiceUpdateParamsConnectProtocolVersion = 1
)

func (r ServiceUpdateParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceUpdateParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceUpdateParamsMetadata struct {
}

func (r ServiceUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Changing the spec of a service is a complex operation. The spec of a service can
// only be updated if the service is in a stopped state. If the service is running,
// it must be stopped first.
type ServiceUpdateParamsSpec struct {
}

func (r ServiceUpdateParamsSpec) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service status updates are only expected from the executing environment. As a
// client of this API you are not expected to provide this field. Updating this
// field requires the `environmentservice:update_status` permission.
type ServiceUpdateParamsStatus struct {
}

func (r ServiceUpdateParamsStatus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServiceDeleteParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceDeleteParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                    `json:"id" format:"uuid"`
	Force                  param.Field[bool]                                      `json:"force"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceDeleteParamsConnectProtocolVersion float64

const (
	ServiceDeleteParamsConnectProtocolVersion1 ServiceDeleteParamsConnectProtocolVersion = 1
)

func (r ServiceDeleteParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceDeleteParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceStartParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceStartParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                   `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceStartParamsConnectProtocolVersion float64

const (
	ServiceStartParamsConnectProtocolVersion1 ServiceStartParamsConnectProtocolVersion = 1
)

func (r ServiceStartParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceStartParamsConnectProtocolVersion1:
		return true
	}
	return false
}

type ServiceStopParams struct {
	// Define the version of the Connect protocol
	ConnectProtocolVersion param.Field[ServiceStopParamsConnectProtocolVersion] `header:"Connect-Protocol-Version,required"`
	ID                     param.Field[string]                                  `json:"id" format:"uuid"`
	// Define the timeout, in ms
	ConnectTimeoutMs param.Field[float64] `header:"Connect-Timeout-Ms"`
}

func (r ServiceStopParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the version of the Connect protocol
type ServiceStopParamsConnectProtocolVersion float64

const (
	ServiceStopParamsConnectProtocolVersion1 ServiceStopParamsConnectProtocolVersion = 1
)

func (r ServiceStopParamsConnectProtocolVersion) IsKnown() bool {
	switch r {
	case ServiceStopParamsConnectProtocolVersion1:
		return true
	}
	return false
}
