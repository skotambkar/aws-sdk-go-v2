// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

const opSetInstanceHealth = "SetInstanceHealth"

// SetInstanceHealthRequest returns a request value for making API operation for
// Auto Scaling.
//
// Sets the health status of the specified instance.
//
// For more information, see Health Checks for Auto Scaling Instances (https://docs.aws.amazon.com/autoscaling/ec2/userguide/healthcheck.html)
// in the Amazon EC2 Auto Scaling User Guide.
//
//    // Example sending a request using SetInstanceHealthRequest.
//    req := client.SetInstanceHealthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/SetInstanceHealth
func (c *Client) SetInstanceHealthRequest(input *types.SetInstanceHealthInput) SetInstanceHealthRequest {
	op := &aws.Operation{
		Name:       opSetInstanceHealth,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.SetInstanceHealthInput{}
	}

	req := c.newRequest(op, input, &types.SetInstanceHealthOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetInstanceHealthRequest{Request: req, Input: input, Copy: c.SetInstanceHealthRequest}
}

// SetInstanceHealthRequest is the request type for the
// SetInstanceHealth API operation.
type SetInstanceHealthRequest struct {
	*aws.Request
	Input *types.SetInstanceHealthInput
	Copy  func(*types.SetInstanceHealthInput) SetInstanceHealthRequest
}

// Send marshals and sends the SetInstanceHealth API request.
func (r SetInstanceHealthRequest) Send(ctx context.Context) (*SetInstanceHealthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetInstanceHealthResponse{
		SetInstanceHealthOutput: r.Request.Data.(*types.SetInstanceHealthOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetInstanceHealthResponse is the response type for the
// SetInstanceHealth API operation.
type SetInstanceHealthResponse struct {
	*types.SetInstanceHealthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetInstanceHealth request.
func (r *SetInstanceHealthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
