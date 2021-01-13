// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The example tests basic map serialization.
func (c *Client) XmlMaps(ctx context.Context, params *XmlMapsInput, optFns ...func(*Options)) (*XmlMapsOutput, error) {
	if params == nil {
		params = &XmlMapsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlMaps", params, optFns, addOperationXmlMapsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlMapsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlMapsInput struct {
	MyMap map[string]types.GreetingStruct
}

type XmlMapsOutput struct {
	MyMap map[string]types.GreetingStruct

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlMapsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpXmlMaps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpXmlMaps{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opXmlMaps(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opXmlMaps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlMaps",
	}
}
