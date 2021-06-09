// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Reserves a contact using specified parameters.
func (c *Client) ReserveContact(ctx context.Context, params *ReserveContactInput, optFns ...func(*Options)) (*ReserveContactOutput, error) {
	if params == nil {
		params = &ReserveContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReserveContact", params, optFns, c.addOperationReserveContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReserveContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ReserveContactInput struct {

	// End time of a contact.
	//
	// This member is required.
	EndTime *time.Time

	// Name of a ground station.
	//
	// This member is required.
	GroundStation *string

	// ARN of a mission profile.
	//
	// This member is required.
	MissionProfileArn *string

	// ARN of a satellite
	//
	// This member is required.
	SatelliteArn *string

	// Start time of a contact.
	//
	// This member is required.
	StartTime *time.Time

	// Tags assigned to a contact.
	Tags map[string]string
}

//
type ReserveContactOutput struct {

	// UUID of a contact.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationReserveContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpReserveContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpReserveContact{}, middleware.After)
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
	if err = addOpReserveContactValidationMiddleware(stack); err != nil {
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
