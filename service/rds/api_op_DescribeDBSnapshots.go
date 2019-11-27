// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opDescribeDBSnapshots = "DescribeDBSnapshots"

// DescribeDBSnapshotsRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Returns information about DB snapshots. This API action supports pagination.
//
//    // Example sending a request using DescribeDBSnapshotsRequest.
//    req := client.DescribeDBSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DescribeDBSnapshots
func (c *Client) DescribeDBSnapshotsRequest(input *types.DescribeDBSnapshotsInput) DescribeDBSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDescribeDBSnapshots,
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
		input = &types.DescribeDBSnapshotsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDBSnapshotsOutput{})
	return DescribeDBSnapshotsRequest{Request: req, Input: input, Copy: c.DescribeDBSnapshotsRequest}
}

// DescribeDBSnapshotsRequest is the request type for the
// DescribeDBSnapshots API operation.
type DescribeDBSnapshotsRequest struct {
	*aws.Request
	Input *types.DescribeDBSnapshotsInput
	Copy  func(*types.DescribeDBSnapshotsInput) DescribeDBSnapshotsRequest
}

// Send marshals and sends the DescribeDBSnapshots API request.
func (r DescribeDBSnapshotsRequest) Send(ctx context.Context) (*DescribeDBSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDBSnapshotsResponse{
		DescribeDBSnapshotsOutput: r.Request.Data.(*types.DescribeDBSnapshotsOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeDBSnapshotsRequestPaginator returns a paginator for DescribeDBSnapshots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeDBSnapshotsRequest(input)
//   p := rds.NewDescribeDBSnapshotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeDBSnapshotsPaginator(req DescribeDBSnapshotsRequest) DescribeDBSnapshotsPaginator {
	return DescribeDBSnapshotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeDBSnapshotsInput
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

// DescribeDBSnapshotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeDBSnapshotsPaginator struct {
	aws.Pager
}

func (p *DescribeDBSnapshotsPaginator) CurrentPage() *types.DescribeDBSnapshotsOutput {
	return p.Pager.CurrentPage().(*types.DescribeDBSnapshotsOutput)
}

// DescribeDBSnapshotsResponse is the response type for the
// DescribeDBSnapshots API operation.
type DescribeDBSnapshotsResponse struct {
	*types.DescribeDBSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDBSnapshots request.
func (r *DescribeDBSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
