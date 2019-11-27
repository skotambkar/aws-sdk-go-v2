// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package batch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
)

const opDescribeJobDefinitions = "DescribeJobDefinitions"

// DescribeJobDefinitionsRequest returns a request value for making API operation for
// AWS Batch.
//
// Describes a list of job definitions. You can specify a status (such as ACTIVE)
// to only return job definitions that match that status.
//
//    // Example sending a request using DescribeJobDefinitionsRequest.
//    req := client.DescribeJobDefinitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/batch-2016-08-10/DescribeJobDefinitions
func (c *Client) DescribeJobDefinitionsRequest(input *types.DescribeJobDefinitionsInput) DescribeJobDefinitionsRequest {
	op := &aws.Operation{
		Name:       opDescribeJobDefinitions,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/describejobdefinitions",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeJobDefinitionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeJobDefinitionsOutput{})
	return DescribeJobDefinitionsRequest{Request: req, Input: input, Copy: c.DescribeJobDefinitionsRequest}
}

// DescribeJobDefinitionsRequest is the request type for the
// DescribeJobDefinitions API operation.
type DescribeJobDefinitionsRequest struct {
	*aws.Request
	Input *types.DescribeJobDefinitionsInput
	Copy  func(*types.DescribeJobDefinitionsInput) DescribeJobDefinitionsRequest
}

// Send marshals and sends the DescribeJobDefinitions API request.
func (r DescribeJobDefinitionsRequest) Send(ctx context.Context) (*DescribeJobDefinitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeJobDefinitionsResponse{
		DescribeJobDefinitionsOutput: r.Request.Data.(*types.DescribeJobDefinitionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeJobDefinitionsRequestPaginator returns a paginator for DescribeJobDefinitions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeJobDefinitionsRequest(input)
//   p := batch.NewDescribeJobDefinitionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeJobDefinitionsPaginator(req DescribeJobDefinitionsRequest) DescribeJobDefinitionsPaginator {
	return DescribeJobDefinitionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeJobDefinitionsInput
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

// DescribeJobDefinitionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeJobDefinitionsPaginator struct {
	aws.Pager
}

func (p *DescribeJobDefinitionsPaginator) CurrentPage() *types.DescribeJobDefinitionsOutput {
	return p.Pager.CurrentPage().(*types.DescribeJobDefinitionsOutput)
}

// DescribeJobDefinitionsResponse is the response type for the
// DescribeJobDefinitions API operation.
type DescribeJobDefinitionsResponse struct {
	*types.DescribeJobDefinitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeJobDefinitions request.
func (r *DescribeJobDefinitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
