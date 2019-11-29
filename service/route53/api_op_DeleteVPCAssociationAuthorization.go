// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/internal/aws_restxml"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opDeleteVPCAssociationAuthorization = "DeleteVPCAssociationAuthorization"

// DeleteVPCAssociationAuthorizationRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Removes authorization to submit an AssociateVPCWithHostedZone request to
// associate a specified VPC with a hosted zone that was created by a different
// account. You must use the account that created the hosted zone to submit
// a DeleteVPCAssociationAuthorization request.
//
// Sending this request only prevents the AWS account that created the VPC from
// associating the VPC with the Amazon Route 53 hosted zone in the future. If
// the VPC is already associated with the hosted zone, DeleteVPCAssociationAuthorization
// won't disassociate the VPC from the hosted zone. If you want to delete an
// existing association, use DisassociateVPCFromHostedZone.
//
//    // Example sending a request using DeleteVPCAssociationAuthorizationRequest.
//    req := client.DeleteVPCAssociationAuthorizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/DeleteVPCAssociationAuthorization
func (c *Client) DeleteVPCAssociationAuthorizationRequest(input *types.DeleteVPCAssociationAuthorizationInput) DeleteVPCAssociationAuthorizationRequest {
	op := &aws.Operation{
		Name:       opDeleteVPCAssociationAuthorization,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/hostedzone/{Id}/deauthorizevpcassociation",
	}

	if input == nil {
		input = &types.DeleteVPCAssociationAuthorizationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteVPCAssociationAuthorizationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restxml.BuildHandler.Name, aws_restxml.DeleteVPCAssociationAuthorizationMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteVPCAssociationAuthorizationRequest{Request: req, Input: input, Copy: c.DeleteVPCAssociationAuthorizationRequest}
}

// DeleteVPCAssociationAuthorizationRequest is the request type for the
// DeleteVPCAssociationAuthorization API operation.
type DeleteVPCAssociationAuthorizationRequest struct {
	*aws.Request
	Input *types.DeleteVPCAssociationAuthorizationInput
	Copy  func(*types.DeleteVPCAssociationAuthorizationInput) DeleteVPCAssociationAuthorizationRequest
}

// Send marshals and sends the DeleteVPCAssociationAuthorization API request.
func (r DeleteVPCAssociationAuthorizationRequest) Send(ctx context.Context) (*DeleteVPCAssociationAuthorizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteVPCAssociationAuthorizationResponse{
		DeleteVPCAssociationAuthorizationOutput: r.Request.Data.(*types.DeleteVPCAssociationAuthorizationOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteVPCAssociationAuthorizationResponse is the response type for the
// DeleteVPCAssociationAuthorization API operation.
type DeleteVPCAssociationAuthorizationResponse struct {
	*types.DeleteVPCAssociationAuthorizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteVPCAssociationAuthorization request.
func (r *DeleteVPCAssociationAuthorizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
