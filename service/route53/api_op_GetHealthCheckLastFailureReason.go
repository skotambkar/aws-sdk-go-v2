// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opGetHealthCheckLastFailureReason = "GetHealthCheckLastFailureReason"

// GetHealthCheckLastFailureReasonRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Gets the reason that a specified health check failed most recently.
//
//    // Example sending a request using GetHealthCheckLastFailureReasonRequest.
//    req := client.GetHealthCheckLastFailureReasonRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/GetHealthCheckLastFailureReason
func (c *Client) GetHealthCheckLastFailureReasonRequest(input *types.GetHealthCheckLastFailureReasonInput) GetHealthCheckLastFailureReasonRequest {
	op := &aws.Operation{
		Name:       opGetHealthCheckLastFailureReason,
		HTTPMethod: "GET",
		HTTPPath:   "/2013-04-01/healthcheck/{HealthCheckId}/lastfailurereason",
	}

	if input == nil {
		input = &types.GetHealthCheckLastFailureReasonInput{}
	}

	req := c.newRequest(op, input, &types.GetHealthCheckLastFailureReasonOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.GetHealthCheckLastFailureReasonMarshaler{Input: input}.GetNamedBuildHandler())

	return GetHealthCheckLastFailureReasonRequest{Request: req, Input: input, Copy: c.GetHealthCheckLastFailureReasonRequest}
}

// GetHealthCheckLastFailureReasonRequest is the request type for the
// GetHealthCheckLastFailureReason API operation.
type GetHealthCheckLastFailureReasonRequest struct {
	*aws.Request
	Input *types.GetHealthCheckLastFailureReasonInput
	Copy  func(*types.GetHealthCheckLastFailureReasonInput) GetHealthCheckLastFailureReasonRequest
}

// Send marshals and sends the GetHealthCheckLastFailureReason API request.
func (r GetHealthCheckLastFailureReasonRequest) Send(ctx context.Context) (*GetHealthCheckLastFailureReasonResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetHealthCheckLastFailureReasonResponse{
		GetHealthCheckLastFailureReasonOutput: r.Request.Data.(*types.GetHealthCheckLastFailureReasonOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetHealthCheckLastFailureReasonResponse is the response type for the
// GetHealthCheckLastFailureReason API operation.
type GetHealthCheckLastFailureReasonResponse struct {
	*types.GetHealthCheckLastFailureReasonOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetHealthCheckLastFailureReason request.
func (r *GetHealthCheckLastFailureReasonResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
