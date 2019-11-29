// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
)

const opListAccelerators = "ListAccelerators"

// ListAcceleratorsRequest returns a request value for making API operation for
// AWS Global Accelerator.
//
// List the accelerators for an AWS account.
//
//    // Example sending a request using ListAcceleratorsRequest.
//    req := client.ListAcceleratorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08/ListAccelerators
func (c *Client) ListAcceleratorsRequest(input *types.ListAcceleratorsInput) ListAcceleratorsRequest {
	op := &aws.Operation{
		Name:       opListAccelerators,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListAcceleratorsInput{}
	}

	req := c.newRequest(op, input, &types.ListAcceleratorsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListAcceleratorsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListAcceleratorsRequest{Request: req, Input: input, Copy: c.ListAcceleratorsRequest}
}

// ListAcceleratorsRequest is the request type for the
// ListAccelerators API operation.
type ListAcceleratorsRequest struct {
	*aws.Request
	Input *types.ListAcceleratorsInput
	Copy  func(*types.ListAcceleratorsInput) ListAcceleratorsRequest
}

// Send marshals and sends the ListAccelerators API request.
func (r ListAcceleratorsRequest) Send(ctx context.Context) (*ListAcceleratorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAcceleratorsResponse{
		ListAcceleratorsOutput: r.Request.Data.(*types.ListAcceleratorsOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAcceleratorsResponse is the response type for the
// ListAccelerators API operation.
type ListAcceleratorsResponse struct {
	*types.ListAcceleratorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAccelerators request.
func (r *ListAcceleratorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
