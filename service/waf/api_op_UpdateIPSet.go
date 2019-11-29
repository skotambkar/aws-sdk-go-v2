// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
)

const opUpdateIPSet = "UpdateIPSet"

// UpdateIPSetRequest returns a request value for making API operation for
// AWS WAF.
//
// Inserts or deletes IPSetDescriptor objects in an IPSet. For each IPSetDescriptor
// object, you specify the following values:
//
//    * Whether to insert or delete the object from the array. If you want to
//    change an IPSetDescriptor object, you delete the existing object and add
//    a new one.
//
//    * The IP address version, IPv4 or IPv6.
//
//    * The IP address in CIDR notation, for example, 192.0.2.0/24 (for the
//    range of IP addresses from 192.0.2.0 to 192.0.2.255) or 192.0.2.44/32
//    (for the individual IP address 192.0.2.44).
//
// AWS WAF supports IPv4 address ranges: /8 and any range between /16 through
// /32. AWS WAF supports IPv6 address ranges: /24, /32, /48, /56, /64, and /128.
// For more information about CIDR notation, see the Wikipedia entry Classless
// Inter-Domain Routing (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
//
// IPv6 addresses can be represented using any of the following formats:
//
//    * 1111:0000:0000:0000:0000:0000:0000:0111/128
//
//    * 1111:0:0:0:0:0:0:0111/128
//
//    * 1111::0111/128
//
//    * 1111::111/128
//
// You use an IPSet to specify which web requests you want to allow or block
// based on the IP addresses that the requests originated from. For example,
// if you're receiving a lot of requests from one or a small number of IP addresses
// and you want to block the requests, you can create an IPSet that specifies
// those IP addresses, and then configure AWS WAF to block the requests.
//
// To create and configure an IPSet, perform the following steps:
//
// Submit a CreateIPSet request.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateIPSet request.
//
// Submit an UpdateIPSet request to specify the IP addresses that you want AWS
// WAF to watch for.
//
// When you update an IPSet, you specify the IP addresses that you want to add
// and/or the IP addresses that you want to delete. If you want to change an
// IP address, you delete the existing IP address and add the new one.
//
// You can insert a maximum of 1000 addresses in a single request.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateIPSetRequest.
//    req := client.UpdateIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-2015-08-24/UpdateIPSet
func (c *Client) UpdateIPSetRequest(input *types.UpdateIPSetInput) UpdateIPSetRequest {
	op := &aws.Operation{
		Name:       opUpdateIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateIPSetInput{}
	}

	req := c.newRequest(op, input, &types.UpdateIPSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateIPSetMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateIPSetRequest{Request: req, Input: input, Copy: c.UpdateIPSetRequest}
}

// UpdateIPSetRequest is the request type for the
// UpdateIPSet API operation.
type UpdateIPSetRequest struct {
	*aws.Request
	Input *types.UpdateIPSetInput
	Copy  func(*types.UpdateIPSetInput) UpdateIPSetRequest
}

// Send marshals and sends the UpdateIPSet API request.
func (r UpdateIPSetRequest) Send(ctx context.Context) (*UpdateIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIPSetResponse{
		UpdateIPSetOutput: r.Request.Data.(*types.UpdateIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIPSetResponse is the response type for the
// UpdateIPSet API operation.
type UpdateIPSetResponse struct {
	*types.UpdateIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIPSet request.
func (r *UpdateIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
