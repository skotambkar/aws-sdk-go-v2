// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workmail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
)

const opDeleteAlias = "DeleteAlias"

// DeleteAliasRequest returns a request value for making API operation for
// Amazon WorkMail.
//
// Remove one or more specified aliases from a set of aliases for a given user.
//
//    // Example sending a request using DeleteAliasRequest.
//    req := client.DeleteAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workmail-2017-10-01/DeleteAlias
func (c *Client) DeleteAliasRequest(input *types.DeleteAliasInput) DeleteAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAliasInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAliasOutput{})
	return DeleteAliasRequest{Request: req, Input: input, Copy: c.DeleteAliasRequest}
}

// DeleteAliasRequest is the request type for the
// DeleteAlias API operation.
type DeleteAliasRequest struct {
	*aws.Request
	Input *types.DeleteAliasInput
	Copy  func(*types.DeleteAliasInput) DeleteAliasRequest
}

// Send marshals and sends the DeleteAlias API request.
func (r DeleteAliasRequest) Send(ctx context.Context) (*DeleteAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAliasResponse{
		DeleteAliasOutput: r.Request.Data.(*types.DeleteAliasOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAliasResponse is the response type for the
// DeleteAlias API operation.
type DeleteAliasResponse struct {
	*types.DeleteAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAlias request.
func (r *DeleteAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
