// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opDeleteAccountAlias = "DeleteAccountAlias"

// DeleteAccountAliasRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified AWS account alias. For information about using an AWS
// account alias, see Using an Alias for Your AWS Account ID (https://docs.aws.amazon.com/IAM/latest/UserGuide/AccountAlias.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeleteAccountAliasRequest.
//    req := client.DeleteAccountAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeleteAccountAlias
func (c *Client) DeleteAccountAliasRequest(input *types.DeleteAccountAliasInput) DeleteAccountAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteAccountAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAccountAliasInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAccountAliasOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteAccountAliasMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteAccountAliasRequest{Request: req, Input: input, Copy: c.DeleteAccountAliasRequest}
}

// DeleteAccountAliasRequest is the request type for the
// DeleteAccountAlias API operation.
type DeleteAccountAliasRequest struct {
	*aws.Request
	Input *types.DeleteAccountAliasInput
	Copy  func(*types.DeleteAccountAliasInput) DeleteAccountAliasRequest
}

// Send marshals and sends the DeleteAccountAlias API request.
func (r DeleteAccountAliasRequest) Send(ctx context.Context) (*DeleteAccountAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAccountAliasResponse{
		DeleteAccountAliasOutput: r.Request.Data.(*types.DeleteAccountAliasOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAccountAliasResponse is the response type for the
// DeleteAccountAlias API operation.
type DeleteAccountAliasResponse struct {
	*types.DeleteAccountAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAccountAlias request.
func (r *DeleteAccountAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
