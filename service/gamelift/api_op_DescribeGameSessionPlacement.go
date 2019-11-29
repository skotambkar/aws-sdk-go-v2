// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opDescribeGameSessionPlacement = "DescribeGameSessionPlacement"

// DescribeGameSessionPlacementRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Retrieves properties and current status of a game session placement request.
// To get game session placement details, specify the placement ID. If successful,
// a GameSessionPlacement object is returned.
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
//    // Example sending a request using DescribeGameSessionPlacementRequest.
//    req := client.DescribeGameSessionPlacementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/DescribeGameSessionPlacement
func (c *Client) DescribeGameSessionPlacementRequest(input *types.DescribeGameSessionPlacementInput) DescribeGameSessionPlacementRequest {
	op := &aws.Operation{
		Name:       opDescribeGameSessionPlacement,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeGameSessionPlacementInput{}
	}

	req := c.newRequest(op, input, &types.DescribeGameSessionPlacementOutput{})
	return DescribeGameSessionPlacementRequest{Request: req, Input: input, Copy: c.DescribeGameSessionPlacementRequest}
}

// DescribeGameSessionPlacementRequest is the request type for the
// DescribeGameSessionPlacement API operation.
type DescribeGameSessionPlacementRequest struct {
	*aws.Request
	Input *types.DescribeGameSessionPlacementInput
	Copy  func(*types.DescribeGameSessionPlacementInput) DescribeGameSessionPlacementRequest
}

// Send marshals and sends the DescribeGameSessionPlacement API request.
func (r DescribeGameSessionPlacementRequest) Send(ctx context.Context) (*DescribeGameSessionPlacementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeGameSessionPlacementResponse{
		DescribeGameSessionPlacementOutput: r.Request.Data.(*types.DescribeGameSessionPlacementOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeGameSessionPlacementResponse is the response type for the
// DescribeGameSessionPlacement API operation.
type DescribeGameSessionPlacementResponse struct {
	*types.DescribeGameSessionPlacementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeGameSessionPlacement request.
func (r *DescribeGameSessionPlacementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
