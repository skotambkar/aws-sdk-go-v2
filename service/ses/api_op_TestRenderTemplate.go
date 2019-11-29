// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const opTestRenderTemplate = "TestRenderTemplate"

// TestRenderTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Creates a preview of the MIME content of an email when provided with a template
// and a set of replacement data.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using TestRenderTemplateRequest.
//    req := client.TestRenderTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/TestRenderTemplate
func (c *Client) TestRenderTemplateRequest(input *types.TestRenderTemplateInput) TestRenderTemplateRequest {
	op := &aws.Operation{
		Name:       opTestRenderTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.TestRenderTemplateInput{}
	}

	req := c.newRequest(op, input, &types.TestRenderTemplateOutput{})
	return TestRenderTemplateRequest{Request: req, Input: input, Copy: c.TestRenderTemplateRequest}
}

// TestRenderTemplateRequest is the request type for the
// TestRenderTemplate API operation.
type TestRenderTemplateRequest struct {
	*aws.Request
	Input *types.TestRenderTemplateInput
	Copy  func(*types.TestRenderTemplateInput) TestRenderTemplateRequest
}

// Send marshals and sends the TestRenderTemplate API request.
func (r TestRenderTemplateRequest) Send(ctx context.Context) (*TestRenderTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TestRenderTemplateResponse{
		TestRenderTemplateOutput: r.Request.Data.(*types.TestRenderTemplateOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TestRenderTemplateResponse is the response type for the
// TestRenderTemplate API operation.
type TestRenderTemplateResponse struct {
	*types.TestRenderTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TestRenderTemplate request.
func (r *TestRenderTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
