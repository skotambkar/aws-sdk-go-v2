// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListProblemsInput struct {
	_ struct{} `type:"structure"`

	// The time when the problem ended, in epoch seconds. If not specified, problems
	// within the past seven days are returned.
	EndTime *time.Time `type:"timestamp"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token to request the next page of results.
	NextToken *string `type:"string"`

	// The name of the resource group.
	ResourceGroupName *string `min:"1" type:"string"`

	// The time when the problem was detected, in epoch seconds. If you don't specify
	// a time frame for the request, problems within the past seven days are returned.
	StartTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s ListProblemsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProblemsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProblemsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ResourceGroupName != nil && len(*s.ResourceGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProblemsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string `type:"string"`

	// The list of problems.
	ProblemList []Problem `type:"list"`
}

// String returns the string representation
func (s ListProblemsOutput) String() string {
	return awsutil.Prettify(s)
}

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
func (c *Client) ListProblemsRequest(input *ListProblemsInput) ListProblemsRequest {
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
		input = &ListProblemsInput{}
	}

	req := c.newRequest(op, input, &ListProblemsOutput{})
	return ListProblemsRequest{Request: req, Input: input, Copy: c.ListProblemsRequest}
}

// ListProblemsRequest is the request type for the
// ListProblems API operation.
type ListProblemsRequest struct {
	*aws.Request
	Input *ListProblemsInput
	Copy  func(*ListProblemsInput) ListProblemsRequest
}

// Send marshals and sends the ListProblems API request.
func (r ListProblemsRequest) Send(ctx context.Context) (*ListProblemsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListProblemsResponse{
		ListProblemsOutput: r.Request.Data.(*ListProblemsOutput),
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
				var inCpy *ListProblemsInput
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

func (p *ListProblemsPaginator) CurrentPage() *ListProblemsOutput {
	return p.Pager.CurrentPage().(*ListProblemsOutput)
}

// ListProblemsResponse is the response type for the
// ListProblems API operation.
type ListProblemsResponse struct {
	*ListProblemsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListProblems request.
func (r *ListProblemsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
