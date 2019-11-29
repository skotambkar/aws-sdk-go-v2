// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
)

const opPutServiceQuotaIncreaseRequestIntoTemplate = "PutServiceQuotaIncreaseRequestIntoTemplate"

// PutServiceQuotaIncreaseRequestIntoTemplateRequest returns a request value for making API operation for
// Service Quotas.
//
// Defines and adds a quota to the service quota template. To add a quota to
// the template, you must provide the ServiceCode, QuotaCode, AwsRegion, and
// DesiredValue. Once you add a quota to the template, use ListServiceQuotaIncreaseRequestsInTemplate
// to see the list of quotas in the template.
//
//    // Example sending a request using PutServiceQuotaIncreaseRequestIntoTemplateRequest.
//    req := client.PutServiceQuotaIncreaseRequestIntoTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/service-quotas-2019-06-24/PutServiceQuotaIncreaseRequestIntoTemplate
func (c *Client) PutServiceQuotaIncreaseRequestIntoTemplateRequest(input *types.PutServiceQuotaIncreaseRequestIntoTemplateInput) PutServiceQuotaIncreaseRequestIntoTemplateRequest {
	op := &aws.Operation{
		Name:       opPutServiceQuotaIncreaseRequestIntoTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutServiceQuotaIncreaseRequestIntoTemplateInput{}
	}

	req := c.newRequest(op, input, &types.PutServiceQuotaIncreaseRequestIntoTemplateOutput{})
	return PutServiceQuotaIncreaseRequestIntoTemplateRequest{Request: req, Input: input, Copy: c.PutServiceQuotaIncreaseRequestIntoTemplateRequest}
}

// PutServiceQuotaIncreaseRequestIntoTemplateRequest is the request type for the
// PutServiceQuotaIncreaseRequestIntoTemplate API operation.
type PutServiceQuotaIncreaseRequestIntoTemplateRequest struct {
	*aws.Request
	Input *types.PutServiceQuotaIncreaseRequestIntoTemplateInput
	Copy  func(*types.PutServiceQuotaIncreaseRequestIntoTemplateInput) PutServiceQuotaIncreaseRequestIntoTemplateRequest
}

// Send marshals and sends the PutServiceQuotaIncreaseRequestIntoTemplate API request.
func (r PutServiceQuotaIncreaseRequestIntoTemplateRequest) Send(ctx context.Context) (*PutServiceQuotaIncreaseRequestIntoTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutServiceQuotaIncreaseRequestIntoTemplateResponse{
		PutServiceQuotaIncreaseRequestIntoTemplateOutput: r.Request.Data.(*types.PutServiceQuotaIncreaseRequestIntoTemplateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutServiceQuotaIncreaseRequestIntoTemplateResponse is the response type for the
// PutServiceQuotaIncreaseRequestIntoTemplate API operation.
type PutServiceQuotaIncreaseRequestIntoTemplateResponse struct {
	*types.PutServiceQuotaIncreaseRequestIntoTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutServiceQuotaIncreaseRequestIntoTemplate request.
func (r *PutServiceQuotaIncreaseRequestIntoTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
