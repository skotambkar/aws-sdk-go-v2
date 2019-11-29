// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package inspector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
)

const opCreateAssessmentTemplate = "CreateAssessmentTemplate"

// CreateAssessmentTemplateRequest returns a request value for making API operation for
// Amazon Inspector.
//
// Creates an assessment template for the assessment target that is specified
// by the ARN of the assessment target. If the service-linked role (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html)
// isn’t already registered, this action also creates and registers a service-linked
// role to grant Amazon Inspector access to AWS Services needed to perform security
// assessments.
//
//    // Example sending a request using CreateAssessmentTemplateRequest.
//    req := client.CreateAssessmentTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/inspector-2016-02-16/CreateAssessmentTemplate
func (c *Client) CreateAssessmentTemplateRequest(input *types.CreateAssessmentTemplateInput) CreateAssessmentTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateAssessmentTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateAssessmentTemplateInput{}
	}

	req := c.newRequest(op, input, &types.CreateAssessmentTemplateOutput{})
	return CreateAssessmentTemplateRequest{Request: req, Input: input, Copy: c.CreateAssessmentTemplateRequest}
}

// CreateAssessmentTemplateRequest is the request type for the
// CreateAssessmentTemplate API operation.
type CreateAssessmentTemplateRequest struct {
	*aws.Request
	Input *types.CreateAssessmentTemplateInput
	Copy  func(*types.CreateAssessmentTemplateInput) CreateAssessmentTemplateRequest
}

// Send marshals and sends the CreateAssessmentTemplate API request.
func (r CreateAssessmentTemplateRequest) Send(ctx context.Context) (*CreateAssessmentTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAssessmentTemplateResponse{
		CreateAssessmentTemplateOutput: r.Request.Data.(*types.CreateAssessmentTemplateOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAssessmentTemplateResponse is the response type for the
// CreateAssessmentTemplate API operation.
type CreateAssessmentTemplateResponse struct {
	*types.CreateAssessmentTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateAssessmentTemplate request.
func (r *CreateAssessmentTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
