// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opCreateLoginProfile = "CreateLoginProfile"

// CreateLoginProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Creates a password for the specified user, giving the user the ability to
// access AWS services through the AWS Management Console. For more information
// about managing passwords, see Managing Passwords (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_ManagingLogins.html)
// in the IAM User Guide.
//
//    // Example sending a request using CreateLoginProfileRequest.
//    req := client.CreateLoginProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/CreateLoginProfile
func (c *Client) CreateLoginProfileRequest(input *types.CreateLoginProfileInput) CreateLoginProfileRequest {
	op := &aws.Operation{
		Name:       opCreateLoginProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateLoginProfileInput{}
	}

	req := c.newRequest(op, input, &types.CreateLoginProfileOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.CreateLoginProfileMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateLoginProfileRequest{Request: req, Input: input, Copy: c.CreateLoginProfileRequest}
}

// CreateLoginProfileRequest is the request type for the
// CreateLoginProfile API operation.
type CreateLoginProfileRequest struct {
	*aws.Request
	Input *types.CreateLoginProfileInput
	Copy  func(*types.CreateLoginProfileInput) CreateLoginProfileRequest
}

// Send marshals and sends the CreateLoginProfile API request.
func (r CreateLoginProfileRequest) Send(ctx context.Context) (*CreateLoginProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateLoginProfileResponse{
		CreateLoginProfileOutput: r.Request.Data.(*types.CreateLoginProfileOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateLoginProfileResponse is the response type for the
// CreateLoginProfile API operation.
type CreateLoginProfileResponse struct {
	*types.CreateLoginProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateLoginProfile request.
func (r *CreateLoginProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
