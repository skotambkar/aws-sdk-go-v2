// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opPutRemediationConfigurations = "PutRemediationConfigurations"

// PutRemediationConfigurationsRequest returns a request value for making API operation for
// AWS Config.
//
// Adds or updates the remediation configuration with a specific AWS Config
// rule with the selected target or action. The API creates the RemediationConfiguration
// object for the AWS Config rule. The AWS Config rule must already exist for
// you to add a remediation configuration. The target (SSM document) must exist
// and have permissions to use the target.
//
//    // Example sending a request using PutRemediationConfigurationsRequest.
//    req := client.PutRemediationConfigurationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutRemediationConfigurations
func (c *Client) PutRemediationConfigurationsRequest(input *types.PutRemediationConfigurationsInput) PutRemediationConfigurationsRequest {
	op := &aws.Operation{
		Name:       opPutRemediationConfigurations,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutRemediationConfigurationsInput{}
	}

	req := c.newRequest(op, input, &types.PutRemediationConfigurationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutRemediationConfigurationsMarshaler{Input: input}.GetNamedBuildHandler())

	return PutRemediationConfigurationsRequest{Request: req, Input: input, Copy: c.PutRemediationConfigurationsRequest}
}

// PutRemediationConfigurationsRequest is the request type for the
// PutRemediationConfigurations API operation.
type PutRemediationConfigurationsRequest struct {
	*aws.Request
	Input *types.PutRemediationConfigurationsInput
	Copy  func(*types.PutRemediationConfigurationsInput) PutRemediationConfigurationsRequest
}

// Send marshals and sends the PutRemediationConfigurations API request.
func (r PutRemediationConfigurationsRequest) Send(ctx context.Context) (*PutRemediationConfigurationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutRemediationConfigurationsResponse{
		PutRemediationConfigurationsOutput: r.Request.Data.(*types.PutRemediationConfigurationsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutRemediationConfigurationsResponse is the response type for the
// PutRemediationConfigurations API operation.
type PutRemediationConfigurationsResponse struct {
	*types.PutRemediationConfigurationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutRemediationConfigurations request.
func (r *PutRemediationConfigurationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
