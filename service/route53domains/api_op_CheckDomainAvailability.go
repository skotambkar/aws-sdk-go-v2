// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53domains

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

const opCheckDomainAvailability = "CheckDomainAvailability"

// CheckDomainAvailabilityRequest returns a request value for making API operation for
// Amazon Route 53 Domains.
//
// This operation checks the availability of one domain name. Note that if the
// availability status of a domain is pending, you must submit another request
// to determine the availability of the domain name.
//
//    // Example sending a request using CheckDomainAvailabilityRequest.
//    req := client.CheckDomainAvailabilityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53domains-2014-05-15/CheckDomainAvailability
func (c *Client) CheckDomainAvailabilityRequest(input *types.CheckDomainAvailabilityInput) CheckDomainAvailabilityRequest {
	op := &aws.Operation{
		Name:       opCheckDomainAvailability,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CheckDomainAvailabilityInput{}
	}

	req := c.newRequest(op, input, &types.CheckDomainAvailabilityOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CheckDomainAvailabilityMarshaler{Input: input}.GetNamedBuildHandler())

	return CheckDomainAvailabilityRequest{Request: req, Input: input, Copy: c.CheckDomainAvailabilityRequest}
}

// CheckDomainAvailabilityRequest is the request type for the
// CheckDomainAvailability API operation.
type CheckDomainAvailabilityRequest struct {
	*aws.Request
	Input *types.CheckDomainAvailabilityInput
	Copy  func(*types.CheckDomainAvailabilityInput) CheckDomainAvailabilityRequest
}

// Send marshals and sends the CheckDomainAvailability API request.
func (r CheckDomainAvailabilityRequest) Send(ctx context.Context) (*CheckDomainAvailabilityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CheckDomainAvailabilityResponse{
		CheckDomainAvailabilityOutput: r.Request.Data.(*types.CheckDomainAvailabilityOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CheckDomainAvailabilityResponse is the response type for the
// CheckDomainAvailability API operation.
type CheckDomainAvailabilityResponse struct {
	*types.CheckDomainAvailabilityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CheckDomainAvailability request.
func (r *CheckDomainAvailabilityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
