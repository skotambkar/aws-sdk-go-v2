// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
)

const opDescribeTapeRecoveryPoints = "DescribeTapeRecoveryPoints"

// DescribeTapeRecoveryPointsRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns a list of virtual tape recovery points that are available for the
// specified tape gateway.
//
// A recovery point is a point-in-time view of a virtual tape at which all the
// data on the virtual tape is consistent. If your gateway crashes, virtual
// tapes that have recovery points can be recovered to a new gateway. This operation
// is only supported in the tape gateway type.
//
//    // Example sending a request using DescribeTapeRecoveryPointsRequest.
//    req := client.DescribeTapeRecoveryPointsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeTapeRecoveryPoints
func (c *Client) DescribeTapeRecoveryPointsRequest(input *types.DescribeTapeRecoveryPointsInput) DescribeTapeRecoveryPointsRequest {
	op := &aws.Operation{
		Name:       opDescribeTapeRecoveryPoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.DescribeTapeRecoveryPointsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTapeRecoveryPointsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeTapeRecoveryPointsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeTapeRecoveryPointsRequest{Request: req, Input: input, Copy: c.DescribeTapeRecoveryPointsRequest}
}

// DescribeTapeRecoveryPointsRequest is the request type for the
// DescribeTapeRecoveryPoints API operation.
type DescribeTapeRecoveryPointsRequest struct {
	*aws.Request
	Input *types.DescribeTapeRecoveryPointsInput
	Copy  func(*types.DescribeTapeRecoveryPointsInput) DescribeTapeRecoveryPointsRequest
}

// Send marshals and sends the DescribeTapeRecoveryPoints API request.
func (r DescribeTapeRecoveryPointsRequest) Send(ctx context.Context) (*DescribeTapeRecoveryPointsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTapeRecoveryPointsResponse{
		DescribeTapeRecoveryPointsOutput: r.Request.Data.(*types.DescribeTapeRecoveryPointsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeTapeRecoveryPointsRequestPaginator returns a paginator for DescribeTapeRecoveryPoints.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeTapeRecoveryPointsRequest(input)
//   p := storagegateway.NewDescribeTapeRecoveryPointsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeTapeRecoveryPointsPaginator(req DescribeTapeRecoveryPointsRequest) DescribeTapeRecoveryPointsPaginator {
	return DescribeTapeRecoveryPointsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.DescribeTapeRecoveryPointsInput
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

// DescribeTapeRecoveryPointsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeTapeRecoveryPointsPaginator struct {
	aws.Pager
}

func (p *DescribeTapeRecoveryPointsPaginator) CurrentPage() *types.DescribeTapeRecoveryPointsOutput {
	return p.Pager.CurrentPage().(*types.DescribeTapeRecoveryPointsOutput)
}

// DescribeTapeRecoveryPointsResponse is the response type for the
// DescribeTapeRecoveryPoints API operation.
type DescribeTapeRecoveryPointsResponse struct {
	*types.DescribeTapeRecoveryPointsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTapeRecoveryPoints request.
func (r *DescribeTapeRecoveryPointsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
