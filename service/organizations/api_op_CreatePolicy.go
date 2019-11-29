// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opCreatePolicy = "CreatePolicy"

// CreatePolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Creates a policy of a specified type that you can attach to a root, an organizational
// unit (OU), or an individual AWS account.
//
// For more information about policies and their use, see Managing Organization
// Policies (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html).
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using CreatePolicyRequest.
//    req := client.CreatePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/CreatePolicy
func (c *Client) CreatePolicyRequest(input *types.CreatePolicyInput) CreatePolicyRequest {
	op := &aws.Operation{
		Name:       opCreatePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreatePolicyInput{}
	}

	req := c.newRequest(op, input, &types.CreatePolicyOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreatePolicyMarshaler{Input: input}.GetNamedBuildHandler())

	return CreatePolicyRequest{Request: req, Input: input, Copy: c.CreatePolicyRequest}
}

// CreatePolicyRequest is the request type for the
// CreatePolicy API operation.
type CreatePolicyRequest struct {
	*aws.Request
	Input *types.CreatePolicyInput
	Copy  func(*types.CreatePolicyInput) CreatePolicyRequest
}

// Send marshals and sends the CreatePolicy API request.
func (r CreatePolicyRequest) Send(ctx context.Context) (*CreatePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePolicyResponse{
		CreatePolicyOutput: r.Request.Data.(*types.CreatePolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePolicyResponse is the response type for the
// CreatePolicy API operation.
type CreatePolicyResponse struct {
	*types.CreatePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePolicy request.
func (r *CreatePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
