// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

const opDeleteMembers = "DeleteMembers"

// DeleteMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Deletes the specified member accounts from Security Hub.
//
//    // Example sending a request using DeleteMembersRequest.
//    req := client.DeleteMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DeleteMembers
func (c *Client) DeleteMembersRequest(input *types.DeleteMembersInput) DeleteMembersRequest {
	op := &aws.Operation{
		Name:       opDeleteMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/members/delete",
	}

	if input == nil {
		input = &types.DeleteMembersInput{}
	}

	req := c.newRequest(op, input, &types.DeleteMembersOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteMembersMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteMembersRequest{Request: req, Input: input, Copy: c.DeleteMembersRequest}
}

// DeleteMembersRequest is the request type for the
// DeleteMembers API operation.
type DeleteMembersRequest struct {
	*aws.Request
	Input *types.DeleteMembersInput
	Copy  func(*types.DeleteMembersInput) DeleteMembersRequest
}

// Send marshals and sends the DeleteMembers API request.
func (r DeleteMembersRequest) Send(ctx context.Context) (*DeleteMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMembersResponse{
		DeleteMembersOutput: r.Request.Data.(*types.DeleteMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMembersResponse is the response type for the
// DeleteMembers API operation.
type DeleteMembersResponse struct {
	*types.DeleteMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMembers request.
func (r *DeleteMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
