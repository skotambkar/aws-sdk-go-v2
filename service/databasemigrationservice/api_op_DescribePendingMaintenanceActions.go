// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
)

const opDescribePendingMaintenanceActions = "DescribePendingMaintenanceActions"

// DescribePendingMaintenanceActionsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// For internal use only
//
//    // Example sending a request using DescribePendingMaintenanceActionsRequest.
//    req := client.DescribePendingMaintenanceActionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribePendingMaintenanceActions
func (c *Client) DescribePendingMaintenanceActionsRequest(input *types.DescribePendingMaintenanceActionsInput) DescribePendingMaintenanceActionsRequest {
	op := &aws.Operation{
		Name:       opDescribePendingMaintenanceActions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribePendingMaintenanceActionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribePendingMaintenanceActionsOutput{})
	return DescribePendingMaintenanceActionsRequest{Request: req, Input: input, Copy: c.DescribePendingMaintenanceActionsRequest}
}

// DescribePendingMaintenanceActionsRequest is the request type for the
// DescribePendingMaintenanceActions API operation.
type DescribePendingMaintenanceActionsRequest struct {
	*aws.Request
	Input *types.DescribePendingMaintenanceActionsInput
	Copy  func(*types.DescribePendingMaintenanceActionsInput) DescribePendingMaintenanceActionsRequest
}

// Send marshals and sends the DescribePendingMaintenanceActions API request.
func (r DescribePendingMaintenanceActionsRequest) Send(ctx context.Context) (*DescribePendingMaintenanceActionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePendingMaintenanceActionsResponse{
		DescribePendingMaintenanceActionsOutput: r.Request.Data.(*types.DescribePendingMaintenanceActionsOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribePendingMaintenanceActionsRequestPaginator returns a paginator for DescribePendingMaintenanceActions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribePendingMaintenanceActionsRequest(input)
//   p := databasemigrationservice.NewDescribePendingMaintenanceActionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribePendingMaintenanceActionsPaginator(req DescribePendingMaintenanceActionsRequest) DescribePendingMaintenanceActionsPaginator {
	return DescribePendingMaintenanceActionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribePendingMaintenanceActionsInput
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

// DescribePendingMaintenanceActionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribePendingMaintenanceActionsPaginator struct {
	aws.Pager
}

func (p *DescribePendingMaintenanceActionsPaginator) CurrentPage() *types.DescribePendingMaintenanceActionsOutput {
	return p.Pager.CurrentPage().(*types.DescribePendingMaintenanceActionsOutput)
}

// DescribePendingMaintenanceActionsResponse is the response type for the
// DescribePendingMaintenanceActions API operation.
type DescribePendingMaintenanceActionsResponse struct {
	*types.DescribePendingMaintenanceActionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePendingMaintenanceActions request.
func (r *DescribePendingMaintenanceActionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
