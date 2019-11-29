// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
)

const opGetDomainStatisticsReport = "GetDomainStatisticsReport"

// GetDomainStatisticsReportRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Retrieve inbox placement and engagement rates for the domains that you use
// to send email.
//
//    // Example sending a request using GetDomainStatisticsReportRequest.
//    req := client.GetDomainStatisticsReportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDomainStatisticsReport
func (c *Client) GetDomainStatisticsReportRequest(input *types.GetDomainStatisticsReportInput) GetDomainStatisticsReportRequest {
	op := &aws.Operation{
		Name:       opGetDomainStatisticsReport,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/deliverability-dashboard/statistics-report/{Domain}",
	}

	if input == nil {
		input = &types.GetDomainStatisticsReportInput{}
	}

	req := c.newRequest(op, input, &types.GetDomainStatisticsReportOutput{})
	return GetDomainStatisticsReportRequest{Request: req, Input: input, Copy: c.GetDomainStatisticsReportRequest}
}

// GetDomainStatisticsReportRequest is the request type for the
// GetDomainStatisticsReport API operation.
type GetDomainStatisticsReportRequest struct {
	*aws.Request
	Input *types.GetDomainStatisticsReportInput
	Copy  func(*types.GetDomainStatisticsReportInput) GetDomainStatisticsReportRequest
}

// Send marshals and sends the GetDomainStatisticsReport API request.
func (r GetDomainStatisticsReportRequest) Send(ctx context.Context) (*GetDomainStatisticsReportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainStatisticsReportResponse{
		GetDomainStatisticsReportOutput: r.Request.Data.(*types.GetDomainStatisticsReportOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainStatisticsReportResponse is the response type for the
// GetDomainStatisticsReport API operation.
type GetDomainStatisticsReportResponse struct {
	*types.GetDomainStatisticsReportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainStatisticsReport request.
func (r *GetDomainStatisticsReportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
