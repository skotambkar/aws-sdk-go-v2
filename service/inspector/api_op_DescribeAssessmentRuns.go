// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the assessment runs that are specified by the ARNs of the assessment
// runs.
func (c *Client) DescribeAssessmentRuns(ctx context.Context, params *DescribeAssessmentRunsInput, optFns ...func(*Options)) (*DescribeAssessmentRunsOutput, error) {
	if params == nil {
		params = &DescribeAssessmentRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssessmentRuns", params, optFns, c.addOperationDescribeAssessmentRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssessmentRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssessmentRunsInput struct {

	// The ARN that specifies the assessment run that you want to describe.
	//
	// This member is required.
	AssessmentRunArns []string
}

type DescribeAssessmentRunsOutput struct {

	// Information about the assessment run.
	//
	// This member is required.
	AssessmentRuns []types.AssessmentRun

	// Assessment run details that cannot be described. An error code is provided for
	// each failed item.
	//
	// This member is required.
	FailedItems map[string]types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribeAssessmentRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAssessmentRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAssessmentRuns{}, middleware.After)
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
	if err = addOpDescribeAssessmentRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssessmentRuns(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAssessmentRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "DescribeAssessmentRuns",
	}
}
