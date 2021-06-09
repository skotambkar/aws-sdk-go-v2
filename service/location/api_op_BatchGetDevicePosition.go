// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A batch request to retrieve all device positions.
func (c *Client) BatchGetDevicePosition(ctx context.Context, params *BatchGetDevicePositionInput, optFns ...func(*Options)) (*BatchGetDevicePositionOutput, error) {
	if params == nil {
		params = &BatchGetDevicePositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetDevicePosition", params, optFns, c.addOperationBatchGetDevicePositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetDevicePositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetDevicePositionInput struct {

	// Devices whose position you want to retrieve.
	//
	// * For example, for two devices:
	// device-ids=DeviceId1&device-ids=DeviceId2
	//
	// This member is required.
	DeviceIds []string

	// The tracker resource retrieving the device position.
	//
	// This member is required.
	TrackerName *string
}

type BatchGetDevicePositionOutput struct {

	// Contains device position details such as the device ID, position, and timestamps
	// for when the position was received and sampled.
	//
	// This member is required.
	DevicePositions []types.DevicePosition

	// Contains error details for each device that failed to send its position to the
	// tracker resource.
	//
	// This member is required.
	Errors []types.BatchGetDevicePositionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationBatchGetDevicePositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetDevicePosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetDevicePosition{}, middleware.After)
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
	if err = addOpBatchGetDevicePositionValidationMiddleware(stack); err != nil {
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
