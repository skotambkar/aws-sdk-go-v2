// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
)

const opListCACertificates = "ListCACertificates"

// ListCACertificatesRequest returns a request value for making API operation for
// AWS IoT.
//
// Lists the CA certificates registered for your AWS account.
//
// The results are paginated with a default page size of 25. You can use the
// returned marker to retrieve additional results.
//
//    // Example sending a request using ListCACertificatesRequest.
//    req := client.ListCACertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListCACertificatesRequest(input *types.ListCACertificatesInput) ListCACertificatesRequest {
	op := &aws.Operation{
		Name:       opListCACertificates,
		HTTPMethod: "GET",
		HTTPPath:   "/cacertificates",
	}

	if input == nil {
		input = &types.ListCACertificatesInput{}
	}

	req := c.newRequest(op, input, &types.ListCACertificatesOutput{})
	return ListCACertificatesRequest{Request: req, Input: input, Copy: c.ListCACertificatesRequest}
}

// ListCACertificatesRequest is the request type for the
// ListCACertificates API operation.
type ListCACertificatesRequest struct {
	*aws.Request
	Input *types.ListCACertificatesInput
	Copy  func(*types.ListCACertificatesInput) ListCACertificatesRequest
}

// Send marshals and sends the ListCACertificates API request.
func (r ListCACertificatesRequest) Send(ctx context.Context) (*ListCACertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCACertificatesResponse{
		ListCACertificatesOutput: r.Request.Data.(*types.ListCACertificatesOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListCACertificatesResponse is the response type for the
// ListCACertificates API operation.
type ListCACertificatesResponse struct {
	*types.ListCACertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCACertificates request.
func (r *ListCACertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
