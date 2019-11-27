// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDeletePortfolioShare = "DeletePortfolioShare"

// DeletePortfolioShareRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Stops sharing the specified portfolio with the specified account or organization
// node. Shares to an organization node can only be deleted by the master account
// of an Organization.
//
//    // Example sending a request using DeletePortfolioShareRequest.
//    req := client.DeletePortfolioShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeletePortfolioShare
func (c *Client) DeletePortfolioShareRequest(input *types.DeletePortfolioShareInput) DeletePortfolioShareRequest {
	op := &aws.Operation{
		Name:       opDeletePortfolioShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeletePortfolioShareInput{}
	}

	req := c.newRequest(op, input, &types.DeletePortfolioShareOutput{})
	return DeletePortfolioShareRequest{Request: req, Input: input, Copy: c.DeletePortfolioShareRequest}
}

// DeletePortfolioShareRequest is the request type for the
// DeletePortfolioShare API operation.
type DeletePortfolioShareRequest struct {
	*aws.Request
	Input *types.DeletePortfolioShareInput
	Copy  func(*types.DeletePortfolioShareInput) DeletePortfolioShareRequest
}

// Send marshals and sends the DeletePortfolioShare API request.
func (r DeletePortfolioShareRequest) Send(ctx context.Context) (*DeletePortfolioShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePortfolioShareResponse{
		DeletePortfolioShareOutput: r.Request.Data.(*types.DeletePortfolioShareOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePortfolioShareResponse is the response type for the
// DeletePortfolioShare API operation.
type DeletePortfolioShareResponse struct {
	*types.DeletePortfolioShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePortfolioShare request.
func (r *DeletePortfolioShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
