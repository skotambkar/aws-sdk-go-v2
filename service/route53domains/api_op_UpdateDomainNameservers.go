// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

const opUpdateDomainNameservers = "UpdateDomainNameservers"

// UpdateDomainNameserversRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation replaces the current set of name servers for the domain with
// the specified set of name servers. If you use Amazon Route 53 as your DNS
// service, specify the four name servers in the delegation set for the hosted
// zone for the domain.
//
// If successful, this operation returns an operation ID that you can use to
// track the progress and completion of the action. If the request is not completed
// successfully, the domain registrant will be notified by email.
//
//    // Example sending a request using UpdateDomainNameserversRequest.
//    req := client.UpdateDomainNameserversRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/UpdateDomainNameservers
func (c *Client) UpdateDomainNameserversRequest(input *types.UpdateDomainNameserversInput) UpdateDomainNameserversRequest {
	op := &aws.Operation{
		Name:       opUpdateDomainNameservers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateDomainNameserversInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDomainNameserversOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateDomainNameserversMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDomainNameserversRequest{Request: req, Input: input, Copy: c.UpdateDomainNameserversRequest}
}

// UpdateDomainNameserversRequest is the request type for the
// UpdateDomainNameservers API operation.
type UpdateDomainNameserversRequest struct {
	*aws.Request
	Input *types.UpdateDomainNameserversInput
	Copy  func(*types.UpdateDomainNameserversInput) UpdateDomainNameserversRequest
}

// Send marshals and sends the UpdateDomainNameservers API request.
func (r UpdateDomainNameserversRequest) Send(ctx context.Context) (*UpdateDomainNameserversResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDomainNameserversResponse{
		UpdateDomainNameserversOutput: r.Request.Data.(*types.UpdateDomainNameserversOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDomainNameserversResponse is the response type for the
// UpdateDomainNameservers API operation.
type UpdateDomainNameserversResponse struct {
	*types.UpdateDomainNameserversOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDomainNameservers request.
func (r *UpdateDomainNameserversResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
