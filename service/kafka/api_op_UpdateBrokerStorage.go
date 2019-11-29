// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
)

const opUpdateBrokerStorage = "UpdateBrokerStorage"

// UpdateBrokerStorageRequest returns a request value for making API operation for
// Managed Streaming for Kafka.
//
// Updates the EBS storage associated with MSK brokers.
//
//    // Example sending a request using UpdateBrokerStorageRequest.
//    req := client.UpdateBrokerStorageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kafka-2018-11-14/UpdateBrokerStorage
func (c *Client) UpdateBrokerStorageRequest(input *types.UpdateBrokerStorageInput) UpdateBrokerStorageRequest {
	op := &aws.Operation{
		Name:       opUpdateBrokerStorage,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{clusterArn}/nodes/storage",
	}

	if input == nil {
		input = &types.UpdateBrokerStorageInput{}
	}

	req := c.newRequest(op, input, &types.UpdateBrokerStorageOutput{})
	return UpdateBrokerStorageRequest{Request: req, Input: input, Copy: c.UpdateBrokerStorageRequest}
}

// UpdateBrokerStorageRequest is the request type for the
// UpdateBrokerStorage API operation.
type UpdateBrokerStorageRequest struct {
	*aws.Request
	Input *types.UpdateBrokerStorageInput
	Copy  func(*types.UpdateBrokerStorageInput) UpdateBrokerStorageRequest
}

// Send marshals and sends the UpdateBrokerStorage API request.
func (r UpdateBrokerStorageRequest) Send(ctx context.Context) (*UpdateBrokerStorageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBrokerStorageResponse{
		UpdateBrokerStorageOutput: r.Request.Data.(*types.UpdateBrokerStorageOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBrokerStorageResponse is the response type for the
// UpdateBrokerStorage API operation.
type UpdateBrokerStorageResponse struct {
	*types.UpdateBrokerStorageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBrokerStorage request.
func (r *UpdateBrokerStorageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
