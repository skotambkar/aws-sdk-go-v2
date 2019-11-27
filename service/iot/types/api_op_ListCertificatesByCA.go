// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input to the ListCertificatesByCA operation.
type ListCertificatesByCAInput struct {
	_ struct{} `type:"structure"`

	// Specifies the order for results. If True, the results are returned in ascending
	// order, based on the creation date.
	AscendingOrder *bool `location:"querystring" locationName:"isAscendingOrder" type:"boolean"`

	// The ID of the CA certificate. This operation will list all registered device
	// certificate that were signed by this CA certificate.
	//
	// CaCertificateId is a required field
	CaCertificateId *string `location:"uri" locationName:"caCertificateId" min:"64" type:"string" required:"true"`

	// The marker for the next set of results.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The result page size.
	PageSize *int64 `location:"querystring" locationName:"pageSize" min:"1" type:"integer"`
}

// String returns the string representation
func (s ListCertificatesByCAInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCertificatesByCAInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListCertificatesByCAInput"}

	if s.CaCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CaCertificateId"))
	}
	if s.CaCertificateId != nil && len(*s.CaCertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CaCertificateId", 64))
	}
	if s.PageSize != nil && *s.PageSize < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PageSize", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output of the ListCertificatesByCA operation.
type ListCertificatesByCAOutput struct {
	_ struct{} `type:"structure"`

	// The device certificates signed by the specified CA certificate.
	Certificates []Certificate `locationName:"certificates" type:"list"`

	// The marker for the next set of results, or null if there are no additional
	// results.
	NextMarker *string `locationName:"nextMarker" type:"string"`
}

// String returns the string representation
func (s ListCertificatesByCAOutput) String() string {
	return awsutil.Prettify(s)
}
