// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
)

const opListSteps = "ListSteps"

// ListStepsRequest returns a request value for making API operation for
// Amazon Elastic MapReduce.
//
// Provides a list of steps for the cluster in reverse order unless you specify
// stepIds with the request.
//
//    // Example sending a request using ListStepsRequest.
//    req := client.ListStepsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticmapreduce-2009-03-31/ListSteps
func (c *Client) ListStepsRequest(input *types.ListStepsInput) ListStepsRequest {
	op := &aws.Operation{
		Name:       opListSteps,
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
		input = &types.ListStepsInput{}
	}

	req := c.newRequest(op, input, &types.ListStepsOutput{})
	return ListStepsRequest{Request: req, Input: input, Copy: c.ListStepsRequest}
}

// ListStepsRequest is the request type for the
// ListSteps API operation.
type ListStepsRequest struct {
	*aws.Request
	Input *types.ListStepsInput
	Copy  func(*types.ListStepsInput) ListStepsRequest
}

// Send marshals and sends the ListSteps API request.
func (r ListStepsRequest) Send(ctx context.Context) (*ListStepsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStepsResponse{
		ListStepsOutput: r.Request.Data.(*types.ListStepsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStepsRequestPaginator returns a paginator for ListSteps.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStepsRequest(input)
//   p := emr.NewListStepsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStepsPaginator(req ListStepsRequest) ListStepsPaginator {
	return ListStepsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListStepsInput
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

// ListStepsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStepsPaginator struct {
	aws.Pager
}

func (p *ListStepsPaginator) CurrentPage() *types.ListStepsOutput {
	return p.Pager.CurrentPage().(*types.ListStepsOutput)
}

// ListStepsResponse is the response type for the
// ListSteps API operation.
type ListStepsResponse struct {
	*types.ListStepsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSteps request.
func (r *ListStepsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
