// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

const opPutObjectLockConfiguration = "PutObjectLockConfiguration"

// PutObjectLockConfigurationRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Places an object lock configuration on the specified bucket. The rule specified
// in the object lock configuration will be applied by default to every new
// object placed in the specified bucket.
//
//    // Example sending a request using PutObjectLockConfigurationRequest.
//    req := client.PutObjectLockConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/PutObjectLockConfiguration
func (c *Client) PutObjectLockConfigurationRequest(input *types.PutObjectLockConfigurationInput) PutObjectLockConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutObjectLockConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?object-lock",
	}

	if input == nil {
		input = &types.PutObjectLockConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.PutObjectLockConfigurationOutput{})
	return PutObjectLockConfigurationRequest{Request: req, Input: input, Copy: c.PutObjectLockConfigurationRequest}
}

// PutObjectLockConfigurationRequest is the request type for the
// PutObjectLockConfiguration API operation.
type PutObjectLockConfigurationRequest struct {
	*aws.Request
	Input *types.PutObjectLockConfigurationInput
	Copy  func(*types.PutObjectLockConfigurationInput) PutObjectLockConfigurationRequest
}

// Send marshals and sends the PutObjectLockConfiguration API request.
func (r PutObjectLockConfigurationRequest) Send(ctx context.Context) (*PutObjectLockConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutObjectLockConfigurationResponse{
		PutObjectLockConfigurationOutput: r.Request.Data.(*types.PutObjectLockConfigurationOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutObjectLockConfigurationResponse is the response type for the
// PutObjectLockConfiguration API operation.
type PutObjectLockConfigurationResponse struct {
	*types.PutObjectLockConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutObjectLockConfiguration request.
func (r *PutObjectLockConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
