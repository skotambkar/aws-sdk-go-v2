// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package gamelift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
)

const opUpdateGameSessionQueue = "UpdateGameSessionQueue"

// UpdateGameSessionQueueRequest returns a request value for making API operation for
// Amazon GameLift.
//
// Updates settings for a game session queue, which determines how new game
// session requests in the queue are processed. To update settings, specify
// the queue name to be updated and provide the new settings. When updating
// destinations, provide a complete list of destinations.
//
//    * CreateGameSessionQueue
//
//    * DescribeGameSessionQueues
//
//    * UpdateGameSessionQueue
//
//    * DeleteGameSessionQueue
//
//    // Example sending a request using UpdateGameSessionQueueRequest.
//    req := client.UpdateGameSessionQueueRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/gamelift-2015-10-01/UpdateGameSessionQueue
func (c *Client) UpdateGameSessionQueueRequest(input *types.UpdateGameSessionQueueInput) UpdateGameSessionQueueRequest {
	op := &aws.Operation{
		Name:       opUpdateGameSessionQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateGameSessionQueueInput{}
	}

	req := c.newRequest(op, input, &types.UpdateGameSessionQueueOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateGameSessionQueueMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateGameSessionQueueRequest{Request: req, Input: input, Copy: c.UpdateGameSessionQueueRequest}
}

// UpdateGameSessionQueueRequest is the request type for the
// UpdateGameSessionQueue API operation.
type UpdateGameSessionQueueRequest struct {
	*aws.Request
	Input *types.UpdateGameSessionQueueInput
	Copy  func(*types.UpdateGameSessionQueueInput) UpdateGameSessionQueueRequest
}

// Send marshals and sends the UpdateGameSessionQueue API request.
func (r UpdateGameSessionQueueRequest) Send(ctx context.Context) (*UpdateGameSessionQueueResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGameSessionQueueResponse{
		UpdateGameSessionQueueOutput: r.Request.Data.(*types.UpdateGameSessionQueueOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGameSessionQueueResponse is the response type for the
// UpdateGameSessionQueue API operation.
type UpdateGameSessionQueueResponse struct {
	*types.UpdateGameSessionQueueOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGameSessionQueue request.
func (r *UpdateGameSessionQueueResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
