// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscalingplans

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
)

const opCreateScalingPlan = "CreateScalingPlan"

// CreateScalingPlanRequest returns a request value for making API operation for
// AWS Auto Scaling Plans.
//
// Creates a scaling plan.
//
//    // Example sending a request using CreateScalingPlanRequest.
//    req := client.CreateScalingPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-plans-2018-01-06/CreateScalingPlan
func (c *Client) CreateScalingPlanRequest(input *types.CreateScalingPlanInput) CreateScalingPlanRequest {
	op := &aws.Operation{
		Name:       opCreateScalingPlan,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateScalingPlanInput{}
	}

	req := c.newRequest(op, input, &types.CreateScalingPlanOutput{})
	return CreateScalingPlanRequest{Request: req, Input: input, Copy: c.CreateScalingPlanRequest}
}

// CreateScalingPlanRequest is the request type for the
// CreateScalingPlan API operation.
type CreateScalingPlanRequest struct {
	*aws.Request
	Input *types.CreateScalingPlanInput
	Copy  func(*types.CreateScalingPlanInput) CreateScalingPlanRequest
}

// Send marshals and sends the CreateScalingPlan API request.
func (r CreateScalingPlanRequest) Send(ctx context.Context) (*CreateScalingPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateScalingPlanResponse{
		CreateScalingPlanOutput: r.Request.Data.(*types.CreateScalingPlanOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateScalingPlanResponse is the response type for the
// CreateScalingPlan API operation.
type CreateScalingPlanResponse struct {
	*types.CreateScalingPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateScalingPlan request.
func (r *CreateScalingPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
