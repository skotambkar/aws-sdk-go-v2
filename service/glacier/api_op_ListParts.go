// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
)

const opListParts = "ListParts"

// ListPartsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation lists the parts of an archive that have been uploaded in a
// specific multipart upload. You can make this request at any time during an
// in-progress multipart upload before you complete the upload (see CompleteMultipartUpload.
// List Parts returns an error for completed uploads. The list returned in the
// List Parts response is sorted by part range.
//
// The List Parts operation supports pagination. By default, this operation
// returns up to 50 uploaded parts in the response. You should always check
// the response for a marker at which to continue the list; if there are no
// more items the marker is null. To return a list of parts that begins at a
// specific part, set the marker request parameter to the value you obtained
// from a previous List Parts request. You can also limit the number of parts
// returned in the response by specifying the limit parameter in the request.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and the underlying REST API, see Working with
// Archives in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html)
// and List Parts (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-list-parts.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using ListPartsRequest.
//    req := client.ListPartsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListPartsRequest(input *types.ListPartsInput) ListPartsRequest {
	op := &aws.Operation{
		Name:       opListParts,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/multipart-uploads/{uploadId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListPartsInput{}
	}

	req := c.newRequest(op, input, &types.ListPartsOutput{})
	return ListPartsRequest{Request: req, Input: input, Copy: c.ListPartsRequest}
}

// ListPartsRequest is the request type for the
// ListParts API operation.
type ListPartsRequest struct {
	*aws.Request
	Input *types.ListPartsInput
	Copy  func(*types.ListPartsInput) ListPartsRequest
}

// Send marshals and sends the ListParts API request.
func (r ListPartsRequest) Send(ctx context.Context) (*ListPartsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListPartsResponse{
		ListPartsOutput: r.Request.Data.(*types.ListPartsOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListPartsRequestPaginator returns a paginator for ListParts.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListPartsRequest(input)
//   p := glacier.NewListPartsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListPartsPaginator(req ListPartsRequest) ListPartsPaginator {
	return ListPartsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListPartsInput
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

// ListPartsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListPartsPaginator struct {
	aws.Pager
}

func (p *ListPartsPaginator) CurrentPage() *types.ListPartsOutput {
	return p.Pager.CurrentPage().(*types.ListPartsOutput)
}

// ListPartsResponse is the response type for the
// ListParts API operation.
type ListPartsResponse struct {
	*types.ListPartsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListParts request.
func (r *ListPartsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
