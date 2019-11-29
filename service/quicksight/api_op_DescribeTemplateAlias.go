// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeTemplateAlias = "DescribeTemplateAlias"

// DescribeTemplateAliasRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes the template aliases of a template.
//
// CLI syntax:
//
// aws quicksight describe-template-alias --aws-account-id 111122223333 --template-id
// 'reports_test_template' --alias-name 'STAGING'
//
//    // Example sending a request using DescribeTemplateAliasRequest.
//    req := client.DescribeTemplateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeTemplateAlias
func (c *Client) DescribeTemplateAliasRequest(input *types.DescribeTemplateAliasInput) DescribeTemplateAliasRequest {
	op := &aws.Operation{
		Name:       opDescribeTemplateAlias,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/aliases/{AliasName}",
	}

	if input == nil {
		input = &types.DescribeTemplateAliasInput{}
	}

	req := c.newRequest(op, input, &types.DescribeTemplateAliasOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeTemplateAliasMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeTemplateAliasRequest{Request: req, Input: input, Copy: c.DescribeTemplateAliasRequest}
}

// DescribeTemplateAliasRequest is the request type for the
// DescribeTemplateAlias API operation.
type DescribeTemplateAliasRequest struct {
	*aws.Request
	Input *types.DescribeTemplateAliasInput
	Copy  func(*types.DescribeTemplateAliasInput) DescribeTemplateAliasRequest
}

// Send marshals and sends the DescribeTemplateAlias API request.
func (r DescribeTemplateAliasRequest) Send(ctx context.Context) (*DescribeTemplateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTemplateAliasResponse{
		DescribeTemplateAliasOutput: r.Request.Data.(*types.DescribeTemplateAliasOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTemplateAliasResponse is the response type for the
// DescribeTemplateAlias API operation.
type DescribeTemplateAliasResponse struct {
	*types.DescribeTemplateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTemplateAlias request.
func (r *DescribeTemplateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
