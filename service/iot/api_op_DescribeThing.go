// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opDescribeThing = "DescribeThing"

// DescribeThingRequest returns a request value for making API operation for
// AWS IoT.
//
// Gets information about the specified thing.
//
//    // Example sending a request using DescribeThingRequest.
//    req := client.DescribeThingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeThingRequest(input *types.DescribeThingInput) DescribeThingRequest {
	op := &aws.Operation{
		Name:       opDescribeThing,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}",
	}

	if input == nil {
		input = &types.DescribeThingInput{}
	}

	req := c.newRequest(op, input, &types.DescribeThingOutput{})
	return DescribeThingRequest{Request: req, Input: input, Copy: c.DescribeThingRequest}
}

// DescribeThingRequest is the request type for the
// DescribeThing API operation.
type DescribeThingRequest struct {
	*aws.Request
	Input *types.DescribeThingInput
	Copy  func(*types.DescribeThingInput) DescribeThingRequest
}

// Send marshals and sends the DescribeThing API request.
func (r DescribeThingRequest) Send(ctx context.Context) (*DescribeThingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeThingResponse{
		DescribeThingOutput: r.Request.Data.(*types.DescribeThingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeThingResponse is the response type for the
// DescribeThing API operation.
type DescribeThingResponse struct {
	*types.DescribeThingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeThing request.
func (r *DescribeThingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
