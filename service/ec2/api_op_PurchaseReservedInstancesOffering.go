// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opPurchaseReservedInstancesOffering = "PurchaseReservedInstancesOffering"

// PurchaseReservedInstancesOfferingRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Purchases a Reserved Instance for use with your account. With Reserved Instances,
// you pay a lower hourly rate compared to On-Demand instance pricing.
//
// Use DescribeReservedInstancesOfferings to get a list of Reserved Instance
// offerings that match your specifications. After you've purchased a Reserved
// Instance, you can check for your new Reserved Instance with DescribeReservedInstances.
//
// To queue a purchase for a future date and time, specify a purchase time.
// If you do not specify a purchase time, the default is the current time.
//
// For more information, see Reserved Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/concepts-on-demand-reserved-instances.html)
// and Reserved Instance Marketplace (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using PurchaseReservedInstancesOfferingRequest.
//    req := client.PurchaseReservedInstancesOfferingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseReservedInstancesOffering
func (c *Client) PurchaseReservedInstancesOfferingRequest(input *types.PurchaseReservedInstancesOfferingInput) PurchaseReservedInstancesOfferingRequest {
	op := &aws.Operation{
		Name:       opPurchaseReservedInstancesOffering,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PurchaseReservedInstancesOfferingInput{}
	}

	req := c.newRequest(op, input, &types.PurchaseReservedInstancesOfferingOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.PurchaseReservedInstancesOfferingMarshaler{Input: input}.GetNamedBuildHandler())

	return PurchaseReservedInstancesOfferingRequest{Request: req, Input: input, Copy: c.PurchaseReservedInstancesOfferingRequest}
}

// PurchaseReservedInstancesOfferingRequest is the request type for the
// PurchaseReservedInstancesOffering API operation.
type PurchaseReservedInstancesOfferingRequest struct {
	*aws.Request
	Input *types.PurchaseReservedInstancesOfferingInput
	Copy  func(*types.PurchaseReservedInstancesOfferingInput) PurchaseReservedInstancesOfferingRequest
}

// Send marshals and sends the PurchaseReservedInstancesOffering API request.
func (r PurchaseReservedInstancesOfferingRequest) Send(ctx context.Context) (*PurchaseReservedInstancesOfferingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseReservedInstancesOfferingResponse{
		PurchaseReservedInstancesOfferingOutput: r.Request.Data.(*types.PurchaseReservedInstancesOfferingOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseReservedInstancesOfferingResponse is the response type for the
// PurchaseReservedInstancesOffering API operation.
type PurchaseReservedInstancesOfferingResponse struct {
	*types.PurchaseReservedInstancesOfferingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseReservedInstancesOffering request.
func (r *PurchaseReservedInstancesOfferingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
