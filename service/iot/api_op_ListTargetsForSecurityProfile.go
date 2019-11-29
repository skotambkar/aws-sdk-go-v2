// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListTargetsForSecurityProfile = "ListTargetsForSecurityProfile"

// ListTargetsForSecurityProfileRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the targets (thing groups) associated with a given Device Defender
// security profile.
//
//    // Example sending a request using ListTargetsForSecurityProfileRequest.
//    req := client.ListTargetsForSecurityProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListTargetsForSecurityProfileRequest(input *types.ListTargetsForSecurityProfileInput) ListTargetsForSecurityProfileRequest {
	op := &aws.Operation{
		Name:       opListTargetsForSecurityProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/security-profiles/{securityProfileName}/targets",
	}

	if input == nil {
		input = &types.ListTargetsForSecurityProfileInput{}
	}

	req := c.newRequest(op, input, &types.ListTargetsForSecurityProfileOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListTargetsForSecurityProfileMarshaler{Input: input}.GetNamedBuildHandler())

	return ListTargetsForSecurityProfileRequest{Request: req, Input: input, Copy: c.ListTargetsForSecurityProfileRequest}
}

// ListTargetsForSecurityProfileRequest is the request type for the
// ListTargetsForSecurityProfile API operation.
type ListTargetsForSecurityProfileRequest struct {
	*aws.Request
	Input *types.ListTargetsForSecurityProfileInput
	Copy  func(*types.ListTargetsForSecurityProfileInput) ListTargetsForSecurityProfileRequest
}

// Send marshals and sends the ListTargetsForSecurityProfile API request.
func (r ListTargetsForSecurityProfileRequest) Send(ctx context.Context) (*ListTargetsForSecurityProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTargetsForSecurityProfileResponse{
		ListTargetsForSecurityProfileOutput: r.Request.Data.(*types.ListTargetsForSecurityProfileOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListTargetsForSecurityProfileResponse is the response type for the
// ListTargetsForSecurityProfile API operation.
type ListTargetsForSecurityProfileResponse struct {
	*types.ListTargetsForSecurityProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTargetsForSecurityProfile request.
func (r *ListTargetsForSecurityProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
