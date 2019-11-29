// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opAddRoleToInstanceProfile = "AddRoleToInstanceProfile"

// AddRoleToInstanceProfileRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Adds the specified IAM role to the specified instance profile. An instance
// profile can contain only one role, and this limit cannot be increased. You
// can remove the existing role and then add a different role to an instance
// profile. You must then wait for the change to appear across all of AWS because
// of eventual consistency (https://en.wikipedia.org/wiki/Eventual_consistency).
// To force the change, you must disassociate the instance profile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DisassociateIamInstanceProfile.html)
// and then associate the instance profile (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociateIamInstanceProfile.html),
// or you can stop your instance and then restart it.
//
// The caller of this API must be granted the PassRole permission on the IAM
// role by a permissions policy.
//
// For more information about roles, go to Working with Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/WorkingWithRoles.html).
// For more information about instance profiles, go to About Instance Profiles
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/AboutInstanceProfiles.html).
//
//    // Example sending a request using AddRoleToInstanceProfileRequest.
//    req := client.AddRoleToInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/AddRoleToInstanceProfile
func (c *Client) AddRoleToInstanceProfileRequest(input *types.AddRoleToInstanceProfileInput) AddRoleToInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opAddRoleToInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AddRoleToInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &types.AddRoleToInstanceProfileOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return AddRoleToInstanceProfileRequest{Request: req, Input: input, Copy: c.AddRoleToInstanceProfileRequest}
}

// AddRoleToInstanceProfileRequest is the request type for the
// AddRoleToInstanceProfile API operation.
type AddRoleToInstanceProfileRequest struct {
	*aws.Request
	Input *types.AddRoleToInstanceProfileInput
	Copy  func(*types.AddRoleToInstanceProfileInput) AddRoleToInstanceProfileRequest
}

// Send marshals and sends the AddRoleToInstanceProfile API request.
func (r AddRoleToInstanceProfileRequest) Send(ctx context.Context) (*AddRoleToInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AddRoleToInstanceProfileResponse{
		AddRoleToInstanceProfileOutput: r.Request.Data.(*types.AddRoleToInstanceProfileOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AddRoleToInstanceProfileResponse is the response type for the
// AddRoleToInstanceProfile API operation.
type AddRoleToInstanceProfileResponse struct {
	*types.AddRoleToInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AddRoleToInstanceProfile request.
func (r *AddRoleToInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
