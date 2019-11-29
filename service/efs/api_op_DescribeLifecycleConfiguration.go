// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
)

const opDescribeLifecycleConfiguration = "DescribeLifecycleConfiguration"

// DescribeLifecycleConfigurationRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Returns the current LifecycleConfiguration object for the specified Amazon
// EFS file system. EFS lifecycle management uses the LifecycleConfiguration
// object to identify which files to move to the EFS Infrequent Access (IA)
// storage class. For a file system without a LifecycleConfiguration object,
// the call returns an empty array in the response.
//
// This operation requires permissions for the elasticfilesystem:DescribeLifecycleConfiguration
// operation.
//
//    // Example sending a request using DescribeLifecycleConfigurationRequest.
//    req := client.DescribeLifecycleConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/DescribeLifecycleConfiguration
func (c *Client) DescribeLifecycleConfigurationRequest(input *types.DescribeLifecycleConfigurationInput) DescribeLifecycleConfigurationRequest {
	op := &aws.Operation{
		Name:       opDescribeLifecycleConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-02-01/file-systems/{FileSystemId}/lifecycle-configuration",
	}

	if input == nil {
		input = &types.DescribeLifecycleConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLifecycleConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeLifecycleConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeLifecycleConfigurationRequest{Request: req, Input: input, Copy: c.DescribeLifecycleConfigurationRequest}
}

// DescribeLifecycleConfigurationRequest is the request type for the
// DescribeLifecycleConfiguration API operation.
type DescribeLifecycleConfigurationRequest struct {
	*aws.Request
	Input *types.DescribeLifecycleConfigurationInput
	Copy  func(*types.DescribeLifecycleConfigurationInput) DescribeLifecycleConfigurationRequest
}

// Send marshals and sends the DescribeLifecycleConfiguration API request.
func (r DescribeLifecycleConfigurationRequest) Send(ctx context.Context) (*DescribeLifecycleConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLifecycleConfigurationResponse{
		DescribeLifecycleConfigurationOutput: r.Request.Data.(*types.DescribeLifecycleConfigurationOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLifecycleConfigurationResponse is the response type for the
// DescribeLifecycleConfiguration API operation.
type DescribeLifecycleConfigurationResponse struct {
	*types.DescribeLifecycleConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLifecycleConfiguration request.
func (r *DescribeLifecycleConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
