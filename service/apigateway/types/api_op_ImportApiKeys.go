// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/enums"
)

// The POST request to import API keys from an external source, such as a CSV-formatted
// file.
type ImportApiKeysInput struct {
	_ struct{} `type:"structure" payload:"Body"`

	// The payload of the POST request to import API keys. For the payload format,
	// see API Key File Format (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-key-file-format.html).
	//
	// Body is a required field
	Body []byte `locationName:"body" type:"blob" required:"true"`

	// A query parameter to indicate whether to rollback ApiKey importation (true)
	// or not (false) when error is encountered.
	FailOnWarnings *bool `location:"querystring" locationName:"failonwarnings" type:"boolean"`

	// A query parameter to specify the input format to imported API keys. Currently,
	// only the csv format is supported.
	//
	// Format is a required field
	Format enums.ApiKeysFormat `location:"querystring" locationName:"format" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s ImportApiKeysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ImportApiKeysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ImportApiKeysInput"}

	if s.Body == nil {
		invalidParams.Add(aws.NewErrParamRequired("Body"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The identifier of an ApiKey used in a UsagePlan.
type ImportApiKeysOutput struct {
	_ struct{} `type:"structure"`

	// A list of all the ApiKey identifiers.
	Ids []string `locationName:"ids" type:"list"`

	// A list of warning messages.
	Warnings []string `locationName:"warnings" type:"list"`
}

// String returns the string representation
func (s ImportApiKeysOutput) String() string {
	return awsutil.Prettify(s)
}
