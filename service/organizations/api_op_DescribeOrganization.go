// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opDescribeOrganization = "DescribeOrganization"

// DescribeOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Retrieves information about the organization that the user's account belongs
// to.
//
// This operation can be called from any account in the organization.
//
// Even if a policy type is shown as available in the organization, you can
// disable it separately at the root level with DisablePolicyType. Use ListRoots
// to see the status of policy types for a specified root.
//
//    // Example sending a request using DescribeOrganizationRequest.
//    req := client.DescribeOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/DescribeOrganization
func (c *Client) DescribeOrganizationRequest(input *types.DescribeOrganizationInput) DescribeOrganizationRequest {
	op := &aws.Operation{
		Name:       opDescribeOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeOrganizationInput{}
	}

	req := c.newRequest(op, input, &types.DescribeOrganizationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeOrganizationMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeOrganizationRequest{Request: req, Input: input, Copy: c.DescribeOrganizationRequest}
}

// DescribeOrganizationRequest is the request type for the
// DescribeOrganization API operation.
type DescribeOrganizationRequest struct {
	*aws.Request
	Input *types.DescribeOrganizationInput
	Copy  func(*types.DescribeOrganizationInput) DescribeOrganizationRequest
}

// Send marshals and sends the DescribeOrganization API request.
func (r DescribeOrganizationRequest) Send(ctx context.Context) (*DescribeOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrganizationResponse{
		DescribeOrganizationOutput: r.Request.Data.(*types.DescribeOrganizationOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOrganizationResponse is the response type for the
// DescribeOrganization API operation.
type DescribeOrganizationResponse struct {
	*types.DescribeOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrganization request.
func (r *DescribeOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
