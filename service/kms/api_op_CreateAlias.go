// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

const opCreateAlias = "CreateAlias"

// CreateAliasRequest returns a request value for making API operation for
// AWS Key Management Service.
//
// Creates a display name for a customer managed customer master key (CMK).
// You can use an alias to identify a CMK in selected operations, such as Encrypt
// and GenerateDataKey.
//
// Each CMK can have multiple aliases, but each alias points to only one CMK.
// The alias name must be unique in the AWS account and region. To simplify
// code that runs in multiple regions, use the same alias name, but point it
// to a different CMK in each region.
//
// Because an alias is not a property of a CMK, you can delete and change the
// aliases of a CMK without affecting the CMK. Also, aliases do not appear in
// the response from the DescribeKey operation. To get the aliases of all CMKs,
// use the ListAliases operation.
//
// The alias name must begin with alias/ followed by a name, such as alias/ExampleAlias.
// It can contain only alphanumeric characters, forward slashes (/), underscores
// (_), and dashes (-). The alias name cannot begin with alias/aws/. The alias/aws/
// prefix is reserved for AWS managed CMKs (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
//
// The alias and the CMK it is mapped to must be in the same AWS account and
// the same region. You cannot perform this operation on an alias in a different
// AWS account.
//
// To map an existing alias to a different CMK, call UpdateAlias.
//
// The result of this operation varies with the key state of the CMK. For details,
// see How Key State Affects Use of a Customer Master Key (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the AWS Key Management Service Developer Guide.
//
//    // Example sending a request using CreateAliasRequest.
//    req := client.CreateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kms-2014-11-01/CreateAlias
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

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateAliasMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
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
