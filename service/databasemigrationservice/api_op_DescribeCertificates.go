// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribeCertificates = "DescribeCertificates"

// DescribeCertificatesRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Provides a description of the certificate.
//
//    // Example sending a request using DescribeCertificatesRequest.
//    req := client.DescribeCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeCertificates
func (c *Client) DescribeCertificatesRequest(input *types.DescribeCertificatesInput) DescribeCertificatesRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeCertificatesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeCertificatesOutput{})
	return DescribeCertificatesRequest{Request: req, Input: input, Copy: c.DescribeCertificatesRequest}
}

// DescribeCertificatesRequest is the request type for the
// DescribeCertificates API operation.
type DescribeCertificatesRequest struct {
	*aws.Request
	Input *types.DescribeCertificatesInput
	Copy  func(*types.DescribeCertificatesInput) DescribeCertificatesRequest
}

// Send marshals and sends the DescribeCertificates API request.
func (r DescribeCertificatesRequest) Send(ctx context.Context) (*DescribeCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificatesResponse{
		DescribeCertificatesOutput: r.Request.Data.(*types.DescribeCertificatesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeCertificatesRequestPaginator returns a paginator for DescribeCertificates.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeCertificatesRequest(input)
//   p := databasemigrationservice.NewDescribeCertificatesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeCertificatesPaginator(req DescribeCertificatesRequest) DescribeCertificatesPaginator {
	return DescribeCertificatesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeCertificatesInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeCertificatesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeCertificatesPaginator struct {
	aws.Pager
}

func (p *DescribeCertificatesPaginator) CurrentPage() *types.DescribeCertificatesOutput {
	return p.Pager.CurrentPage().(*types.DescribeCertificatesOutput)
}

// DescribeCertificatesResponse is the response type for the
// DescribeCertificates API operation.
type DescribeCertificatesResponse struct {
	*types.DescribeCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificates request.
func (r *DescribeCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
