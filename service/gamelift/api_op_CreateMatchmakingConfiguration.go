// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opCreateMatchmakingConfiguration = "CreateMatchmakingConfiguration"

// CreateMatchmakingConfigurationRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Defines a new matchmaking configuration for use with FlexMatch. A matchmaking
// configuration sets out guidelines for matching players and getting the matches
// into games. You can set up multiple matchmaking configurations to handle
// the scenarios needed for your game. Each matchmaking ticket (StartMatchmaking
// or StartMatchBackfill) specifies a configuration for the match and provides
// player attributes to support the configuration being used.
//
// To create a matchmaking configuration, at a minimum you must specify the
// following: configuration name; a rule set that governs how to evaluate players
// and find acceptable matches; a game session queue to use when placing a new
// game session for the match; and the maximum time allowed for a matchmaking
// attempt.
//
// There are two ways to track the progress of matchmaking tickets: (1) polling
// ticket status with DescribeMatchmaking; or (2) receiving notifications with
// Amazon Simple Notification Service (SNS). To use notifications, you first
// need to set up an SNS topic to receive the notifications, and provide the
// topic ARN in the matchmaking configuration. Since notifications promise only
// "best effort" delivery, we recommend calling DescribeMatchmaking if no notifications
// are received within 30 seconds.
//
// Learn more
//
//  Design a FlexMatch Matchmaker (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-configuration.html)
//
//  Setting up Notifications for Matchmaking (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-notification.html)
//
// Related operations
//
//    * CreateMatchmakingConfiguration
//
//    * DescribeMatchmakingConfigurations
//
//    * UpdateMatchmakingConfiguration
//
//    * DeleteMatchmakingConfiguration
//
//    * CreateMatchmakingRuleSet
//
//    * DescribeMatchmakingRuleSets
//
//    * ValidateMatchmakingRuleSet
//
//    * DeleteMatchmakingRuleSet
//
//    // Example sending a request using CreateMatchmakingConfigurationRequest.
//    req := client.CreateMatchmakingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/CreateMatchmakingConfiguration
func (c *Client) CreateMatchmakingConfigurationRequest(input *types.CreateMatchmakingConfigurationInput) CreateMatchmakingConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateMatchmakingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateMatchmakingConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.CreateMatchmakingConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateMatchmakingConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateMatchmakingConfigurationRequest{Request: req, Input: input, Copy: c.CreateMatchmakingConfigurationRequest}
}

// CreateMatchmakingConfigurationRequest is the request type for the
// CreateMatchmakingConfiguration API operation.
type CreateMatchmakingConfigurationRequest struct {
	*aws.Request
	Input *types.CreateMatchmakingConfigurationInput
	Copy  func(*types.CreateMatchmakingConfigurationInput) CreateMatchmakingConfigurationRequest
}

// Send marshals and sends the CreateMatchmakingConfiguration API request.
func (r CreateMatchmakingConfigurationRequest) Send(ctx context.Context) (*CreateMatchmakingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMatchmakingConfigurationResponse{
		CreateMatchmakingConfigurationOutput: r.Request.Data.(*types.CreateMatchmakingConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMatchmakingConfigurationResponse is the response type for the
// CreateMatchmakingConfiguration API operation.
type CreateMatchmakingConfigurationResponse struct {
	*types.CreateMatchmakingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMatchmakingConfiguration request.
func (r *CreateMatchmakingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
