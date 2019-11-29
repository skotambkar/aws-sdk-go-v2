// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
)

const opDescribeEventBus = "DescribeEventBus"

// DescribeEventBusRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// Displays details about an event bus in your account. This can include the
// external AWS accounts that are permitted to write events to your default
// event bus, and the associated policy. For custom event buses and partner
// event buses, it displays the name, ARN, policy, state, and creation time.
//
// To enable your account to receive events from other accounts on its default
// event bus, use PutPermission.
//
// For more information about partner event buses, see CreateEventBus.
//
//    // Example sending a request using DescribeEventBusRequest.
//    req := client.DescribeEventBusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/DescribeEventBus
func (c *Client) DescribeEventBusRequest(input *types.DescribeEventBusInput) DescribeEventBusRequest {
	op := &aws.Operation{
		Name:       opDescribeEventBus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEventBusInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEventBusOutput{})
	return DescribeEventBusRequest{Request: req, Input: input, Copy: c.DescribeEventBusRequest}
}

// DescribeEventBusRequest is the request type for the
// DescribeEventBus API operation.
type DescribeEventBusRequest struct {
	*aws.Request
	Input *types.DescribeEventBusInput
	Copy  func(*types.DescribeEventBusInput) DescribeEventBusRequest
}

// Send marshals and sends the DescribeEventBus API request.
func (r DescribeEventBusRequest) Send(ctx context.Context) (*DescribeEventBusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventBusResponse{
		DescribeEventBusOutput: r.Request.Data.(*types.DescribeEventBusOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventBusResponse is the response type for the
// DescribeEventBus API operation.
type DescribeEventBusResponse struct {
	*types.DescribeEventBusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventBus request.
func (r *DescribeEventBusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
