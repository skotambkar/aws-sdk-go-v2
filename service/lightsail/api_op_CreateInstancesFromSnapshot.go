// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates one or more new instances from a manual or automatic snapshot of an
// instance. The create instances from snapshot operation supports tag-based access
// control via request tags and resource tags applied to the resource identified by
// instance snapshot name. For more information, see the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateInstancesFromSnapshot(ctx context.Context, params *CreateInstancesFromSnapshotInput, optFns ...func(*Options)) (*CreateInstancesFromSnapshotOutput, error) {
	if params == nil {
		params = &CreateInstancesFromSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstancesFromSnapshot", params, optFns, c.addOperationCreateInstancesFromSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstancesFromSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstancesFromSnapshotInput struct {

	// The Availability Zone where you want to create your instances. Use the following
	// formatting: us-east-2a (case sensitive). You can get a list of Availability
	// Zones by using the get regions
	// (http://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRegions.html)
	// operation. Be sure to add the include Availability Zones parameter to your
	// request.
	//
	// This member is required.
	AvailabilityZone *string

	// The bundle of specification information for your virtual private server (or
	// instance), including the pricing plan (e.g., micro_1_0).
	//
	// This member is required.
	BundleId *string

	// The names for your new instances.
	//
	// This member is required.
	InstanceNames []string

	// An array of objects representing the add-ons to enable for the new instance.
	AddOns []types.AddOnRequest

	// An object containing information about one or more disk mappings.
	AttachedDiskMapping map[string][]types.DiskMap

	// The name of the instance snapshot on which you are basing your new instances.
	// Use the get instance snapshots operation to return information about your
	// existing snapshots. Constraint:
	//
	// * This parameter cannot be defined together
	// with the source instance name parameter. The instance snapshot name and source
	// instance name parameters are mutually exclusive.
	InstanceSnapshotName *string

	// The IP address type for the instance. The possible values are ipv4 for IPv4
	// only, and dualstack for IPv4 and IPv6. The default value is dualstack.
	IpAddressType types.IpAddressType

	// The name for your key pair.
	KeyPairName *string

	// The date of the automatic snapshot to use for the new instance. Use the get auto
	// snapshots operation to identify the dates of the available automatic snapshots.
	// Constraints:
	//
	// * Must be specified in YYYY-MM-DD format.
	//
	// * This parameter cannot
	// be defined together with the use latest restorable auto snapshot parameter. The
	// restore date and use latest restorable auto snapshot parameters are mutually
	// exclusive.
	//
	// * Define this parameter only when creating a new instance from an
	// automatic snapshot. For more information, see the Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	RestoreDate *string

	// The name of the source instance from which the source automatic snapshot was
	// created. Constraints:
	//
	// * This parameter cannot be defined together with the
	// instance snapshot name parameter. The source instance name and instance snapshot
	// name parameters are mutually exclusive.
	//
	// * Define this parameter only when
	// creating a new instance from an automatic snapshot. For more information, see
	// the Lightsail Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	SourceInstanceName *string

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []types.Tag

	// A Boolean value to indicate whether to use the latest available automatic
	// snapshot. Constraints:
	//
	// * This parameter cannot be defined together with the
	// restore date parameter. The use latest restorable auto snapshot and restore date
	// parameters are mutually exclusive.
	//
	// * Define this parameter only when creating a
	// new instance from an automatic snapshot. For more information, see the Lightsail
	// Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-configuring-automatic-snapshots).
	UseLatestRestorableAutoSnapshot *bool

	// You can create a launch script that configures a server with additional user
	// data. For example, apt-get -y update. Depending on the machine image you choose,
	// the command to get software on your instance varies. Amazon Linux and CentOS use
	// yum, Debian and Ubuntu use apt-get, and FreeBSD uses pkg. For a complete list,
	// see the Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/getting-started/article/compare-options-choose-lightsail-instance-image).
	UserData *string
}

type CreateInstancesFromSnapshotOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateInstancesFromSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstancesFromSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstancesFromSnapshot{}, middleware.After)
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
	if err = addOpCreateInstancesFromSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstancesFromSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstancesFromSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateInstancesFromSnapshot",
	}
}
