// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets one or more parameters of the specified parameter group to their default
// values and sets the source values of the parameters to "engine-default". To
// reset the entire parameter group specify the ResetAllParameters parameter. For
// parameter changes to take effect you must reboot any associated clusters.
func (c *Client) ResetClusterParameterGroup(ctx context.Context, params *ResetClusterParameterGroupInput, optFns ...func(*Options)) (*ResetClusterParameterGroupOutput, error) {
	if params == nil {
		params = &ResetClusterParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetClusterParameterGroup", params, optFns, c.addOperationResetClusterParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetClusterParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ResetClusterParameterGroupInput struct {

	// The name of the cluster parameter group to be reset.
	//
	// This member is required.
	ParameterGroupName *string

	// An array of names of parameters to be reset. If ResetAllParameters option is not
	// used, then at least one parameter name must be supplied. Constraints: A maximum
	// of 20 parameters can be reset in a single request.
	Parameters []types.Parameter

	// If true, all parameters in the specified parameter group will be reset to their
	// default values. Default: true
	ResetAllParameters bool
}

//
type ResetClusterParameterGroupOutput struct {

	// The name of the cluster parameter group.
	ParameterGroupName *string

	// The status of the parameter group. For example, if you made a change to a
	// parameter group name-value pair, then the change could be pending a reboot of an
	// associated cluster.
	ParameterGroupStatus *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationResetClusterParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResetClusterParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResetClusterParameterGroup{}, middleware.After)
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
	if err = addOpResetClusterParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetClusterParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetClusterParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ResetClusterParameterGroup",
	}
}
