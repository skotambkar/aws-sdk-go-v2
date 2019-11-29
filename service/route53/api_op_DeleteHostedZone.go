// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opDeleteHostedZone = "DeleteHostedZone"

// DeleteHostedZoneRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Deletes a hosted zone.
//
// If the hosted zone was created by another service, such as AWS Cloud Map,
// see Deleting Public Hosted Zones That Were Created by Another Service (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DeleteHostedZone.html#delete-public-hosted-zone-created-by-another-service)
// in the Amazon Route 53 Developer Guide for information about how to delete
// it. (The process is the same for public and private hosted zones that were
// created by another service.)
//
// If you want to keep your domain registration but you want to stop routing
// internet traffic to your website or web application, we recommend that you
// delete resource record sets in the hosted zone instead of deleting the hosted
// zone.
//
// If you delete a hosted zone, you can't undelete it. You must create a new
// hosted zone and update the name servers for your domain registration, which
// can require up to 48 hours to take effect. (If you delegated responsibility
// for a subdomain to a hosted zone and you delete the child hosted zone, you
// must update the name servers in the parent hosted zone.) In addition, if
// you delete a hosted zone, someone could hijack the domain and route traffic
// to their own resources using your domain name.
//
// If you want to avoid the monthly charge for the hosted zone, you can transfer
// DNS service for the domain to a free DNS service. When you transfer DNS service,
// you have to update the name servers for the domain registration. If the domain
// is registered with Route 53, see UpdateDomainNameservers (https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_UpdateDomainNameservers.html)
// for information about how to replace Route 53 name servers with name servers
// for the new DNS service. If the domain is registered with another registrar,
// use the method provided by the registrar to update name servers for the domain
// registration. For more information, perform an internet search on "free DNS
// service."
//
// You can delete a hosted zone only if it contains only the default SOA record
// and NS resource record sets. If the hosted zone contains other resource record
// sets, you must delete them before you can delete the hosted zone. If you
// try to delete a hosted zone that contains other resource record sets, the
// request fails, and Route 53 returns a HostedZoneNotEmpty error. For information
// about deleting records from your hosted zone, see ChangeResourceRecordSets
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html).
//
// To verify that the hosted zone has been deleted, do one of the following:
//
//    * Use the GetHostedZone action to request information about the hosted
//    zone.
//
//    * Use the ListHostedZones action to get a list of the hosted zones associated
//    with the current AWS account.
//
//    // Example sending a request using DeleteHostedZoneRequest.
//    req := client.DeleteHostedZoneRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteHostedZone
func (c *Client) DeleteHostedZoneRequest(input *types.DeleteHostedZoneInput) DeleteHostedZoneRequest {
	op := &aws.Operation{
		Name:       opDeleteHostedZone,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2013-04-01/hostedzone/{Id}",
	}

	if input == nil {
		input = &types.DeleteHostedZoneInput{}
	}

	req := c.newRequest(op, input, &types.DeleteHostedZoneOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.DeleteHostedZoneMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteHostedZoneRequest{Request: req, Input: input, Copy: c.DeleteHostedZoneRequest}
}

// DeleteHostedZoneRequest is the request type for the
// DeleteHostedZone API operation.
type DeleteHostedZoneRequest struct {
	*aws.Request
	Input *types.DeleteHostedZoneInput
	Copy  func(*types.DeleteHostedZoneInput) DeleteHostedZoneRequest
}

// Send marshals and sends the DeleteHostedZone API request.
func (r DeleteHostedZoneRequest) Send(ctx context.Context) (*DeleteHostedZoneResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteHostedZoneResponse{
		DeleteHostedZoneOutput: r.Request.Data.(*types.DeleteHostedZoneOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteHostedZoneResponse is the response type for the
// DeleteHostedZone API operation.
type DeleteHostedZoneResponse struct {
	*types.DeleteHostedZoneOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteHostedZone request.
func (r *DeleteHostedZoneResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
