// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
)

const opDeleteTableVersion = "DeleteTableVersion"

// DeleteTableVersionRequest returns a request value for making API operation for
// AWS Glue.
//
// Deletes a specified version of a table.
//
//    // Example sending a request using DeleteTableVersionRequest.
//    req := client.DeleteTableVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/DeleteTableVersion
func (c *Client) DeleteTableVersionRequest(input *types.DeleteTableVersionInput) DeleteTableVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteTableVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteTableVersionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteTableVersionOutput{})
	return DeleteTableVersionRequest{Request: req, Input: input, Copy: c.DeleteTableVersionRequest}
}

// DeleteTableVersionRequest is the request type for the
// DeleteTableVersion API operation.
type DeleteTableVersionRequest struct {
	*aws.Request
	Input *types.DeleteTableVersionInput
	Copy  func(*types.DeleteTableVersionInput) DeleteTableVersionRequest
}

// Send marshals and sends the DeleteTableVersion API request.
func (r DeleteTableVersionRequest) Send(ctx context.Context) (*DeleteTableVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTableVersionResponse{
		DeleteTableVersionOutput: r.Request.Data.(*types.DeleteTableVersionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTableVersionResponse is the response type for the
// DeleteTableVersion API operation.
type DeleteTableVersionResponse struct {
	*types.DeleteTableVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTableVersion request.
func (r *DeleteTableVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
