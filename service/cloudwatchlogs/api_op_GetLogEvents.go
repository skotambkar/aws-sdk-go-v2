// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

const opGetLogEvents = "GetLogEvents"

// GetLogEventsRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists log events from the specified log stream. You can list all the log
// events or filter using a time range.
//
// By default, this operation returns as many log events as can fit in a response
// size of 1MB (up to 10,000 log events). You can get additional log events
// by specifying one of the tokens in a subsequent call.
//
//    // Example sending a request using GetLogEventsRequest.
//    req := client.GetLogEventsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/GetLogEvents
func (c *Client) GetLogEventsRequest(input *types.GetLogEventsInput) GetLogEventsRequest {
	op := &aws.Operation{
		Name:       opGetLogEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextForwardToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetLogEventsInput{}
	}

	req := c.newRequest(op, input, &types.GetLogEventsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetLogEventsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetLogEventsRequest{Request: req, Input: input, Copy: c.GetLogEventsRequest}
}

// GetLogEventsRequest is the request type for the
// GetLogEvents API operation.
type GetLogEventsRequest struct {
	*aws.Request
	Input *types.GetLogEventsInput
	Copy  func(*types.GetLogEventsInput) GetLogEventsRequest
}

// Send marshals and sends the GetLogEvents API request.
func (r GetLogEventsRequest) Send(ctx context.Context) (*GetLogEventsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetLogEventsResponse{
		GetLogEventsOutput: r.Request.Data.(*types.GetLogEventsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetLogEventsRequestPaginator returns a paginator for GetLogEvents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetLogEventsRequest(input)
//   p := cloudwatchlogs.NewGetLogEventsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetLogEventsPaginator(req GetLogEventsRequest) GetLogEventsPaginator {
	return GetLogEventsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetLogEventsInput
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

// GetLogEventsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetLogEventsPaginator struct {
	aws.Pager
}

func (p *GetLogEventsPaginator) CurrentPage() *types.GetLogEventsOutput {
	return p.Pager.CurrentPage().(*types.GetLogEventsOutput)
}

// GetLogEventsResponse is the response type for the
// GetLogEvents API operation.
type GetLogEventsResponse struct {
	*types.GetLogEventsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetLogEvents request.
func (r *GetLogEventsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
