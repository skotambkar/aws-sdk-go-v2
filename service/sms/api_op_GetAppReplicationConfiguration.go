// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opGetAppReplicationConfiguration = "GetAppReplicationConfiguration"

// GetAppReplicationConfigurationRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Retrieves an application replication configuration associatd with an application.
//
//    // Example sending a request using GetAppReplicationConfigurationRequest.
//    req := client.GetAppReplicationConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/GetAppReplicationConfiguration
func (c *Client) GetAppReplicationConfigurationRequest(input *types.GetAppReplicationConfigurationInput) GetAppReplicationConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetAppReplicationConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetAppReplicationConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.GetAppReplicationConfigurationOutput{})
	return GetAppReplicationConfigurationRequest{Request: req, Input: input, Copy: c.GetAppReplicationConfigurationRequest}
}

// GetAppReplicationConfigurationRequest is the request type for the
// GetAppReplicationConfiguration API operation.
type GetAppReplicationConfigurationRequest struct {
	*aws.Request
	Input *types.GetAppReplicationConfigurationInput
	Copy  func(*types.GetAppReplicationConfigurationInput) GetAppReplicationConfigurationRequest
}

// Send marshals and sends the GetAppReplicationConfiguration API request.
func (r GetAppReplicationConfigurationRequest) Send(ctx context.Context) (*GetAppReplicationConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAppReplicationConfigurationResponse{
		GetAppReplicationConfigurationOutput: r.Request.Data.(*types.GetAppReplicationConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAppReplicationConfigurationResponse is the response type for the
// GetAppReplicationConfiguration API operation.
type GetAppReplicationConfigurationResponse struct {
	*types.GetAppReplicationConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAppReplicationConfiguration request.
func (r *GetAppReplicationConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
