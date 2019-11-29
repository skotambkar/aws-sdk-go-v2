// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
)

const opListInstances = "ListInstances"

// ListInstancesRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides information for all active EC2 instances and EC2 instances terminated
// in the last 30 days, up to a maximum of 2,000. EC2 instances in any of the
// following states are considered active: AWAITING_FULFILLMENT, PROVISIONING,
// BOOTSTRAPPING, RUNNING.
//
//    // Example sending a request using ListInstancesRequest.
//    req := client.ListInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListInstances
func (c *Client) ListInstancesRequest(input *types.ListInstancesInput) ListInstancesRequest {
	op := &aws.Operation{
		Name:       opListInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListInstancesInput{}
	}

	req := c.newRequest(op, input, &types.ListInstancesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListInstancesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListInstancesRequest{Request: req, Input: input, Copy: c.ListInstancesRequest}
}

// ListInstancesRequest is the request type for the
// ListInstances API operation.
type ListInstancesRequest struct {
	*aws.Request
	Input *types.ListInstancesInput
	Copy  func(*types.ListInstancesInput) ListInstancesRequest
}

// Send marshals and sends the ListInstances API request.
func (r ListInstancesRequest) Send(ctx context.Context) (*ListInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListInstancesResponse{
		ListInstancesOutput: r.Request.Data.(*types.ListInstancesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListInstancesRequestPaginator returns a paginator for ListInstances.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListInstancesRequest(input)
//   p := emr.NewListInstancesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListInstancesPaginator(req ListInstancesRequest) ListInstancesPaginator {
	return ListInstancesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListInstancesInput
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

// ListInstancesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListInstancesPaginator struct {
	aws.Pager
}

func (p *ListInstancesPaginator) CurrentPage() *types.ListInstancesOutput {
	return p.Pager.CurrentPage().(*types.ListInstancesOutput)
}

// ListInstancesResponse is the response type for the
// ListInstances API operation.
type ListInstancesResponse struct {
	*types.ListInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListInstances request.
func (r *ListInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
