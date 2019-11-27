// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codedeploy

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
)

const opListDeploymentConfigs = "ListDeploymentConfigs"

// ListDeploymentConfigsRequest returns a request value for making API operation for
// AWS CodeDeploy.
//
// Lists the deployment configurations with the IAM user or AWS account.
//
//    // Example sending a request using ListDeploymentConfigsRequest.
//    req := client.ListDeploymentConfigsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codedeploy-2014-10-06/ListDeploymentConfigs
func (c *Client) ListDeploymentConfigsRequest(input *types.ListDeploymentConfigsInput) ListDeploymentConfigsRequest {
	op := &aws.Operation{
		Name:       opListDeploymentConfigs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListDeploymentConfigsInput{}
	}

	req := c.newRequest(op, input, &types.ListDeploymentConfigsOutput{})
	return ListDeploymentConfigsRequest{Request: req, Input: input, Copy: c.ListDeploymentConfigsRequest}
}

// ListDeploymentConfigsRequest is the request type for the
// ListDeploymentConfigs API operation.
type ListDeploymentConfigsRequest struct {
	*aws.Request
	Input *types.ListDeploymentConfigsInput
	Copy  func(*types.ListDeploymentConfigsInput) ListDeploymentConfigsRequest
}

// Send marshals and sends the ListDeploymentConfigs API request.
func (r ListDeploymentConfigsRequest) Send(ctx context.Context) (*ListDeploymentConfigsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDeploymentConfigsResponse{
		ListDeploymentConfigsOutput: r.Request.Data.(*types.ListDeploymentConfigsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDeploymentConfigsRequestPaginator returns a paginator for ListDeploymentConfigs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDeploymentConfigsRequest(input)
//   p := codedeploy.NewListDeploymentConfigsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDeploymentConfigsPaginator(req ListDeploymentConfigsRequest) ListDeploymentConfigsPaginator {
	return ListDeploymentConfigsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListDeploymentConfigsInput
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

// ListDeploymentConfigsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDeploymentConfigsPaginator struct {
	aws.Pager
}

func (p *ListDeploymentConfigsPaginator) CurrentPage() *types.ListDeploymentConfigsOutput {
	return p.Pager.CurrentPage().(*types.ListDeploymentConfigsOutput)
}

// ListDeploymentConfigsResponse is the response type for the
// ListDeploymentConfigs API operation.
type ListDeploymentConfigsResponse struct {
	*types.ListDeploymentConfigsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDeploymentConfigs request.
func (r *ListDeploymentConfigsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
