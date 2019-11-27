// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
)

const opDescribeEmergencyContactSettings = "DescribeEmergencyContactSettings"

// DescribeEmergencyContactSettingsRequest returns a request value for making API operation for
// AWS Shield.
//
// Lists the email addresses that the DRT can use to contact you during a suspected
// attack.
//
//    // Example sending a request using DescribeEmergencyContactSettingsRequest.
//    req := client.DescribeEmergencyContactSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeEmergencyContactSettings
func (c *Client) DescribeEmergencyContactSettingsRequest(input *types.DescribeEmergencyContactSettingsInput) DescribeEmergencyContactSettingsRequest {
	op := &aws.Operation{
		Name:       opDescribeEmergencyContactSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeEmergencyContactSettingsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeEmergencyContactSettingsOutput{})
	return DescribeEmergencyContactSettingsRequest{Request: req, Input: input, Copy: c.DescribeEmergencyContactSettingsRequest}
}

// DescribeEmergencyContactSettingsRequest is the request type for the
// DescribeEmergencyContactSettings API operation.
type DescribeEmergencyContactSettingsRequest struct {
	*aws.Request
	Input *types.DescribeEmergencyContactSettingsInput
	Copy  func(*types.DescribeEmergencyContactSettingsInput) DescribeEmergencyContactSettingsRequest
}

// Send marshals and sends the DescribeEmergencyContactSettings API request.
func (r DescribeEmergencyContactSettingsRequest) Send(ctx context.Context) (*DescribeEmergencyContactSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEmergencyContactSettingsResponse{
		DescribeEmergencyContactSettingsOutput: r.Request.Data.(*types.DescribeEmergencyContactSettingsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEmergencyContactSettingsResponse is the response type for the
// DescribeEmergencyContactSettings API operation.
type DescribeEmergencyContactSettingsResponse struct {
	*types.DescribeEmergencyContactSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEmergencyContactSettings request.
func (r *DescribeEmergencyContactSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
