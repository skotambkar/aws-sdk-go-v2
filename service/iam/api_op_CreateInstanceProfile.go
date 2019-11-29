// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opCreateInstanceProfile = "CreateInstanceProfile"

// CreateInstanceProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a new instance profile. For information about instance profiles,
// go to About Instance Profiles (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
//
// For information about the number of instance profiles you can create, see
// Limitations on IAM Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateInstanceProfileRequest.
//    req := client.CreateInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateInstanceProfile
func (c *Client) CreateInstanceProfileRequest(input *types.CreateInstanceProfileInput) CreateInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opCreateInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &types.CreateInstanceProfileOutput{})
	return CreateInstanceProfileRequest{Request: req, Input: input, Copy: c.CreateInstanceProfileRequest}
}

// CreateInstanceProfileRequest is the request type for the
// CreateInstanceProfile API operation.
type CreateInstanceProfileRequest struct {
	*aws.Request
	Input *types.CreateInstanceProfileInput
	Copy  func(*types.CreateInstanceProfileInput) CreateInstanceProfileRequest
}

// Send marshals and sends the CreateInstanceProfile API request.
func (r CreateInstanceProfileRequest) Send(ctx context.Context) (*CreateInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateInstanceProfileResponse{
		CreateInstanceProfileOutput: r.Request.Data.(*types.CreateInstanceProfileOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateInstanceProfileResponse is the response type for the
// CreateInstanceProfile API operation.
type CreateInstanceProfileResponse struct {
	*types.CreateInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateInstanceProfile request.
func (r *CreateInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
