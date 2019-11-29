// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
)

const opListPoliciesForTarget = "ListPoliciesForTarget"

// ListPoliciesForTargetRequest returns a request value for making API operation for
// AWS Organizations.
//
// Lists the policies that are directly attached to the specified target root,
// organizational unit (OU), or account. You must specify the policy type that
// you want included in the returned list.
//
// Always check the NextToken response parameter for a null value when calling
// a List* operation. These operations can occasionally return an empty set
// of results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListPoliciesForTargetRequest.
//    req := client.ListPoliciesForTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListPoliciesForTarget
func (c *Client) ListPoliciesForTargetRequest(input *types.ListPoliciesForTargetInput) ListPoliciesForTargetRequest {
	op := &aws.Operation{
		Name:       opListPoliciesForTarget,
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
		input = &types.ListPoliciesForTargetInput{}
	}

	req := c.newRequest(op, input, &types.ListPoliciesForTargetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListPoliciesForTargetMarshaler{Input: input}.GetNamedBuildHandler())

	return ListPoliciesForTargetRequest{Request: req, Input: input, Copy: c.ListPoliciesForTargetRequest}
}

// ListPoliciesForTargetRequest is the request type for the
// ListPoliciesForTarget API operation.
type ListPoliciesForTargetRequest struct {
	*aws.Request
	Input *types.ListPoliciesForTargetInput
	Copy  func(*types.ListPoliciesForTargetInput) ListPoliciesForTargetRequest
}

// Send marshals and sends the ListPoliciesForTarget API request.
func (r ListPoliciesForTargetRequest) Send(ctx context.Context) (*ListPoliciesForTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPoliciesForTargetResponse{
		ListPoliciesForTargetOutput: r.Request.Data.(*types.ListPoliciesForTargetOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPoliciesForTargetRequestPaginator returns a paginator for ListPoliciesForTarget.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPoliciesForTargetRequest(input)
//   p := organizations.NewListPoliciesForTargetRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPoliciesForTargetPaginator(req ListPoliciesForTargetRequest) ListPoliciesForTargetPaginator {
	return ListPoliciesForTargetPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPoliciesForTargetInput
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

// ListPoliciesForTargetPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPoliciesForTargetPaginator struct {
	aws.Pager
}

func (p *ListPoliciesForTargetPaginator) CurrentPage() *types.ListPoliciesForTargetOutput {
	return p.Pager.CurrentPage().(*types.ListPoliciesForTargetOutput)
}

// ListPoliciesForTargetResponse is the response type for the
// ListPoliciesForTarget API operation.
type ListPoliciesForTargetResponse struct {
	*types.ListPoliciesForTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListPoliciesForTarget request.
func (r *ListPoliciesForTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
