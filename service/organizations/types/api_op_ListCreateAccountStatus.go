// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/organizations/enums"
)

type ListCreateAccountStatusInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Use this to limit the number of results you want included per
	// page in the response. If you do not include this parameter, it defaults to
	// a value that is specific to the operation. If additional items exist beyond
	// the maximum you specify, the NextToken response element is present and has
	// a value (is not null). Include that value as the NextToken request parameter
	// in the next call to the operation to get the next part of the results. Note
	// that Organizations might return fewer results than the maximum even when
	// there are more results available. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// Use this parameter if you receive a NextToken response in a previous request
	// that indicates that there is more output available. Set it to the value of
	// the previous call's NextToken response to indicate where the output should
	// continue from.
	NextToken *string `type:"string"`

	// A list of one or more states that you want included in the response. If this
	// parameter isn't present, all requests are included in the response.
	States []enums.CreateAccountState `type:"list"`
}

// String returns the string representation
func (s ListCreateAccountStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCreateAccountStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCreateAccountStatusInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListCreateAccountStatusOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects with details about the requests. Certain elements, such
	// as the accountId number, are present in the output only after the account
	// has been successfully created.
	CreateAccountStatuses []CreateAccountStatus `type:"list"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListCreateAccountStatusOutput) String() string {
	return awsutil.Prettify(s)
}
