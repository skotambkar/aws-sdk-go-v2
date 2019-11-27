// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
)

const opDeleteMLModel = "DeleteMLModel"

// DeleteMLModelRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Assigns the DELETED status to an MLModel, rendering it unusable.
//
// After using the DeleteMLModel operation, you can use the GetMLModel operation
// to verify that the status of the MLModel changed to DELETED.
//
// Caution: The result of the DeleteMLModel operation is irreversible.
//
//    // Example sending a request using DeleteMLModelRequest.
//    req := client.DeleteMLModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteMLModelRequest(input *types.DeleteMLModelInput) DeleteMLModelRequest {
	op := &aws.Operation{
		Name:       opDeleteMLModel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteMLModelInput{}
	}

	req := c.newRequest(op, input, &types.DeleteMLModelOutput{})
	return DeleteMLModelRequest{Request: req, Input: input, Copy: c.DeleteMLModelRequest}
}

// DeleteMLModelRequest is the request type for the
// DeleteMLModel API operation.
type DeleteMLModelRequest struct {
	*aws.Request
	Input *types.DeleteMLModelInput
	Copy  func(*types.DeleteMLModelInput) DeleteMLModelRequest
}

// Send marshals and sends the DeleteMLModel API request.
func (r DeleteMLModelRequest) Send(ctx context.Context) (*DeleteMLModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteMLModelResponse{
		DeleteMLModelOutput: r.Request.Data.(*types.DeleteMLModelOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteMLModelResponse is the response type for the
// DeleteMLModel API operation.
type DeleteMLModelResponse struct {
	*types.DeleteMLModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteMLModel request.
func (r *DeleteMLModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
