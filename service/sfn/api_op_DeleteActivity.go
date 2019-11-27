// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfn

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
)

const opDeleteActivity = "DeleteActivity"

// DeleteActivityRequest returns a request value for making API operation for
// AWS Step Functions.
//
// Deletes an activity.
//
//    // Example sending a request using DeleteActivityRequest.
//    req := client.DeleteActivityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/states-2016-11-23/DeleteActivity
func (c *Client) DeleteActivityRequest(input *types.DeleteActivityInput) DeleteActivityRequest {
	op := &aws.Operation{
		Name:       opDeleteActivity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteActivityInput{}
	}

	req := c.newRequest(op, input, &types.DeleteActivityOutput{})
	return DeleteActivityRequest{Request: req, Input: input, Copy: c.DeleteActivityRequest}
}

// DeleteActivityRequest is the request type for the
// DeleteActivity API operation.
type DeleteActivityRequest struct {
	*aws.Request
	Input *types.DeleteActivityInput
	Copy  func(*types.DeleteActivityInput) DeleteActivityRequest
}

// Send marshals and sends the DeleteActivity API request.
func (r DeleteActivityRequest) Send(ctx context.Context) (*DeleteActivityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteActivityResponse{
		DeleteActivityOutput: r.Request.Data.(*types.DeleteActivityOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteActivityResponse is the response type for the
// DeleteActivity API operation.
type DeleteActivityResponse struct {
	*types.DeleteActivityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteActivity request.
func (r *DeleteActivityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
