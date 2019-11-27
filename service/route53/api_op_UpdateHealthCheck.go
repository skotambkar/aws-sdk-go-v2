// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opUpdateHealthCheck = "UpdateHealthCheck"

// UpdateHealthCheckRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Updates an existing health check. Note that some values can't be updated.
//
// For more information about updating health checks, see Creating, Updating,
// and Deleting Health Checks (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html)
// in the Amazon Route 53 Developer Guide.
//
//    // Example sending a request using UpdateHealthCheckRequest.
//    req := client.UpdateHealthCheckRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/UpdateHealthCheck
func (c *Client) UpdateHealthCheckRequest(input *types.UpdateHealthCheckInput) UpdateHealthCheckRequest {
	op := &aws.Operation{
		Name:       opUpdateHealthCheck,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/healthcheck/{HealthCheckId}",
	}

	if input == nil {
		input = &types.UpdateHealthCheckInput{}
	}

	req := c.newRequest(op, input, &types.UpdateHealthCheckOutput{})
	return UpdateHealthCheckRequest{Request: req, Input: input, Copy: c.UpdateHealthCheckRequest}
}

// UpdateHealthCheckRequest is the request type for the
// UpdateHealthCheck API operation.
type UpdateHealthCheckRequest struct {
	*aws.Request
	Input *types.UpdateHealthCheckInput
	Copy  func(*types.UpdateHealthCheckInput) UpdateHealthCheckRequest
}

// Send marshals and sends the UpdateHealthCheck API request.
func (r UpdateHealthCheckRequest) Send(ctx context.Context) (*UpdateHealthCheckResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateHealthCheckResponse{
		UpdateHealthCheckOutput: r.Request.Data.(*types.UpdateHealthCheckOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateHealthCheckResponse is the response type for the
// UpdateHealthCheck API operation.
type UpdateHealthCheckResponse struct {
	*types.UpdateHealthCheckOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateHealthCheck request.
func (r *UpdateHealthCheckResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
