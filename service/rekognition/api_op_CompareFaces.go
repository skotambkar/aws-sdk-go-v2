// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Compares a face in the source input image with each of the 100 largest faces
// detected in the target input image. If the source image contains multiple faces,
// the service detects the largest face and compares it with each face detected in
// the target image. CompareFaces uses machine learning algorithms, which are
// probabilistic. A false negative is an incorrect prediction that a face in the
// target image has a low similarity confidence score when compared to the face in
// the source image. To reduce the probability of false negatives, we recommend
// that you compare the target image against multiple source images. If you plan to
// use CompareFaces to make a decision that impacts an individual's rights,
// privacy, or access to services, we recommend that you pass the result to a human
// for review and further validation before taking action. You pass the input and
// target images either as base64-encoded image bytes or as references to images in
// an Amazon S3 bucket. If you use the AWS CLI to call Amazon Rekognition
// operations, passing image bytes isn't supported. The image must be formatted as
// a PNG or JPEG file. In response, the operation returns an array of face matches
// ordered by similarity score in descending order. For each face match, the
// response provides a bounding box of the face, facial landmarks, pose details
// (pitch, role, and yaw), quality (brightness and sharpness), and confidence value
// (indicating the level of confidence that the bounding box contains a face). The
// response also provides a similarity score, which indicates how closely the faces
// match. By default, only faces with a similarity score of greater than or equal
// to 80% are returned in the response. You can change this value by specifying the
// SimilarityThreshold parameter. CompareFaces also returns an array of faces that
// don't match the source image. For each face, it returns a bounding box,
// confidence value, landmarks, pose details, and quality. The response also
// returns information about the face in the source image, including the bounding
// box of the face and confidence value. The QualityFilter input parameter allows
// you to filter out detected faces that don’t meet a required quality bar. The
// quality bar is based on a variety of common use cases. Use QualityFilter to set
// the quality bar by specifying LOW, MEDIUM, or HIGH. If you do not want to filter
// detected faces, specify NONE. The default value is NONE. If the image doesn't
// contain Exif metadata, CompareFaces returns orientation information for the
// source and target images. Use these values to display the images with the
// correct image orientation. If no faces are detected in the source or target
// images, CompareFaces returns an InvalidParameterException error. This is a
// stateless API operation. That is, data returned by this operation doesn't
// persist. For an example, see Comparing Faces in Images in the Amazon Rekognition
// Developer Guide. This operation requires permissions to perform the
// rekognition:CompareFaces action.
func (c *Client) CompareFaces(ctx context.Context, params *CompareFacesInput, optFns ...func(*Options)) (*CompareFacesOutput, error) {
	if params == nil {
		params = &CompareFacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompareFaces", params, optFns, c.addOperationCompareFacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompareFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompareFacesInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	SourceImage *types.Image

	// The target image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	TargetImage *types.Image

	// A filter that specifies a quality bar for how much filtering is done to identify
	// faces. Filtered faces aren't compared. If you specify AUTO, Amazon Rekognition
	// chooses the quality bar. If you specify LOW, MEDIUM, or HIGH, filtering removes
	// all faces that don’t meet the chosen quality bar. The quality bar is based on a
	// variety of common use cases. Low-quality detections can occur for a number of
	// reasons. Some examples are an object that's misidentified as a face, a face
	// that's too blurry, or a face with a pose that's too extreme to use. If you
	// specify NONE, no filtering is performed. The default value is NONE. To use
	// quality filtering, the collection you are using must be associated with version
	// 3 of the face model or higher.
	QualityFilter types.QualityFilter

	// The minimum level of confidence in the face matches that a match must meet to be
	// included in the FaceMatches array.
	SimilarityThreshold *float32
}

type CompareFacesOutput struct {

	// An array of faces in the target image that match the source image face. Each
	// CompareFacesMatch object provides the bounding box, the confidence level that
	// the bounding box contains a face, and the similarity score for the face in the
	// bounding box and the face in the source image.
	FaceMatches []types.CompareFacesMatch

	// The face in the source image that was used for comparison.
	SourceImageFace *types.ComparedSourceImageFace

	// The value of SourceImageOrientationCorrection is always null. If the input image
	// is in .jpeg format, it might contain exchangeable image file format (Exif)
	// metadata that includes the image's orientation. Amazon Rekognition uses this
	// orientation information to perform image correction. The bounding box
	// coordinates are translated to represent object locations after the orientation
	// information in the Exif metadata is used to correct the image orientation.
	// Images in .png format don't contain Exif metadata. Amazon Rekognition doesn’t
	// perform image correction for images in .png format and .jpeg images without
	// orientation information in the image Exif metadata. The bounding box coordinates
	// aren't translated and represent the object locations before the image is
	// rotated.
	SourceImageOrientationCorrection types.OrientationCorrection

	// The value of TargetImageOrientationCorrection is always null. If the input image
	// is in .jpeg format, it might contain exchangeable image file format (Exif)
	// metadata that includes the image's orientation. Amazon Rekognition uses this
	// orientation information to perform image correction. The bounding box
	// coordinates are translated to represent object locations after the orientation
	// information in the Exif metadata is used to correct the image orientation.
	// Images in .png format don't contain Exif metadata. Amazon Rekognition doesn’t
	// perform image correction for images in .png format and .jpeg images without
	// orientation information in the image Exif metadata. The bounding box coordinates
	// aren't translated and represent the object locations before the image is
	// rotated.
	TargetImageOrientationCorrection types.OrientationCorrection

	// An array of faces in the target image that did not match the source image face.
	UnmatchedFaces []types.ComparedFace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationCompareFacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCompareFaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCompareFaces{}, middleware.After)
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
	if err = addOpCompareFacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompareFaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCompareFaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CompareFaces",
	}
}
