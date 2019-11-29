// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

const opGetAlias = "GetAlias"

// GetAliasRequest returns a request value for making API operation for
// AWS Lambda.
//
// Returns details about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/versioning-aliases.html).
//
//    // Example sending a request using GetAliasRequest.
//    req := client.GetAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/GetAlias
func (c *Client) GetAliasRequest(input *types.GetAliasInput) GetAliasRequest {
	op := &aws.Operation{
		Name:       opGetAlias,
		HTTPMethod: "GET",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/aliases/{Name}",
	}

	if input == nil {
		input = &types.GetAliasInput{}
	}

	req := c.newRequest(op, input, &types.GetAliasOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetAliasMarshaler{Input: input}.GetNamedBuildHandler())

	return GetAliasRequest{Request: req, Input: input, Copy: c.GetAliasRequest}
}

// GetAliasRequest is the request type for the
// GetAlias API operation.
type GetAliasRequest struct {
	*aws.Request
	Input *types.GetAliasInput
	Copy  func(*types.GetAliasInput) GetAliasRequest
}

// Send marshals and sends the GetAlias API request.
func (r GetAliasRequest) Send(ctx context.Context) (*GetAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetAliasResponse{
		GetAliasOutput: r.Request.Data.(*types.GetAliasOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetAliasResponse is the response type for the
// GetAlias API operation.
type GetAliasResponse struct {
	*types.GetAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetAlias request.
func (r *GetAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
