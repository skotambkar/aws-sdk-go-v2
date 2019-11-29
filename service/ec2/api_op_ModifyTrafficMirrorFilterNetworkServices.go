// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyTrafficMirrorFilterNetworkServices = "ModifyTrafficMirrorFilterNetworkServices"

// ModifyTrafficMirrorFilterNetworkServicesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Allows or restricts mirroring network services.
//
// By default, Amazon DNS network services are not eligible for Traffic Mirror.
// Use AddNetworkServices to add network services to a Traffic Mirror filter.
// When a network service is added to the Traffic Mirror filter, all traffic
// related to that network service will be mirrored. When you no longer want
// to mirror network services, use RemoveNetworkServices to remove the network
// services from the Traffic Mirror filter.
//
// For information about filter rule properties, see Network Services (https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html)
// in the Traffic Mirroring User Guide .
//
//    // Example sending a request using ModifyTrafficMirrorFilterNetworkServicesRequest.
//    req := client.ModifyTrafficMirrorFilterNetworkServicesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyTrafficMirrorFilterNetworkServices
func (c *Client) ModifyTrafficMirrorFilterNetworkServicesRequest(input *types.ModifyTrafficMirrorFilterNetworkServicesInput) ModifyTrafficMirrorFilterNetworkServicesRequest {
	op := &aws.Operation{
		Name:       opModifyTrafficMirrorFilterNetworkServices,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyTrafficMirrorFilterNetworkServicesInput{}
	}

	req := c.newRequest(op, input, &types.ModifyTrafficMirrorFilterNetworkServicesOutput{})
	return ModifyTrafficMirrorFilterNetworkServicesRequest{Request: req, Input: input, Copy: c.ModifyTrafficMirrorFilterNetworkServicesRequest}
}

// ModifyTrafficMirrorFilterNetworkServicesRequest is the request type for the
// ModifyTrafficMirrorFilterNetworkServices API operation.
type ModifyTrafficMirrorFilterNetworkServicesRequest struct {
	*aws.Request
	Input *types.ModifyTrafficMirrorFilterNetworkServicesInput
	Copy  func(*types.ModifyTrafficMirrorFilterNetworkServicesInput) ModifyTrafficMirrorFilterNetworkServicesRequest
}

// Send marshals and sends the ModifyTrafficMirrorFilterNetworkServices API request.
func (r ModifyTrafficMirrorFilterNetworkServicesRequest) Send(ctx context.Context) (*ModifyTrafficMirrorFilterNetworkServicesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyTrafficMirrorFilterNetworkServicesResponse{
		ModifyTrafficMirrorFilterNetworkServicesOutput: r.Request.Data.(*types.ModifyTrafficMirrorFilterNetworkServicesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyTrafficMirrorFilterNetworkServicesResponse is the response type for the
// ModifyTrafficMirrorFilterNetworkServices API operation.
type ModifyTrafficMirrorFilterNetworkServicesResponse struct {
	*types.ModifyTrafficMirrorFilterNetworkServicesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyTrafficMirrorFilterNetworkServices request.
func (r *ModifyTrafficMirrorFilterNetworkServicesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
