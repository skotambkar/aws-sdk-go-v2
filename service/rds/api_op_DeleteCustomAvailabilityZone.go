// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/rds/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opDeleteCustomAvailabilityZone = "DeleteCustomAvailabilityZone"

// DeleteCustomAvailabilityZoneRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes a custom Availability Zone (AZ).
//
// A custom AZ is an on-premises AZ that is integrated with a VMware vSphere
// cluster.
//
// For more information about RDS on VMware, see the RDS on VMware User Guide.
// (https://docs.aws.amazon.com/AmazonRDS/latest/RDSonVMwareUserGuide/rds-on-vmware.html)
//
//    // Example sending a request using DeleteCustomAvailabilityZoneRequest.
//    req := client.DeleteCustomAvailabilityZoneRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteCustomAvailabilityZone
func (c *Client) DeleteCustomAvailabilityZoneRequest(input *types.DeleteCustomAvailabilityZoneInput) DeleteCustomAvailabilityZoneRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomAvailabilityZone,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteCustomAvailabilityZoneInput{}
	}

	req := c.newRequest(op, input, &types.DeleteCustomAvailabilityZoneOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteCustomAvailabilityZoneMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteCustomAvailabilityZoneRequest{Request: req, Input: input, Copy: c.DeleteCustomAvailabilityZoneRequest}
}

// DeleteCustomAvailabilityZoneRequest is the request type for the
// DeleteCustomAvailabilityZone API operation.
type DeleteCustomAvailabilityZoneRequest struct {
	*aws.Request
	Input *types.DeleteCustomAvailabilityZoneInput
	Copy  func(*types.DeleteCustomAvailabilityZoneInput) DeleteCustomAvailabilityZoneRequest
}

// Send marshals and sends the DeleteCustomAvailabilityZone API request.
func (r DeleteCustomAvailabilityZoneRequest) Send(ctx context.Context) (*DeleteCustomAvailabilityZoneResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomAvailabilityZoneResponse{
		DeleteCustomAvailabilityZoneOutput: r.Request.Data.(*types.DeleteCustomAvailabilityZoneOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomAvailabilityZoneResponse is the response type for the
// DeleteCustomAvailabilityZone API operation.
type DeleteCustomAvailabilityZoneResponse struct {
	*types.DeleteCustomAvailabilityZoneOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomAvailabilityZone request.
func (r *DeleteCustomAvailabilityZoneResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
