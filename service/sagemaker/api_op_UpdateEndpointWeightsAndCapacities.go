// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opUpdateEndpointWeightsAndCapacities = "UpdateEndpointWeightsAndCapacities"

// UpdateEndpointWeightsAndCapacitiesRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Updates variant weight of one or more variants associated with an existing
// endpoint, or capacity of one variant associated with an existing endpoint.
// When it receives the request, Amazon SageMaker sets the endpoint status to
// Updating. After updating the endpoint, it sets the status to InService. To
// check the status of an endpoint, use the DescribeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/API_DescribeEndpoint.html)
// API.
//
//    // Example sending a request using UpdateEndpointWeightsAndCapacitiesRequest.
//    req := client.UpdateEndpointWeightsAndCapacitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/UpdateEndpointWeightsAndCapacities
func (c *Client) UpdateEndpointWeightsAndCapacitiesRequest(input *types.UpdateEndpointWeightsAndCapacitiesInput) UpdateEndpointWeightsAndCapacitiesRequest {
	op := &aws.Operation{
		Name:       opUpdateEndpointWeightsAndCapacities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateEndpointWeightsAndCapacitiesInput{}
	}

	req := c.newRequest(op, input, &types.UpdateEndpointWeightsAndCapacitiesOutput{})
	return UpdateEndpointWeightsAndCapacitiesRequest{Request: req, Input: input, Copy: c.UpdateEndpointWeightsAndCapacitiesRequest}
}

// UpdateEndpointWeightsAndCapacitiesRequest is the request type for the
// UpdateEndpointWeightsAndCapacities API operation.
type UpdateEndpointWeightsAndCapacitiesRequest struct {
	*aws.Request
	Input *types.UpdateEndpointWeightsAndCapacitiesInput
	Copy  func(*types.UpdateEndpointWeightsAndCapacitiesInput) UpdateEndpointWeightsAndCapacitiesRequest
}

// Send marshals and sends the UpdateEndpointWeightsAndCapacities API request.
func (r UpdateEndpointWeightsAndCapacitiesRequest) Send(ctx context.Context) (*UpdateEndpointWeightsAndCapacitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEndpointWeightsAndCapacitiesResponse{
		UpdateEndpointWeightsAndCapacitiesOutput: r.Request.Data.(*types.UpdateEndpointWeightsAndCapacitiesOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEndpointWeightsAndCapacitiesResponse is the response type for the
// UpdateEndpointWeightsAndCapacities API operation.
type UpdateEndpointWeightsAndCapacitiesResponse struct {
	*types.UpdateEndpointWeightsAndCapacitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEndpointWeightsAndCapacities request.
func (r *UpdateEndpointWeightsAndCapacitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
