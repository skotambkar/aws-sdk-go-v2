// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
)

const opAssociateServiceQuotaTemplate = "AssociateServiceQuotaTemplate"

// AssociateServiceQuotaTemplateRequest returns a request value for making API operation for
// Service Quotas.
//
// Associates the Service Quotas template with your organization so that when
// new accounts are created in your organization, the template submits increase
// requests for the specified service quotas. Use the Service Quotas template
// to request an increase for any adjustable quota value. After you define the
// Service Quotas template, use this operation to associate, or enable, the
// template.
//
//    // Example sending a request using AssociateServiceQuotaTemplateRequest.
//    req := client.AssociateServiceQuotaTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/AssociateServiceQuotaTemplate
func (c *Client) AssociateServiceQuotaTemplateRequest(input *types.AssociateServiceQuotaTemplateInput) AssociateServiceQuotaTemplateRequest {
	op := &aws.Operation{
		Name:       opAssociateServiceQuotaTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateServiceQuotaTemplateInput{}
	}

	req := c.newRequest(op, input, &types.AssociateServiceQuotaTemplateOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateServiceQuotaTemplateMarshaler{Input: input}.GetNamedBuildHandler())

	return AssociateServiceQuotaTemplateRequest{Request: req, Input: input, Copy: c.AssociateServiceQuotaTemplateRequest}
}

// AssociateServiceQuotaTemplateRequest is the request type for the
// AssociateServiceQuotaTemplate API operation.
type AssociateServiceQuotaTemplateRequest struct {
	*aws.Request
	Input *types.AssociateServiceQuotaTemplateInput
	Copy  func(*types.AssociateServiceQuotaTemplateInput) AssociateServiceQuotaTemplateRequest
}

// Send marshals and sends the AssociateServiceQuotaTemplate API request.
func (r AssociateServiceQuotaTemplateRequest) Send(ctx context.Context) (*AssociateServiceQuotaTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateServiceQuotaTemplateResponse{
		AssociateServiceQuotaTemplateOutput: r.Request.Data.(*types.AssociateServiceQuotaTemplateOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateServiceQuotaTemplateResponse is the response type for the
// AssociateServiceQuotaTemplate API operation.
type AssociateServiceQuotaTemplateResponse struct {
	*types.AssociateServiceQuotaTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateServiceQuotaTemplate request.
func (r *AssociateServiceQuotaTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
