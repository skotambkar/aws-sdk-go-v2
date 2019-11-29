// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opPutConformancePack = "PutConformancePack"

// PutConformancePackRequest returns a request value for making API operation for
// AWS Config.
//
// Creates or updates a conformance pack. A conformance pack is a collection
// of AWS Config rules that can be easily deployed in an account and a region.
//
// This API creates a service linked role AWSServiceRoleForConfigConforms in
// your account. The service linked role is created only when the role does
// not exist in your account. AWS Config verifies the existence of role with
// GetRole action.
//
// You must specify either the TemplateS3Uri or the TemplateBody parameter,
// but not both. If you provide both AWS Config uses the TemplateS3Uri parameter
// and ignores the TemplateBody parameter.
//
//    // Example sending a request using PutConformancePackRequest.
//    req := client.PutConformancePackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/PutConformancePack
func (c *Client) PutConformancePackRequest(input *types.PutConformancePackInput) PutConformancePackRequest {
	op := &aws.Operation{
		Name:       opPutConformancePack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutConformancePackInput{}
	}

	req := c.newRequest(op, input, &types.PutConformancePackOutput{})
	return PutConformancePackRequest{Request: req, Input: input, Copy: c.PutConformancePackRequest}
}

// PutConformancePackRequest is the request type for the
// PutConformancePack API operation.
type PutConformancePackRequest struct {
	*aws.Request
	Input *types.PutConformancePackInput
	Copy  func(*types.PutConformancePackInput) PutConformancePackRequest
}

// Send marshals and sends the PutConformancePack API request.
func (r PutConformancePackRequest) Send(ctx context.Context) (*PutConformancePackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutConformancePackResponse{
		PutConformancePackOutput: r.Request.Data.(*types.PutConformancePackOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutConformancePackResponse is the response type for the
// PutConformancePack API operation.
type PutConformancePackResponse struct {
	*types.PutConformancePackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutConformancePack request.
func (r *PutConformancePackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
