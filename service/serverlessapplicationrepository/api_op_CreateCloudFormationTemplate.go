// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository/types"
)

const opCreateCloudFormationTemplate = "CreateCloudFormationTemplate"

// CreateCloudFormationTemplateRequest returns a request value for making API operation for
// AWSServerlessApplicationRepository.
//
// Creates an AWS CloudFormation template.
//
//    // Example sending a request using CreateCloudFormationTemplateRequest.
//    req := client.CreateCloudFormationTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/serverlessrepo-2017-09-08/CreateCloudFormationTemplate
func (c *Client) CreateCloudFormationTemplateRequest(input *types.CreateCloudFormationTemplateInput) CreateCloudFormationTemplateRequest {
	op := &aws.Operation{
		Name:       opCreateCloudFormationTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/applications/{applicationId}/templates",
	}

	if input == nil {
		input = &types.CreateCloudFormationTemplateInput{}
	}

	req := c.newRequest(op, input, &types.CreateCloudFormationTemplateOutput{})
	return CreateCloudFormationTemplateRequest{Request: req, Input: input, Copy: c.CreateCloudFormationTemplateRequest}
}

// CreateCloudFormationTemplateRequest is the request type for the
// CreateCloudFormationTemplate API operation.
type CreateCloudFormationTemplateRequest struct {
	*aws.Request
	Input *types.CreateCloudFormationTemplateInput
	Copy  func(*types.CreateCloudFormationTemplateInput) CreateCloudFormationTemplateRequest
}

// Send marshals and sends the CreateCloudFormationTemplate API request.
func (r CreateCloudFormationTemplateRequest) Send(ctx context.Context) (*CreateCloudFormationTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCloudFormationTemplateResponse{
		CreateCloudFormationTemplateOutput: r.Request.Data.(*types.CreateCloudFormationTemplateOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCloudFormationTemplateResponse is the response type for the
// CreateCloudFormationTemplate API operation.
type CreateCloudFormationTemplateResponse struct {
	*types.CreateCloudFormationTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCloudFormationTemplate request.
func (r *CreateCloudFormationTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
