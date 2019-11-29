// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opPurchaseScheduledInstances = "PurchaseScheduledInstances"

// PurchaseScheduledInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Purchases the Scheduled Instances with the specified schedule.
//
// Scheduled Instances enable you to purchase Amazon EC2 compute capacity by
// the hour for a one-year term. Before you can purchase a Scheduled Instance,
// you must call DescribeScheduledInstanceAvailability to check for available
// schedules and obtain a purchase token. After you purchase a Scheduled Instance,
// you must call RunScheduledInstances during each scheduled time period.
//
// After you purchase a Scheduled Instance, you can't cancel, modify, or resell
// your purchase.
//
//    // Example sending a request using PurchaseScheduledInstancesRequest.
//    req := client.PurchaseScheduledInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/PurchaseScheduledInstances
func (c *Client) PurchaseScheduledInstancesRequest(input *types.PurchaseScheduledInstancesInput) PurchaseScheduledInstancesRequest {
	op := &aws.Operation{
		Name:       opPurchaseScheduledInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PurchaseScheduledInstancesInput{}
	}

	req := c.newRequest(op, input, &types.PurchaseScheduledInstancesOutput{})
	return PurchaseScheduledInstancesRequest{Request: req, Input: input, Copy: c.PurchaseScheduledInstancesRequest}
}

// PurchaseScheduledInstancesRequest is the request type for the
// PurchaseScheduledInstances API operation.
type PurchaseScheduledInstancesRequest struct {
	*aws.Request
	Input *types.PurchaseScheduledInstancesInput
	Copy  func(*types.PurchaseScheduledInstancesInput) PurchaseScheduledInstancesRequest
}

// Send marshals and sends the PurchaseScheduledInstances API request.
func (r PurchaseScheduledInstancesRequest) Send(ctx context.Context) (*PurchaseScheduledInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PurchaseScheduledInstancesResponse{
		PurchaseScheduledInstancesOutput: r.Request.Data.(*types.PurchaseScheduledInstancesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PurchaseScheduledInstancesResponse is the response type for the
// PurchaseScheduledInstances API operation.
type PurchaseScheduledInstancesResponse struct {
	*types.PurchaseScheduledInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PurchaseScheduledInstances request.
func (r *PurchaseScheduledInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
