// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opDeleteLoggingConfiguration = "DeleteLoggingConfiguration"

// DeleteLoggingConfigurationRequest returns a request value for making API operation for
// AWS WAF.
//
// Permanently deletes the LoggingConfiguration from the specified web ACL.
//
//    // Example sending a request using DeleteLoggingConfigurationRequest.
//    req := client.DeleteLoggingConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/DeleteLoggingConfiguration
func (c *Client) DeleteLoggingConfigurationRequest(input *types.DeleteLoggingConfigurationInput) DeleteLoggingConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteLoggingConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteLoggingConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteLoggingConfigurationOutput{})
	return DeleteLoggingConfigurationRequest{Request: req, Input: input, Copy: c.DeleteLoggingConfigurationRequest}
}

// DeleteLoggingConfigurationRequest is the request type for the
// DeleteLoggingConfiguration API operation.
type DeleteLoggingConfigurationRequest struct {
	*aws.Request
	Input *types.DeleteLoggingConfigurationInput
	Copy  func(*types.DeleteLoggingConfigurationInput) DeleteLoggingConfigurationRequest
}

// Send marshals and sends the DeleteLoggingConfiguration API request.
func (r DeleteLoggingConfigurationRequest) Send(ctx context.Context) (*DeleteLoggingConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLoggingConfigurationResponse{
		DeleteLoggingConfigurationOutput: r.Request.Data.(*types.DeleteLoggingConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLoggingConfigurationResponse is the response type for the
// DeleteLoggingConfiguration API operation.
type DeleteLoggingConfigurationResponse struct {
	*types.DeleteLoggingConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLoggingConfiguration request.
func (r *DeleteLoggingConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
