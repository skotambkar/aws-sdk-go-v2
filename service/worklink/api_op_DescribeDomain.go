// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
)

const opDescribeDomain = "DescribeDomain"

// DescribeDomainRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Provides information about the domain.
//
//    // Example sending a request using DescribeDomainRequest.
//    req := client.DescribeDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/DescribeDomain
func (c *Client) DescribeDomainRequest(input *types.DescribeDomainInput) DescribeDomainRequest {
	op := &aws.Operation{
		Name:       opDescribeDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/describeDomain",
	}

	if input == nil {
		input = &types.DescribeDomainInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDomainOutput{})
	return DescribeDomainRequest{Request: req, Input: input, Copy: c.DescribeDomainRequest}
}

// DescribeDomainRequest is the request type for the
// DescribeDomain API operation.
type DescribeDomainRequest struct {
	*aws.Request
	Input *types.DescribeDomainInput
	Copy  func(*types.DescribeDomainInput) DescribeDomainRequest
}

// Send marshals and sends the DescribeDomain API request.
func (r DescribeDomainRequest) Send(ctx context.Context) (*DescribeDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDomainResponse{
		DescribeDomainOutput: r.Request.Data.(*types.DescribeDomainOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDomainResponse is the response type for the
// DescribeDomain API operation.
type DescribeDomainResponse struct {
	*types.DescribeDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDomain request.
func (r *DescribeDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
