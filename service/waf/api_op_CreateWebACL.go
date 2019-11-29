// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opCreateWebACL = "CreateWebACL"

// CreateWebACLRequest returns a request value for making API operation for
// AWS WAF.
//
// Creates a WebACL, which contains the Rules that identify the CloudFront web
// requests that you want to allow, block, or count. AWS WAF evaluates Rules
// in order based on the value of Priority for each Rule.
//
// You also specify a default action, either ALLOW or BLOCK. If a web request
// doesn't match any of the Rules in a WebACL, AWS WAF responds to the request
// with the default action.
//
// To create and configure a WebACL, perform the following steps:
//
// Create and update the ByteMatchSet objects and other predicates that you
// want to include in Rules. For more information, see CreateByteMatchSet, UpdateByteMatchSet,
// CreateIPSet, UpdateIPSet, CreateSqlInjectionMatchSet, and UpdateSqlInjectionMatchSet.
//
// Create and update the Rules that you want to include in the WebACL. For more
// information, see CreateRule and UpdateRule.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateWebACL request.
//
// Submit a CreateWebACL request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateWebACL request.
//
// Submit an UpdateWebACL request to specify the Rules that you want to include
// in the WebACL, to specify the default action, and to associate the WebACL
// with a CloudFront distribution.
//
// For more information about how to use the AWS WAF API, see the AWS WAF Developer
// Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using CreateWebACLRequest.
//    req := client.CreateWebACLRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/CreateWebACL
func (c *Client) CreateWebACLRequest(input *types.CreateWebACLInput) CreateWebACLRequest {
	op := &aws.Operation{
		Name:       opCreateWebACL,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateWebACLInput{}
	}

	req := c.newRequest(op, input, &types.CreateWebACLOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateWebACLMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateWebACLRequest{Request: req, Input: input, Copy: c.CreateWebACLRequest}
}

// CreateWebACLRequest is the request type for the
// CreateWebACL API operation.
type CreateWebACLRequest struct {
	*aws.Request
	Input *types.CreateWebACLInput
	Copy  func(*types.CreateWebACLInput) CreateWebACLRequest
}

// Send marshals and sends the CreateWebACL API request.
func (r CreateWebACLRequest) Send(ctx context.Context) (*CreateWebACLResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateWebACLResponse{
		CreateWebACLOutput: r.Request.Data.(*types.CreateWebACLOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateWebACLResponse is the response type for the
// CreateWebACL API operation.
type CreateWebACLResponse struct {
	*types.CreateWebACLOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateWebACL request.
func (r *CreateWebACLResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
