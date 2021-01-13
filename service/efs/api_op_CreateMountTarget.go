// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a mount target for a file system. You can then mount the file system on
// EC2 instances by using the mount target. You can create one mount target in each
// Availability Zone in your VPC. All EC2 instances in a VPC within a given
// Availability Zone share a single mount target for a given file system. If you
// have multiple subnets in an Availability Zone, you create a mount target in one
// of the subnets. EC2 instances do not need to be in the same subnet as the mount
// target in order to access their file system. For more information, see Amazon
// EFS: How it Works (https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html).
// In the request, you also specify a file system ID for which you are creating the
// mount target and the file system's lifecycle state must be available. For more
// information, see DescribeFileSystems. In the request, you also provide a subnet
// ID, which determines the following:
//
// * VPC in which Amazon EFS creates the mount
// target
//
// * Availability Zone in which Amazon EFS creates the mount target
//
// * IP
// address range from which Amazon EFS selects the IP address of the mount target
// (if you don't specify an IP address in the request)
//
// After creating the mount
// target, Amazon EFS returns a response that includes, a MountTargetId and an
// IpAddress. You use this IP address when mounting the file system in an EC2
// instance. You can also use the mount target's DNS name when mounting the file
// system. The EC2 instance on which you mount the file system by using the mount
// target can resolve the mount target's DNS name to its IP address. For more
// information, see How it Works: Implementation Overview
// (https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html#how-it-works-implementation).
// Note that you can create mount targets for a file system in only one VPC, and
// there can be only one mount target per Availability Zone. That is, if the file
// system already has one or more mount targets created for it, the subnet
// specified in the request to add another mount target must meet the following
// requirements:
//
// * Must belong to the same VPC as the subnets of the existing
// mount targets
//
// * Must not be in the same Availability Zone as any of the subnets
// of the existing mount targets
//
// If the request satisfies the requirements, Amazon
// EFS does the following:
//
// * Creates a new mount target in the specified
// subnet.
//
// * Also creates a new network interface in the subnet as follows:
//
// * If
// the request provides an IpAddress, Amazon EFS assigns that IP address to the
// network interface. Otherwise, Amazon EFS assigns a free address in the subnet
// (in the same way that the Amazon EC2 CreateNetworkInterface call does when a
// request does not specify a primary private IP address).
//
// * If the request
// provides SecurityGroups, this network interface is associated with those
// security groups. Otherwise, it belongs to the default security group for the
// subnet's VPC.
//
// * Assigns the description Mount target fsmt-id for file system
// fs-id  where  fsmt-id  is the mount target ID, and  fs-id  is the
// FileSystemId.
//
// * Sets the requesterManaged property of the network interface to
// true, and the requesterId value to EFS.
//
// Each Amazon EFS mount target has one
// corresponding requester-managed EC2 network interface. After the network
// interface is created, Amazon EFS sets the NetworkInterfaceId field in the mount
// target's description to the network interface ID, and the IpAddress field to its
// address. If network interface creation fails, the entire CreateMountTarget
// operation fails.
//
// The CreateMountTarget call returns only after creating the
// network interface, but while the mount target state is still creating, you can
// check the mount target creation status by calling the DescribeMountTargets
// operation, which among other things returns the mount target state. We recommend
// that you create a mount target in each of the Availability Zones. There are cost
// considerations for using a file system in an Availability Zone through a mount
// target created in another Availability Zone. For more information, see Amazon
// EFS (http://aws.amazon.com/efs/). In addition, by always using a mount target
// local to the instance's Availability Zone, you eliminate a partial failure
// scenario. If the Availability Zone in which your mount target is created goes
// down, then you can't access your file system through that mount target. This
// operation requires permissions for the following action on the file system:
//
// *
// elasticfilesystem:CreateMountTarget
//
// This operation also requires permissions
// for the following Amazon EC2 actions:
//
// * ec2:DescribeSubnets
//
// *
// ec2:DescribeNetworkInterfaces
//
// * ec2:CreateNetworkInterface
func (c *Client) CreateMountTarget(ctx context.Context, params *CreateMountTargetInput, optFns ...func(*Options)) (*CreateMountTargetOutput, error) {
	if params == nil {
		params = &CreateMountTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMountTarget", params, optFns, addOperationCreateMountTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMountTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateMountTargetInput struct {

	// The ID of the file system for which to create the mount target.
	//
	// This member is required.
	FileSystemId *string

	// The ID of the subnet to add the mount target in.
	//
	// This member is required.
	SubnetId *string

	// Valid IPv4 address within the address range of the specified subnet.
	IpAddress *string

	// Up to five VPC security group IDs, of the form sg-xxxxxxxx. These must be for
	// the same VPC as subnet specified.
	SecurityGroups []string
}

// Provides a description of a mount target.
type CreateMountTargetOutput struct {

	// The ID of the file system for which the mount target is intended.
	//
	// This member is required.
	FileSystemId *string

	// Lifecycle state of the mount target.
	//
	// This member is required.
	LifeCycleState types.LifeCycleState

	// System-assigned mount target ID.
	//
	// This member is required.
	MountTargetId *string

	// The ID of the mount target's subnet.
	//
	// This member is required.
	SubnetId *string

	// The unique and consistent identifier of the Availability Zone (AZ) that the
	// mount target resides in. For example, use1-az1 is an AZ ID for the us-east-1
	// Region and it has the same location in every AWS account.
	AvailabilityZoneId *string

	// The name of the Availability Zone (AZ) that the mount target resides in. AZs are
	// independently mapped to names for each AWS account. For example, the
	// Availability Zone us-east-1a for your AWS account might not be the same location
	// as us-east-1a for another AWS account.
	AvailabilityZoneName *string

	// Address at which the file system can be mounted by using the mount target.
	IpAddress *string

	// The ID of the network interface that Amazon EFS created when it created the
	// mount target.
	NetworkInterfaceId *string

	// AWS account ID that owns the resource.
	OwnerId *string

	// The Virtual Private Cloud (VPC) ID that the mount target is configured in.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMountTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMountTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMountTarget{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateMountTargetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMountTarget(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateMountTarget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "CreateMountTarget",
	}
}
