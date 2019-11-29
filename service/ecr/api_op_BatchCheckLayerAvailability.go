// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecr/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

const opBatchCheckLayerAvailability = "BatchCheckLayerAvailability"

// BatchCheckLayerAvailabilityRequest returns a request value for making API operation for
// Amazon EC2 Container Registry.
//
// Check the availability of multiple image layers in a specified registry and
// repository.
//
// This operation is used by the Amazon ECR proxy, and it is not intended for
// general use by customers for pulling and pushing images. In most cases, you
// should use the docker CLI to pull, tag, and push images.
//
//    // Example sending a request using BatchCheckLayerAvailabilityRequest.
//    req := client.BatchCheckLayerAvailabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ecr-2015-09-21/BatchCheckLayerAvailability
func (c *Client) BatchCheckLayerAvailabilityRequest(input *types.BatchCheckLayerAvailabilityInput) BatchCheckLayerAvailabilityRequest {
	op := &aws.Operation{
		Name:       opBatchCheckLayerAvailability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.BatchCheckLayerAvailabilityInput{}
	}

	req := c.newRequest(op, input, &types.BatchCheckLayerAvailabilityOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.BatchCheckLayerAvailabilityMarshaler{Input: input}.GetNamedBuildHandler())

	return BatchCheckLayerAvailabilityRequest{Request: req, Input: input, Copy: c.BatchCheckLayerAvailabilityRequest}
}

// BatchCheckLayerAvailabilityRequest is the request type for the
// BatchCheckLayerAvailability API operation.
type BatchCheckLayerAvailabilityRequest struct {
	*aws.Request
	Input *types.BatchCheckLayerAvailabilityInput
	Copy  func(*types.BatchCheckLayerAvailabilityInput) BatchCheckLayerAvailabilityRequest
}

// Send marshals and sends the BatchCheckLayerAvailability API request.
func (r BatchCheckLayerAvailabilityRequest) Send(ctx context.Context) (*BatchCheckLayerAvailabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchCheckLayerAvailabilityResponse{
		BatchCheckLayerAvailabilityOutput: r.Request.Data.(*types.BatchCheckLayerAvailabilityOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchCheckLayerAvailabilityResponse is the response type for the
// BatchCheckLayerAvailability API operation.
type BatchCheckLayerAvailabilityResponse struct {
	*types.BatchCheckLayerAvailabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchCheckLayerAvailability request.
func (r *BatchCheckLayerAvailabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
