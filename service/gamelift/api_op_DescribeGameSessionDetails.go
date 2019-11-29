// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opDescribeGameSessionDetails = "DescribeGameSessionDetails"

// DescribeGameSessionDetailsRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves properties, including the protection policy in force, for one or
// more game sessions. This action can be used in several ways: (1) provide
// a GameSessionId or GameSessionArn to request details for a specific game
// session; (2) provide either a FleetId or an AliasId to request properties
// for all game sessions running on a fleet.
//
// To get game session record(s), specify just one of the following: game session
// ID, fleet ID, or alias ID. You can filter this request by game session status.
// Use the pagination parameters to retrieve results as a set of sequential
// pages. If successful, a GameSessionDetail object is returned for each session
// matching the request.
//
//    * CreateGameSession
//
//    * DescribeGameSessions
//
//    * DescribeGameSessionDetails
//
//    * SearchGameSessions
//
//    * UpdateGameSession
//
//    * GetGameSessionLogUrl
//
//    * Game session placements StartGameSessionPlacement DescribeGameSessionPlacement
//    StopGameSessionPlacement
//
//    // Example sending a request using DescribeGameSessionDetailsRequest.
//    req := client.DescribeGameSessionDetailsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeGameSessionDetails
func (c *Client) DescribeGameSessionDetailsRequest(input *types.DescribeGameSessionDetailsInput) DescribeGameSessionDetailsRequest {
	op := &aws.Operation{
		Name:       opDescribeGameSessionDetails,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeGameSessionDetailsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeGameSessionDetailsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeGameSessionDetailsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeGameSessionDetailsRequest{Request: req, Input: input, Copy: c.DescribeGameSessionDetailsRequest}
}

// DescribeGameSessionDetailsRequest is the request type for the
// DescribeGameSessionDetails API operation.
type DescribeGameSessionDetailsRequest struct {
	*aws.Request
	Input *types.DescribeGameSessionDetailsInput
	Copy  func(*types.DescribeGameSessionDetailsInput) DescribeGameSessionDetailsRequest
}

// Send marshals and sends the DescribeGameSessionDetails API request.
func (r DescribeGameSessionDetailsRequest) Send(ctx context.Context) (*DescribeGameSessionDetailsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGameSessionDetailsResponse{
		DescribeGameSessionDetailsOutput: r.Request.Data.(*types.DescribeGameSessionDetailsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGameSessionDetailsResponse is the response type for the
// DescribeGameSessionDetails API operation.
type DescribeGameSessionDetailsResponse struct {
	*types.DescribeGameSessionDetailsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGameSessionDetails request.
func (r *DescribeGameSessionDetailsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
