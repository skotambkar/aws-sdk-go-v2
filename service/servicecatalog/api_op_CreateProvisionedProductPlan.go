// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
)

const opCreateProvisionedProductPlan = "CreateProvisionedProductPlan"

// CreateProvisionedProductPlanRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a plan. A plan includes the list of resources to be created (when
// provisioning a new product) or modified (when updating a provisioned product)
// when the plan is executed.
//
// You can create one plan per provisioned product. To create a plan for an
// existing provisioned product, the product status must be AVAILBLE or TAINTED.
//
// To view the resource changes in the change set, use DescribeProvisionedProductPlan.
// To create or modify the provisioned product, use ExecuteProvisionedProductPlan.
//
//    // Example sending a request using CreateProvisionedProductPlanRequest.
//    req := client.CreateProvisionedProductPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreateProvisionedProductPlan
func (c *Client) CreateProvisionedProductPlanRequest(input *types.CreateProvisionedProductPlanInput) CreateProvisionedProductPlanRequest {
	op := &aws.Operation{
		Name:       opCreateProvisionedProductPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateProvisionedProductPlanInput{}
	}

	req := c.newRequest(op, input, &types.CreateProvisionedProductPlanOutput{})
	return CreateProvisionedProductPlanRequest{Request: req, Input: input, Copy: c.CreateProvisionedProductPlanRequest}
}

// CreateProvisionedProductPlanRequest is the request type for the
// CreateProvisionedProductPlan API operation.
type CreateProvisionedProductPlanRequest struct {
	*aws.Request
	Input *types.CreateProvisionedProductPlanInput
	Copy  func(*types.CreateProvisionedProductPlanInput) CreateProvisionedProductPlanRequest
}

// Send marshals and sends the CreateProvisionedProductPlan API request.
func (r CreateProvisionedProductPlanRequest) Send(ctx context.Context) (*CreateProvisionedProductPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateProvisionedProductPlanResponse{
		CreateProvisionedProductPlanOutput: r.Request.Data.(*types.CreateProvisionedProductPlanOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateProvisionedProductPlanResponse is the response type for the
// CreateProvisionedProductPlan API operation.
type CreateProvisionedProductPlanResponse struct {
	*types.CreateProvisionedProductPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateProvisionedProductPlan request.
func (r *CreateProvisionedProductPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
