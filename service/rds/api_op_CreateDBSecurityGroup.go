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

// Creates a new DB security group. DB security groups control access to a DB
// instance. A DB security group controls access to EC2-Classic DB instances that
// are not in a VPC.
func (c *Client) CreateDBSecurityGroup(ctx context.Context, params *CreateDBSecurityGroupInput, optFns ...func(*Options)) (*CreateDBSecurityGroupOutput, error) {
	if params == nil {
		params = &CreateDBSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDBSecurityGroup", params, optFns, c.addOperationCreateDBSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDBSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateDBSecurityGroupInput struct {

	// The description for the DB security group.
	//
	// This member is required.
	DBSecurityGroupDescription *string

	// The name for the DB security group. This value is stored as a lowercase string.
	// Constraints:
	//
	// * Must be 1 to 255 letters, numbers, or hyphens.
	//
	// * First
	// character must be a letter
	//
	// * Can't end with a hyphen or contain two consecutive
	// hyphens
	//
	// * Must not be "Default"
	//
	// Example: mysecuritygroup
	//
	// This member is required.
	DBSecurityGroupName *string

	// Tags to assign to the DB security group.
	Tags []types.Tag
}

type CreateDBSecurityGroupOutput struct {

	// Contains the details for an Amazon RDS DB security group. This data type is used
	// as a response element in the DescribeDBSecurityGroups action.
	DBSecurityGroup *types.DBSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateDBSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateDBSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateDBSecurityGroup{}, middleware.After)
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
	if err = addOpCreateDBSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDBSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDBSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "CreateDBSecurityGroup",
	}
}
