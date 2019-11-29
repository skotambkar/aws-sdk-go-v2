// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opCreateNetworkAclEntry = "CreateNetworkAclEntry"

// CreateNetworkAclEntryRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Creates an entry (a rule) in a network ACL with the specified rule number.
// Each network ACL has a set of numbered ingress rules and a separate set of
// numbered egress rules. When determining whether a packet should be allowed
// in or out of a subnet associated with the ACL, we process the entries in
// the ACL according to the rule numbers, in ascending order. Each network ACL
// has a set of ingress rules and a separate set of egress rules.
//
// We recommend that you leave room between the rule numbers (for example, 100,
// 110, 120, ...), and not number them one right after the other (for example,
// 101, 102, 103, ...). This makes it easier to add a rule between existing
// ones without having to renumber the rules.
//
// After you add an entry, you can't modify it; you must either replace it,
// or create an entry and delete the old one.
//
// For more information about network ACLs, see Network ACLs (https://docs.aws.amazon.com/vpc/latest/userguide/VPC_ACLs.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using CreateNetworkAclEntryRequest.
//    req := client.CreateNetworkAclEntryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/CreateNetworkAclEntry
func (c *Client) CreateNetworkAclEntryRequest(input *types.CreateNetworkAclEntryInput) CreateNetworkAclEntryRequest {
	op := &aws.Operation{
		Name:       opCreateNetworkAclEntry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateNetworkAclEntryInput{}
	}

	req := c.newRequest(op, input, &types.CreateNetworkAclEntryOutput{})
	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return CreateNetworkAclEntryRequest{Request: req, Input: input, Copy: c.CreateNetworkAclEntryRequest}
}

// CreateNetworkAclEntryRequest is the request type for the
// CreateNetworkAclEntry API operation.
type CreateNetworkAclEntryRequest struct {
	*aws.Request
	Input *types.CreateNetworkAclEntryInput
	Copy  func(*types.CreateNetworkAclEntryInput) CreateNetworkAclEntryRequest
}

// Send marshals and sends the CreateNetworkAclEntry API request.
func (r CreateNetworkAclEntryRequest) Send(ctx context.Context) (*CreateNetworkAclEntryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNetworkAclEntryResponse{
		CreateNetworkAclEntryOutput: r.Request.Data.(*types.CreateNetworkAclEntryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNetworkAclEntryResponse is the response type for the
// CreateNetworkAclEntry API operation.
type CreateNetworkAclEntryResponse struct {
	*types.CreateNetworkAclEntryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNetworkAclEntry request.
func (r *CreateNetworkAclEntryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
