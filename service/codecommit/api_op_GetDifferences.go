// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
)

const opGetDifferences = "GetDifferences"

// GetDifferencesRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Returns information about the differences in a valid commit specifier (such
// as a branch, tag, HEAD, commit ID or other fully qualified reference). Results
// can be limited to a specified path.
//
//    // Example sending a request using GetDifferencesRequest.
//    req := client.GetDifferencesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/GetDifferences
func (c *Client) GetDifferencesRequest(input *types.GetDifferencesInput) GetDifferencesRequest {
	op := &aws.Operation{
		Name:       opGetDifferences,
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
		input = &types.GetDifferencesInput{}
	}

	req := c.newRequest(op, input, &types.GetDifferencesOutput{})
	return GetDifferencesRequest{Request: req, Input: input, Copy: c.GetDifferencesRequest}
}

// GetDifferencesRequest is the request type for the
// GetDifferences API operation.
type GetDifferencesRequest struct {
	*aws.Request
	Input *types.GetDifferencesInput
	Copy  func(*types.GetDifferencesInput) GetDifferencesRequest
}

// Send marshals and sends the GetDifferences API request.
func (r GetDifferencesRequest) Send(ctx context.Context) (*GetDifferencesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDifferencesResponse{
		GetDifferencesOutput: r.Request.Data.(*types.GetDifferencesOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDifferencesRequestPaginator returns a paginator for GetDifferences.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDifferencesRequest(input)
//   p := codecommit.NewGetDifferencesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDifferencesPaginator(req GetDifferencesRequest) GetDifferencesPaginator {
	return GetDifferencesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetDifferencesInput
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

// GetDifferencesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDifferencesPaginator struct {
	aws.Pager
}

func (p *GetDifferencesPaginator) CurrentPage() *types.GetDifferencesOutput {
	return p.Pager.CurrentPage().(*types.GetDifferencesOutput)
}

// GetDifferencesResponse is the response type for the
// GetDifferences API operation.
type GetDifferencesResponse struct {
	*types.GetDifferencesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDifferences request.
func (r *GetDifferencesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
