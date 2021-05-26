// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Dynamically decreases the number of replicas in a Redis (cluster mode disabled)
// replication group or the number of replica nodes in one or more node groups
// (shards) of a Redis (cluster mode enabled) replication group. This operation is
// performed with no cluster down time.
func (c *Client) DecreaseReplicaCount(ctx context.Context, params *DecreaseReplicaCountInput, optFns ...func(*Options)) (*DecreaseReplicaCountOutput, error) {
	if params == nil {
		params = &DecreaseReplicaCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DecreaseReplicaCount", params, optFns, c.addOperationDecreaseReplicaCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DecreaseReplicaCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DecreaseReplicaCountInput struct {

	// If True, the number of replica nodes is decreased immediately.
	// ApplyImmediately=False is not currently supported.
	//
	// This member is required.
	ApplyImmediately bool

	// The id of the replication group from which you want to remove replica nodes.
	//
	// This member is required.
	ReplicationGroupId *string

	// The number of read replica nodes you want at the completion of this operation.
	// For Redis (cluster mode disabled) replication groups, this is the number of
	// replica nodes in the replication group. For Redis (cluster mode enabled)
	// replication groups, this is the number of replica nodes in each of the
	// replication group's node groups. The minimum number of replicas in a shard or
	// replication group is:
	//
	// * Redis (cluster mode disabled)
	//
	// * If Multi-AZ is
	// enabled: 1
	//
	// * If Multi-AZ is not enabled: 0
	//
	// * Redis (cluster mode enabled): 0
	// (though you will not be able to failover to a replica if your primary node
	// fails)
	NewReplicaCount *int32

	// A list of ConfigureShard objects that can be used to configure each shard in a
	// Redis (cluster mode enabled) replication group. The ConfigureShard has three
	// members: NewReplicaCount, NodeGroupId, and PreferredAvailabilityZones.
	ReplicaConfiguration []types.ConfigureShard

	// A list of the node ids to remove from the replication group or node group
	// (shard).
	ReplicasToRemove []string
}

type DecreaseReplicaCountOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDecreaseReplicaCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDecreaseReplicaCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDecreaseReplicaCount{}, middleware.After)
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
	if err = addOpDecreaseReplicaCountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDecreaseReplicaCount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDecreaseReplicaCount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DecreaseReplicaCount",
	}
}
