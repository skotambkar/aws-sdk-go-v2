// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opDescribeInstallationMedia = "DescribeInstallationMedia"

// DescribeInstallationMediaRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Describes the available installation media for on-premises, bring your own
// media (BYOM) DB engines, such as Microsoft SQL Server.
//
//    // Example sending a request using DescribeInstallationMediaRequest.
//    req := client.DescribeInstallationMediaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeInstallationMedia
func (c *Client) DescribeInstallationMediaRequest(input *types.DescribeInstallationMediaInput) DescribeInstallationMediaRequest {
	op := &aws.Operation{
		Name:       opDescribeInstallationMedia,
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
		input = &types.DescribeInstallationMediaInput{}
	}

	req := c.newRequest(op, input, &types.DescribeInstallationMediaOutput{})
	return DescribeInstallationMediaRequest{Request: req, Input: input, Copy: c.DescribeInstallationMediaRequest}
}

// DescribeInstallationMediaRequest is the request type for the
// DescribeInstallationMedia API operation.
type DescribeInstallationMediaRequest struct {
	*aws.Request
	Input *types.DescribeInstallationMediaInput
	Copy  func(*types.DescribeInstallationMediaInput) DescribeInstallationMediaRequest
}

// Send marshals and sends the DescribeInstallationMedia API request.
func (r DescribeInstallationMediaRequest) Send(ctx context.Context) (*DescribeInstallationMediaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstallationMediaResponse{
		DescribeInstallationMediaOutput: r.Request.Data.(*types.DescribeInstallationMediaOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeInstallationMediaRequestPaginator returns a paginator for DescribeInstallationMedia.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeInstallationMediaRequest(input)
//   p := rds.NewDescribeInstallationMediaRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeInstallationMediaPaginator(req DescribeInstallationMediaRequest) DescribeInstallationMediaPaginator {
	return DescribeInstallationMediaPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeInstallationMediaInput
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

// DescribeInstallationMediaPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeInstallationMediaPaginator struct {
	aws.Pager
}

func (p *DescribeInstallationMediaPaginator) CurrentPage() *types.DescribeInstallationMediaOutput {
	return p.Pager.CurrentPage().(*types.DescribeInstallationMediaOutput)
}

// DescribeInstallationMediaResponse is the response type for the
// DescribeInstallationMedia API operation.
type DescribeInstallationMediaResponse struct {
	*types.DescribeInstallationMediaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstallationMedia request.
func (r *DescribeInstallationMediaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
