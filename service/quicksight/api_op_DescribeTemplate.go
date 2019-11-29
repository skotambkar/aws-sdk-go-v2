// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeTemplate = "DescribeTemplate"

// DescribeTemplateRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes a template's metadata.
//
// CLI syntax:
//
// aws quicksight describe-template --aws-account-id 111122223333 --template-id
// reports_test_template
//
// aws quicksight describe-template --aws-account-id 111122223333 --template-id
// reports_test_template --version-number-2
//
// aws quicksight describe-template --aws-account-id 111122223333 --template-id
// reports_test_template --alias-name '\$LATEST'
//
// Users can explicitly describe the latest version of the dashboard by passing
// $LATEST to the alias-name parameter. $LATEST is an internally supported alias,
// which points to the latest version of the dashboard.
//
//    // Example sending a request using DescribeTemplateRequest.
//    req := client.DescribeTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeTemplate
func (c *Client) DescribeTemplateRequest(input *types.DescribeTemplateInput) DescribeTemplateRequest {
	op := &aws.Operation{
		Name:       opDescribeTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}",
	}

	if input == nil {
		input = &types.DescribeTemplateInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTemplateOutput{})
	return DescribeTemplateRequest{Request: req, Input: input, Copy: c.DescribeTemplateRequest}
}

// DescribeTemplateRequest is the request type for the
// DescribeTemplate API operation.
type DescribeTemplateRequest struct {
	*aws.Request
	Input *types.DescribeTemplateInput
	Copy  func(*types.DescribeTemplateInput) DescribeTemplateRequest
}

// Send marshals and sends the DescribeTemplate API request.
func (r DescribeTemplateRequest) Send(ctx context.Context) (*DescribeTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTemplateResponse{
		DescribeTemplateOutput: r.Request.Data.(*types.DescribeTemplateOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTemplateResponse is the response type for the
// DescribeTemplate API operation.
type DescribeTemplateResponse struct {
	*types.DescribeTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTemplate request.
func (r *DescribeTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
