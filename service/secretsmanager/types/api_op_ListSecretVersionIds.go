// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSecretVersionIdsInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Specifies that you want the results to include versions that do
	// not have any staging labels attached to them. Such versions are considered
	// deprecated and are subject to deletion by Secrets Manager as needed.
	IncludeDeprecated *bool `type:"boolean"`

	// (Optional) Limits the number of results that you want to include in the response.
	// If you don't include this parameter, it defaults to a value that's specific
	// to the operation. If additional items exist beyond the maximum you specify,
	// the NextToken response element is present and has a value (isn't null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Secrets Manager might return
	// fewer results than the maximum even when there are more results available.
	// You should check NextToken after every operation to ensure that you receive
	// all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// (Optional) Use this parameter in a request if you receive a NextToken response
	// in a previous request that indicates that there's more output available.
	// In a subsequent call, set it to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string `min:"1" type:"string"`

	// The identifier for the secret containing the versions you want to list. You
	// can specify either the Amazon Resource Name (ARN) or the friendly name of
	// the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListSecretVersionIdsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSecretVersionIdsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSecretVersionIdsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSecretVersionIdsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the secret.
	//
	// Secrets Manager automatically adds several random characters to the name
	// at the end of the ARN when you initially create a secret. This affects only
	// the ARN and not the actual friendly name. This ensures that if you create
	// a new secret with the same name as an old secret that you previously deleted,
	// then users with access to the old secret don't automatically get access to
	// the new secret because the ARNs are different.
	ARN *string `min:"20" type:"string"`

	// The friendly name of the secret.
	Name *string `min:"1" type:"string"`

	// If present in the response, this value indicates that there's more output
	// available than what's included in the current response. This can occur even
	// when the response includes no values at all, such as when you ask for a filtered
	// view of a very long list. Use this value in the NextToken request parameter
	// in a subsequent call to the operation to continue processing and get the
	// next part of the output. You should repeat this until the NextToken response
	// element comes back empty (as null).
	NextToken *string `min:"1" type:"string"`

	// The list of the currently available versions of the specified secret.
	Versions []SecretVersionsListEntry `type:"list"`
}

// String returns the string representation
func (s ListSecretVersionIdsOutput) String() string {
	return awsutil.Prettify(s)
}