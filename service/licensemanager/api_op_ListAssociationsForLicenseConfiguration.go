// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
)

const opListAssociationsForLicenseConfiguration = "ListAssociationsForLicenseConfiguration"

// ListAssociationsForLicenseConfigurationRequest returns a request value for making API operation for
// AWS License Manager.
//
// Lists the resource associations for a license configuration. Resource associations
// need not consume licenses from a license configuration. For example, an AMI
// or a stopped instance may not consume a license (depending on the license
// rules). Use this operation to find all resources associated with a license
// configuration.
//
//    // Example sending a request using ListAssociationsForLicenseConfigurationRequest.
//    req := client.ListAssociationsForLicenseConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/ListAssociationsForLicenseConfiguration
func (c *Client) ListAssociationsForLicenseConfigurationRequest(input *types.ListAssociationsForLicenseConfigurationInput) ListAssociationsForLicenseConfigurationRequest {
	op := &aws.Operation{
		Name:       opListAssociationsForLicenseConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListAssociationsForLicenseConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.ListAssociationsForLicenseConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListAssociationsForLicenseConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return ListAssociationsForLicenseConfigurationRequest{Request: req, Input: input, Copy: c.ListAssociationsForLicenseConfigurationRequest}
}

// ListAssociationsForLicenseConfigurationRequest is the request type for the
// ListAssociationsForLicenseConfiguration API operation.
type ListAssociationsForLicenseConfigurationRequest struct {
	*aws.Request
	Input *types.ListAssociationsForLicenseConfigurationInput
	Copy  func(*types.ListAssociationsForLicenseConfigurationInput) ListAssociationsForLicenseConfigurationRequest
}

// Send marshals and sends the ListAssociationsForLicenseConfiguration API request.
func (r ListAssociationsForLicenseConfigurationRequest) Send(ctx context.Context) (*ListAssociationsForLicenseConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAssociationsForLicenseConfigurationResponse{
		ListAssociationsForLicenseConfigurationOutput: r.Request.Data.(*types.ListAssociationsForLicenseConfigurationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAssociationsForLicenseConfigurationResponse is the response type for the
// ListAssociationsForLicenseConfiguration API operation.
type ListAssociationsForLicenseConfigurationResponse struct {
	*types.ListAssociationsForLicenseConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAssociationsForLicenseConfiguration request.
func (r *ListAssociationsForLicenseConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
