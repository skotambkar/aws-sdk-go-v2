// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

const opGetCallerIdentity = "GetCallerIdentity"

// GetCallerIdentityRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns details about the IAM user or role whose credentials are used to
// call the operation.
//
// No permissions are required to perform this operation. If an administrator
// adds a policy to your IAM user or role that explicitly denies access to the
// sts:GetCallerIdentity action, you can still perform this operation. Permissions
// are not required because the same information is returned when an IAM user
// or role is denied access. To view an example response, see I Am Not Authorized
// to Perform: iam:DeleteVirtualMFADevice (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_access-denied-delete-mfa).
//
//    // Example sending a request using GetCallerIdentityRequest.
//    req := client.GetCallerIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/GetCallerIdentity
func (c *Client) GetCallerIdentityRequest(input *types.GetCallerIdentityInput) GetCallerIdentityRequest {
	op := &aws.Operation{
		Name:       opGetCallerIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetCallerIdentityInput{}
	}

	req := c.newRequest(op, input, &types.GetCallerIdentityOutput{})
	return GetCallerIdentityRequest{Request: req, Input: input, Copy: c.GetCallerIdentityRequest}
}

// GetCallerIdentityRequest is the request type for the
// GetCallerIdentity API operation.
type GetCallerIdentityRequest struct {
	*aws.Request
	Input *types.GetCallerIdentityInput
	Copy  func(*types.GetCallerIdentityInput) GetCallerIdentityRequest
}

// Send marshals and sends the GetCallerIdentity API request.
func (r GetCallerIdentityRequest) Send(ctx context.Context) (*GetCallerIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCallerIdentityResponse{
		GetCallerIdentityOutput: r.Request.Data.(*types.GetCallerIdentityOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCallerIdentityResponse is the response type for the
// GetCallerIdentity API operation.
type GetCallerIdentityResponse struct {
	*types.GetCallerIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCallerIdentity request.
func (r *GetCallerIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
