// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/textract/enums"
)

type AnalyzeDocumentInput struct {
	_ struct{} `type:"structure"`

	// The input document as base64-encoded bytes or an Amazon S3 object. If you
	// use the AWS CLI to call Amazon Textract operations, you can't pass image
	// bytes. The document must be an image in JPG or PNG format.
	//
	// If you are using an AWS SDK to call Amazon Textract, you might not need to
	// base64-encode image bytes passed using the Bytes field.
	//
	// Document is a required field
	Document *Document `type:"structure" required:"true"`

	// A list of the types of analysis to perform. Add TABLES to the list to return
	// information about the tables detected in the input document. Add FORMS to
	// return detected fields and the associated text. To perform both types of
	// analysis, add TABLES and FORMS to FeatureTypes.
	//
	// FeatureTypes is a required field
	FeatureTypes []enums.FeatureType `type:"list" required:"true"`
}

// String returns the string representation
func (s AnalyzeDocumentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AnalyzeDocumentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AnalyzeDocumentInput"}

	if s.Document == nil {
		invalidParams.Add(aws.NewErrParamRequired("Document"))
	}

	if s.FeatureTypes == nil {
		invalidParams.Add(aws.NewErrParamRequired("FeatureTypes"))
	}
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			invalidParams.AddNested("Document", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AnalyzeDocumentOutput struct {
	_ struct{} `type:"structure"`

	// The text that's detected and analyzed by AnalyzeDocument.
	Blocks []Block `type:"list"`

	// Metadata about the analyzed document. An example is the number of pages.
	DocumentMetadata *DocumentMetadata `type:"structure"`
}

// String returns the string representation
func (s AnalyzeDocumentOutput) String() string {
	return awsutil.Prettify(s)
}
