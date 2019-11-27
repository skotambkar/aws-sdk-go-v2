// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package budgets

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
)

const opCreateBudget = "CreateBudget"

// CreateBudgetRequest returns a request value for making API operation for
// AWS Budgets.
//
// Creates a budget and, if included, notifications and subscribers.
//
// Only one of BudgetLimit or PlannedBudgetLimits can be present in the syntax
// at one time. Use the syntax that matches your case. The Request Syntax section
// shows the BudgetLimit syntax. For PlannedBudgetLimits, see the Examples (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_budgets_CreateBudget.html#API_CreateBudget_Examples)
// section.
//
//    // Example sending a request using CreateBudgetRequest.
//    req := client.CreateBudgetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateBudgetRequest(input *types.CreateBudgetInput) CreateBudgetRequest {
	op := &aws.Operation{
		Name:       opCreateBudget,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateBudgetInput{}
	}

	req := c.newRequest(op, input, &types.CreateBudgetOutput{})
	return CreateBudgetRequest{Request: req, Input: input, Copy: c.CreateBudgetRequest}
}

// CreateBudgetRequest is the request type for the
// CreateBudget API operation.
type CreateBudgetRequest struct {
	*aws.Request
	Input *types.CreateBudgetInput
	Copy  func(*types.CreateBudgetInput) CreateBudgetRequest
}

// Send marshals and sends the CreateBudget API request.
func (r CreateBudgetRequest) Send(ctx context.Context) (*CreateBudgetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateBudgetResponse{
		CreateBudgetOutput: r.Request.Data.(*types.CreateBudgetOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateBudgetResponse is the response type for the
// CreateBudget API operation.
type CreateBudgetResponse struct {
	*types.CreateBudgetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateBudget request.
func (r *CreateBudgetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
