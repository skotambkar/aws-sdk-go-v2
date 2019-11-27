// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

const opRegisterDomain = "RegisterDomain"

// RegisterDomainRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation registers a domain. Domains are registered either by Amazon
// Registrar (for .com, .net, and .org domains) or by our registrar associate,
// Gandi (for all other domains). For some top-level domains (TLDs), this operation
// requires extra parameters.
//
// When you register a domain, Amazon Route 53 does the following:
//
//    * Creates a Amazon Route 53 hosted zone that has the same name as the
//    domain. Amazon Route 53 assigns four name servers to your hosted zone
//    and automatically updates your domain registration with the names of these
//    name servers.
//
//    * Enables autorenew, so your domain registration will renew automatically
//    each year. We'll notify you in advance of the renewal date so you can
//    choose whether to renew the registration.
//
//    * Optionally enables privacy protection, so WHOIS queries return contact
//    information either for Amazon Registrar (for .com, .net, and .org domains)
//    or for our registrar associate, Gandi (for all other TLDs). If you don't
//    enable privacy protection, WHOIS queries return the information that you
//    entered for the registrant, admin, and tech contacts.
//
//    * If registration is successful, returns an operation ID that you can
//    use to track the progress and completion of the action. If the request
//    is not completed successfully, the domain registrant is notified by email.
//
//    * Charges your AWS account an amount based on the top-level domain. For
//    more information, see Amazon Route 53 Pricing (http://aws.amazon.com/route53/pricing/).
//
//    // Example sending a request using RegisterDomainRequest.
//    req := client.RegisterDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/RegisterDomain
func (c *Client) RegisterDomainRequest(input *types.RegisterDomainInput) RegisterDomainRequest {
	op := &aws.Operation{
		Name:       opRegisterDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RegisterDomainInput{}
	}

	req := c.newRequest(op, input, &types.RegisterDomainOutput{})
	return RegisterDomainRequest{Request: req, Input: input, Copy: c.RegisterDomainRequest}
}

// RegisterDomainRequest is the request type for the
// RegisterDomain API operation.
type RegisterDomainRequest struct {
	*aws.Request
	Input *types.RegisterDomainInput
	Copy  func(*types.RegisterDomainInput) RegisterDomainRequest
}

// Send marshals and sends the RegisterDomain API request.
func (r RegisterDomainRequest) Send(ctx context.Context) (*RegisterDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterDomainResponse{
		RegisterDomainOutput: r.Request.Data.(*types.RegisterDomainOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterDomainResponse is the response type for the
// RegisterDomain API operation.
type RegisterDomainResponse struct {
	*types.RegisterDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterDomain request.
func (r *RegisterDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
