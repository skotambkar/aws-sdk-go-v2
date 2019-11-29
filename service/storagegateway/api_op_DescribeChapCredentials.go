// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDescribeChapCredentials = "DescribeChapCredentials"

// DescribeChapCredentialsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns an array of Challenge-Handshake Authentication Protocol (CHAP) credentials
// information for a specified iSCSI target, one for each target-initiator pair.
// This operation is supported in the volume and tape gateway types.
//
//    // Example sending a request using DescribeChapCredentialsRequest.
//    req := client.DescribeChapCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeChapCredentials
func (c *Client) DescribeChapCredentialsRequest(input *types.DescribeChapCredentialsInput) DescribeChapCredentialsRequest {
	op := &aws.Operation{
		Name:       opDescribeChapCredentials,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeChapCredentialsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeChapCredentialsOutput{})
	return DescribeChapCredentialsRequest{Request: req, Input: input, Copy: c.DescribeChapCredentialsRequest}
}

// DescribeChapCredentialsRequest is the request type for the
// DescribeChapCredentials API operation.
type DescribeChapCredentialsRequest struct {
	*aws.Request
	Input *types.DescribeChapCredentialsInput
	Copy  func(*types.DescribeChapCredentialsInput) DescribeChapCredentialsRequest
}

// Send marshals and sends the DescribeChapCredentials API request.
func (r DescribeChapCredentialsRequest) Send(ctx context.Context) (*DescribeChapCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeChapCredentialsResponse{
		DescribeChapCredentialsOutput: r.Request.Data.(*types.DescribeChapCredentialsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeChapCredentialsResponse is the response type for the
// DescribeChapCredentials API operation.
type DescribeChapCredentialsResponse struct {
	*types.DescribeChapCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeChapCredentials request.
func (r *DescribeChapCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
