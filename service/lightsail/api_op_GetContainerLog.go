// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the log events of a container of your Amazon Lightsail container
// service. If your container service has more than one node (i.e., a scale greater
// than 1), then the log events that are returned for the specified container are
// merged from all nodes on your container service. Container logs are retained for
// a certain amount of time. For more information, see Amazon Lightsail endpoints
// and quotas (https://docs.aws.amazon.com/general/latest/gr/lightsail.html) in the
// AWS General Reference.
func (c *Client) GetContainerLog(ctx context.Context, params *GetContainerLogInput, optFns ...func(*Options)) (*GetContainerLogOutput, error) {
	if params == nil {
		params = &GetContainerLogInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContainerLog", params, optFns, addOperationGetContainerLogMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContainerLogOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContainerLogInput struct {

	// The name of the container that is either running or previously ran on the
	// container service for which to return a log.
	//
	// This member is required.
	ContainerName *string

	// The name of the container service for which to get a container log.
	//
	// This member is required.
	ServiceName *string

	// The end of the time interval for which to get log data. Constraints:
	//
	// *
	// Specified in Coordinated Universal Time (UTC).
	//
	// * Specified in the Unix time
	// format. For example, if you wish to use an end time of October 1, 2018, at 9 PM
	// UTC, specify 1538427600 as the end time.
	//
	// You can convert a human-friendly time
	// to Unix time format using a converter like Epoch converter
	// (https://www.epochconverter.com/).
	EndTime *time.Time

	// The pattern to use to filter the returned log events to a specific term. The
	// following are a few examples of filter patterns that you can specify:
	//
	// * To
	// return all log events, specify a filter pattern of "".
	//
	// * To exclude log events
	// that contain the ERROR term, and return all other log events, specify a filter
	// pattern of "-ERROR".
	//
	// * To return log events that contain the ERROR term,
	// specify a filter pattern of "ERROR".
	//
	// * To return log events that contain both
	// the ERROR and Exception terms, specify a filter pattern of "ERROR Exception".
	//
	// *
	// To return log events that contain the ERROR or the Exception term, specify a
	// filter pattern of "?ERROR ?Exception".
	FilterPattern *string

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetContainerLog request. If your results are
	// paginated, the response will return a next page token that you can specify as
	// the page token in a subsequent request.
	PageToken *string

	// The start of the time interval for which to get log data. Constraints:
	//
	// *
	// Specified in Coordinated Universal Time (UTC).
	//
	// * Specified in the Unix time
	// format. For example, if you wish to use a start time of October 1, 2018, at 8 PM
	// UTC, specify 1538424000 as the start time.
	//
	// You can convert a human-friendly
	// time to Unix time format using a converter like Epoch converter
	// (https://www.epochconverter.com/).
	StartTime *time.Time
}

type GetContainerLogOutput struct {

	// An array of objects that describe the log events of a container.
	LogEvents []types.ContainerServiceLogEvent

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetContainerLog request and specify the next
	// page token using the pageToken parameter.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetContainerLogMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetContainerLog{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetContainerLog{}, middleware.After)
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
	if err = addOpGetContainerLogValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContainerLog(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContainerLog(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetContainerLog",
	}
}
