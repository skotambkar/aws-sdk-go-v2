// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opListCreatedArtifacts = "ListCreatedArtifacts"

// ListCreatedArtifactsRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Lists the created artifacts attached to a given migration task in an update
// stream. This API has the following traits:
//
//    * Gets the list of the created artifacts while migration is taking place.
//
//    * Shows the artifacts created by the migration tool that was associated
//    by the AssociateCreatedArtifact API.
//
//    * Lists created artifacts in a paginated interface.
//
//    // Example sending a request using ListCreatedArtifactsRequest.
//    req := client.ListCreatedArtifactsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/ListCreatedArtifacts
func (c *Client) ListCreatedArtifactsRequest(input *types.ListCreatedArtifactsInput) ListCreatedArtifactsRequest {
	op := &aws.Operation{
		Name:       opListCreatedArtifacts,
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
		input = &types.ListCreatedArtifactsInput{}
	}

	req := c.newRequest(op, input, &types.ListCreatedArtifactsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListCreatedArtifactsMarshaler{Input: input}.GetNamedBuildHandler())

	return ListCreatedArtifactsRequest{Request: req, Input: input, Copy: c.ListCreatedArtifactsRequest}
}

// ListCreatedArtifactsRequest is the request type for the
// ListCreatedArtifacts API operation.
type ListCreatedArtifactsRequest struct {
	*aws.Request
	Input *types.ListCreatedArtifactsInput
	Copy  func(*types.ListCreatedArtifactsInput) ListCreatedArtifactsRequest
}

// Send marshals and sends the ListCreatedArtifacts API request.
func (r ListCreatedArtifactsRequest) Send(ctx context.Context) (*ListCreatedArtifactsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListCreatedArtifactsResponse{
		ListCreatedArtifactsOutput: r.Request.Data.(*types.ListCreatedArtifactsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListCreatedArtifactsRequestPaginator returns a paginator for ListCreatedArtifacts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListCreatedArtifactsRequest(input)
//   p := migrationhub.NewListCreatedArtifactsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListCreatedArtifactsPaginator(req ListCreatedArtifactsRequest) ListCreatedArtifactsPaginator {
	return ListCreatedArtifactsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListCreatedArtifactsInput
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

// ListCreatedArtifactsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListCreatedArtifactsPaginator struct {
	aws.Pager
}

func (p *ListCreatedArtifactsPaginator) CurrentPage() *types.ListCreatedArtifactsOutput {
	return p.Pager.CurrentPage().(*types.ListCreatedArtifactsOutput)
}

// ListCreatedArtifactsResponse is the response type for the
// ListCreatedArtifacts API operation.
type ListCreatedArtifactsResponse struct {
	*types.ListCreatedArtifactsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListCreatedArtifacts request.
func (r *ListCreatedArtifactsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
