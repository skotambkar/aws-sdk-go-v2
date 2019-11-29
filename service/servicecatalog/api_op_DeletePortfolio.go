// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDeletePortfolio = "DeletePortfolio"

// DeletePortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Deletes the specified portfolio.
//
// You cannot delete a portfolio if it was shared with you or if it has associated
// products, users, constraints, or shared accounts.
//
//    // Example sending a request using DeletePortfolioRequest.
//    req := client.DeletePortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeletePortfolio
func (c *Client) DeletePortfolioRequest(input *types.DeletePortfolioInput) DeletePortfolioRequest {
	op := &aws.Operation{
		Name:       opDeletePortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeletePortfolioInput{}
	}

	req := c.newRequest(op, input, &types.DeletePortfolioOutput{})
	return DeletePortfolioRequest{Request: req, Input: input, Copy: c.DeletePortfolioRequest}
}

// DeletePortfolioRequest is the request type for the
// DeletePortfolio API operation.
type DeletePortfolioRequest struct {
	*aws.Request
	Input *types.DeletePortfolioInput
	Copy  func(*types.DeletePortfolioInput) DeletePortfolioRequest
}

// Send marshals and sends the DeletePortfolio API request.
func (r DeletePortfolioRequest) Send(ctx context.Context) (*DeletePortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePortfolioResponse{
		DeletePortfolioOutput: r.Request.Data.(*types.DeletePortfolioOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePortfolioResponse is the response type for the
// DeletePortfolio API operation.
type DeletePortfolioResponse struct {
	*types.DeletePortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePortfolio request.
func (r *DeletePortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
