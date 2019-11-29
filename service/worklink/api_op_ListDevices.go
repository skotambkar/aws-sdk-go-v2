// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
)

const opListDevices = "ListDevices"

// ListDevicesRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Retrieves a list of devices registered with the specified fleet.
//
//    // Example sending a request using ListDevicesRequest.
//    req := client.ListDevicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListDevices
func (c *Client) ListDevicesRequest(input *types.ListDevicesInput) ListDevicesRequest {
	op := &aws.Operation{
		Name:       opListDevices,
		HTTPMethod: "POST",
		HTTPPath:   "/listDevices",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.ListDevicesInput{}
	}

	req := c.newRequest(op, input, &types.ListDevicesOutput{})
	return ListDevicesRequest{Request: req, Input: input, Copy: c.ListDevicesRequest}
}

// ListDevicesRequest is the request type for the
// ListDevices API operation.
type ListDevicesRequest struct {
	*aws.Request
	Input *types.ListDevicesInput
	Copy  func(*types.ListDevicesInput) ListDevicesRequest
}

// Send marshals and sends the ListDevices API request.
func (r ListDevicesRequest) Send(ctx context.Context) (*ListDevicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDevicesResponse{
		ListDevicesOutput: r.Request.Data.(*types.ListDevicesOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListDevicesRequestPaginator returns a paginator for ListDevices.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListDevicesRequest(input)
//   p := worklink.NewListDevicesRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListDevicesPaginator(req ListDevicesRequest) ListDevicesPaginator {
	return ListDevicesPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.ListDevicesInput
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

// ListDevicesPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListDevicesPaginator struct {
	aws.Pager
}

func (p *ListDevicesPaginator) CurrentPage() *types.ListDevicesOutput {
	return p.Pager.CurrentPage().(*types.ListDevicesOutput)
}

// ListDevicesResponse is the response type for the
// ListDevices API operation.
type ListDevicesResponse struct {
	*types.ListDevicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDevices request.
func (r *ListDevicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
