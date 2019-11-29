// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opGetCrawlerMetrics = "GetCrawlerMetrics"

// GetCrawlerMetricsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves metrics about specified crawlers.
//
//    // Example sending a request using GetCrawlerMetricsRequest.
//    req := client.GetCrawlerMetricsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCrawlerMetrics
func (c *Client) GetCrawlerMetricsRequest(input *types.GetCrawlerMetricsInput) GetCrawlerMetricsRequest {
	op := &aws.Operation{
		Name:       opGetCrawlerMetrics,
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
		input = &types.GetCrawlerMetricsInput{}
	}

	req := c.newRequest(op, input, &types.GetCrawlerMetricsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetCrawlerMetricsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetCrawlerMetricsRequest{Request: req, Input: input, Copy: c.GetCrawlerMetricsRequest}
}

// GetCrawlerMetricsRequest is the request type for the
// GetCrawlerMetrics API operation.
type GetCrawlerMetricsRequest struct {
	*aws.Request
	Input *types.GetCrawlerMetricsInput
	Copy  func(*types.GetCrawlerMetricsInput) GetCrawlerMetricsRequest
}

// Send marshals and sends the GetCrawlerMetrics API request.
func (r GetCrawlerMetricsRequest) Send(ctx context.Context) (*GetCrawlerMetricsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCrawlerMetricsResponse{
		GetCrawlerMetricsOutput: r.Request.Data.(*types.GetCrawlerMetricsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetCrawlerMetricsRequestPaginator returns a paginator for GetCrawlerMetrics.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetCrawlerMetricsRequest(input)
//   p := glue.NewGetCrawlerMetricsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetCrawlerMetricsPaginator(req GetCrawlerMetricsRequest) GetCrawlerMetricsPaginator {
	return GetCrawlerMetricsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetCrawlerMetricsInput
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

// GetCrawlerMetricsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetCrawlerMetricsPaginator struct {
	aws.Pager
}

func (p *GetCrawlerMetricsPaginator) CurrentPage() *types.GetCrawlerMetricsOutput {
	return p.Pager.CurrentPage().(*types.GetCrawlerMetricsOutput)
}

// GetCrawlerMetricsResponse is the response type for the
// GetCrawlerMetrics API operation.
type GetCrawlerMetricsResponse struct {
	*types.GetCrawlerMetricsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCrawlerMetrics request.
func (r *GetCrawlerMetricsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
