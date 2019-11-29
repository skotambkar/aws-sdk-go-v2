// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

const opGetMembers = "GetMembers"

// GetMembersRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Returns the details on the Security Hub member accounts that the account
// IDs specify.
//
//    // Example sending a request using GetMembersRequest.
//    req := client.GetMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/GetMembers
func (c *Client) GetMembersRequest(input *types.GetMembersInput) GetMembersRequest {
	op := &aws.Operation{
		Name:       opGetMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/members/get",
	}

	if input == nil {
		input = &types.GetMembersInput{}
	}

	req := c.newRequest(op, input, &types.GetMembersOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetMembersMarshaler{Input: input}.GetNamedBuildHandler())

	return GetMembersRequest{Request: req, Input: input, Copy: c.GetMembersRequest}
}

// GetMembersRequest is the request type for the
// GetMembers API operation.
type GetMembersRequest struct {
	*aws.Request
	Input *types.GetMembersInput
	Copy  func(*types.GetMembersInput) GetMembersRequest
}

// Send marshals and sends the GetMembers API request.
func (r GetMembersRequest) Send(ctx context.Context) (*GetMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetMembersResponse{
		GetMembersOutput: r.Request.Data.(*types.GetMembersOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetMembersResponse is the response type for the
// GetMembers API operation.
type GetMembersResponse struct {
	*types.GetMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetMembers request.
func (r *GetMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
