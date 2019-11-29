// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opResolveRoom = "ResolveRoom"

// ResolveRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Determines the details for the room from which a skill request was invoked.
// This operation is used by skill developers.
//
//    // Example sending a request using ResolveRoomRequest.
//    req := client.ResolveRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/ResolveRoom
func (c *Client) ResolveRoomRequest(input *types.ResolveRoomInput) ResolveRoomRequest {
	op := &aws.Operation{
		Name:       opResolveRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResolveRoomInput{}
	}

	req := c.newRequest(op, input, &types.ResolveRoomOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ResolveRoomMarshaler{Input: input}.GetNamedBuildHandler())

	return ResolveRoomRequest{Request: req, Input: input, Copy: c.ResolveRoomRequest}
}

// ResolveRoomRequest is the request type for the
// ResolveRoom API operation.
type ResolveRoomRequest struct {
	*aws.Request
	Input *types.ResolveRoomInput
	Copy  func(*types.ResolveRoomInput) ResolveRoomRequest
}

// Send marshals and sends the ResolveRoom API request.
func (r ResolveRoomRequest) Send(ctx context.Context) (*ResolveRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResolveRoomResponse{
		ResolveRoomOutput: r.Request.Data.(*types.ResolveRoomOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResolveRoomResponse is the response type for the
// ResolveRoom API operation.
type ResolveRoomResponse struct {
	*types.ResolveRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResolveRoom request.
func (r *ResolveRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
