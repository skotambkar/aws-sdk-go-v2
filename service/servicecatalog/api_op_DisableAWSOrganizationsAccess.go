// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDisableAWSOrganizationsAccess = "DisableAWSOrganizationsAccess"

// DisableAWSOrganizationsAccessRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disable portfolio sharing through AWS Organizations feature. This feature
// will not delete your current shares but it will prevent you from creating
// new shares throughout your organization. Current shares will not be in sync
// with your organization structure if it changes after calling this API. This
// API can only be called by the master account in the organization.
//
//    // Example sending a request using DisableAWSOrganizationsAccessRequest.
//    req := client.DisableAWSOrganizationsAccessRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DisableAWSOrganizationsAccess
func (c *Client) DisableAWSOrganizationsAccessRequest(input *types.DisableAWSOrganizationsAccessInput) DisableAWSOrganizationsAccessRequest {
	op := &aws.Operation{
		Name:       opDisableAWSOrganizationsAccess,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisableAWSOrganizationsAccessInput{}
	}

	req := c.newRequest(op, input, &types.DisableAWSOrganizationsAccessOutput{})
	return DisableAWSOrganizationsAccessRequest{Request: req, Input: input, Copy: c.DisableAWSOrganizationsAccessRequest}
}

// DisableAWSOrganizationsAccessRequest is the request type for the
// DisableAWSOrganizationsAccess API operation.
type DisableAWSOrganizationsAccessRequest struct {
	*aws.Request
	Input *types.DisableAWSOrganizationsAccessInput
	Copy  func(*types.DisableAWSOrganizationsAccessInput) DisableAWSOrganizationsAccessRequest
}

// Send marshals and sends the DisableAWSOrganizationsAccess API request.
func (r DisableAWSOrganizationsAccessRequest) Send(ctx context.Context) (*DisableAWSOrganizationsAccessResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableAWSOrganizationsAccessResponse{
		DisableAWSOrganizationsAccessOutput: r.Request.Data.(*types.DisableAWSOrganizationsAccessOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableAWSOrganizationsAccessResponse is the response type for the
// DisableAWSOrganizationsAccess API operation.
type DisableAWSOrganizationsAccessResponse struct {
	*types.DisableAWSOrganizationsAccessOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableAWSOrganizationsAccess request.
func (r *DisableAWSOrganizationsAccessResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
