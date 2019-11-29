// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opRemoveAccountFromOrganization = "RemoveAccountFromOrganization"

// RemoveAccountFromOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Removes the specified account from the organization.
//
// The removed account becomes a standalone account that isn't a member of any
// organization. It's no longer subject to any policies and is responsible for
// its own bill payments. The organization's master account is no longer charged
// for any expenses accrued by the member account after it's removed from the
// organization.
//
// This operation can be called only from the organization's master account.
// Member accounts can remove themselves with LeaveOrganization instead.
//
// You can remove an account from your organization only if the account is configured
// with the information required to operate as a standalone account. When you
// create an account in an organization using the AWS Organizations console,
// API, or CLI commands, the information required of standalone accounts is
// not automatically collected. For an account that you want to make standalone,
// you must accept the end user license agreement (EULA), choose a support plan,
// provide and verify the required contact information, and provide a current
// payment method. AWS uses the payment method to charge for any billable (not
// free tier) AWS activity that occurs while the account isn't attached to an
// organization. To remove an account that doesn't yet have this information,
// you must sign in as the member account and follow the steps at To leave an
// organization when all required account information has not yet been provided
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_remove.html#leave-without-all-info)
// in the AWS Organizations User Guide.
//
//    // Example sending a request using RemoveAccountFromOrganizationRequest.
//    req := client.RemoveAccountFromOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/RemoveAccountFromOrganization
func (c *Client) RemoveAccountFromOrganizationRequest(input *types.RemoveAccountFromOrganizationInput) RemoveAccountFromOrganizationRequest {
	op := &aws.Operation{
		Name:       opRemoveAccountFromOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RemoveAccountFromOrganizationInput{}
	}

	req := c.newRequest(op, input, &types.RemoveAccountFromOrganizationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RemoveAccountFromOrganizationMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return RemoveAccountFromOrganizationRequest{Request: req, Input: input, Copy: c.RemoveAccountFromOrganizationRequest}
}

// RemoveAccountFromOrganizationRequest is the request type for the
// RemoveAccountFromOrganization API operation.
type RemoveAccountFromOrganizationRequest struct {
	*aws.Request
	Input *types.RemoveAccountFromOrganizationInput
	Copy  func(*types.RemoveAccountFromOrganizationInput) RemoveAccountFromOrganizationRequest
}

// Send marshals and sends the RemoveAccountFromOrganization API request.
func (r RemoveAccountFromOrganizationRequest) Send(ctx context.Context) (*RemoveAccountFromOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RemoveAccountFromOrganizationResponse{
		RemoveAccountFromOrganizationOutput: r.Request.Data.(*types.RemoveAccountFromOrganizationOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RemoveAccountFromOrganizationResponse is the response type for the
// RemoveAccountFromOrganization API operation.
type RemoveAccountFromOrganizationResponse struct {
	*types.RemoveAccountFromOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RemoveAccountFromOrganization request.
func (r *RemoveAccountFromOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
