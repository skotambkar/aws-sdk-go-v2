// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchlogs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

const opDescribeLogStreams = "DescribeLogStreams"

// DescribeLogStreamsRequest returns a request value for making API operation for
// Amazon CloudWatch Logs.
//
// Lists the log streams for the specified log group. You can list all the log
// streams or filter the results by prefix. You can also control how the results
// are ordered.
//
// This operation has a limit of five transactions per second, after which transactions
// are throttled.
//
//    // Example sending a request using DescribeLogStreamsRequest.
//    req := client.DescribeLogStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28/DescribeLogStreams
func (c *Client) DescribeLogStreamsRequest(input *types.DescribeLogStreamsInput) DescribeLogStreamsRequest {
	op := &aws.Operation{
		Name:       opDescribeLogStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeLogStreamsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeLogStreamsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeLogStreamsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeLogStreamsRequest{Request: req, Input: input, Copy: c.DescribeLogStreamsRequest}
}

// DescribeLogStreamsRequest is the request type for the
// DescribeLogStreams API operation.
type DescribeLogStreamsRequest struct {
	*aws.Request
	Input *types.DescribeLogStreamsInput
	Copy  func(*types.DescribeLogStreamsInput) DescribeLogStreamsRequest
}

// Send marshals and sends the DescribeLogStreams API request.
func (r DescribeLogStreamsRequest) Send(ctx context.Context) (*DescribeLogStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLogStreamsResponse{
		DescribeLogStreamsOutput: r.Request.Data.(*types.DescribeLogStreamsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeLogStreamsRequestPaginator returns a paginator for DescribeLogStreams.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeLogStreamsRequest(input)
//   p := cloudwatchlogs.NewDescribeLogStreamsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeLogStreamsPaginator(req DescribeLogStreamsRequest) DescribeLogStreamsPaginator {
	return DescribeLogStreamsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeLogStreamsInput
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

// DescribeLogStreamsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeLogStreamsPaginator struct {
	aws.Pager
}

func (p *DescribeLogStreamsPaginator) CurrentPage() *types.DescribeLogStreamsOutput {
	return p.Pager.CurrentPage().(*types.DescribeLogStreamsOutput)
}

// DescribeLogStreamsResponse is the response type for the
// DescribeLogStreams API operation.
type DescribeLogStreamsResponse struct {
	*types.DescribeLogStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLogStreams request.
func (r *DescribeLogStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
