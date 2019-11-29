// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
)

const opUpdateClusterConfiguration = "UpdateClusterConfiguration"

// UpdateClusterConfigurationRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Updates the cluster with the configuration that is specified in the request
// body.
//
//    // Example sending a request using UpdateClusterConfigurationRequest.
//    req := client.UpdateClusterConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateClusterConfiguration
func (c *Client) UpdateClusterConfigurationRequest(input *types.UpdateClusterConfigurationInput) UpdateClusterConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateClusterConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{clusterArn}/configuration",
	}

	if input == nil {
		input = &types.UpdateClusterConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.UpdateClusterConfigurationOutput{})
	return UpdateClusterConfigurationRequest{Request: req, Input: input, Copy: c.UpdateClusterConfigurationRequest}
}

// UpdateClusterConfigurationRequest is the request type for the
// UpdateClusterConfiguration API operation.
type UpdateClusterConfigurationRequest struct {
	*aws.Request
	Input *types.UpdateClusterConfigurationInput
	Copy  func(*types.UpdateClusterConfigurationInput) UpdateClusterConfigurationRequest
}

// Send marshals and sends the UpdateClusterConfiguration API request.
func (r UpdateClusterConfigurationRequest) Send(ctx context.Context) (*UpdateClusterConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterConfigurationResponse{
		UpdateClusterConfigurationOutput: r.Request.Data.(*types.UpdateClusterConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterConfigurationResponse is the response type for the
// UpdateClusterConfiguration API operation.
type UpdateClusterConfigurationResponse struct {
	*types.UpdateClusterConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClusterConfiguration request.
func (r *UpdateClusterConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
