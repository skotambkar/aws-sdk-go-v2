// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package securityhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
)

const opDisableSecurityHub = "DisableSecurityHub"

// DisableSecurityHubRequest returns a request value for making API operation for
// AWS SecurityHub.
//
// Disables Security Hub in your account only in the current Region. To disable
// Security Hub in all Regions, you must submit one request per Region where
// you have enabled Security Hub. When you disable Security Hub for a master
// account, it doesn't disable Security Hub for any associated member accounts.
//
// When you disable Security Hub, your existing findings and insights and any
// Security Hub configuration settings are deleted after 90 days and can't be
// recovered. Any standards that were enabled are disabled, and your master
// and member account associations are removed. If you want to save your existing
// findings, you must export them before you disable Security Hub.
//
//    // Example sending a request using DisableSecurityHubRequest.
//    req := client.DisableSecurityHubRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/securityhub-2018-10-26/DisableSecurityHub
func (c *Client) DisableSecurityHubRequest(input *types.DisableSecurityHubInput) DisableSecurityHubRequest {
	op := &aws.Operation{
		Name:       opDisableSecurityHub,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts",
	}

	if input == nil {
		input = &types.DisableSecurityHubInput{}
	}

	req := c.newRequest(op, input, &types.DisableSecurityHubOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DisableSecurityHubMarshaler{Input: input}.GetNamedBuildHandler())

	return DisableSecurityHubRequest{Request: req, Input: input, Copy: c.DisableSecurityHubRequest}
}

// DisableSecurityHubRequest is the request type for the
// DisableSecurityHub API operation.
type DisableSecurityHubRequest struct {
	*aws.Request
	Input *types.DisableSecurityHubInput
	Copy  func(*types.DisableSecurityHubInput) DisableSecurityHubRequest
}

// Send marshals and sends the DisableSecurityHub API request.
func (r DisableSecurityHubRequest) Send(ctx context.Context) (*DisableSecurityHubResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableSecurityHubResponse{
		DisableSecurityHubOutput: r.Request.Data.(*types.DisableSecurityHubOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableSecurityHubResponse is the response type for the
// DisableSecurityHub API operation.
type DisableSecurityHubResponse struct {
	*types.DisableSecurityHubOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableSecurityHub request.
func (r *DisableSecurityHubResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
