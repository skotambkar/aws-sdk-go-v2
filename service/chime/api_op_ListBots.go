// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
)

const opListBots = "ListBots"

// ListBotsRequest returns a request value for making API operation for
// Amazon Chime.
//
// Lists the bots associated with the administrator's Amazon Chime Enterprise
// account ID.
//
//    // Example sending a request using ListBotsRequest.
//    req := client.ListBotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/ListBots
func (c *Client) ListBotsRequest(input *types.ListBotsInput) ListBotsRequest {
	op := &aws.Operation{
		Name:       opListBots,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{accountId}/bots",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListBotsInput{}
	}

	req := c.newRequest(op, input, &types.ListBotsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.ListBotsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListBotsRequest{Request: req, Input: input, Copy: c.ListBotsRequest}
}

// ListBotsRequest is the request type for the
// ListBots API operation.
type ListBotsRequest struct {
	*aws.Request
	Input *types.ListBotsInput
	Copy  func(*types.ListBotsInput) ListBotsRequest
}

// Send marshals and sends the ListBots API request.
func (r ListBotsRequest) Send(ctx context.Context) (*ListBotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBotsResponse{
		ListBotsOutput: r.Request.Data.(*types.ListBotsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBotsRequestPaginator returns a paginator for ListBots.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBotsRequest(input)
//   p := chime.NewListBotsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBotsPaginator(req ListBotsRequest) ListBotsPaginator {
	return ListBotsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListBotsInput
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

// ListBotsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBotsPaginator struct {
	aws.Pager
}

func (p *ListBotsPaginator) CurrentPage() *types.ListBotsOutput {
	return p.Pager.CurrentPage().(*types.ListBotsOutput)
}

// ListBotsResponse is the response type for the
// ListBots API operation.
type ListBotsResponse struct {
	*types.ListBotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBots request.
func (r *ListBotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
