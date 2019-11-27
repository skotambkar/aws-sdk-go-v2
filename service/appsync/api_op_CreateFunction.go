// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opCreateFunction = "CreateFunction"

// CreateFunctionRequest returns a request value for making API operation for
// AWS AppSync.
//
// Creates a Function object.
//
// A function is a reusable entity. Multiple functions can be used to compose
// the resolver logic.
//
//    // Example sending a request using CreateFunctionRequest.
//    req := client.CreateFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateFunction
func (c *Client) CreateFunctionRequest(input *types.CreateFunctionInput) CreateFunctionRequest {
	op := &aws.Operation{
		Name:       opCreateFunction,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/functions",
	}

	if input == nil {
		input = &types.CreateFunctionInput{}
	}

	req := c.newRequest(op, input, &types.CreateFunctionOutput{})
	return CreateFunctionRequest{Request: req, Input: input, Copy: c.CreateFunctionRequest}
}

// CreateFunctionRequest is the request type for the
// CreateFunction API operation.
type CreateFunctionRequest struct {
	*aws.Request
	Input *types.CreateFunctionInput
	Copy  func(*types.CreateFunctionInput) CreateFunctionRequest
}

// Send marshals and sends the CreateFunction API request.
func (r CreateFunctionRequest) Send(ctx context.Context) (*CreateFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateFunctionResponse{
		CreateFunctionOutput: r.Request.Data.(*types.CreateFunctionOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateFunctionResponse is the response type for the
// CreateFunction API operation.
type CreateFunctionResponse struct {
	*types.CreateFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateFunction request.
func (r *CreateFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
