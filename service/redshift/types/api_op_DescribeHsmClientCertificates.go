// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHsmClientCertificatesInput struct {
	_ struct{} `type:"structure"`

	// The identifier of a specific HSM client certificate for which you want information.
	// If no identifier is specified, information is returned for all HSM client
	// certificates owned by your AWS customer account.
	HsmClientCertificateIdentifier *string `type:"string"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeHsmClientCertificates request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A tag key or keys for which you want to return all matching HSM client certificates
	// that are associated with the specified key or keys. For example, suppose
	// that you have HSM client certificates that are tagged with keys called owner
	// and environment. If you specify both of these tag keys in the request, Amazon
	// Redshift returns a response with the HSM client certificates that have either
	// or both of these tag keys associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching HSM client
	// certificates that are associated with the specified tag value or values.
	// For example, suppose that you have HSM client certificates that are tagged
	// with values called admin and test. If you specify both of these tag values
	// in the request, Amazon Redshift returns a response with the HSM client certificates
	// that have either or both of these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeHsmClientCertificatesInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeHsmClientCertificatesOutput struct {
	_ struct{} `type:"structure"`

	// A list of the identifiers for one or more HSM client certificates used by
	// Amazon Redshift clusters to store and retrieve database encryption keys in
	// an HSM.
	HsmClientCertificates []HsmClientCertificate `locationNameList:"HsmClientCertificate" type:"list"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeHsmClientCertificatesOutput) String() string {
	return awsutil.Prettify(s)
}
