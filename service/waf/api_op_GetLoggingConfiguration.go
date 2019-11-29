// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opGetLoggingConfiguration = "GetLoggingConfiguration"

// GetLoggingConfigurationRequest returns a request value for making API operation for
// AWS WAF.
//
// Returns the LoggingConfiguration for the specified web ACL.
//
//    // Example sending a request using GetLoggingConfigurationRequest.
//    req := client.GetLoggingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/GetLoggingConfiguration
func (c *Client) GetLoggingConfigurationRequest(input *types.GetLoggingConfigurationInput) GetLoggingConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetLoggingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetLoggingConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.GetLoggingConfigurationOutput{})
	return GetLoggingConfigurationRequest{Request: req, Input: input, Copy: c.GetLoggingConfigurationRequest}
}

// GetLoggingConfigurationRequest is the request type for the
// GetLoggingConfiguration API operation.
type GetLoggingConfigurationRequest struct {
	*aws.Request
	Input *types.GetLoggingConfigurationInput
	Copy  func(*types.GetLoggingConfigurationInput) GetLoggingConfigurationRequest
}

// Send marshals and sends the GetLoggingConfiguration API request.
func (r GetLoggingConfigurationRequest) Send(ctx context.Context) (*GetLoggingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLoggingConfigurationResponse{
		GetLoggingConfigurationOutput: r.Request.Data.(*types.GetLoggingConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetLoggingConfigurationResponse is the response type for the
// GetLoggingConfiguration API operation.
type GetLoggingConfigurationResponse struct {
	*types.GetLoggingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLoggingConfiguration request.
func (r *GetLoggingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
