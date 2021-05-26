// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the specified product. This operation is run with
// administrator access.
func (c *Client) DescribeProductAsAdmin(ctx context.Context, params *DescribeProductAsAdminInput, optFns ...func(*Options)) (*DescribeProductAsAdminOutput, error) {
	if params == nil {
		params = &DescribeProductAsAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProductAsAdmin", params, optFns, c.addOperationDescribeProductAsAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProductAsAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProductAsAdminInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The product identifier.
	Id *string

	// The product name.
	Name *string

	// The unique identifier of the shared portfolio that the specified product is
	// associated with. You can provide this parameter to retrieve the shared
	// TagOptions associated with the product. If this parameter is provided and if
	// TagOptions sharing is enabled in the portfolio share, the API returns both local
	// and shared TagOptions associated with the product. Otherwise only local
	// TagOptions will be returned.
	SourcePortfolioId *string
}

type DescribeProductAsAdminOutput struct {

	// Information about the associated budgets.
	Budgets []types.BudgetDetail

	// Information about the product view.
	ProductViewDetail *types.ProductViewDetail

	// Information about the provisioning artifacts (also known as versions) for the
	// specified product.
	ProvisioningArtifactSummaries []types.ProvisioningArtifactSummary

	// Information about the TagOptions associated with the product.
	TagOptions []types.TagOptionDetail

	// Information about the tags associated with the product.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeProductAsAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProductAsAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProductAsAdmin{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProductAsAdmin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeProductAsAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeProductAsAdmin",
	}
}
