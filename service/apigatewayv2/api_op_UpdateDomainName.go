// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewayv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
)

const opUpdateDomainName = "UpdateDomainName"

// UpdateDomainNameRequest returns a request value for making API operation for
// AmazonApiGatewayV2.
//
// Updates a domain name.
//
//    // Example sending a request using UpdateDomainNameRequest.
//    req := client.UpdateDomainNameRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/apigatewayv2-2018-11-29/UpdateDomainName
func (c *Client) UpdateDomainNameRequest(input *types.UpdateDomainNameInput) UpdateDomainNameRequest {
	op := &aws.Operation{
		Name:       opUpdateDomainName,
		HTTPMethod: "PATCH",
		HTTPPath:   "/v2/domainnames/{domainName}",
	}

	if input == nil {
		input = &types.UpdateDomainNameInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDomainNameOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateDomainNameMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDomainNameRequest{Request: req, Input: input, Copy: c.UpdateDomainNameRequest}
}

// UpdateDomainNameRequest is the request type for the
// UpdateDomainName API operation.
type UpdateDomainNameRequest struct {
	*aws.Request
	Input *types.UpdateDomainNameInput
	Copy  func(*types.UpdateDomainNameInput) UpdateDomainNameRequest
}

// Send marshals and sends the UpdateDomainName API request.
func (r UpdateDomainNameRequest) Send(ctx context.Context) (*UpdateDomainNameResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDomainNameResponse{
		UpdateDomainNameOutput: r.Request.Data.(*types.UpdateDomainNameOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDomainNameResponse is the response type for the
// UpdateDomainName API operation.
type UpdateDomainNameResponse struct {
	*types.UpdateDomainNameOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDomainName request.
func (r *UpdateDomainNameResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
