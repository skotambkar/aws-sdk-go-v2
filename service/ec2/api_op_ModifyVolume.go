// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyVolume = "ModifyVolume"

// ModifyVolumeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// You can modify several parameters of an existing EBS volume, including volume
// size, volume type, and IOPS capacity. If your EBS volume is attached to a
// current-generation EC2 instance type, you may be able to apply these changes
// without stopping the instance or detaching the volume from it. For more information
// about modifying an EBS volume running Linux, see Modifying the Size, IOPS,
// or Type of an EBS Volume on Linux (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html).
// For more information about modifying an EBS volume running Windows, see Modifying
// the Size, IOPS, or Type of an EBS Volume on Windows (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html).
//
// When you complete a resize operation on your volume, you need to extend the
// volume's file-system size to take advantage of the new storage capacity.
// For information about extending a Linux file system, see Extending a Linux
// File System (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#recognize-expanded-volume-linux).
// For information about extending a Windows file system, see Extending a Windows
// File System (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html#recognize-expanded-volume-windows).
//
// You can use CloudWatch Events to check the status of a modification to an
// EBS volume. For information about CloudWatch Events, see the Amazon CloudWatch
// Events User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/).
// You can also track the status of a modification using DescribeVolumesModifications.
// For information about tracking status changes using either method, see Monitoring
// Volume Modifications (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#monitoring_mods).
//
// With previous-generation instance types, resizing an EBS volume may require
// detaching and reattaching the volume or stopping and restarting the instance.
// For more information, see Modifying the Size, IOPS, or Type of an EBS Volume
// on Linux (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html)
// and Modifying the Size, IOPS, or Type of an EBS Volume on Windows (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html).
//
// If you reach the maximum volume modification rate per volume limit, you will
// need to wait at least six hours before applying further modifications to
// the affected EBS volume.
//
//    // Example sending a request using ModifyVolumeRequest.
//    req := client.ModifyVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVolume
func (c *Client) ModifyVolumeRequest(input *types.ModifyVolumeInput) ModifyVolumeRequest {
	op := &aws.Operation{
		Name:       opModifyVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyVolumeInput{}
	}

	req := c.newRequest(op, input, &types.ModifyVolumeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ModifyVolumeMarshaler{Input: input}.GetNamedBuildHandler())

	return ModifyVolumeRequest{Request: req, Input: input, Copy: c.ModifyVolumeRequest}
}

// ModifyVolumeRequest is the request type for the
// ModifyVolume API operation.
type ModifyVolumeRequest struct {
	*aws.Request
	Input *types.ModifyVolumeInput
	Copy  func(*types.ModifyVolumeInput) ModifyVolumeRequest
}

// Send marshals and sends the ModifyVolume API request.
func (r ModifyVolumeRequest) Send(ctx context.Context) (*ModifyVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVolumeResponse{
		ModifyVolumeOutput: r.Request.Data.(*types.ModifyVolumeOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVolumeResponse is the response type for the
// ModifyVolume API operation.
type ModifyVolumeResponse struct {
	*types.ModifyVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVolume request.
func (r *ModifyVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
