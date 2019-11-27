// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package acmpca

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
)

const opDescribeCertificateAuthorityAuditReport = "DescribeCertificateAuthorityAuditReport"

// DescribeCertificateAuthorityAuditReportRequest returns a request value for making API operation for
// AWS Certificate Manager Private Certificate Authority.
//
// Lists information about a specific audit report created by calling the CreateCertificateAuthorityAuditReport
// action. Audit information is created every time the certificate authority
// (CA) private key is used. The private key is used when you call the IssueCertificate
// action or the RevokeCertificate action.
//
//    // Example sending a request using DescribeCertificateAuthorityAuditReportRequest.
//    req := client.DescribeCertificateAuthorityAuditReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/acm-pca-2017-08-22/DescribeCertificateAuthorityAuditReport
func (c *Client) DescribeCertificateAuthorityAuditReportRequest(input *types.DescribeCertificateAuthorityAuditReportInput) DescribeCertificateAuthorityAuditReportRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificateAuthorityAuditReport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeCertificateAuthorityAuditReportInput{}
	}

	req := c.newRequest(op, input, &types.DescribeCertificateAuthorityAuditReportOutput{})
	return DescribeCertificateAuthorityAuditReportRequest{Request: req, Input: input, Copy: c.DescribeCertificateAuthorityAuditReportRequest}
}

// DescribeCertificateAuthorityAuditReportRequest is the request type for the
// DescribeCertificateAuthorityAuditReport API operation.
type DescribeCertificateAuthorityAuditReportRequest struct {
	*aws.Request
	Input *types.DescribeCertificateAuthorityAuditReportInput
	Copy  func(*types.DescribeCertificateAuthorityAuditReportInput) DescribeCertificateAuthorityAuditReportRequest
}

// Send marshals and sends the DescribeCertificateAuthorityAuditReport API request.
func (r DescribeCertificateAuthorityAuditReportRequest) Send(ctx context.Context) (*DescribeCertificateAuthorityAuditReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificateAuthorityAuditReportResponse{
		DescribeCertificateAuthorityAuditReportOutput: r.Request.Data.(*types.DescribeCertificateAuthorityAuditReportOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificateAuthorityAuditReportResponse is the response type for the
// DescribeCertificateAuthorityAuditReport API operation.
type DescribeCertificateAuthorityAuditReportResponse struct {
	*types.DescribeCertificateAuthorityAuditReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificateAuthorityAuditReport request.
func (r *DescribeCertificateAuthorityAuditReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
