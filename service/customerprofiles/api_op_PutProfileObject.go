// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds additional objects to customer profiles of a given ObjectType. When adding
// a specific profile object, like a Contact Trace Record (CTR), an inferred
// profile can get created if it is not mapped to an existing profile. The
// resulting profile will only have a phone number populated in the standard
// ProfileObject. Any additional CTRs with the same phone number will be mapped to
// the same inferred profile. When a ProfileObject is created and if a
// ProfileObjectType already exists for the ProfileObject, it will provide data to
// a standard profile depending on the ProfileObjectType definition.
// PutProfileObject needs an ObjectType, which can be created using
// PutProfileObjectType.
func (c *Client) PutProfileObject(ctx context.Context, params *PutProfileObjectInput, optFns ...func(*Options)) (*PutProfileObjectOutput, error) {
	if params == nil {
		params = &PutProfileObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutProfileObject", params, optFns, addOperationPutProfileObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutProfileObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutProfileObjectInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// A string that is serialized from a JSON object.
	//
	// This member is required.
	Object *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string
}

type PutProfileObjectOutput struct {

	// The unique identifier of the profile object generated by the service.
	ProfileObjectUniqueKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutProfileObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutProfileObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutProfileObject{}, middleware.After)
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
	if err = addOpPutProfileObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutProfileObject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutProfileObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "PutProfileObject",
	}
}
