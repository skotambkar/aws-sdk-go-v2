// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/enums"
)

type BatchDetectSentimentInput struct {
	_ struct{} `type:"structure"`

	// The language of the input documents. You can specify any of the primary languages
	// supported by Amazon Comprehend: German ("de"), English ("en"), Spanish ("es"),
	// French ("fr"), Italian ("it"), or Portuguese ("pt"). All documents must be
	// in the same language.
	//
	// LanguageCode is a required field
	LanguageCode enums.LanguageCode `type:"string" required:"true" enum:"true"`

	// A list containing the text of the input documents. The list can contain a
	// maximum of 25 documents. Each document must contain fewer that 5,000 bytes
	// of UTF-8 encoded characters.
	//
	// TextList is a required field
	TextList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectSentimentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDetectSentimentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDetectSentimentInput"}
	if len(s.LanguageCode) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("LanguageCode"))
	}

	if s.TextList == nil {
		invalidParams.Add(aws.NewErrParamRequired("TextList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDetectSentimentOutput struct {
	_ struct{} `type:"structure"`

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order
	// of the documents in the input list. If there are no errors in the batch,
	// the ErrorList is empty.
	//
	// ErrorList is a required field
	ErrorList []BatchItemError `type:"list" required:"true"`

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the documents
	// in the input list. If all of the documents contain an error, the ResultList
	// is empty.
	//
	// ResultList is a required field
	ResultList []BatchDetectSentimentItemResult `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDetectSentimentOutput) String() string {
	return awsutil.Prettify(s)
}
