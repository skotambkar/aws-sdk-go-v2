// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opListMigrationTasks = "ListMigrationTasks"

// ListMigrationTasksRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Lists all, or filtered by resource name, migration tasks associated with
// the user account making this call. This API has the following traits:
//
//    * Can show a summary list of the most recent migration tasks.
//
//    * Can show a summary list of migration tasks associated with a given discovered
//    resource.
//
//    * Lists migration tasks in a paginated interface.
//
//    // Example sending a request using ListMigrationTasksRequest.
//    req := client.ListMigrationTasksRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListMigrationTasks
func (c *Client) ListMigrationTasksRequest(input *types.ListMigrationTasksInput) ListMigrationTasksRequest {
	op := &aws.Operation{
		Name:       opListMigrationTasks,
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
		input = &types.ListMigrationTasksInput{}
	}

	req := c.newRequest(op, input, &types.ListMigrationTasksOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListMigrationTasksMarshaler{Input: input}.GetNamedBuildHandler())

	return ListMigrationTasksRequest{Request: req, Input: input, Copy: c.ListMigrationTasksRequest}
}

// ListMigrationTasksRequest is the request type for the
// ListMigrationTasks API operation.
type ListMigrationTasksRequest struct {
	*aws.Request
	Input *types.ListMigrationTasksInput
	Copy  func(*types.ListMigrationTasksInput) ListMigrationTasksRequest
}

// Send marshals and sends the ListMigrationTasks API request.
func (r ListMigrationTasksRequest) Send(ctx context.Context) (*ListMigrationTasksResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListMigrationTasksResponse{
		ListMigrationTasksOutput: r.Request.Data.(*types.ListMigrationTasksOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListMigrationTasksRequestPaginator returns a paginator for ListMigrationTasks.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListMigrationTasksRequest(input)
//   p := migrationhub.NewListMigrationTasksRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListMigrationTasksPaginator(req ListMigrationTasksRequest) ListMigrationTasksPaginator {
	return ListMigrationTasksPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListMigrationTasksInput
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

// ListMigrationTasksPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListMigrationTasksPaginator struct {
	aws.Pager
}

func (p *ListMigrationTasksPaginator) CurrentPage() *types.ListMigrationTasksOutput {
	return p.Pager.CurrentPage().(*types.ListMigrationTasksOutput)
}

// ListMigrationTasksResponse is the response type for the
// ListMigrationTasks API operation.
type ListMigrationTasksResponse struct {
	*types.ListMigrationTasksOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListMigrationTasks request.
func (r *ListMigrationTasksResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
