// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
)

const opListBackupPlanVersions = "ListBackupPlanVersions"

// ListBackupPlanVersionsRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns version metadata of your backup plans, including Amazon Resource
// Names (ARNs), backup plan IDs, creation and deletion dates, plan names, and
// version IDs.
//
//    // Example sending a request using ListBackupPlanVersionsRequest.
//    req := client.ListBackupPlanVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/ListBackupPlanVersions
func (c *Client) ListBackupPlanVersionsRequest(input *types.ListBackupPlanVersionsInput) ListBackupPlanVersionsRequest {
	op := &aws.Operation{
		Name:       opListBackupPlanVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/versions/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListBackupPlanVersionsInput{}
	}

	req := c.newRequest(op, input, &types.ListBackupPlanVersionsOutput{})
	return ListBackupPlanVersionsRequest{Request: req, Input: input, Copy: c.ListBackupPlanVersionsRequest}
}

// ListBackupPlanVersionsRequest is the request type for the
// ListBackupPlanVersions API operation.
type ListBackupPlanVersionsRequest struct {
	*aws.Request
	Input *types.ListBackupPlanVersionsInput
	Copy  func(*types.ListBackupPlanVersionsInput) ListBackupPlanVersionsRequest
}

// Send marshals and sends the ListBackupPlanVersions API request.
func (r ListBackupPlanVersionsRequest) Send(ctx context.Context) (*ListBackupPlanVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListBackupPlanVersionsResponse{
		ListBackupPlanVersionsOutput: r.Request.Data.(*types.ListBackupPlanVersionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListBackupPlanVersionsRequestPaginator returns a paginator for ListBackupPlanVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListBackupPlanVersionsRequest(input)
//   p := backup.NewListBackupPlanVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListBackupPlanVersionsPaginator(req ListBackupPlanVersionsRequest) ListBackupPlanVersionsPaginator {
	return ListBackupPlanVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListBackupPlanVersionsInput
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

// ListBackupPlanVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListBackupPlanVersionsPaginator struct {
	aws.Pager
}

func (p *ListBackupPlanVersionsPaginator) CurrentPage() *types.ListBackupPlanVersionsOutput {
	return p.Pager.CurrentPage().(*types.ListBackupPlanVersionsOutput)
}

// ListBackupPlanVersionsResponse is the response type for the
// ListBackupPlanVersions API operation.
type ListBackupPlanVersionsResponse struct {
	*types.ListBackupPlanVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListBackupPlanVersions request.
func (r *ListBackupPlanVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
