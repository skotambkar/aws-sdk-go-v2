// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opResolveAlias = "ResolveAlias"

// ResolveAliasRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves the fleet ID that a specified alias is currently pointing to.
//
//    * CreateAlias
//
//    * ListAliases
//
//    * DescribeAlias
//
//    * UpdateAlias
//
//    * DeleteAlias
//
//    * ResolveAlias
//
//    // Example sending a request using ResolveAliasRequest.
//    req := client.ResolveAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/ResolveAlias
func (c *Client) ResolveAliasRequest(input *types.ResolveAliasInput) ResolveAliasRequest {
	op := &aws.Operation{
		Name:       opResolveAlias,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResolveAliasInput{}
	}

	req := c.newRequest(op, input, &types.ResolveAliasOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ResolveAliasMarshaler{Input: input}.GetNamedBuildHandler())

	return ResolveAliasRequest{Request: req, Input: input, Copy: c.ResolveAliasRequest}
}

// ResolveAliasRequest is the request type for the
// ResolveAlias API operation.
type ResolveAliasRequest struct {
	*aws.Request
	Input *types.ResolveAliasInput
	Copy  func(*types.ResolveAliasInput) ResolveAliasRequest
}

// Send marshals and sends the ResolveAlias API request.
func (r ResolveAliasRequest) Send(ctx context.Context) (*ResolveAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResolveAliasResponse{
		ResolveAliasOutput: r.Request.Data.(*types.ResolveAliasOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResolveAliasResponse is the response type for the
// ResolveAlias API operation.
type ResolveAliasResponse struct {
	*types.ResolveAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResolveAlias request.
func (r *ResolveAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
