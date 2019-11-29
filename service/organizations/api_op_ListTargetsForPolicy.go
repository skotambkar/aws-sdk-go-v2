// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opListTargetsForPolicy = "ListTargetsForPolicy"

// ListTargetsForPolicyRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists all the roots, organizational units (OUs), and accounts that the specified
// policy is attached to.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListTargetsForPolicyRequest.
//    req := client.ListTargetsForPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListTargetsForPolicy
func (c *Client) ListTargetsForPolicyRequest(input *types.ListTargetsForPolicyInput) ListTargetsForPolicyRequest {
	op := &aws.Operation{
		Name:       opListTargetsForPolicy,
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
		input = &types.ListTargetsForPolicyInput{}
	}

	req := c.newRequest(op, input, &types.ListTargetsForPolicyOutput{})
	return ListTargetsForPolicyRequest{Request: req, Input: input, Copy: c.ListTargetsForPolicyRequest}
}

// ListTargetsForPolicyRequest is the request type for the
// ListTargetsForPolicy API operation.
type ListTargetsForPolicyRequest struct {
	*aws.Request
	Input *types.ListTargetsForPolicyInput
	Copy  func(*types.ListTargetsForPolicyInput) ListTargetsForPolicyRequest
}

// Send marshals and sends the ListTargetsForPolicy API request.
func (r ListTargetsForPolicyRequest) Send(ctx context.Context) (*ListTargetsForPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTargetsForPolicyResponse{
		ListTargetsForPolicyOutput: r.Request.Data.(*types.ListTargetsForPolicyOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTargetsForPolicyRequestPaginator returns a paginator for ListTargetsForPolicy.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTargetsForPolicyRequest(input)
//   p := organizations.NewListTargetsForPolicyRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTargetsForPolicyPaginator(req ListTargetsForPolicyRequest) ListTargetsForPolicyPaginator {
	return ListTargetsForPolicyPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTargetsForPolicyInput
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

// ListTargetsForPolicyPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTargetsForPolicyPaginator struct {
	aws.Pager
}

func (p *ListTargetsForPolicyPaginator) CurrentPage() *types.ListTargetsForPolicyOutput {
	return p.Pager.CurrentPage().(*types.ListTargetsForPolicyOutput)
}

// ListTargetsForPolicyResponse is the response type for the
// ListTargetsForPolicy API operation.
type ListTargetsForPolicyResponse struct {
	*types.ListTargetsForPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTargetsForPolicy request.
func (r *ListTargetsForPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
