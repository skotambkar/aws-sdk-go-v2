// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opDeleteRoomMembership = "DeleteRoomMembership"

// DeleteRoomMembershipRequest returns a request value for making API operation for
// Amazon Chime.
//
// Removes a member from a chat room.
//
//    // Example sending a request using DeleteRoomMembershipRequest.
//    req := client.DeleteRoomMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/DeleteRoomMembership
func (c *Client) DeleteRoomMembershipRequest(input *types.DeleteRoomMembershipInput) DeleteRoomMembershipRequest {
	op := &aws.Operation{
		Name:       opDeleteRoomMembership,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{accountId}/rooms/{roomId}/memberships/{memberId}",
	}

	if input == nil {
		input = &types.DeleteRoomMembershipInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRoomMembershipOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRoomMembershipRequest{Request: req, Input: input, Copy: c.DeleteRoomMembershipRequest}
}

// DeleteRoomMembershipRequest is the request type for the
// DeleteRoomMembership API operation.
type DeleteRoomMembershipRequest struct {
	*aws.Request
	Input *types.DeleteRoomMembershipInput
	Copy  func(*types.DeleteRoomMembershipInput) DeleteRoomMembershipRequest
}

// Send marshals and sends the DeleteRoomMembership API request.
func (r DeleteRoomMembershipRequest) Send(ctx context.Context) (*DeleteRoomMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRoomMembershipResponse{
		DeleteRoomMembershipOutput: r.Request.Data.(*types.DeleteRoomMembershipOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRoomMembershipResponse is the response type for the
// DeleteRoomMembership API operation.
type DeleteRoomMembershipResponse struct {
	*types.DeleteRoomMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRoomMembership request.
func (r *DeleteRoomMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
