// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribeEndpoints = "DescribeEndpoints"

// DescribeEndpointsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns information about the endpoints for your account in the current region.
//
//    // Example sending a request using DescribeEndpointsRequest.
//    req := client.DescribeEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeEndpoints
func (c *Client) DescribeEndpointsRequest(input *types.DescribeEndpointsInput) DescribeEndpointsRequest {
	op := &aws.Operation{
		Name:       opDescribeEndpoints,
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
		input = &types.DescribeEndpointsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEndpointsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeEndpointsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeEndpointsRequest{Request: req, Input: input, Copy: c.DescribeEndpointsRequest}
}

// DescribeEndpointsRequest is the request type for the
// DescribeEndpoints API operation.
type DescribeEndpointsRequest struct {
	*aws.Request
	Input *types.DescribeEndpointsInput
	Copy  func(*types.DescribeEndpointsInput) DescribeEndpointsRequest
}

// Send marshals and sends the DescribeEndpoints API request.
func (r DescribeEndpointsRequest) Send(ctx context.Context) (*DescribeEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEndpointsResponse{
		DescribeEndpointsOutput: r.Request.Data.(*types.DescribeEndpointsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeEndpointsRequestPaginator returns a paginator for DescribeEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeEndpointsRequest(input)
//   p := databasemigrationservice.NewDescribeEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeEndpointsPaginator(req DescribeEndpointsRequest) DescribeEndpointsPaginator {
	return DescribeEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeEndpointsInput
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

// DescribeEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeEndpointsPaginator struct {
	aws.Pager
}

func (p *DescribeEndpointsPaginator) CurrentPage() *types.DescribeEndpointsOutput {
	return p.Pager.CurrentPage().(*types.DescribeEndpointsOutput)
}

// DescribeEndpointsResponse is the response type for the
// DescribeEndpoints API operation.
type DescribeEndpointsResponse struct {
	*types.DescribeEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEndpoints request.
func (r *DescribeEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
