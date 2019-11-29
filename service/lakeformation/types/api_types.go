// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/enums"
)

// A list of failures when performing a batch grant or batch revoke operation.
type BatchPermissionsFailureEntry struct {
	_ struct{} `type:"structure"`

	// An error message that applies to the failure of the entry.
	Error *ErrorDetail `type:"structure"`

	// An identifier for an entry of the batch request.
	RequestEntry *BatchPermissionsRequestEntry `type:"structure"`
}

// String returns the string representation
func (s BatchPermissionsFailureEntry) String() string {
	return awsutil.Prettify(s)
}

// A permission to a resource granted by batch operation to the principal.
type BatchPermissionsRequestEntry struct {
	_ struct{} `type:"structure"`

	// A unique identifier for the batch permissions request entry.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The permissions to be granted.
	Permissions []enums.Permission `type:"list"`

	// Indicates if the option to pass permissions is granted.
	PermissionsWithGrantOption []enums.Permission `type:"list"`

	// The principal to be granted a permission.
	Principal *DataLakePrincipal `type:"structure"`

	// The resource to which the principal is to be granted a permission.
	Resource *Resource `type:"structure"`
}

// String returns the string representation
func (s BatchPermissionsRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchPermissionsRequestEntry) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchPermissionsRequestEntry"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.Principal != nil {
		if err := s.Principal.Validate(); err != nil {
			invalidParams.AddNested("Principal", err.(aws.ErrInvalidParams))
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			invalidParams.AddNested("Resource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure for the catalog object.
type CatalogResource struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CatalogResource) String() string {
	return awsutil.Prettify(s)
}

// A wildcard object, consisting of an optional list of excluded column names
// or indexes.
type ColumnWildcard struct {
	_ struct{} `type:"structure"`

	// Excludes column names. Any column with this name will be excluded.
	ExcludedColumnNames []string `type:"list"`
}

// String returns the string representation
func (s ColumnWildcard) String() string {
	return awsutil.Prettify(s)
}

// The AWS Lake Formation principal.
type DataLakePrincipal struct {
	_ struct{} `type:"structure"`

	// An identifier for the AWS Lake Formation principal.
	DataLakePrincipalIdentifier *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DataLakePrincipal) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DataLakePrincipal) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DataLakePrincipal"}
	if s.DataLakePrincipalIdentifier != nil && len(*s.DataLakePrincipalIdentifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataLakePrincipalIdentifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The AWS Lake Formation principal.
type DataLakeSettings struct {
	_ struct{} `type:"structure"`

	// A list of up to three principal permissions entries for default create database
	// permissions.
	CreateDatabaseDefaultPermissions []PrincipalPermissions `type:"list"`

	// A list of up to three principal permissions entries for default create table
	// permissions.
	CreateTableDefaultPermissions []PrincipalPermissions `type:"list"`

	// A list of AWS Lake Formation principals.
	DataLakeAdmins []DataLakePrincipal `type:"list"`
}

// String returns the string representation
func (s DataLakeSettings) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DataLakeSettings) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DataLakeSettings"}
	if s.CreateDatabaseDefaultPermissions != nil {
		for i, v := range s.CreateDatabaseDefaultPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CreateDatabaseDefaultPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.CreateTableDefaultPermissions != nil {
		for i, v := range s.CreateTableDefaultPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "CreateTableDefaultPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.DataLakeAdmins != nil {
		for i, v := range s.DataLakeAdmins {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DataLakeAdmins", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure for a data location object where permissions are granted or revoked.
type DataLocationResource struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that uniquely identifies the data location
	// resource.
	//
	// ResourceArn is a required field
	ResourceArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DataLocationResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DataLocationResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DataLocationResource"}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure for the database object.
type DatabaseResource struct {
	_ struct{} `type:"structure"`

	// The name of the database resource. Unique to the Data Catalog.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DatabaseResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DatabaseResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DatabaseResource"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains details about an error.
type ErrorDetail struct {
	_ struct{} `type:"structure"`

	// The code associated with this error.
	ErrorCode *string `min:"1" type:"string"`

	// A message describing the error.
	ErrorMessage *string `type:"string"`
}

// String returns the string representation
func (s ErrorDetail) String() string {
	return awsutil.Prettify(s)
}

// This structure describes the filtering of columns in a table based on a filter
// condition.
type FilterCondition struct {
	_ struct{} `type:"structure"`

	// The comparison operator used in the filter condition.
	ComparisonOperator enums.ComparisonOperator `type:"string" enum:"true"`

	// The field to filter in the filter condition.
	Field enums.FieldNameString `type:"string" enum:"true"`

	// A string with values used in evaluating the filter condition.
	StringValueList []string `type:"list"`
}

// String returns the string representation
func (s FilterCondition) String() string {
	return awsutil.Prettify(s)
}

// Permissions granted to a principal.
type PrincipalPermissions struct {
	_ struct{} `type:"structure"`

	// The permissions that are granted to the principal.
	Permissions []enums.Permission `type:"list"`

	// The principal who is granted permissions.
	Principal *DataLakePrincipal `type:"structure"`
}

// String returns the string representation
func (s PrincipalPermissions) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PrincipalPermissions) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PrincipalPermissions"}
	if s.Principal != nil {
		if err := s.Principal.Validate(); err != nil {
			invalidParams.AddNested("Principal", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The permissions granted or revoked on a resource.
type PrincipalResourcePermissions struct {
	_ struct{} `type:"structure"`

	// The permissions to be granted or revoked on the resource.
	Permissions []enums.Permission `type:"list"`

	// Indicates whether to grant the ability to grant permissions (as a subset
	// of permissions granted).
	PermissionsWithGrantOption []enums.Permission `type:"list"`

	// The Data Lake principal to be granted or revoked permissions.
	Principal *DataLakePrincipal `type:"structure"`

	// The resource where permissions are to be granted or revoked.
	Resource *Resource `type:"structure"`
}

// String returns the string representation
func (s PrincipalResourcePermissions) String() string {
	return awsutil.Prettify(s)
}

// A structure for the resource.
type Resource struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	Catalog *CatalogResource `type:"structure"`

	// The location of an Amazon S3 path where permissions are granted or revoked.
	DataLocation *DataLocationResource `type:"structure"`

	// The database for the resource. Unique to the Data Catalog. A database is
	// a set of associated table definitions organized into a logical group. You
	// can Grant and Revoke database permissions to a principal.
	Database *DatabaseResource `type:"structure"`

	// The table for the resource. A table is a metadata definition that represents
	// your data. You can Grant and Revoke table privileges to a principal.
	Table *TableResource `type:"structure"`

	// The table with columns for the resource. A principal with permissions to
	// this resource can select metadata from the columns of a table in the Data
	// Catalog and the underlying data in Amazon S3.
	TableWithColumns *TableWithColumnsResource `type:"structure"`
}

// String returns the string representation
func (s Resource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Resource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Resource"}
	if s.DataLocation != nil {
		if err := s.DataLocation.Validate(); err != nil {
			invalidParams.AddNested("DataLocation", err.(aws.ErrInvalidParams))
		}
	}
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			invalidParams.AddNested("Database", err.(aws.ErrInvalidParams))
		}
	}
	if s.Table != nil {
		if err := s.Table.Validate(); err != nil {
			invalidParams.AddNested("Table", err.(aws.ErrInvalidParams))
		}
	}
	if s.TableWithColumns != nil {
		if err := s.TableWithColumns.Validate(); err != nil {
			invalidParams.AddNested("TableWithColumns", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure containing information about an AWS Lake Formation resource.
type ResourceInfo struct {
	_ struct{} `type:"structure"`

	// The date and time the resource was last modified.
	LastModified *time.Time `type:"timestamp"`

	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string `type:"string"`

	// The IAM role that registered a resource.
	RoleArn *string `type:"string"`
}

// String returns the string representation
func (s ResourceInfo) String() string {
	return awsutil.Prettify(s)
}

// A structure for the table object. A table is a metadata definition that represents
// your data. You can Grant and Revoke table privileges to a principal.
type TableResource struct {
	_ struct{} `type:"structure"`

	// The name of the database for the table. Unique to a Data Catalog. A database
	// is a set of associated table definitions organized into a logical group.
	// You can Grant and Revoke database privileges to a principal.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// The name of the table.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TableResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TableResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TableResource"}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure for a table with columns object. This object is only used when
// granting a SELECT permission.
//
// This object must take a value for at least one of ColumnsNames, ColumnsIndexes,
// or ColumnsWildcard.
type TableWithColumnsResource struct {
	_ struct{} `type:"structure"`

	// The list of column names for the table. At least one of ColumnNames or ColumnWildcard
	// is required.
	ColumnNames []string `type:"list"`

	// A wildcard specified by a ColumnWildcard object. At least one of ColumnNames
	// or ColumnWildcard is required.
	ColumnWildcard *ColumnWildcard `type:"structure"`

	// The name of the database for the table with columns resource. Unique to the
	// Data Catalog. A database is a set of associated table definitions organized
	// into a logical group. You can Grant and Revoke database privileges to a principal.
	DatabaseName *string `min:"1" type:"string"`

	// The name of the table resource. A table is a metadata definition that represents
	// your data. You can Grant and Revoke table privileges to a principal.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s TableWithColumnsResource) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TableWithColumnsResource) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TableWithColumnsResource"}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
