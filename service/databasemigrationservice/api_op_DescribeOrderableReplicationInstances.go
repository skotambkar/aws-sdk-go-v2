// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribeOrderableReplicationInstances = "DescribeOrderableReplicationInstances"

// DescribeOrderableReplicationInstancesRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns information about the replication instance types that can be created
// in the specified region.
//
//    // Example sending a request using DescribeOrderableReplicationInstancesRequest.
//    req := client.DescribeOrderableReplicationInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeOrderableReplicationInstances
func (c *Client) DescribeOrderableReplicationInstancesRequest(input *types.DescribeOrderableReplicationInstancesInput) DescribeOrderableReplicationInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeOrderableReplicationInstances,
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
		input = &types.DescribeOrderableReplicationInstancesInput{}
	}

	req := c.newRequest(op, input, &types.DescribeOrderableReplicationInstancesOutput{})
	return DescribeOrderableReplicationInstancesRequest{Request: req, Input: input, Copy: c.DescribeOrderableReplicationInstancesRequest}
}

// DescribeOrderableReplicationInstancesRequest is the request type for the
// DescribeOrderableReplicationInstances API operation.
type DescribeOrderableReplicationInstancesRequest struct {
	*aws.Request
	Input *types.DescribeOrderableReplicationInstancesInput
	Copy  func(*types.DescribeOrderableReplicationInstancesInput) DescribeOrderableReplicationInstancesRequest
}

// Send marshals and sends the DescribeOrderableReplicationInstances API request.
func (r DescribeOrderableReplicationInstancesRequest) Send(ctx context.Context) (*DescribeOrderableReplicationInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOrderableReplicationInstancesResponse{
		DescribeOrderableReplicationInstancesOutput: r.Request.Data.(*types.DescribeOrderableReplicationInstancesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeOrderableReplicationInstancesRequestPaginator returns a paginator for DescribeOrderableReplicationInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeOrderableReplicationInstancesRequest(input)
//   p := databasemigrationservice.NewDescribeOrderableReplicationInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeOrderableReplicationInstancesPaginator(req DescribeOrderableReplicationInstancesRequest) DescribeOrderableReplicationInstancesPaginator {
	return DescribeOrderableReplicationInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeOrderableReplicationInstancesInput
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

// DescribeOrderableReplicationInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeOrderableReplicationInstancesPaginator struct {
	aws.Pager
}

func (p *DescribeOrderableReplicationInstancesPaginator) CurrentPage() *types.DescribeOrderableReplicationInstancesOutput {
	return p.Pager.CurrentPage().(*types.DescribeOrderableReplicationInstancesOutput)
}

// DescribeOrderableReplicationInstancesResponse is the response type for the
// DescribeOrderableReplicationInstances API operation.
type DescribeOrderableReplicationInstancesResponse struct {
	*types.DescribeOrderableReplicationInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOrderableReplicationInstances request.
func (r *DescribeOrderableReplicationInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
