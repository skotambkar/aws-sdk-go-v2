// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opListModelPackages = "ListModelPackages"

// ListModelPackagesRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists the model packages that have been created.
//
//    // Example sending a request using ListModelPackagesRequest.
//    req := client.ListModelPackagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListModelPackages
func (c *Client) ListModelPackagesRequest(input *types.ListModelPackagesInput) ListModelPackagesRequest {
	op := &aws.Operation{
		Name:       opListModelPackages,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ListModelPackagesInput{}
	}

	req := c.newRequest(op, input, &types.ListModelPackagesOutput{})
	return ListModelPackagesRequest{Request: req, Input: input, Copy: c.ListModelPackagesRequest}
}

// ListModelPackagesRequest is the request type for the
// ListModelPackages API operation.
type ListModelPackagesRequest struct {
	*aws.Request
	Input *types.ListModelPackagesInput
	Copy  func(*types.ListModelPackagesInput) ListModelPackagesRequest
}

// Send marshals and sends the ListModelPackages API request.
func (r ListModelPackagesRequest) Send(ctx context.Context) (*ListModelPackagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListModelPackagesResponse{
		ListModelPackagesOutput: r.Request.Data.(*types.ListModelPackagesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListModelPackagesResponse is the response type for the
// ListModelPackages API operation.
type ListModelPackagesResponse struct {
	*types.ListModelPackagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListModelPackages request.
func (r *ListModelPackagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
