// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package licensemanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
)

const opGetLicenseConfiguration = "GetLicenseConfiguration"

// GetLicenseConfigurationRequest returns a request value for making API operation for
// AWS License Manager.
//
// Returns a detailed description of a license configuration.
//
//    // Example sending a request using GetLicenseConfigurationRequest.
//    req := client.GetLicenseConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/license-manager-2018-08-01/GetLicenseConfiguration
func (c *Client) GetLicenseConfigurationRequest(input *types.GetLicenseConfigurationInput) GetLicenseConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetLicenseConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetLicenseConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.GetLicenseConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetLicenseConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLicenseConfigurationRequest{Request: req, Input: input, Copy: c.GetLicenseConfigurationRequest}
}

// GetLicenseConfigurationRequest is the request type for the
// GetLicenseConfiguration API operation.
type GetLicenseConfigurationRequest struct {
	*aws.Request
	Input *types.GetLicenseConfigurationInput
	Copy  func(*types.GetLicenseConfigurationInput) GetLicenseConfigurationRequest
}

// Send marshals and sends the GetLicenseConfiguration API request.
func (r GetLicenseConfigurationRequest) Send(ctx context.Context) (*GetLicenseConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLicenseConfigurationResponse{
		GetLicenseConfigurationOutput: r.Request.Data.(*types.GetLicenseConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLicenseConfigurationResponse is the response type for the
// GetLicenseConfiguration API operation.
type GetLicenseConfigurationResponse struct {
	*types.GetLicenseConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLicenseConfiguration request.
func (r *GetLicenseConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
