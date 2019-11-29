// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

const opTransferDomain = "TransferDomain"

// TransferDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation transfers a domain from another registrar to Amazon Route
// 53. When the transfer is complete, the domain is registered either with Amazon
// Registrar (for .com, .net, and .org domains) or with our registrar associate,
// Gandi (for all other TLDs).
//
// For transfer requirements, a detailed procedure, and information about viewing
// the status of a domain transfer, see Transferring Registration for a Domain
// to Amazon Route 53 (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-transfer-to-route-53.html)
// in the Amazon Route 53 Developer Guide.
//
// If the registrar for your domain is also the DNS service provider for the
// domain, we highly recommend that you consider transferring your DNS service
// to Amazon Route 53 or to another DNS service provider before you transfer
// your registration. Some registrars provide free DNS service when you purchase
// a domain registration. When you transfer the registration, the previous registrar
// will not renew your domain registration and could end your DNS service at
// any time.
//
// If the registrar for your domain is also the DNS service provider for the
// domain and you don't transfer DNS service to another provider, your website,
// email, and the web applications associated with the domain might become unavailable.
//
// If the transfer is successful, this method returns an operation ID that you
// can use to track the progress and completion of the action. If the transfer
// doesn't complete successfully, the domain registrant will be notified by
// email.
//
//    // Example sending a request using TransferDomainRequest.
//    req := client.TransferDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/TransferDomain
func (c *Client) TransferDomainRequest(input *types.TransferDomainInput) TransferDomainRequest {
	op := &aws.Operation{
		Name:       opTransferDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.TransferDomainInput{}
	}

	req := c.newRequest(op, input, &types.TransferDomainOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.TransferDomainMarshaler{Input: input}.GetNamedBuildHandler())

	return TransferDomainRequest{Request: req, Input: input, Copy: c.TransferDomainRequest}
}

// TransferDomainRequest is the request type for the
// TransferDomain API operation.
type TransferDomainRequest struct {
	*aws.Request
	Input *types.TransferDomainInput
	Copy  func(*types.TransferDomainInput) TransferDomainRequest
}

// Send marshals and sends the TransferDomain API request.
func (r TransferDomainRequest) Send(ctx context.Context) (*TransferDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TransferDomainResponse{
		TransferDomainOutput: r.Request.Data.(*types.TransferDomainOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TransferDomainResponse is the response type for the
// TransferDomain API operation.
type TransferDomainResponse struct {
	*types.TransferDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TransferDomain request.
func (r *TransferDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
