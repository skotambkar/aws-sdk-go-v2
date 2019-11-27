// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codestar

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codestar/types"
)

const opListUserProfiles = "ListUserProfiles"

// ListUserProfilesRequest returns a request value for making API operation for
// AWS CodeStar.
//
// Lists all the user profiles configured for your AWS account in AWS CodeStar.
//
//    // Example sending a request using ListUserProfilesRequest.
//    req := client.ListUserProfilesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codestar-2017-04-19/ListUserProfiles
func (c *Client) ListUserProfilesRequest(input *types.ListUserProfilesInput) ListUserProfilesRequest {
	op := &aws.Operation{
		Name:       opListUserProfiles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListUserProfilesInput{}
	}

	req := c.newRequest(op, input, &types.ListUserProfilesOutput{})
	return ListUserProfilesRequest{Request: req, Input: input, Copy: c.ListUserProfilesRequest}
}

// ListUserProfilesRequest is the request type for the
// ListUserProfiles API operation.
type ListUserProfilesRequest struct {
	*aws.Request
	Input *types.ListUserProfilesInput
	Copy  func(*types.ListUserProfilesInput) ListUserProfilesRequest
}

// Send marshals and sends the ListUserProfiles API request.
func (r ListUserProfilesRequest) Send(ctx context.Context) (*ListUserProfilesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListUserProfilesResponse{
		ListUserProfilesOutput: r.Request.Data.(*types.ListUserProfilesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListUserProfilesResponse is the response type for the
// ListUserProfiles API operation.
type ListUserProfilesResponse struct {
	*types.ListUserProfilesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListUserProfiles request.
func (r *ListUserProfilesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
