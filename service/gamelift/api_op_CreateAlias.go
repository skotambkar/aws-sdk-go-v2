// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opCreateAlias = "CreateAlias"

// CreateAliasRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Creates an alias for a fleet. In most situations, you can use an alias ID
// in place of a fleet ID. By using a fleet alias instead of a specific fleet
// ID, you can switch gameplay and players to a new fleet without changing your
// game client or other game components. For example, for games in production,
// using an alias allows you to seamlessly redirect your player base to a new
// game server update.
//
// Amazon GameLift supports two types of routing strategies for aliases: simple
// and terminal. A simple alias points to an active fleet. A terminal alias
// is used to display messaging or link to a URL instead of routing players
// to an active fleet. For example, you might use a terminal alias when a game
// version is no longer supported and you want to direct players to an upgrade
// site.
//
// To create a fleet alias, specify an alias name, routing strategy, and optional
// description. Each simple alias can point to only one fleet, but a fleet can
// have multiple aliases. If successful, a new alias record is returned, including
// an alias ID, which you can reference when creating a game session. You can
// reassign an alias to another fleet by calling UpdateAlias.
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
//    // Example sending a request using CreateAliasRequest.
//    req := client.CreateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateAlias
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
