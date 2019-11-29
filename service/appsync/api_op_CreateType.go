// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appsync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
)

const opCreateType = "CreateType"

// CreateTypeRequest returns a request value for making API operation for
// AWS AppSync.
//
// Creates a Type object.
//
//    // Example sending a request using CreateTypeRequest.
//    req := client.CreateTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appsync-2017-07-25/CreateType
func (c *Client) CreateTypeRequest(input *types.CreateTypeInput) CreateTypeRequest {
	op := &aws.Operation{
		Name:       opCreateType,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/apis/{apiId}/types",
	}

	if input == nil {
		input = &types.CreateTypeInput{}
	}

	req := c.newRequest(op, input, &types.CreateTypeOutput{})
	return CreateTypeRequest{Request: req, Input: input, Copy: c.CreateTypeRequest}
}

// CreateTypeRequest is the request type for the
// CreateType API operation.
type CreateTypeRequest struct {
	*aws.Request
	Input *types.CreateTypeInput
	Copy  func(*types.CreateTypeInput) CreateTypeRequest
}

// Send marshals and sends the CreateType API request.
func (r CreateTypeRequest) Send(ctx context.Context) (*CreateTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTypeResponse{
		CreateTypeOutput: r.Request.Data.(*types.CreateTypeOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTypeResponse is the response type for the
// CreateType API operation.
type CreateTypeResponse struct {
	*types.CreateTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateType request.
func (r *CreateTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
