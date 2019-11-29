// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package alexaforbusiness

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
)

const opAssociateSkillGroupWithRoom = "AssociateSkillGroupWithRoom"

// AssociateSkillGroupWithRoomRequest returns a request value for making API operation for
// Alexa For Business.
//
// Associates a skill group with a given room. This enables all skills in the
// associated skill group on all devices in the room.
//
//    // Example sending a request using AssociateSkillGroupWithRoomRequest.
//    req := client.AssociateSkillGroupWithRoomRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/alexaforbusiness-2017-11-09/AssociateSkillGroupWithRoom
func (c *Client) AssociateSkillGroupWithRoomRequest(input *types.AssociateSkillGroupWithRoomInput) AssociateSkillGroupWithRoomRequest {
	op := &aws.Operation{
		Name:       opAssociateSkillGroupWithRoom,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateSkillGroupWithRoomInput{}
	}

	req := c.newRequest(op, input, &types.AssociateSkillGroupWithRoomOutput{})
	return AssociateSkillGroupWithRoomRequest{Request: req, Input: input, Copy: c.AssociateSkillGroupWithRoomRequest}
}

// AssociateSkillGroupWithRoomRequest is the request type for the
// AssociateSkillGroupWithRoom API operation.
type AssociateSkillGroupWithRoomRequest struct {
	*aws.Request
	Input *types.AssociateSkillGroupWithRoomInput
	Copy  func(*types.AssociateSkillGroupWithRoomInput) AssociateSkillGroupWithRoomRequest
}

// Send marshals and sends the AssociateSkillGroupWithRoom API request.
func (r AssociateSkillGroupWithRoomRequest) Send(ctx context.Context) (*AssociateSkillGroupWithRoomResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateSkillGroupWithRoomResponse{
		AssociateSkillGroupWithRoomOutput: r.Request.Data.(*types.AssociateSkillGroupWithRoomOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateSkillGroupWithRoomResponse is the response type for the
// AssociateSkillGroupWithRoom API operation.
type AssociateSkillGroupWithRoomResponse struct {
	*types.AssociateSkillGroupWithRoomOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateSkillGroupWithRoom request.
func (r *AssociateSkillGroupWithRoomResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
