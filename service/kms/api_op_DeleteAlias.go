// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opDeleteAlias = "DeleteAlias"

// DeleteAliasRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Deletes the specified alias. You cannot perform this operation on an alias
// in a different AWS account.
//
// Because an alias is not a property of a CMK, you can delete and change the
// aliases of a CMK without affecting the CMK. Also, aliases do not appear in
// the response from the DescribeKey operation. To get the aliases of all CMKs,
// use the ListAliases operation.
//
// Each CMK can have multiple aliases. To change the alias of a CMK, use DeleteAlias
// to delete the current alias and CreateAlias to create a new alias. To associate
// an existing alias with a different customer master key (CMK), call UpdateAlias.
//
//    // Example sending a request using DeleteAliasRequest.
//    req := client.DeleteAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/DeleteAlias
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
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
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
