// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a job that terminates specific launched EC2 Test and Cutover instances.
// This command will not work for any Source Server with a lifecycle.state of
// TESTING, CUTTING_OVER, or CUTOVER.
func (c *Client) TerminateTargetInstances(ctx context.Context, params *TerminateTargetInstancesInput, optFns ...func(*Options)) (*TerminateTargetInstancesOutput, error) {
	if params == nil {
		params = &TerminateTargetInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateTargetInstances", params, optFns, c.addOperationTerminateTargetInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateTargetInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateTargetInstancesInput struct {

	// Terminate Target instance by Source Server IDs.
	//
	// This member is required.
	SourceServerIDs []string

	// Terminate Target instance by Tags.
	Tags map[string]string
}

type TerminateTargetInstancesOutput struct {

	// Terminate Target instance Job response.
	Job *types.Job

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationTerminateTargetInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTerminateTargetInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTerminateTargetInstances{}, middleware.After)
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
	if err = addOpTerminateTargetInstancesValidationMiddleware(stack); err != nil {
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
