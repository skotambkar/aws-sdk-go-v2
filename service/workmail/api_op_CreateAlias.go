// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opCreateAlias = "CreateAlias"

// CreateAliasRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Adds an alias to the set of a given member (user or group) of Amazon WorkMail.
//
//    // Example sending a request using CreateAliasRequest.
//    req := client.CreateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/CreateAlias
func (c *Client) CreateAliasRequest(input *types.CreateAliasInput) CreateAliasRequest {
	op := &aws.Operation{
		Name:       opCreateAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateAliasInput{}
	}

	req := c.newRequest(op, input, &types.CreateAliasOutput{})
	return CreateAliasRequest{Request: req, Input: input, Copy: c.CreateAliasRequest}
}

// CreateAliasRequest is the request type for the
// CreateAlias API operation.
type CreateAliasRequest struct {
	*aws.Request
	Input *types.CreateAliasInput
	Copy  func(*types.CreateAliasInput) CreateAliasRequest
}

// Send marshals and sends the CreateAlias API request.
func (r CreateAliasRequest) Send(ctx context.Context) (*CreateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAliasResponse{
		CreateAliasOutput: r.Request.Data.(*types.CreateAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAliasResponse is the response type for the
// CreateAlias API operation.
type CreateAliasResponse struct {
	*types.CreateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAlias request.
func (r *CreateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
