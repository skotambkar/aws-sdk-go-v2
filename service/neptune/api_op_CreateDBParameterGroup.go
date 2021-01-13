// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new DB parameter group. A DB parameter group is initially created with
// the default parameters for the database engine used by the DB instance. To
// provide custom values for any of the parameters, you must modify the group after
// creating it using ModifyDBParameterGroup. Once you've created a DB parameter
// group, you need to associate it with your DB instance using ModifyDBInstance.
// When you associate a new DB parameter group with a running DB instance, you need
// to reboot the DB instance without failover for the new DB parameter group and
// associated settings to take effect. After you create a DB parameter group, you
// should wait at least 5 minutes before creating your first DB instance that uses
// that DB parameter group as the default parameter group. This allows Amazon
// Neptune to fully complete the create action before the parameter group is used
// as the default for a new DB instance. This is especially important for
// parameters that are critical when creating the default database for a DB
// instance, such as the character set for the default database defined by the
// character_set_database parameter. You can use the Parameter Groups option of the
// Amazon Neptune console or the DescribeDBParameters command to verify that your
// DB parameter group has been created or modified.
func (c *Client) CreateDBParameterGroup(ctx context.Context, params *CreateDBParameterGroupInput, optFns ...func(*Options)) (*CreateDBParameterGroupOutput, error) {
	if params == nil {
		params = &CreateDBParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBParameterGroup", params, optFns, addOperationCreateDBParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDBParameterGroupInput struct {

	// The DB parameter group family name. A DB parameter group can be associated with
	// one and only one DB parameter group family, and can be applied only to a DB
	// instance running a database engine and engine version compatible with that DB
	// parameter group family.
	//
	// This member is required.
	DBParameterGroupFamily *string

	// The name of the DB parameter group. Constraints:
	//
	// * Must be 1 to 255 letters,
	// numbers, or hyphens.
	//
	// * First character must be a letter
	//
	// * Cannot end with a
	// hyphen or contain two consecutive hyphens
	//
	// This value is stored as a lowercase
	// string.
	//
	// This member is required.
	DBParameterGroupName *string

	// The description for the DB parameter group.
	//
	// This member is required.
	Description *string

	// The tags to be assigned to the new DB parameter group.
	Tags []types.Tag
}

type CreateDBParameterGroupOutput struct {

	// Contains the details of an Amazon Neptune DB parameter group. This data type is
	// used as a response element in the DescribeDBParameterGroups action.
	DBParameterGroup *types.DBParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBParameterGroup{}, middleware.After)
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
	if err = addOpCreateDBParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBParameterGroup",
	}
}
