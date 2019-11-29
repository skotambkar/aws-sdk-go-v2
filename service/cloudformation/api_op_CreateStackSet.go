// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opCreateStackSet = "CreateStackSet"

// CreateStackSetRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Creates a stack set.
//
//    // Example sending a request using CreateStackSetRequest.
//    req := client.CreateStackSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateStackSet
func (c *Client) CreateStackSetRequest(input *types.CreateStackSetInput) CreateStackSetRequest {
	op := &aws.Operation{
		Name:       opCreateStackSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateStackSetInput{}
	}

	req := c.newRequest(op, input, &types.CreateStackSetOutput{})
	return CreateStackSetRequest{Request: req, Input: input, Copy: c.CreateStackSetRequest}
}

// CreateStackSetRequest is the request type for the
// CreateStackSet API operation.
type CreateStackSetRequest struct {
	*aws.Request
	Input *types.CreateStackSetInput
	Copy  func(*types.CreateStackSetInput) CreateStackSetRequest
}

// Send marshals and sends the CreateStackSet API request.
func (r CreateStackSetRequest) Send(ctx context.Context) (*CreateStackSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStackSetResponse{
		CreateStackSetOutput: r.Request.Data.(*types.CreateStackSetOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStackSetResponse is the response type for the
// CreateStackSet API operation.
type CreateStackSetResponse struct {
	*types.CreateStackSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStackSet request.
func (r *CreateStackSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
