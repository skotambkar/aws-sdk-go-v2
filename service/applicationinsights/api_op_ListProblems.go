// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
)

const opListProblems = "ListProblems"

// ListProblemsRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Lists the problems with your application.
//
//    // Example sending a request using ListProblemsRequest.
//    req := client.ListProblemsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/ListProblems
func (c *Client) ListProblemsRequest(input *types.ListProblemsInput) ListProblemsRequest {
	op := &aws.Operation{
		Name:       opListProblems,
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
		input = &types.ListProblemsInput{}
	}

	req := c.newRequest(op, input, &types.ListProblemsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListProblemsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListProblemsRequest{Request: req, Input: input, Copy: c.ListProblemsRequest}
}

// ListProblemsRequest is the request type for the
// ListProblems API operation.
type ListProblemsRequest struct {
	*aws.Request
	Input *types.ListProblemsInput
	Copy  func(*types.ListProblemsInput) ListProblemsRequest
}

// Send marshals and sends the ListProblems API request.
func (r ListProblemsRequest) Send(ctx context.Context) (*ListProblemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProblemsResponse{
		ListProblemsOutput: r.Request.Data.(*types.ListProblemsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListProblemsRequestPaginator returns a paginator for ListProblems.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListProblemsRequest(input)
//   p := applicationinsights.NewListProblemsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListProblemsPaginator(req ListProblemsRequest) ListProblemsPaginator {
	return ListProblemsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListProblemsInput
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

// ListProblemsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListProblemsPaginator struct {
	aws.Pager
}

func (p *ListProblemsPaginator) CurrentPage() *types.ListProblemsOutput {
	return p.Pager.CurrentPage().(*types.ListProblemsOutput)
}

// ListProblemsResponse is the response type for the
// ListProblems API operation.
type ListProblemsResponse struct {
	*types.ListProblemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProblems request.
func (r *ListProblemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
