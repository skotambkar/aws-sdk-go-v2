// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opDeleteEmailTemplate = "DeleteEmailTemplate"

// DeleteEmailTemplateRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes a message template that was designed for use in messages that were
// sent through the email channel.
//
//    // Example sending a request using DeleteEmailTemplateRequest.
//    req := client.DeleteEmailTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteEmailTemplate
func (c *Client) DeleteEmailTemplateRequest(input *types.DeleteEmailTemplateInput) DeleteEmailTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteEmailTemplate,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/templates/{template-name}/email",
	}

	if input == nil {
		input = &types.DeleteEmailTemplateInput{}
	}

	req := c.newRequest(op, input, &types.DeleteEmailTemplateOutput{})
	return DeleteEmailTemplateRequest{Request: req, Input: input, Copy: c.DeleteEmailTemplateRequest}
}

// DeleteEmailTemplateRequest is the request type for the
// DeleteEmailTemplate API operation.
type DeleteEmailTemplateRequest struct {
	*aws.Request
	Input *types.DeleteEmailTemplateInput
	Copy  func(*types.DeleteEmailTemplateInput) DeleteEmailTemplateRequest
}

// Send marshals and sends the DeleteEmailTemplate API request.
func (r DeleteEmailTemplateRequest) Send(ctx context.Context) (*DeleteEmailTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEmailTemplateResponse{
		DeleteEmailTemplateOutput: r.Request.Data.(*types.DeleteEmailTemplateOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEmailTemplateResponse is the response type for the
// DeleteEmailTemplate API operation.
type DeleteEmailTemplateResponse struct {
	*types.DeleteEmailTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEmailTemplate request.
func (r *DeleteEmailTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
