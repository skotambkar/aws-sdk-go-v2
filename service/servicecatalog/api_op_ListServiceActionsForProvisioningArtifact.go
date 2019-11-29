// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opListServiceActionsForProvisioningArtifact = "ListServiceActionsForProvisioningArtifact"

// ListServiceActionsForProvisioningArtifactRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Returns a paginated list of self-service actions associated with the specified
// Product ID and Provisioning Artifact ID.
//
//    // Example sending a request using ListServiceActionsForProvisioningArtifactRequest.
//    req := client.ListServiceActionsForProvisioningArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/ListServiceActionsForProvisioningArtifact
func (c *Client) ListServiceActionsForProvisioningArtifactRequest(input *types.ListServiceActionsForProvisioningArtifactInput) ListServiceActionsForProvisioningArtifactRequest {
	op := &aws.Operation{
		Name:       opListServiceActionsForProvisioningArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListServiceActionsForProvisioningArtifactInput{}
	}

	req := c.newRequest(op, input, &types.ListServiceActionsForProvisioningArtifactOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.ListServiceActionsForProvisioningArtifactMarshaler{Input: input}.GetNamedBuildHandler())

	return ListServiceActionsForProvisioningArtifactRequest{Request: req, Input: input, Copy: c.ListServiceActionsForProvisioningArtifactRequest}
}

// ListServiceActionsForProvisioningArtifactRequest is the request type for the
// ListServiceActionsForProvisioningArtifact API operation.
type ListServiceActionsForProvisioningArtifactRequest struct {
	*aws.Request
	Input *types.ListServiceActionsForProvisioningArtifactInput
	Copy  func(*types.ListServiceActionsForProvisioningArtifactInput) ListServiceActionsForProvisioningArtifactRequest
}

// Send marshals and sends the ListServiceActionsForProvisioningArtifact API request.
func (r ListServiceActionsForProvisioningArtifactRequest) Send(ctx context.Context) (*ListServiceActionsForProvisioningArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListServiceActionsForProvisioningArtifactResponse{
		ListServiceActionsForProvisioningArtifactOutput: r.Request.Data.(*types.ListServiceActionsForProvisioningArtifactOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListServiceActionsForProvisioningArtifactRequestPaginator returns a paginator for ListServiceActionsForProvisioningArtifact.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListServiceActionsForProvisioningArtifactRequest(input)
//   p := servicecatalog.NewListServiceActionsForProvisioningArtifactRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListServiceActionsForProvisioningArtifactPaginator(req ListServiceActionsForProvisioningArtifactRequest) ListServiceActionsForProvisioningArtifactPaginator {
	return ListServiceActionsForProvisioningArtifactPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListServiceActionsForProvisioningArtifactInput
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

// ListServiceActionsForProvisioningArtifactPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListServiceActionsForProvisioningArtifactPaginator struct {
	aws.Pager
}

func (p *ListServiceActionsForProvisioningArtifactPaginator) CurrentPage() *types.ListServiceActionsForProvisioningArtifactOutput {
	return p.Pager.CurrentPage().(*types.ListServiceActionsForProvisioningArtifactOutput)
}

// ListServiceActionsForProvisioningArtifactResponse is the response type for the
// ListServiceActionsForProvisioningArtifact API operation.
type ListServiceActionsForProvisioningArtifactResponse struct {
	*types.ListServiceActionsForProvisioningArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListServiceActionsForProvisioningArtifact request.
func (r *ListServiceActionsForProvisioningArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
