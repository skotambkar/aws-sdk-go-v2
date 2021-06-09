// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes stack instances for the specified accounts, in the specified Regions.
func (c *Client) DeleteStackInstances(ctx context.Context, params *DeleteStackInstancesInput, optFns ...func(*Options)) (*DeleteStackInstancesOutput, error) {
	if params == nil {
		params = &DeleteStackInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteStackInstances", params, optFns, c.addOperationDeleteStackInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteStackInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteStackInstancesInput struct {

	// The Regions where you want to delete stack set instances.
	//
	// This member is required.
	Regions []string

	// Removes the stack instances from the specified stack set, but doesn't delete the
	// stacks. You can't reassociate a retained stack or add an existing, saved stack
	// to a new stack set. For more information, see Stack set operation options
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options).
	//
	// This member is required.
	RetainStacks bool

	// The name or unique ID of the stack set that you want to delete stack instances
	// for.
	//
	// This member is required.
	StackSetName *string

	// [Self-managed permissions] The names of the AWS accounts that you want to delete
	// stack instances for. You can specify Accounts or DeploymentTargets, but not
	// both.
	Accounts []string

	// [Service-managed permissions] Specifies whether you are acting as an account
	// administrator in the organization's management account or as a delegated
	// administrator in a member account. By default, SELF is specified. Use SELF for
	// stack sets with self-managed permissions.
	//
	// * If you are signed in to the
	// management account, specify SELF.
	//
	// * If you are signed in to a delegated
	// administrator account, specify DELEGATED_ADMIN. Your AWS account must be
	// registered as a delegated administrator in the management account. For more
	// information, see Register a delegated administrator
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-orgs-delegated-admin.html)
	// in the AWS CloudFormation User Guide.
	CallAs types.CallAs

	// [Service-managed permissions] The AWS Organizations accounts from which to
	// delete stack instances. You can specify Accounts or DeploymentTargets, but not
	// both.
	DeploymentTargets *types.DeploymentTargets

	// The unique identifier for this stack set operation. If you don't specify an
	// operation ID, the SDK generates one automatically. The operation ID also
	// functions as an idempotency token, to ensure that AWS CloudFormation performs
	// the stack set operation only once, even if you retry the request multiple times.
	// You can retry stack set operation requests to ensure that AWS CloudFormation
	// successfully received them. Repeating this stack set operation with a new
	// operation ID retries all stack instances whose status is OUTDATED.
	OperationId *string

	// Preferences for how AWS CloudFormation performs this stack set operation.
	OperationPreferences *types.StackSetOperationPreferences
}

type DeleteStackInstancesOutput struct {

	// The unique identifier for this stack set operation.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDeleteStackInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteStackInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteStackInstances{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteStackInstancesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteStackInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStackInstances(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteStackInstances struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteStackInstances) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteStackInstances) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteStackInstancesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteStackInstancesInput ")
	}

	if input.OperationId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.OperationId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteStackInstancesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteStackInstances{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteStackInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeleteStackInstances",
	}
}
