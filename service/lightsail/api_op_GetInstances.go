// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
)

const opGetInstances = "GetInstances"

// GetInstancesRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about all Amazon Lightsail virtual private servers, or
// instances.
//
//    // Example sending a request using GetInstancesRequest.
//    req := client.GetInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetInstances
func (c *Client) GetInstancesRequest(input *types.GetInstancesInput) GetInstancesRequest {
	op := &aws.Operation{
		Name:       opGetInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetInstancesInput{}
	}

	req := c.newRequest(op, input, &types.GetInstancesOutput{})
	return GetInstancesRequest{Request: req, Input: input, Copy: c.GetInstancesRequest}
}

// GetInstancesRequest is the request type for the
// GetInstances API operation.
type GetInstancesRequest struct {
	*aws.Request
	Input *types.GetInstancesInput
	Copy  func(*types.GetInstancesInput) GetInstancesRequest
}

// Send marshals and sends the GetInstances API request.
func (r GetInstancesRequest) Send(ctx context.Context) (*GetInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetInstancesResponse{
		GetInstancesOutput: r.Request.Data.(*types.GetInstancesOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetInstancesResponse is the response type for the
// GetInstances API operation.
type GetInstancesResponse struct {
	*types.GetInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetInstances request.
func (r *GetInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
