// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opCreateReusableDelegationSet = "CreateReusableDelegationSet"

// CreateReusableDelegationSetRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Creates a delegation set (a group of four name servers) that can be reused
// by multiple hosted zones. If a hosted zoned ID is specified, CreateReusableDelegationSet
// marks the delegation set associated with that zone as reusable.
//
// You can't associate a reusable delegation set with a private hosted zone.
//
// For information about using a reusable delegation set to configure white
// label name servers, see Configuring White Label Name Servers (http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/white-label-name-servers.html).
//
// The process for migrating existing hosted zones to use a reusable delegation
// set is comparable to the process for configuring white label name servers.
// You need to perform the following steps:
//
// Create a reusable delegation set.
//
// Recreate hosted zones, and reduce the TTL to 60 seconds or less.
//
// Recreate resource record sets in the new hosted zones.
//
// Change the registrar's name servers to use the name servers for the new hosted
// zones.
//
// Monitor traffic for the website or application.
//
// Change TTLs back to their original values.
//
// If you want to migrate existing hosted zones to use a reusable delegation
// set, the existing hosted zones can't use any of the name servers that are
// assigned to the reusable delegation set. If one or more hosted zones do use
// one or more name servers that are assigned to the reusable delegation set,
// you can do one of the following:
//
//    * For small numbers of hosted zones—up to a few hundred—it's relatively
//    easy to create reusable delegation sets until you get one that has four
//    name servers that don't overlap with any of the name servers in your hosted
//    zones.
//
//    * For larger numbers of hosted zones, the easiest solution is to use more
//    than one reusable delegation set.
//
//    * For larger numbers of hosted zones, you can also migrate hosted zones
//    that have overlapping name servers to hosted zones that don't have overlapping
//    name servers, then migrate the hosted zones again to use the reusable
//    delegation set.
//
//    // Example sending a request using CreateReusableDelegationSetRequest.
//    req := client.CreateReusableDelegationSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateReusableDelegationSet
func (c *Client) CreateReusableDelegationSetRequest(input *types.CreateReusableDelegationSetInput) CreateReusableDelegationSetRequest {
	op := &aws.Operation{
		Name:       opCreateReusableDelegationSet,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/delegationset",
	}

	if input == nil {
		input = &types.CreateReusableDelegationSetInput{}
	}

	req := c.newRequest(op, input, &types.CreateReusableDelegationSetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.CreateReusableDelegationSetMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateReusableDelegationSetRequest{Request: req, Input: input, Copy: c.CreateReusableDelegationSetRequest}
}

// CreateReusableDelegationSetRequest is the request type for the
// CreateReusableDelegationSet API operation.
type CreateReusableDelegationSetRequest struct {
	*aws.Request
	Input *types.CreateReusableDelegationSetInput
	Copy  func(*types.CreateReusableDelegationSetInput) CreateReusableDelegationSetRequest
}

// Send marshals and sends the CreateReusableDelegationSet API request.
func (r CreateReusableDelegationSetRequest) Send(ctx context.Context) (*CreateReusableDelegationSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateReusableDelegationSetResponse{
		CreateReusableDelegationSetOutput: r.Request.Data.(*types.CreateReusableDelegationSetOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateReusableDelegationSetResponse is the response type for the
// CreateReusableDelegationSet API operation.
type CreateReusableDelegationSetResponse struct {
	*types.CreateReusableDelegationSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateReusableDelegationSet request.
func (r *CreateReusableDelegationSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
