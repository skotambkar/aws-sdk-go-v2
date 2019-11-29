// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opCreateStack = "CreateStack"

// CreateStackRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Creates a stack as specified in the template. After the call completes successfully,
// the stack creation starts. You can check the status of the stack via the
// DescribeStacks API.
//
//    // Example sending a request using CreateStackRequest.
//    req := client.CreateStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/CreateStack
func (c *Client) CreateStackRequest(input *types.CreateStackInput) CreateStackRequest {
	op := &aws.Operation{
		Name:       opCreateStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateStackInput{}
	}

	req := c.newRequest(op, input, &types.CreateStackOutput{})
	return CreateStackRequest{Request: req, Input: input, Copy: c.CreateStackRequest}
}

// CreateStackRequest is the request type for the
// CreateStack API operation.
type CreateStackRequest struct {
	*aws.Request
	Input *types.CreateStackInput
	Copy  func(*types.CreateStackInput) CreateStackRequest
}

// Send marshals and sends the CreateStack API request.
func (r CreateStackRequest) Send(ctx context.Context) (*CreateStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateStackResponse{
		CreateStackOutput: r.Request.Data.(*types.CreateStackOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateStackResponse is the response type for the
// CreateStack API operation.
type CreateStackResponse struct {
	*types.CreateStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateStack request.
func (r *CreateStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
