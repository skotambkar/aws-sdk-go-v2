// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a dataset group, which holds a collection of related datasets. You can
// add datasets to the dataset group when you create the dataset group, or later by
// using the UpdateDatasetGroup operation. After creating a dataset group and
// adding datasets, you use the dataset group when you create a predictor. For more
// information, see howitworks-datasets-groups. To get a list of all your datasets
// groups, use the ListDatasetGroups operation. The Status of a dataset group must
// be ACTIVE before you can use the dataset group to create a predictor. To get the
// status, use the DescribeDatasetGroup operation.
func (c *Client) CreateDatasetGroup(ctx context.Context, params *CreateDatasetGroupInput, optFns ...func(*Options)) (*CreateDatasetGroupOutput, error) {
	if params == nil {
		params = &CreateDatasetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatasetGroup", params, optFns, addOperationCreateDatasetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatasetGroupInput struct {

	// A name for the dataset group.
	//
	// This member is required.
	DatasetGroupName *string

	// The domain associated with the dataset group. When you add a dataset to a
	// dataset group, this value and the value specified for the Domain parameter of
	// the CreateDataset operation must match. The Domain and DatasetType that you
	// choose determine the fields that must be present in training data that you
	// import to a dataset. For example, if you choose the RETAIL domain and
	// TARGET_TIME_SERIES as the DatasetType, Amazon Forecast requires that item_id,
	// timestamp, and demand fields are present in your data. For more information, see
	// howitworks-datasets-groups.
	//
	// This member is required.
	Domain types.Domain

	// An array of Amazon Resource Names (ARNs) of the datasets that you want to
	// include in the dataset group.
	DatasetArns []string

	// The optional metadata that you apply to the dataset group to help you categorize
	// and organize them. Each tag consists of a key and an optional value, both of
	// which you define. The following basic restrictions apply to tags:
	//
	// * Maximum
	// number of tags per resource - 50.
	//
	// * For each resource, each tag key must be
	// unique, and each tag key can have only one value.
	//
	// * Maximum key length - 128
	// Unicode characters in UTF-8.
	//
	// * Maximum value length - 256 Unicode characters in
	// UTF-8.
	//
	// * If your tagging schema is used across multiple services and resources,
	// remember that other services may have restrictions on allowed characters.
	// Generally allowed characters are: letters, numbers, and spaces representable in
	// UTF-8, and the following characters: + - = . _ : / @.
	//
	// * Tag keys and values are
	// case sensitive.
	//
	// * Do not use aws:, AWS:, or any upper or lowercase combination
	// of such as a prefix for keys as it is reserved for AWS use. You cannot edit or
	// delete tag keys with this prefix. Values can have this prefix. If a tag value
	// has aws as its prefix but the key does not, then Forecast considers it to be a
	// user tag and will count against the limit of 50 tags. Tags with only the key
	// prefix of aws do not count against your tags per resource limit.
	Tags []types.Tag
}

type CreateDatasetGroupOutput struct {

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDatasetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDatasetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDatasetGroup{}, middleware.After)
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
	if err = addOpCreateDatasetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatasetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatasetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "CreateDatasetGroup",
	}
}
