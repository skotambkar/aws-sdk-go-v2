// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opUpdateRoom = "UpdateRoom"

// UpdateRoomRequest returns a request value for making API operation for
// Amazon Chime.
//
// Updates room details, such as the room name.
//
//    // Example sending a request using UpdateRoomRequest.
//    req := client.UpdateRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/UpdateRoom
func (c *Client) UpdateRoomRequest(input *types.UpdateRoomInput) UpdateRoomRequest {
	op := &aws.Operation{
		Name:       opUpdateRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/rooms/{roomId}",
	}

	if input == nil {
		input = &types.UpdateRoomInput{}
	}

	req := c.newRequest(op, input, &types.UpdateRoomOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateRoomMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateRoomRequest{Request: req, Input: input, Copy: c.UpdateRoomRequest}
}

// UpdateRoomRequest is the request type for the
// UpdateRoom API operation.
type UpdateRoomRequest struct {
	*aws.Request
	Input *types.UpdateRoomInput
	Copy  func(*types.UpdateRoomInput) UpdateRoomRequest
}

// Send marshals and sends the UpdateRoom API request.
func (r UpdateRoomRequest) Send(ctx context.Context) (*UpdateRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRoomResponse{
		UpdateRoomOutput: r.Request.Data.(*types.UpdateRoomOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRoomResponse is the response type for the
// UpdateRoom API operation.
type UpdateRoomResponse struct {
	*types.UpdateRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRoom request.
func (r *UpdateRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
