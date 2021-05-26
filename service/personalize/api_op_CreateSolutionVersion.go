// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Trains or retrains an active solution. A solution is created using the
// CreateSolution operation and must be in the ACTIVE state before calling
// CreateSolutionVersion. A new version of the solution is created every time you
// call this operation. Status A solution version can be in one of the following
// states:
//
// * CREATE PENDING
//
// * CREATE IN_PROGRESS
//
// * ACTIVE
//
// * CREATE FAILED
//
// *
// CREATE STOPPING
//
// * CREATE STOPPED
//
// To get the status of the version, call
// DescribeSolutionVersion. Wait until the status shows as ACTIVE before calling
// CreateCampaign. If the status shows as CREATE FAILED, the response includes a
// failureReason key, which describes why the job failed. Related APIs
//
// *
// ListSolutionVersions
//
// * DescribeSolutionVersion
//
// * ListSolutions
//
// *
// CreateSolution
//
// * DescribeSolution
//
// * DeleteSolution
func (c *Client) CreateSolutionVersion(ctx context.Context, params *CreateSolutionVersionInput, optFns ...func(*Options)) (*CreateSolutionVersionOutput, error) {
	if params == nil {
		params = &CreateSolutionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSolutionVersion", params, optFns, c.addOperationCreateSolutionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSolutionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolutionVersionInput struct {

	// The Amazon Resource Name (ARN) of the solution containing the training
	// configuration information.
	//
	// This member is required.
	SolutionArn *string

	// The scope of training to be performed when creating the solution version. The
	// FULL option trains the solution version based on the entirety of the input
	// solution's training data, while the UPDATE option processes only the data that
	// has changed in comparison to the input solution. Choose UPDATE when you want to
	// incrementally update your solution version instead of creating an entirely new
	// one. The UPDATE option can only be used when you already have an active solution
	// version created from the input solution using the FULL option and the input
	// solution was trained with the User-Personalization
	// (https://docs.aws.amazon.com/personalize/latest/dg/native-recipe-new-item-USER_PERSONALIZATION.html)
	// recipe or the HRNN-Coldstart
	// (https://docs.aws.amazon.com/personalize/latest/dg/native-recipe-hrnn-coldstart.html)
	// recipe.
	TrainingMode types.TrainingMode
}

type CreateSolutionVersionOutput struct {

	// The ARN of the new solution version.
	SolutionVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCreateSolutionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSolutionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSolutionVersion{}, middleware.After)
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
	if err = addOpCreateSolutionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolutionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSolutionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateSolutionVersion",
	}
}
