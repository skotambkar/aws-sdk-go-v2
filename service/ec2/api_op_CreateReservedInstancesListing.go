// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateReservedInstancesListing = "CreateReservedInstancesListing"

// CreateReservedInstancesListingRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates a listing for Amazon EC2 Standard Reserved Instances to be sold in
// the Reserved Instance Marketplace. You can submit one Standard Reserved Instance
// listing at a time. To get a list of your Standard Reserved Instances, you
// can use the DescribeReservedInstances operation.
//
// Only Standard Reserved Instances can be sold in the Reserved Instance Marketplace.
// Convertible Reserved Instances cannot be sold.
//
// The Reserved Instance Marketplace matches sellers who want to resell Standard
// Reserved Instance capacity that they no longer need with buyers who want
// to purchase additional capacity. Reserved Instances bought and sold through
// the Reserved Instance Marketplace work like any other Reserved Instances.
//
// To sell your Standard Reserved Instances, you must first register as a seller
// in the Reserved Instance Marketplace. After completing the registration process,
// you can create a Reserved Instance Marketplace listing of some or all of
// your Standard Reserved Instances, and specify the upfront price to receive
// for them. Your Standard Reserved Instance listings then become available
// for purchase. To view the details of your Standard Reserved Instance listing,
// you can use the DescribeReservedInstancesListings operation.
//
// For more information, see Reserved Instance Marketplace (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ri-market-general.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using CreateReservedInstancesListingRequest.
//    req := client.CreateReservedInstancesListingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateReservedInstancesListing
func (c *Client) CreateReservedInstancesListingRequest(input *types.CreateReservedInstancesListingInput) CreateReservedInstancesListingRequest {
	op := &aws.Operation{
		Name:       opCreateReservedInstancesListing,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateReservedInstancesListingInput{}
	}

	req := c.newRequest(op, input, &types.CreateReservedInstancesListingOutput{})
	return CreateReservedInstancesListingRequest{Request: req, Input: input, Copy: c.CreateReservedInstancesListingRequest}
}

// CreateReservedInstancesListingRequest is the request type for the
// CreateReservedInstancesListing API operation.
type CreateReservedInstancesListingRequest struct {
	*aws.Request
	Input *types.CreateReservedInstancesListingInput
	Copy  func(*types.CreateReservedInstancesListingInput) CreateReservedInstancesListingRequest
}

// Send marshals and sends the CreateReservedInstancesListing API request.
func (r CreateReservedInstancesListingRequest) Send(ctx context.Context) (*CreateReservedInstancesListingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReservedInstancesListingResponse{
		CreateReservedInstancesListingOutput: r.Request.Data.(*types.CreateReservedInstancesListingOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReservedInstancesListingResponse is the response type for the
// CreateReservedInstancesListing API operation.
type CreateReservedInstancesListingResponse struct {
	*types.CreateReservedInstancesListingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReservedInstancesListing request.
func (r *CreateReservedInstancesListingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
