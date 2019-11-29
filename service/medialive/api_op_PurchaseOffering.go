// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
)

const opPurchaseOffering = "PurchaseOffering"

// PurchaseOfferingRequest returns a request value for making API operation for
// AWS Elemental MediaLive.
//
// Purchase an offering and create a reservation.
//
//    // Example sending a request using PurchaseOfferingRequest.
//    req := client.PurchaseOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/medialive-2017-10-14/PurchaseOffering
func (c *Client) PurchaseOfferingRequest(input *types.PurchaseOfferingInput) PurchaseOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/prod/offerings/{offeringId}/purchase",
	}

	if input == nil {
		input = &types.PurchaseOfferingInput{}
	}

	req := c.newRequest(op, input, &types.PurchaseOfferingOutput{})
	return PurchaseOfferingRequest{Request: req, Input: input, Copy: c.PurchaseOfferingRequest}
}

// PurchaseOfferingRequest is the request type for the
// PurchaseOffering API operation.
type PurchaseOfferingRequest struct {
	*aws.Request
	Input *types.PurchaseOfferingInput
	Copy  func(*types.PurchaseOfferingInput) PurchaseOfferingRequest
}

// Send marshals and sends the PurchaseOffering API request.
func (r PurchaseOfferingRequest) Send(ctx context.Context) (*PurchaseOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseOfferingResponse{
		PurchaseOfferingOutput: r.Request.Data.(*types.PurchaseOfferingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseOfferingResponse is the response type for the
// PurchaseOffering API operation.
type PurchaseOfferingResponse struct {
	*types.PurchaseOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseOffering request.
func (r *PurchaseOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
