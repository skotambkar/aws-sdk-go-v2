// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opListTypeRegistrations = "ListTypeRegistrations"

// ListTypeRegistrationsRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Returns a list of registration tokens for the specified type.
//
//    // Example sending a request using ListTypeRegistrationsRequest.
//    req := client.ListTypeRegistrationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/ListTypeRegistrations
func (c *Client) ListTypeRegistrationsRequest(input *types.ListTypeRegistrationsInput) ListTypeRegistrationsRequest {
	op := &aws.Operation{
		Name:       opListTypeRegistrations,
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
		input = &types.ListTypeRegistrationsInput{}
	}

	req := c.newRequest(op, input, &types.ListTypeRegistrationsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.ListTypeRegistrationsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListTypeRegistrationsRequest{Request: req, Input: input, Copy: c.ListTypeRegistrationsRequest}
}

// ListTypeRegistrationsRequest is the request type for the
// ListTypeRegistrations API operation.
type ListTypeRegistrationsRequest struct {
	*aws.Request
	Input *types.ListTypeRegistrationsInput
	Copy  func(*types.ListTypeRegistrationsInput) ListTypeRegistrationsRequest
}

// Send marshals and sends the ListTypeRegistrations API request.
func (r ListTypeRegistrationsRequest) Send(ctx context.Context) (*ListTypeRegistrationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListTypeRegistrationsResponse{
		ListTypeRegistrationsOutput: r.Request.Data.(*types.ListTypeRegistrationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListTypeRegistrationsRequestPaginator returns a paginator for ListTypeRegistrations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListTypeRegistrationsRequest(input)
//   p := cloudformation.NewListTypeRegistrationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListTypeRegistrationsPaginator(req ListTypeRegistrationsRequest) ListTypeRegistrationsPaginator {
	return ListTypeRegistrationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListTypeRegistrationsInput
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

// ListTypeRegistrationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListTypeRegistrationsPaginator struct {
	aws.Pager
}

func (p *ListTypeRegistrationsPaginator) CurrentPage() *types.ListTypeRegistrationsOutput {
	return p.Pager.CurrentPage().(*types.ListTypeRegistrationsOutput)
}

// ListTypeRegistrationsResponse is the response type for the
// ListTypeRegistrations API operation.
type ListTypeRegistrationsResponse struct {
	*types.ListTypeRegistrationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListTypeRegistrations request.
func (r *ListTypeRegistrationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
