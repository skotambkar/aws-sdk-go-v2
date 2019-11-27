// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opDisassociatePrincipalFromPortfolio = "DisassociatePrincipalFromPortfolio"

// DisassociatePrincipalFromPortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Disassociates a previously associated principal ARN from a specified portfolio.
//
//    // Example sending a request using DisassociatePrincipalFromPortfolioRequest.
//    req := client.DisassociatePrincipalFromPortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DisassociatePrincipalFromPortfolio
func (c *Client) DisassociatePrincipalFromPortfolioRequest(input *types.DisassociatePrincipalFromPortfolioInput) DisassociatePrincipalFromPortfolioRequest {
	op := &aws.Operation{
		Name:       opDisassociatePrincipalFromPortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociatePrincipalFromPortfolioInput{}
	}

	req := c.newRequest(op, input, &types.DisassociatePrincipalFromPortfolioOutput{})
	return DisassociatePrincipalFromPortfolioRequest{Request: req, Input: input, Copy: c.DisassociatePrincipalFromPortfolioRequest}
}

// DisassociatePrincipalFromPortfolioRequest is the request type for the
// DisassociatePrincipalFromPortfolio API operation.
type DisassociatePrincipalFromPortfolioRequest struct {
	*aws.Request
	Input *types.DisassociatePrincipalFromPortfolioInput
	Copy  func(*types.DisassociatePrincipalFromPortfolioInput) DisassociatePrincipalFromPortfolioRequest
}

// Send marshals and sends the DisassociatePrincipalFromPortfolio API request.
func (r DisassociatePrincipalFromPortfolioRequest) Send(ctx context.Context) (*DisassociatePrincipalFromPortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociatePrincipalFromPortfolioResponse{
		DisassociatePrincipalFromPortfolioOutput: r.Request.Data.(*types.DisassociatePrincipalFromPortfolioOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociatePrincipalFromPortfolioResponse is the response type for the
// DisassociatePrincipalFromPortfolio API operation.
type DisassociatePrincipalFromPortfolioResponse struct {
	*types.DisassociatePrincipalFromPortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociatePrincipalFromPortfolio request.
func (r *DisassociatePrincipalFromPortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
