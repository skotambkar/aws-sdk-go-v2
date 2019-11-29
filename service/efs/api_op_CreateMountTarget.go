// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
)

const opCreateMountTarget = "CreateMountTarget"

// CreateMountTargetRequest returns a request value for making API operation for
// Amazon Elastic File System.
//
// Creates a mount target for a file system. You can then mount the file system
// on EC2 instances by using the mount target.
//
// You can create one mount target in each Availability Zone in your VPC. All
// EC2 instances in a VPC within a given Availability Zone share a single mount
// target for a given file system. If you have multiple subnets in an Availability
// Zone, you create a mount target in one of the subnets. EC2 instances do not
// need to be in the same subnet as the mount target in order to access their
// file system. For more information, see Amazon EFS: How it Works (https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html).
//
// In the request, you also specify a file system ID for which you are creating
// the mount target and the file system's lifecycle state must be available.
// For more information, see DescribeFileSystems.
//
// In the request, you also provide a subnet ID, which determines the following:
//
//    * VPC in which Amazon EFS creates the mount target
//
//    * Availability Zone in which Amazon EFS creates the mount target
//
//    * IP address range from which Amazon EFS selects the IP address of the
//    mount target (if you don't specify an IP address in the request)
//
// After creating the mount target, Amazon EFS returns a response that includes,
// a MountTargetId and an IpAddress. You use this IP address when mounting the
// file system in an EC2 instance. You can also use the mount target's DNS name
// when mounting the file system. The EC2 instance on which you mount the file
// system by using the mount target can resolve the mount target's DNS name
// to its IP address. For more information, see How it Works: Implementation
// Overview (https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html#how-it-works-implementation).
//
// Note that you can create mount targets for a file system in only one VPC,
// and there can be only one mount target per Availability Zone. That is, if
// the file system already has one or more mount targets created for it, the
// subnet specified in the request to add another mount target must meet the
// following requirements:
//
//    * Must belong to the same VPC as the subnets of the existing mount targets
//
//    * Must not be in the same Availability Zone as any of the subnets of the
//    existing mount targets
//
// If the request satisfies the requirements, Amazon EFS does the following:
//
//    * Creates a new mount target in the specified subnet.
//
//    * Also creates a new network interface in the subnet as follows: If the
//    request provides an IpAddress, Amazon EFS assigns that IP address to the
//    network interface. Otherwise, Amazon EFS assigns a free address in the
//    subnet (in the same way that the Amazon EC2 CreateNetworkInterface call
//    does when a request does not specify a primary private IP address). If
//    the request provides SecurityGroups, this network interface is associated
//    with those security groups. Otherwise, it belongs to the default security
//    group for the subnet's VPC. Assigns the description Mount target fsmt-id
//    for file system fs-id where fsmt-id is the mount target ID, and fs-id
//    is the FileSystemId. Sets the requesterManaged property of the network
//    interface to true, and the requesterId value to EFS. Each Amazon EFS mount
//    target has one corresponding requester-managed EC2 network interface.
//    After the network interface is created, Amazon EFS sets the NetworkInterfaceId
//    field in the mount target's description to the network interface ID, and
//    the IpAddress field to its address. If network interface creation fails,
//    the entire CreateMountTarget operation fails.
//
// The CreateMountTarget call returns only after creating the network interface,
// but while the mount target state is still creating, you can check the mount
// target creation status by calling the DescribeMountTargets operation, which
// among other things returns the mount target state.
//
// We recommend that you create a mount target in each of the Availability Zones.
// There are cost considerations for using a file system in an Availability
// Zone through a mount target created in another Availability Zone. For more
// information, see Amazon EFS (http://aws.amazon.com/efs/). In addition, by
// always using a mount target local to the instance's Availability Zone, you
// eliminate a partial failure scenario. If the Availability Zone in which your
// mount target is created goes down, then you can't access your file system
// through that mount target.
//
// This operation requires permissions for the following action on the file
// system:
//
//    * elasticfilesystem:CreateMountTarget
//
// This operation also requires permissions for the following Amazon EC2 actions:
//
//    * ec2:DescribeSubnets
//
//    * ec2:DescribeNetworkInterfaces
//
//    * ec2:CreateNetworkInterface
//
//    // Example sending a request using CreateMountTargetRequest.
//    req := client.CreateMountTargetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticfilesystem-2015-02-01/CreateMountTarget
func (c *Client) CreateMountTargetRequest(input *types.CreateMountTargetInput) CreateMountTargetRequest {
	op := &aws.Operation{
		Name:       opCreateMountTarget,
		HTTPMethod: "POST",
		HTTPPath:   "/2015-02-01/mount-targets",
	}

	if input == nil {
		input = &types.CreateMountTargetInput{}
	}

	req := c.newRequest(op, input, &types.CreateMountTargetOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateMountTargetMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateMountTargetRequest{Request: req, Input: input, Copy: c.CreateMountTargetRequest}
}

// CreateMountTargetRequest is the request type for the
// CreateMountTarget API operation.
type CreateMountTargetRequest struct {
	*aws.Request
	Input *types.CreateMountTargetInput
	Copy  func(*types.CreateMountTargetInput) CreateMountTargetRequest
}

// Send marshals and sends the CreateMountTarget API request.
func (r CreateMountTargetRequest) Send(ctx context.Context) (*CreateMountTargetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMountTargetResponse{
		CreateMountTargetOutput: r.Request.Data.(*types.CreateMountTargetOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMountTargetResponse is the response type for the
// CreateMountTarget API operation.
type CreateMountTargetResponse struct {
	*types.CreateMountTargetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMountTarget request.
func (r *CreateMountTargetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
