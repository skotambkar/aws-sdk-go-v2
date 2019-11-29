// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package robomaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
)

const opListFleets = "ListFleets"

// ListFleetsRequest returns a request value for making API operation for
// AWS RoboMaker.
//
// Returns a list of fleets. You can optionally provide filters to retrieve
// specific fleets.
//
//    // Example sending a request using ListFleetsRequest.
//    req := client.ListFleetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/robomaker-2018-06-29/ListFleets
func (c *Client) ListFleetsRequest(input *types.ListFleetsInput) ListFleetsRequest {
	op := &aws.Operation{
		Name:       opListFleets,
		HTTPMethod: "POST",
		HTTPPath:   "/listFleets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListFleetsInput{}
	}

	req := c.newRequest(op, input, &types.ListFleetsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListFleetsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListFleetsRequest{Request: req, Input: input, Copy: c.ListFleetsRequest}
}

// ListFleetsRequest is the request type for the
// ListFleets API operation.
type ListFleetsRequest struct {
	*aws.Request
	Input *types.ListFleetsInput
	Copy  func(*types.ListFleetsInput) ListFleetsRequest
}

// Send marshals and sends the ListFleets API request.
func (r ListFleetsRequest) Send(ctx context.Context) (*ListFleetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFleetsResponse{
		ListFleetsOutput: r.Request.Data.(*types.ListFleetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFleetsRequestPaginator returns a paginator for ListFleets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFleetsRequest(input)
//   p := robomaker.NewListFleetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFleetsPaginator(req ListFleetsRequest) ListFleetsPaginator {
	return ListFleetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListFleetsInput
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

// ListFleetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFleetsPaginator struct {
	aws.Pager
}

func (p *ListFleetsPaginator) CurrentPage() *types.ListFleetsOutput {
	return p.Pager.CurrentPage().(*types.ListFleetsOutput)
}

// ListFleetsResponse is the response type for the
// ListFleets API operation.
type ListFleetsResponse struct {
	*types.ListFleetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFleets request.
func (r *ListFleetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
