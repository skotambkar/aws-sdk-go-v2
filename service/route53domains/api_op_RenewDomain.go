// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

const opRenewDomain = "RenewDomain"

// RenewDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation renews a domain for the specified number of years. The cost
// of renewing your domain is billed to your AWS account.
//
// We recommend that you renew your domain several weeks before the expiration
// date. Some TLD registries delete domains before the expiration date if you
// haven't renewed far enough in advance. For more information about renewing
// domain registration, see Renewing Registration for a Domain (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-renew.html)
// in the Amazon Route 53 Developer Guide.
//
//    // Example sending a request using RenewDomainRequest.
//    req := client.RenewDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/RenewDomain
func (c *Client) RenewDomainRequest(input *types.RenewDomainInput) RenewDomainRequest {
	op := &aws.Operation{
		Name:       opRenewDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RenewDomainInput{}
	}

	req := c.newRequest(op, input, &types.RenewDomainOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.RenewDomainMarshaler{Input: input}.GetNamedBuildHandler())

	return RenewDomainRequest{Request: req, Input: input, Copy: c.RenewDomainRequest}
}

// RenewDomainRequest is the request type for the
// RenewDomain API operation.
type RenewDomainRequest struct {
	*aws.Request
	Input *types.RenewDomainInput
	Copy  func(*types.RenewDomainInput) RenewDomainRequest
}

// Send marshals and sends the RenewDomain API request.
func (r RenewDomainRequest) Send(ctx context.Context) (*RenewDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RenewDomainResponse{
		RenewDomainOutput: r.Request.Data.(*types.RenewDomainOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RenewDomainResponse is the response type for the
// RenewDomain API operation.
type RenewDomainResponse struct {
	*types.RenewDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RenewDomain request.
func (r *RenewDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
