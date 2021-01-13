// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a new premigration assessment run for one or more individual assessments
// of a migration task. The assessments that you can specify depend on the source
// and target database engine and the migration type defined for the given task. To
// run this operation, your migration task must already be created. After you run
// this operation, you can review the status of each individual assessment. You can
// also run the migration task manually after the assessment run and its individual
// assessments complete.
func (c *Client) StartReplicationTaskAssessmentRun(ctx context.Context, params *StartReplicationTaskAssessmentRunInput, optFns ...func(*Options)) (*StartReplicationTaskAssessmentRunOutput, error) {
	if params == nil {
		params = &StartReplicationTaskAssessmentRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartReplicationTaskAssessmentRun", params, optFns, addOperationStartReplicationTaskAssessmentRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartReplicationTaskAssessmentRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type StartReplicationTaskAssessmentRunInput struct {

	// Unique name to identify the assessment run.
	//
	// This member is required.
	AssessmentRunName *string

	// Amazon Resource Name (ARN) of the migration task associated with the
	// premigration assessment run that you want to start.
	//
	// This member is required.
	ReplicationTaskArn *string

	// Amazon S3 bucket where you want AWS DMS to store the results of this assessment
	// run.
	//
	// This member is required.
	ResultLocationBucket *string

	// ARN of a service role needed to start the assessment run.
	//
	// This member is required.
	ServiceAccessRoleArn *string

	// Space-separated list of names for specific individual assessments that you want
	// to exclude. These names come from the default list of individual assessments
	// that AWS DMS supports for the associated migration task. This task is specified
	// by ReplicationTaskArn. You can't set a value for Exclude if you also set a value
	// for IncludeOnly in the API operation. To identify the names of the default
	// individual assessments that AWS DMS supports for the associated migration task,
	// run the DescribeApplicableIndividualAssessments operation using its own
	// ReplicationTaskArn request parameter.
	Exclude []string

	// Space-separated list of names for specific individual assessments that you want
	// to include. These names come from the default list of individual assessments
	// that AWS DMS supports for the associated migration task. This task is specified
	// by ReplicationTaskArn. You can't set a value for IncludeOnly if you also set a
	// value for Exclude in the API operation. To identify the names of the default
	// individual assessments that AWS DMS supports for the associated migration task,
	// run the DescribeApplicableIndividualAssessments operation using its own
	// ReplicationTaskArn request parameter.
	IncludeOnly []string

	// Encryption mode that you can specify to encrypt the results of this assessment
	// run. If you don't specify this request parameter, AWS DMS stores the assessment
	// run results without encryption. You can specify one of the options following:
	//
	// *
	// "SSE_S3" – The server-side encryption provided as a default by Amazon S3.
	//
	// *
	// "SSE_KMS" – AWS Key Management Service (AWS KMS) encryption. This encryption can
	// use either a custom KMS encryption key that you specify or the default KMS
	// encryption key that DMS provides.
	ResultEncryptionMode *string

	// ARN of a custom KMS encryption key that you specify when you set
	// ResultEncryptionMode to "SSE_KMS".
	ResultKmsKeyArn *string

	// Folder within an Amazon S3 bucket where you want AWS DMS to store the results of
	// this assessment run.
	ResultLocationFolder *string
}

//
type StartReplicationTaskAssessmentRunOutput struct {

	// The premigration assessment run that was started.
	ReplicationTaskAssessmentRun *types.ReplicationTaskAssessmentRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartReplicationTaskAssessmentRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartReplicationTaskAssessmentRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartReplicationTaskAssessmentRun{}, middleware.After)
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
	if err = addOpStartReplicationTaskAssessmentRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartReplicationTaskAssessmentRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartReplicationTaskAssessmentRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "StartReplicationTaskAssessmentRun",
	}
}
