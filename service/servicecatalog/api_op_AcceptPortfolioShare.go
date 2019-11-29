// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opAcceptPortfolioShare = "AcceptPortfolioShare"

// AcceptPortfolioShareRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Accepts an offer to share the specified portfolio.
//
//    // Example sending a request using AcceptPortfolioShareRequest.
//    req := client.AcceptPortfolioShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/AcceptPortfolioShare
func (c *Client) AcceptPortfolioShareRequest(input *types.AcceptPortfolioShareInput) AcceptPortfolioShareRequest {
	op := &aws.Operation{
		Name:       opAcceptPortfolioShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AcceptPortfolioShareInput{}
	}

	req := c.newRequest(op, input, &types.AcceptPortfolioShareOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AcceptPortfolioShareMarshaler{Input: input}.GetNamedBuildHandler())

	return AcceptPortfolioShareRequest{Request: req, Input: input, Copy: c.AcceptPortfolioShareRequest}
}

// AcceptPortfolioShareRequest is the request type for the
// AcceptPortfolioShare API operation.
type AcceptPortfolioShareRequest struct {
	*aws.Request
	Input *types.AcceptPortfolioShareInput
	Copy  func(*types.AcceptPortfolioShareInput) AcceptPortfolioShareRequest
}

// Send marshals and sends the AcceptPortfolioShare API request.
func (r AcceptPortfolioShareRequest) Send(ctx context.Context) (*AcceptPortfolioShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptPortfolioShareResponse{
		AcceptPortfolioShareOutput: r.Request.Data.(*types.AcceptPortfolioShareOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptPortfolioShareResponse is the response type for the
// AcceptPortfolioShare API operation.
type AcceptPortfolioShareResponse struct {
	*types.AcceptPortfolioShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptPortfolioShare request.
func (r *AcceptPortfolioShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
