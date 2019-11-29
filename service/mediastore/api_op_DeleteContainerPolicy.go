// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
)

const opDeleteContainerPolicy = "DeleteContainerPolicy"

// DeleteContainerPolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Deletes the access policy that is associated with the specified container.
//
//    // Example sending a request using DeleteContainerPolicyRequest.
//    req := client.DeleteContainerPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteContainerPolicy
func (c *Client) DeleteContainerPolicyRequest(input *types.DeleteContainerPolicyInput) DeleteContainerPolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteContainerPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteContainerPolicyInput{}
	}

	req := c.newRequest(op, input, &types.DeleteContainerPolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteContainerPolicyMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteContainerPolicyRequest{Request: req, Input: input, Copy: c.DeleteContainerPolicyRequest}
}

// DeleteContainerPolicyRequest is the request type for the
// DeleteContainerPolicy API operation.
type DeleteContainerPolicyRequest struct {
	*aws.Request
	Input *types.DeleteContainerPolicyInput
	Copy  func(*types.DeleteContainerPolicyInput) DeleteContainerPolicyRequest
}

// Send marshals and sends the DeleteContainerPolicy API request.
func (r DeleteContainerPolicyRequest) Send(ctx context.Context) (*DeleteContainerPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteContainerPolicyResponse{
		DeleteContainerPolicyOutput: r.Request.Data.(*types.DeleteContainerPolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteContainerPolicyResponse is the response type for the
// DeleteContainerPolicy API operation.
type DeleteContainerPolicyResponse struct {
	*types.DeleteContainerPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteContainerPolicy request.
func (r *DeleteContainerPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
