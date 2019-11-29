// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
)

const opListSamples = "ListSamples"

// ListSamplesRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about samples, given an AWS Device Farm job ARN.
//
//    // Example sending a request using ListSamplesRequest.
//    req := client.ListSamplesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/ListSamples
func (c *Client) ListSamplesRequest(input *types.ListSamplesInput) ListSamplesRequest {
	op := &aws.Operation{
		Name:       opListSamples,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListSamplesInput{}
	}

	req := c.newRequest(op, input, &types.ListSamplesOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListSamplesMarshaler{Input: input}.GetNamedBuildHandler())

	return ListSamplesRequest{Request: req, Input: input, Copy: c.ListSamplesRequest}
}

// ListSamplesRequest is the request type for the
// ListSamples API operation.
type ListSamplesRequest struct {
	*aws.Request
	Input *types.ListSamplesInput
	Copy  func(*types.ListSamplesInput) ListSamplesRequest
}

// Send marshals and sends the ListSamples API request.
func (r ListSamplesRequest) Send(ctx context.Context) (*ListSamplesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSamplesResponse{
		ListSamplesOutput: r.Request.Data.(*types.ListSamplesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSamplesRequestPaginator returns a paginator for ListSamples.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSamplesRequest(input)
//   p := devicefarm.NewListSamplesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSamplesPaginator(req ListSamplesRequest) ListSamplesPaginator {
	return ListSamplesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListSamplesInput
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

// ListSamplesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSamplesPaginator struct {
	aws.Pager
}

func (p *ListSamplesPaginator) CurrentPage() *types.ListSamplesOutput {
	return p.Pager.CurrentPage().(*types.ListSamplesOutput)
}

// ListSamplesResponse is the response type for the
// ListSamples API operation.
type ListSamplesResponse struct {
	*types.ListSamplesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSamples request.
func (r *ListSamplesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
