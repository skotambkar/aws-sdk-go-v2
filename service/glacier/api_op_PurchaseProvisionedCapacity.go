// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/glacier/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
)

const opPurchaseProvisionedCapacity = "PurchaseProvisionedCapacity"

// PurchaseProvisionedCapacityRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation purchases a provisioned capacity unit for an AWS account.
//
//    // Example sending a request using PurchaseProvisionedCapacityRequest.
//    req := client.PurchaseProvisionedCapacityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) PurchaseProvisionedCapacityRequest(input *types.PurchaseProvisionedCapacityInput) PurchaseProvisionedCapacityRequest {
	op := &aws.Operation{
		Name:       opPurchaseProvisionedCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/{accountId}/provisioned-capacity",
	}

	if input == nil {
		input = &types.PurchaseProvisionedCapacityInput{}
	}

	req := c.newRequest(op, input, &types.PurchaseProvisionedCapacityOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.PurchaseProvisionedCapacityMarshaler{Input: input}.GetNamedBuildHandler())

	return PurchaseProvisionedCapacityRequest{Request: req, Input: input, Copy: c.PurchaseProvisionedCapacityRequest}
}

// PurchaseProvisionedCapacityRequest is the request type for the
// PurchaseProvisionedCapacity API operation.
type PurchaseProvisionedCapacityRequest struct {
	*aws.Request
	Input *types.PurchaseProvisionedCapacityInput
	Copy  func(*types.PurchaseProvisionedCapacityInput) PurchaseProvisionedCapacityRequest
}

// Send marshals and sends the PurchaseProvisionedCapacity API request.
func (r PurchaseProvisionedCapacityRequest) Send(ctx context.Context) (*PurchaseProvisionedCapacityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseProvisionedCapacityResponse{
		PurchaseProvisionedCapacityOutput: r.Request.Data.(*types.PurchaseProvisionedCapacityOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseProvisionedCapacityResponse is the response type for the
// PurchaseProvisionedCapacity API operation.
type PurchaseProvisionedCapacityResponse struct {
	*types.PurchaseProvisionedCapacityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseProvisionedCapacity request.
func (r *PurchaseProvisionedCapacityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
