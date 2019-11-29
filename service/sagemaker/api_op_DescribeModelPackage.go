// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opDescribeModelPackage = "DescribeModelPackage"

// DescribeModelPackageRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Returns a description of the specified model package, which is used to create
// Amazon SageMaker models or list them on AWS Marketplace.
//
// To create models in Amazon SageMaker, buyers can subscribe to model packages
// listed on AWS Marketplace.
//
//    // Example sending a request using DescribeModelPackageRequest.
//    req := client.DescribeModelPackageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeModelPackage
func (c *Client) DescribeModelPackageRequest(input *types.DescribeModelPackageInput) DescribeModelPackageRequest {
	op := &aws.Operation{
		Name:       opDescribeModelPackage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeModelPackageInput{}
	}

	req := c.newRequest(op, input, &types.DescribeModelPackageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeModelPackageMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeModelPackageRequest{Request: req, Input: input, Copy: c.DescribeModelPackageRequest}
}

// DescribeModelPackageRequest is the request type for the
// DescribeModelPackage API operation.
type DescribeModelPackageRequest struct {
	*aws.Request
	Input *types.DescribeModelPackageInput
	Copy  func(*types.DescribeModelPackageInput) DescribeModelPackageRequest
}

// Send marshals and sends the DescribeModelPackage API request.
func (r DescribeModelPackageRequest) Send(ctx context.Context) (*DescribeModelPackageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeModelPackageResponse{
		DescribeModelPackageOutput: r.Request.Data.(*types.DescribeModelPackageOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeModelPackageResponse is the response type for the
// DescribeModelPackage API operation.
type DescribeModelPackageResponse struct {
	*types.DescribeModelPackageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeModelPackage request.
func (r *DescribeModelPackageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
