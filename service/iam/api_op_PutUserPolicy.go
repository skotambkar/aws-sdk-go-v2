// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/iam/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opPutUserPolicy = "PutUserPolicy"

// PutUserPolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Adds or updates an inline policy document that is embedded in the specified
// IAM user.
//
// An IAM user can also have a managed policy attached to it. To attach a managed
// policy to a user, use AttachUserPolicy. To create a new managed policy, use
// CreatePolicy. For information about policies, see Managed Policies and Inline
// Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
// For information about limits on the number of inline policies that you can
// embed in a user, see Limitations on IAM Entities (https://docs.aws.amazon.com/IAM/latest/UserGuide/LimitationsOnEntities.html)
// in the IAM User Guide.
//
// Because policy documents can be large, you should use POST rather than GET
// when calling PutUserPolicy. For general information about using the Query
// API with IAM, go to Making Query Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html)
// in the IAM User Guide.
//
//    // Example sending a request using PutUserPolicyRequest.
//    req := client.PutUserPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/PutUserPolicy
func (c *Client) PutUserPolicyRequest(input *types.PutUserPolicyInput) PutUserPolicyRequest {
	op := &aws.Operation{
		Name:       opPutUserPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutUserPolicyInput{}
	}

	req := c.newRequest(op, input, &types.PutUserPolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.PutUserPolicyMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutUserPolicyRequest{Request: req, Input: input, Copy: c.PutUserPolicyRequest}
}

// PutUserPolicyRequest is the request type for the
// PutUserPolicy API operation.
type PutUserPolicyRequest struct {
	*aws.Request
	Input *types.PutUserPolicyInput
	Copy  func(*types.PutUserPolicyInput) PutUserPolicyRequest
}

// Send marshals and sends the PutUserPolicy API request.
func (r PutUserPolicyRequest) Send(ctx context.Context) (*PutUserPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutUserPolicyResponse{
		PutUserPolicyOutput: r.Request.Data.(*types.PutUserPolicyOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutUserPolicyResponse is the response type for the
// PutUserPolicy API operation.
type PutUserPolicyResponse struct {
	*types.PutUserPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutUserPolicy request.
func (r *PutUserPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
