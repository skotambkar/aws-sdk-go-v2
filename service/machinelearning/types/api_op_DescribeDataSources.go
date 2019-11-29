// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/enums"
)

type DescribeDataSourcesInput struct {
	_ struct{} `type:"structure"`

	// The equal to operator. The DataSource results will have FilterVariable values
	// that exactly match the value specified with EQ.
	EQ *string `type:"string"`

	// Use one of the following variables to filter a list of DataSource:
	//
	//    * CreatedAt - Sets the search criteria to DataSource creation dates.
	//
	//    * Status - Sets the search criteria to DataSource statuses.
	//
	//    * Name - Sets the search criteria to the contents of DataSource Name.
	//
	//    * DataUri - Sets the search criteria to the URI of data files used to
	//    create the DataSource. The URI can identify either a file or an Amazon
	//    Simple Storage Service (Amazon S3) bucket or directory.
	//
	//    * IAMUser - Sets the search criteria to the user account that invoked
	//    the DataSource creation.
	FilterVariable enums.DataSourceFilterVariable `type:"string" enum:"true"`

	// The greater than or equal to operator. The DataSource results will have FilterVariable
	// values that are greater than or equal to the value specified with GE.
	GE *string `type:"string"`

	// The greater than operator. The DataSource results will have FilterVariable
	// values that are greater than the value specified with GT.
	GT *string `type:"string"`

	// The less than or equal to operator. The DataSource results will have FilterVariable
	// values that are less than or equal to the value specified with LE.
	LE *string `type:"string"`

	// The less than operator. The DataSource results will have FilterVariable values
	// that are less than the value specified with LT.
	LT *string `type:"string"`

	// The maximum number of DataSource to include in the result.
	Limit *int64 `min:"1" type:"integer"`

	// The not equal to operator. The DataSource results will have FilterVariable
	// values not equal to the value specified with NE.
	NE *string `type:"string"`

	// The ID of the page in the paginated results.
	NextToken *string `type:"string"`

	// A string that is found at the beginning of a variable, such as Name or Id.
	//
	// For example, a DataSource could have the Name 2014-09-09-HolidayGiftMailer.
	// To search for this DataSource, select Name for the FilterVariable and any
	// of the following strings for the Prefix:
	//
	//    * 2014-09
	//
	//    * 2014-09-09
	//
	//    * 2014-09-09-Holiday
	Prefix *string `type:"string"`

	// A two-value parameter that determines the sequence of the resulting list
	// of DataSource.
	//
	//    * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	//    * dsc - Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by FilterVariable.
	SortOrder enums.SortOrder `type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeDataSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDataSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDataSourcesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the query results from a DescribeDataSources operation. The content
// is essentially a list of DataSource.
type DescribeDataSourcesOutput struct {
	_ struct{} `type:"structure"`

	// An ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string `type:"string"`

	// A list of DataSource that meet the search criteria.
	Results []DataSource `type:"list"`
}

// String returns the string representation
func (s DescribeDataSourcesOutput) String() string {
	return awsutil.Prettify(s)
}