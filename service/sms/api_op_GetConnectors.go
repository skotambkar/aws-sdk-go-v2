// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sms

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sms/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
)

const opGetConnectors = "GetConnectors"

// GetConnectorsRequest returns a request value for making API operation for
// AWS Server Migration Service.
//
// Describes the connectors registered with the AWS SMS.
//
//    // Example sending a request using GetConnectorsRequest.
//    req := client.GetConnectorsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sms-2016-10-24/GetConnectors
func (c *Client) GetConnectorsRequest(input *types.GetConnectorsInput) GetConnectorsRequest {
	op := &aws.Operation{
		Name:       opGetConnectors,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetConnectorsInput{}
	}

	req := c.newRequest(op, input, &types.GetConnectorsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.GetConnectorsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetConnectorsRequest{Request: req, Input: input, Copy: c.GetConnectorsRequest}
}

// GetConnectorsRequest is the request type for the
// GetConnectors API operation.
type GetConnectorsRequest struct {
	*aws.Request
	Input *types.GetConnectorsInput
	Copy  func(*types.GetConnectorsInput) GetConnectorsRequest
}

// Send marshals and sends the GetConnectors API request.
func (r GetConnectorsRequest) Send(ctx context.Context) (*GetConnectorsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetConnectorsResponse{
		GetConnectorsOutput: r.Request.Data.(*types.GetConnectorsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetConnectorsRequestPaginator returns a paginator for GetConnectors.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetConnectorsRequest(input)
//   p := sms.NewGetConnectorsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetConnectorsPaginator(req GetConnectorsRequest) GetConnectorsPaginator {
	return GetConnectorsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetConnectorsInput
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

// GetConnectorsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetConnectorsPaginator struct {
	aws.Pager
}

func (p *GetConnectorsPaginator) CurrentPage() *types.GetConnectorsOutput {
	return p.Pager.CurrentPage().(*types.GetConnectorsOutput)
}

// GetConnectorsResponse is the response type for the
// GetConnectors API operation.
type GetConnectorsResponse struct {
	*types.GetConnectorsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetConnectors request.
func (r *GetConnectorsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
