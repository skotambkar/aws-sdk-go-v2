// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opGetBucketAccelerateConfiguration = "GetBucketAccelerateConfiguration"

// GetBucketAccelerateConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the accelerate configuration of a bucket.
//
//    // Example sending a request using GetBucketAccelerateConfigurationRequest.
//    req := client.GetBucketAccelerateConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketAccelerateConfiguration
func (c *Client) GetBucketAccelerateConfigurationRequest(input *types.GetBucketAccelerateConfigurationInput) GetBucketAccelerateConfigurationRequest {
	op := &aws.Operation{
		Name:       opGetBucketAccelerateConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?accelerate",
	}

	if input == nil {
		input = &types.GetBucketAccelerateConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.GetBucketAccelerateConfigurationOutput{})
	return GetBucketAccelerateConfigurationRequest{Request: req, Input: input, Copy: c.GetBucketAccelerateConfigurationRequest}
}

// GetBucketAccelerateConfigurationRequest is the request type for the
// GetBucketAccelerateConfiguration API operation.
type GetBucketAccelerateConfigurationRequest struct {
	*aws.Request
	Input *types.GetBucketAccelerateConfigurationInput
	Copy  func(*types.GetBucketAccelerateConfigurationInput) GetBucketAccelerateConfigurationRequest
}

// Send marshals and sends the GetBucketAccelerateConfiguration API request.
func (r GetBucketAccelerateConfigurationRequest) Send(ctx context.Context) (*GetBucketAccelerateConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketAccelerateConfigurationResponse{
		GetBucketAccelerateConfigurationOutput: r.Request.Data.(*types.GetBucketAccelerateConfigurationOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketAccelerateConfigurationResponse is the response type for the
// GetBucketAccelerateConfiguration API operation.
type GetBucketAccelerateConfigurationResponse struct {
	*types.GetBucketAccelerateConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketAccelerateConfiguration request.
func (r *GetBucketAccelerateConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
