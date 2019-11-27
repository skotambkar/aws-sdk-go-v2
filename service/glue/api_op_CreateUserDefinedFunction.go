// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opCreateUserDefinedFunction = "CreateUserDefinedFunction"

// CreateUserDefinedFunctionRequest returns a request value for making API operation for
// AWS Glue.
//
// Creates a new function definition in the Data Catalog.
//
//    // Example sending a request using CreateUserDefinedFunctionRequest.
//    req := client.CreateUserDefinedFunctionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/CreateUserDefinedFunction
func (c *Client) CreateUserDefinedFunctionRequest(input *types.CreateUserDefinedFunctionInput) CreateUserDefinedFunctionRequest {
	op := &aws.Operation{
		Name:       opCreateUserDefinedFunction,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateUserDefinedFunctionInput{}
	}

	req := c.newRequest(op, input, &types.CreateUserDefinedFunctionOutput{})
	return CreateUserDefinedFunctionRequest{Request: req, Input: input, Copy: c.CreateUserDefinedFunctionRequest}
}

// CreateUserDefinedFunctionRequest is the request type for the
// CreateUserDefinedFunction API operation.
type CreateUserDefinedFunctionRequest struct {
	*aws.Request
	Input *types.CreateUserDefinedFunctionInput
	Copy  func(*types.CreateUserDefinedFunctionInput) CreateUserDefinedFunctionRequest
}

// Send marshals and sends the CreateUserDefinedFunction API request.
func (r CreateUserDefinedFunctionRequest) Send(ctx context.Context) (*CreateUserDefinedFunctionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserDefinedFunctionResponse{
		CreateUserDefinedFunctionOutput: r.Request.Data.(*types.CreateUserDefinedFunctionOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserDefinedFunctionResponse is the response type for the
// CreateUserDefinedFunction API operation.
type CreateUserDefinedFunctionResponse struct {
	*types.CreateUserDefinedFunctionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUserDefinedFunction request.
func (r *CreateUserDefinedFunctionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
