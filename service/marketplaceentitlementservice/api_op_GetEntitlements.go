// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package marketplaceentitlementservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice/types"
)

const opGetEntitlements = "GetEntitlements"

// GetEntitlementsRequest returns a request value for making API operation for
// AWS Marketplace Entitlement Service.
//
// GetEntitlements retrieves entitlement values for a given product. The results
// can be filtered based on customer identifier or product dimensions.
//
//    // Example sending a request using GetEntitlementsRequest.
//    req := client.GetEntitlementsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/entitlement.marketplace-2017-01-11/GetEntitlements
func (c *Client) GetEntitlementsRequest(input *types.GetEntitlementsInput) GetEntitlementsRequest {
	op := &aws.Operation{
		Name:       opGetEntitlements,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetEntitlementsInput{}
	}

	req := c.newRequest(op, input, &types.GetEntitlementsOutput{})
	return GetEntitlementsRequest{Request: req, Input: input, Copy: c.GetEntitlementsRequest}
}

// GetEntitlementsRequest is the request type for the
// GetEntitlements API operation.
type GetEntitlementsRequest struct {
	*aws.Request
	Input *types.GetEntitlementsInput
	Copy  func(*types.GetEntitlementsInput) GetEntitlementsRequest
}

// Send marshals and sends the GetEntitlements API request.
func (r GetEntitlementsRequest) Send(ctx context.Context) (*GetEntitlementsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetEntitlementsResponse{
		GetEntitlementsOutput: r.Request.Data.(*types.GetEntitlementsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetEntitlementsResponse is the response type for the
// GetEntitlements API operation.
type GetEntitlementsResponse struct {
	*types.GetEntitlementsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetEntitlements request.
func (r *GetEntitlementsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
