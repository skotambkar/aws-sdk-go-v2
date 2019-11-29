// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

const opListInvitations = "ListInvitations"

// ListInvitationsRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Lists all Security Hub membership invitations that were sent to the current
// AWS account.
//
//    // Example sending a request using ListInvitationsRequest.
//    req := client.ListInvitationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/ListInvitations
func (c *Client) ListInvitationsRequest(input *types.ListInvitationsInput) ListInvitationsRequest {
	op := &aws.Operation{
		Name:       opListInvitations,
		HTTPMethod: "GET",
		HTTPPath:   "/invitations",
	}

	if input == nil {
		input = &types.ListInvitationsInput{}
	}

	req := c.newRequest(op, input, &types.ListInvitationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListInvitationsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListInvitationsRequest{Request: req, Input: input, Copy: c.ListInvitationsRequest}
}

// ListInvitationsRequest is the request type for the
// ListInvitations API operation.
type ListInvitationsRequest struct {
	*aws.Request
	Input *types.ListInvitationsInput
	Copy  func(*types.ListInvitationsInput) ListInvitationsRequest
}

// Send marshals and sends the ListInvitations API request.
func (r ListInvitationsRequest) Send(ctx context.Context) (*ListInvitationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInvitationsResponse{
		ListInvitationsOutput: r.Request.Data.(*types.ListInvitationsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListInvitationsResponse is the response type for the
// ListInvitations API operation.
type ListInvitationsResponse struct {
	*types.ListInvitationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInvitations request.
func (r *ListInvitationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
