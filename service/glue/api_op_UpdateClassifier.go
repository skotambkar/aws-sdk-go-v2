// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opUpdateClassifier = "UpdateClassifier"

// UpdateClassifierRequest returns a request value for making API operation for
// AWS Glue.
//
// Modifies an existing classifier (a GrokClassifier, an XMLClassifier, a JsonClassifier,
// or a CsvClassifier, depending on which field is present).
//
//    // Example sending a request using UpdateClassifierRequest.
//    req := client.UpdateClassifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/UpdateClassifier
func (c *Client) UpdateClassifierRequest(input *types.UpdateClassifierInput) UpdateClassifierRequest {
	op := &aws.Operation{
		Name:       opUpdateClassifier,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateClassifierInput{}
	}

	req := c.newRequest(op, input, &types.UpdateClassifierOutput{})
	return UpdateClassifierRequest{Request: req, Input: input, Copy: c.UpdateClassifierRequest}
}

// UpdateClassifierRequest is the request type for the
// UpdateClassifier API operation.
type UpdateClassifierRequest struct {
	*aws.Request
	Input *types.UpdateClassifierInput
	Copy  func(*types.UpdateClassifierInput) UpdateClassifierRequest
}

// Send marshals and sends the UpdateClassifier API request.
func (r UpdateClassifierRequest) Send(ctx context.Context) (*UpdateClassifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClassifierResponse{
		UpdateClassifierOutput: r.Request.Data.(*types.UpdateClassifierOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClassifierResponse is the response type for the
// UpdateClassifier API operation.
type UpdateClassifierResponse struct {
	*types.UpdateClassifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClassifier request.
func (r *UpdateClassifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
