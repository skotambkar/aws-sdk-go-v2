// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package clouddirectory

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
)

const opLookupPolicy = "LookupPolicy"

// LookupPolicyRequest returns a request value for making API operation for
// Amazon CloudDirectory.
//
// Lists all policies from the root of the Directory to the object specified.
// If there are no policies present, an empty list is returned. If policies
// are present, and if some objects don't have the policies attached, it returns
// the ObjectIdentifier for such objects. If policies are present, it returns
// ObjectIdentifier, policyId, and policyType. Paths that don't lead to the
// root from the target object are ignored. For more information, see Policies
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/key_concepts_directory.html#key_concepts_policies).
//
//    // Example sending a request using LookupPolicyRequest.
//    req := client.LookupPolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/clouddirectory-2017-01-11/LookupPolicy
func (c *Client) LookupPolicyRequest(input *types.LookupPolicyInput) LookupPolicyRequest {
	op := &aws.Operation{
		Name:       opLookupPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/amazonclouddirectory/2017-01-11/policy/lookup",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.LookupPolicyInput{}
	}

	req := c.newRequest(op, input, &types.LookupPolicyOutput{})
	return LookupPolicyRequest{Request: req, Input: input, Copy: c.LookupPolicyRequest}
}

// LookupPolicyRequest is the request type for the
// LookupPolicy API operation.
type LookupPolicyRequest struct {
	*aws.Request
	Input *types.LookupPolicyInput
	Copy  func(*types.LookupPolicyInput) LookupPolicyRequest
}

// Send marshals and sends the LookupPolicy API request.
func (r LookupPolicyRequest) Send(ctx context.Context) (*LookupPolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &LookupPolicyResponse{
		LookupPolicyOutput: r.Request.Data.(*types.LookupPolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewLookupPolicyRequestPaginator returns a paginator for LookupPolicy.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.LookupPolicyRequest(input)
//   p := clouddirectory.NewLookupPolicyRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewLookupPolicyPaginator(req LookupPolicyRequest) LookupPolicyPaginator {
	return LookupPolicyPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.LookupPolicyInput
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

// LookupPolicyPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type LookupPolicyPaginator struct {
	aws.Pager
}

func (p *LookupPolicyPaginator) CurrentPage() *types.LookupPolicyOutput {
	return p.Pager.CurrentPage().(*types.LookupPolicyOutput)
}

// LookupPolicyResponse is the response type for the
// LookupPolicy API operation.
type LookupPolicyResponse struct {
	*types.LookupPolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// LookupPolicy request.
func (r *LookupPolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
