// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

const opGetCredentialReport = "GetCredentialReport"

// GetCredentialReportRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Retrieves a credential report for the AWS account. For more information about
// the credential report, see Getting Credential Reports (https://docs.aws.amazon.com/IAM/latest/UserGuide/credential-reports.html)
// in the IAM User Guide.
//
//    // Example sending a request using GetCredentialReportRequest.
//    req := client.GetCredentialReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/GetCredentialReport
func (c *Client) GetCredentialReportRequest(input *types.GetCredentialReportInput) GetCredentialReportRequest {
	op := &aws.Operation{
		Name:       opGetCredentialReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetCredentialReportInput{}
	}

	req := c.newRequest(op, input, &types.GetCredentialReportOutput{})
	return GetCredentialReportRequest{Request: req, Input: input, Copy: c.GetCredentialReportRequest}
}

// GetCredentialReportRequest is the request type for the
// GetCredentialReport API operation.
type GetCredentialReportRequest struct {
	*aws.Request
	Input *types.GetCredentialReportInput
	Copy  func(*types.GetCredentialReportInput) GetCredentialReportRequest
}

// Send marshals and sends the GetCredentialReport API request.
func (r GetCredentialReportRequest) Send(ctx context.Context) (*GetCredentialReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCredentialReportResponse{
		GetCredentialReportOutput: r.Request.Data.(*types.GetCredentialReportOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetCredentialReportResponse is the response type for the
// GetCredentialReport API operation.
type GetCredentialReportResponse struct {
	*types.GetCredentialReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCredentialReport request.
func (r *GetCredentialReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
