// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opCreateDevicePool = "CreateDevicePool"

// CreateDevicePoolRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Creates a device pool.
//
//    // Example sending a request using CreateDevicePoolRequest.
//    req := client.CreateDevicePoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateDevicePool
func (c *Client) CreateDevicePoolRequest(input *types.CreateDevicePoolInput) CreateDevicePoolRequest {
	op := &aws.Operation{
		Name:       opCreateDevicePool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDevicePoolInput{}
	}

	req := c.newRequest(op, input, &types.CreateDevicePoolOutput{})
	return CreateDevicePoolRequest{Request: req, Input: input, Copy: c.CreateDevicePoolRequest}
}

// CreateDevicePoolRequest is the request type for the
// CreateDevicePool API operation.
type CreateDevicePoolRequest struct {
	*aws.Request
	Input *types.CreateDevicePoolInput
	Copy  func(*types.CreateDevicePoolInput) CreateDevicePoolRequest
}

// Send marshals and sends the CreateDevicePool API request.
func (r CreateDevicePoolRequest) Send(ctx context.Context) (*CreateDevicePoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDevicePoolResponse{
		CreateDevicePoolOutput: r.Request.Data.(*types.CreateDevicePoolOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDevicePoolResponse is the response type for the
// CreateDevicePool API operation.
type CreateDevicePoolResponse struct {
	*types.CreateDevicePoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDevicePool request.
func (r *CreateDevicePoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
