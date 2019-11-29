// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
)

const opGetReservedNodeExchangeOfferings = "GetReservedNodeExchangeOfferings"

// GetReservedNodeExchangeOfferingsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns an array of DC2 ReservedNodeOfferings that matches the payment type,
// term, and usage price of the given DC1 reserved node.
//
//    // Example sending a request using GetReservedNodeExchangeOfferingsRequest.
//    req := client.GetReservedNodeExchangeOfferingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/GetReservedNodeExchangeOfferings
func (c *Client) GetReservedNodeExchangeOfferingsRequest(input *types.GetReservedNodeExchangeOfferingsInput) GetReservedNodeExchangeOfferingsRequest {
	op := &aws.Operation{
		Name:       opGetReservedNodeExchangeOfferings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetReservedNodeExchangeOfferingsInput{}
	}

	req := c.newRequest(op, input, &types.GetReservedNodeExchangeOfferingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.GetReservedNodeExchangeOfferingsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetReservedNodeExchangeOfferingsRequest{Request: req, Input: input, Copy: c.GetReservedNodeExchangeOfferingsRequest}
}

// GetReservedNodeExchangeOfferingsRequest is the request type for the
// GetReservedNodeExchangeOfferings API operation.
type GetReservedNodeExchangeOfferingsRequest struct {
	*aws.Request
	Input *types.GetReservedNodeExchangeOfferingsInput
	Copy  func(*types.GetReservedNodeExchangeOfferingsInput) GetReservedNodeExchangeOfferingsRequest
}

// Send marshals and sends the GetReservedNodeExchangeOfferings API request.
func (r GetReservedNodeExchangeOfferingsRequest) Send(ctx context.Context) (*GetReservedNodeExchangeOfferingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetReservedNodeExchangeOfferingsResponse{
		GetReservedNodeExchangeOfferingsOutput: r.Request.Data.(*types.GetReservedNodeExchangeOfferingsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetReservedNodeExchangeOfferingsResponse is the response type for the
// GetReservedNodeExchangeOfferings API operation.
type GetReservedNodeExchangeOfferingsResponse struct {
	*types.GetReservedNodeExchangeOfferingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetReservedNodeExchangeOfferings request.
func (r *GetReservedNodeExchangeOfferingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
