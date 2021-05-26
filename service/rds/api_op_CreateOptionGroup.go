// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new option group. You can create up to 20 option groups.
func (c *Client) CreateOptionGroup(ctx context.Context, params *CreateOptionGroupInput, optFns ...func(*Options)) (*CreateOptionGroupOutput, error) {
	if params == nil {
		params = &CreateOptionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOptionGroup", params, optFns, c.addOperationCreateOptionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOptionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateOptionGroupInput struct {

	// Specifies the name of the engine that this option group should be associated
	// with. Valid Values:
	//
	// * mariadb
	//
	// * mysql
	//
	// * oracle-ee
	//
	// * oracle-se2
	//
	// *
	// oracle-se1
	//
	// * oracle-se
	//
	// * postgres
	//
	// * sqlserver-ee
	//
	// * sqlserver-se
	//
	// *
	// sqlserver-ex
	//
	// * sqlserver-web
	//
	// This member is required.
	EngineName *string

	// Specifies the major version of the engine that this option group should be
	// associated with.
	//
	// This member is required.
	MajorEngineVersion *string

	// The description of the option group.
	//
	// This member is required.
	OptionGroupDescription *string

	// Specifies the name of the option group to be created. Constraints:
	//
	// * Must be 1
	// to 255 letters, numbers, or hyphens
	//
	// * First character must be a letter
	//
	// * Can't
	// end with a hyphen or contain two consecutive hyphens
	//
	// Example: myoptiongroup
	//
	// This member is required.
	OptionGroupName *string

	// Tags to assign to the option group.
	Tags []types.Tag
}

type CreateOptionGroupOutput struct {

	//
	OptionGroup *types.OptionGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateOptionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateOptionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateOptionGroup{}, middleware.After)
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
	if err = addOpCreateOptionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOptionGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOptionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateOptionGroup",
	}
}
