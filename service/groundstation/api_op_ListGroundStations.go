// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package groundstation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
)

const opListGroundStations = "ListGroundStations"

// ListGroundStationsRequest returns a request value for making API operation for
// AWS Ground Station.
//
// Returns a list of ground stations.
//
//    // Example sending a request using ListGroundStationsRequest.
//    req := client.ListGroundStationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/groundstation-2019-05-23/ListGroundStations
func (c *Client) ListGroundStationsRequest(input *types.ListGroundStationsInput) ListGroundStationsRequest {
	op := &aws.Operation{
		Name:       opListGroundStations,
		HTTPMethod: "GET",
		HTTPPath:   "/groundstation",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListGroundStationsInput{}
	}

	req := c.newRequest(op, input, &types.ListGroundStationsOutput{})
	return ListGroundStationsRequest{Request: req, Input: input, Copy: c.ListGroundStationsRequest}
}

// ListGroundStationsRequest is the request type for the
// ListGroundStations API operation.
type ListGroundStationsRequest struct {
	*aws.Request
	Input *types.ListGroundStationsInput
	Copy  func(*types.ListGroundStationsInput) ListGroundStationsRequest
}

// Send marshals and sends the ListGroundStations API request.
func (r ListGroundStationsRequest) Send(ctx context.Context) (*ListGroundStationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListGroundStationsResponse{
		ListGroundStationsOutput: r.Request.Data.(*types.ListGroundStationsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListGroundStationsRequestPaginator returns a paginator for ListGroundStations.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListGroundStationsRequest(input)
//   p := groundstation.NewListGroundStationsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListGroundStationsPaginator(req ListGroundStationsRequest) ListGroundStationsPaginator {
	return ListGroundStationsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListGroundStationsInput
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

// ListGroundStationsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListGroundStationsPaginator struct {
	aws.Pager
}

func (p *ListGroundStationsPaginator) CurrentPage() *types.ListGroundStationsOutput {
	return p.Pager.CurrentPage().(*types.ListGroundStationsOutput)
}

// ListGroundStationsResponse is the response type for the
// ListGroundStations API operation.
type ListGroundStationsResponse struct {
	*types.ListGroundStationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListGroundStations request.
func (r *ListGroundStationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
