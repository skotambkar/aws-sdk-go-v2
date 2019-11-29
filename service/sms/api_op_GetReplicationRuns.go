// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opGetReplicationRuns = "GetReplicationRuns"

// GetReplicationRunsRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Describes the replication runs for the specified replication job.
//
//    // Example sending a request using GetReplicationRunsRequest.
//    req := client.GetReplicationRunsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/GetReplicationRuns
func (c *Client) GetReplicationRunsRequest(input *types.GetReplicationRunsInput) GetReplicationRunsRequest {
	op := &aws.Operation{
		Name:       opGetReplicationRuns,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetReplicationRunsInput{}
	}

	req := c.newRequest(op, input, &types.GetReplicationRunsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetReplicationRunsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetReplicationRunsRequest{Request: req, Input: input, Copy: c.GetReplicationRunsRequest}
}

// GetReplicationRunsRequest is the request type for the
// GetReplicationRuns API operation.
type GetReplicationRunsRequest struct {
	*aws.Request
	Input *types.GetReplicationRunsInput
	Copy  func(*types.GetReplicationRunsInput) GetReplicationRunsRequest
}

// Send marshals and sends the GetReplicationRuns API request.
func (r GetReplicationRunsRequest) Send(ctx context.Context) (*GetReplicationRunsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReplicationRunsResponse{
		GetReplicationRunsOutput: r.Request.Data.(*types.GetReplicationRunsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetReplicationRunsRequestPaginator returns a paginator for GetReplicationRuns.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetReplicationRunsRequest(input)
//   p := sms.NewGetReplicationRunsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetReplicationRunsPaginator(req GetReplicationRunsRequest) GetReplicationRunsPaginator {
	return GetReplicationRunsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetReplicationRunsInput
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

// GetReplicationRunsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetReplicationRunsPaginator struct {
	aws.Pager
}

func (p *GetReplicationRunsPaginator) CurrentPage() *types.GetReplicationRunsOutput {
	return p.Pager.CurrentPage().(*types.GetReplicationRunsOutput)
}

// GetReplicationRunsResponse is the response type for the
// GetReplicationRuns API operation.
type GetReplicationRunsResponse struct {
	*types.GetReplicationRunsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReplicationRuns request.
func (r *GetReplicationRunsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
