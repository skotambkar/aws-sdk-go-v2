// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opGetDedicatedIps = "GetDedicatedIps"

// GetDedicatedIpsRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// List the dedicated IP addresses that are associated with your Amazon Pinpoint
// account.
//
//    // Example sending a request using GetDedicatedIpsRequest.
//    req := client.GetDedicatedIpsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDedicatedIps
func (c *Client) GetDedicatedIpsRequest(input *types.GetDedicatedIpsInput) GetDedicatedIpsRequest {
	op := &aws.Operation{
		Name:       opGetDedicatedIps,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/dedicated-ips",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "PageSize",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &types.GetDedicatedIpsInput{}
	}

	req := c.newRequest(op, input, &types.GetDedicatedIpsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetDedicatedIpsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetDedicatedIpsRequest{Request: req, Input: input, Copy: c.GetDedicatedIpsRequest}
}

// GetDedicatedIpsRequest is the request type for the
// GetDedicatedIps API operation.
type GetDedicatedIpsRequest struct {
	*aws.Request
	Input *types.GetDedicatedIpsInput
	Copy  func(*types.GetDedicatedIpsInput) GetDedicatedIpsRequest
}

// Send marshals and sends the GetDedicatedIps API request.
func (r GetDedicatedIpsRequest) Send(ctx context.Context) (*GetDedicatedIpsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDedicatedIpsResponse{
		GetDedicatedIpsOutput: r.Request.Data.(*types.GetDedicatedIpsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetDedicatedIpsRequestPaginator returns a paginator for GetDedicatedIps.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetDedicatedIpsRequest(input)
//   p := pinpointemail.NewGetDedicatedIpsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetDedicatedIpsPaginator(req GetDedicatedIpsRequest) GetDedicatedIpsPaginator {
	return GetDedicatedIpsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *types.GetDedicatedIpsInput
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

// GetDedicatedIpsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetDedicatedIpsPaginator struct {
	aws.Pager
}

func (p *GetDedicatedIpsPaginator) CurrentPage() *types.GetDedicatedIpsOutput {
	return p.Pager.CurrentPage().(*types.GetDedicatedIpsOutput)
}

// GetDedicatedIpsResponse is the response type for the
// GetDedicatedIps API operation.
type GetDedicatedIpsResponse struct {
	*types.GetDedicatedIpsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDedicatedIps request.
func (r *GetDedicatedIpsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
