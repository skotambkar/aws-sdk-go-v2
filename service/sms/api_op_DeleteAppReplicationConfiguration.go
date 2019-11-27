// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opDeleteAppReplicationConfiguration = "DeleteAppReplicationConfiguration"

// DeleteAppReplicationConfigurationRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Deletes existing replication configuration for an application.
//
//    // Example sending a request using DeleteAppReplicationConfigurationRequest.
//    req := client.DeleteAppReplicationConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/DeleteAppReplicationConfiguration
func (c *Client) DeleteAppReplicationConfigurationRequest(input *types.DeleteAppReplicationConfigurationInput) DeleteAppReplicationConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteAppReplicationConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteAppReplicationConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAppReplicationConfigurationOutput{})
	return DeleteAppReplicationConfigurationRequest{Request: req, Input: input, Copy: c.DeleteAppReplicationConfigurationRequest}
}

// DeleteAppReplicationConfigurationRequest is the request type for the
// DeleteAppReplicationConfiguration API operation.
type DeleteAppReplicationConfigurationRequest struct {
	*aws.Request
	Input *types.DeleteAppReplicationConfigurationInput
	Copy  func(*types.DeleteAppReplicationConfigurationInput) DeleteAppReplicationConfigurationRequest
}

// Send marshals and sends the DeleteAppReplicationConfiguration API request.
func (r DeleteAppReplicationConfigurationRequest) Send(ctx context.Context) (*DeleteAppReplicationConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAppReplicationConfigurationResponse{
		DeleteAppReplicationConfigurationOutput: r.Request.Data.(*types.DeleteAppReplicationConfigurationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAppReplicationConfigurationResponse is the response type for the
// DeleteAppReplicationConfiguration API operation.
type DeleteAppReplicationConfigurationResponse struct {
	*types.DeleteAppReplicationConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteAppReplicationConfiguration request.
func (r *DeleteAppReplicationConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
