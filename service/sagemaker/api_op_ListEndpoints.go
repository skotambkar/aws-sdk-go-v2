// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opListEndpoints = "ListEndpoints"

// ListEndpointsRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Lists endpoints.
//
//    // Example sending a request using ListEndpointsRequest.
//    req := client.ListEndpointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/ListEndpoints
func (c *Client) ListEndpointsRequest(input *types.ListEndpointsInput) ListEndpointsRequest {
	op := &aws.Operation{
		Name:       opListEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListEndpointsInput{}
	}

	req := c.newRequest(op, input, &types.ListEndpointsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListEndpointsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListEndpointsRequest{Request: req, Input: input, Copy: c.ListEndpointsRequest}
}

// ListEndpointsRequest is the request type for the
// ListEndpoints API operation.
type ListEndpointsRequest struct {
	*aws.Request
	Input *types.ListEndpointsInput
	Copy  func(*types.ListEndpointsInput) ListEndpointsRequest
}

// Send marshals and sends the ListEndpoints API request.
func (r ListEndpointsRequest) Send(ctx context.Context) (*ListEndpointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListEndpointsResponse{
		ListEndpointsOutput: r.Request.Data.(*types.ListEndpointsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListEndpointsRequestPaginator returns a paginator for ListEndpoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListEndpointsRequest(input)
//   p := sagemaker.NewListEndpointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListEndpointsPaginator(req ListEndpointsRequest) ListEndpointsPaginator {
	return ListEndpointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListEndpointsInput
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

// ListEndpointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListEndpointsPaginator struct {
	aws.Pager
}

func (p *ListEndpointsPaginator) CurrentPage() *types.ListEndpointsOutput {
	return p.Pager.CurrentPage().(*types.ListEndpointsOutput)
}

// ListEndpointsResponse is the response type for the
// ListEndpoints API operation.
type ListEndpointsResponse struct {
	*types.ListEndpointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListEndpoints request.
func (r *ListEndpointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
