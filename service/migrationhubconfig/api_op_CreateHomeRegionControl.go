// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubconfig

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API sets up the home region for the calling account only.
func (c *Client) CreateHomeRegionControl(ctx context.Context, params *CreateHomeRegionControlInput, optFns ...func(*Options)) (*CreateHomeRegionControlOutput, error) {
	if params == nil {
		params = &CreateHomeRegionControlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHomeRegionControl", params, optFns, addOperationCreateHomeRegionControlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHomeRegionControlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateHomeRegionControlInput struct {

	// The name of the home region of the calling account.
	//
	// This member is required.
	HomeRegion *string

	// The account for which this command sets up a home region control. The Target is
	// always of type ACCOUNT.
	//
	// This member is required.
	Target *types.Target

	// Optional Boolean flag to indicate whether any effect should take place. It tests
	// whether the caller has permission to make the call.
	DryRun bool
}

type CreateHomeRegionControlOutput struct {

	// This object is the HomeRegionControl object that's returned by a successful call
	// to CreateHomeRegionControl.
	HomeRegionControl *types.HomeRegionControl

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHomeRegionControlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHomeRegionControl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHomeRegionControl{}, middleware.After)
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
	if err = addOpCreateHomeRegionControlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHomeRegionControl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateHomeRegionControl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "CreateHomeRegionControl",
	}
}
