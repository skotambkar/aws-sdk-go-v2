// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/athena/enums"
)

// Information about the columns in a query execution result.
type ColumnInfo struct {
	_ struct{} `type:"structure"`

	// Indicates whether values in the column are case-sensitive.
	CaseSensitive *bool `type:"boolean"`

	// The catalog to which the query results belong.
	CatalogName *string `type:"string"`

	// A column label.
	Label *string `type:"string"`

	// The name of the column.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Indicates the column's nullable status.
	Nullable enums.ColumnNullable `type:"string" enum:"true"`

	// For DECIMAL data types, specifies the total number of digits, up to 38. For
	// performance reasons, we recommend up to 18 digits.
	Precision *int64 `type:"integer"`

	// For DECIMAL data types, specifies the total number of digits in the fractional
	// part of the value. Defaults to 0.
	Scale *int64 `type:"integer"`

	// The schema name (database name) to which the query results belong.
	SchemaName *string `type:"string"`

	// The table name for the query results.
	TableName *string `type:"string"`

	// The data type of the column.
	//
	// Type is a required field
	Type *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ColumnInfo) String() string {
	return awsutil.Prettify(s)
}

// A piece of data (a field in the table).
type Datum struct {
	_ struct{} `type:"structure"`

	// The value of the datum.
	VarCharValue *string `type:"string"`
}

// String returns the string representation
func (s Datum) String() string {
	return awsutil.Prettify(s)
}

// If query results are encrypted in Amazon S3, indicates the encryption option
// used (for example, SSE-KMS or CSE-KMS) and key information.
type EncryptionConfiguration struct {
	_ struct{} `type:"structure"`

	// Indicates whether Amazon S3 server-side encryption with Amazon S3-managed
	// keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or
	// client-side encryption with KMS-managed keys (CSE-KMS) is used.
	//
	// If a query runs in a workgroup and the workgroup overrides client-side settings,
	// then the workgroup's setting for encryption is used. It specifies whether
	// query results must be encrypted, for all queries that run in this workgroup.
	//
	// EncryptionOption is a required field
	EncryptionOption enums.EncryptionOption `type:"string" required:"true" enum:"true"`

	// For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID.
	KmsKey *string `type:"string"`
}

// String returns the string representation
func (s EncryptionConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EncryptionConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EncryptionConfiguration"}
	if len(s.EncryptionOption) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("EncryptionOption"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A query, where QueryString is the list of SQL query statements that comprise
// the query.
type NamedQuery struct {
	_ struct{} `type:"structure"`

	// The database to which the query belongs.
	//
	// Database is a required field
	Database *string `min:"1" type:"string" required:"true"`

	// The query description.
	Description *string `min:"1" type:"string"`

	// The query name.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The unique identifier of the query.
	NamedQueryId *string `type:"string"`

	// The SQL query statements that comprise the query.
	//
	// QueryString is a required field
	QueryString *string `min:"1" type:"string" required:"true"`

	// The name of the workgroup that contains the named query.
	WorkGroup *string `type:"string"`
}

// String returns the string representation
func (s NamedQuery) String() string {
	return awsutil.Prettify(s)
}

// Information about a single instance of a query execution.
type QueryExecution struct {
	_ struct{} `type:"structure"`

	// The SQL query statements which the query execution ran.
	Query *string `min:"1" type:"string"`

	// The database in which the query execution occurred.
	QueryExecutionContext *QueryExecutionContext `type:"structure"`

	// The unique identifier for each query execution.
	QueryExecutionId *string `type:"string"`

	// The location in Amazon S3 where query results were stored and the encryption
	// option, if any, used for query results. These are known as "client-side settings".
	// If workgroup settings override client-side settings, then the query uses
	// the location for the query results and the encryption configuration that
	// are specified for the workgroup.
	ResultConfiguration *ResultConfiguration `type:"structure"`

	// The type of query statement that was run. DDL indicates DDL query statements.
	// DML indicates DML (Data Manipulation Language) query statements, such as
	// CREATE TABLE AS SELECT. UTILITY indicates query statements other than DDL
	// and DML, such as SHOW CREATE TABLE, or DESCRIBE <table>.
	StatementType enums.StatementType `type:"string" enum:"true"`

	// The location of a manifest file that tracks file locations generated by the
	// query, the amount of data scanned by the query, and the amount of time that
	// it took the query to run.
	Statistics *QueryExecutionStatistics `type:"structure"`

	// The completion date, current state, submission time, and state change reason
	// (if applicable) for the query execution.
	Status *QueryExecutionStatus `type:"structure"`

	// The name of the workgroup in which the query ran.
	WorkGroup *string `type:"string"`
}

// String returns the string representation
func (s QueryExecution) String() string {
	return awsutil.Prettify(s)
}

// The database in which the query execution occurs.
type QueryExecutionContext struct {
	_ struct{} `type:"structure"`

	// The name of the database.
	Database *string `min:"1" type:"string"`
}

// String returns the string representation
func (s QueryExecutionContext) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryExecutionContext) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "QueryExecutionContext"}
	if s.Database != nil && len(*s.Database) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Database", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The location of a manifest file that tracks file locations generated by the
// query, the amount of data scanned by the query, and the amount of time that
// it took the query to run.
type QueryExecutionStatistics struct {
	_ struct{} `type:"structure"`

	// The location and file name of a data manifest file. The manifest file is
	// saved to the Athena query results location in Amazon S3. It tracks files
	// that the query wrote to Amazon S3. If the query fails, the manifest file
	// also tracks files that the query intended to write. The manifest is useful
	// for identifying orphaned files resulting from a failed query. For more information,
	// see Working with Query Output Files (https://docs.aws.amazon.com/athena/latest/ug/querying.html)
	// in the Amazon Athena User Guide.
	DataManifestLocation *string `type:"string"`

	// The number of bytes in the data that was queried.
	DataScannedInBytes *int64 `type:"long"`

	// The number of milliseconds that the query took to execute.
	EngineExecutionTimeInMillis *int64 `type:"long"`
}

// String returns the string representation
func (s QueryExecutionStatistics) String() string {
	return awsutil.Prettify(s)
}

// The completion date, current state, submission time, and state change reason
// (if applicable) for the query execution.
type QueryExecutionStatus struct {
	_ struct{} `type:"structure"`

	// The date and time that the query completed.
	CompletionDateTime *time.Time `type:"timestamp"`

	// The state of query execution. QUEUED state is listed but is not used by Athena
	// and is reserved for future use. RUNNING indicates that the query has been
	// submitted to the service, and Athena will execute the query as soon as resources
	// are available. SUCCEEDED indicates that the query completed without errors.
	// FAILED indicates that the query experienced an error and did not complete
	// processing. CANCELLED indicates that a user input interrupted query execution.
	State enums.QueryExecutionState `type:"string" enum:"true"`

	// Further detail about the status of the query.
	StateChangeReason *string `type:"string"`

	// The date and time that the query was submitted.
	SubmissionDateTime *time.Time `type:"timestamp"`
}

// String returns the string representation
func (s QueryExecutionStatus) String() string {
	return awsutil.Prettify(s)
}

// The location in Amazon S3 where query results are stored and the encryption
// option, if any, used for query results. These are known as "client-side settings".
// If workgroup settings override client-side settings, then the query uses
// the workgroup settings.
type ResultConfiguration struct {
	_ struct{} `type:"structure"`

	// If query results are encrypted in Amazon S3, indicates the encryption option
	// used (for example, SSE-KMS or CSE-KMS) and key information. This is a client-side
	// setting. If workgroup settings override client-side settings, then the query
	// uses the encryption configuration that is specified for the workgroup, and
	// also uses the location for storing query results specified in the workgroup.
	// See WorkGroupConfiguration$EnforceWorkGroupConfiguration and Workgroup Settings
	// Override Client-Side Settings (https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html).
	EncryptionConfiguration *EncryptionConfiguration `type:"structure"`

	// The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/.
	// To run the query, you must specify the query results location using one of
	// the ways: either for individual queries using either this setting (client-side),
	// or in the workgroup, using WorkGroupConfiguration. If none of them is set,
	// Athena issues an error that no output location is provided. For more information,
	// see Query Results (https://docs.aws.amazon.com/athena/latest/ug/querying.html).
	// If workgroup settings override client-side settings, then the query uses
	// the settings specified for the workgroup. See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	OutputLocation *string `type:"string"`
}

// String returns the string representation
func (s ResultConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResultConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResultConfiguration"}
	if s.EncryptionConfiguration != nil {
		if err := s.EncryptionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The information about the updates in the query results, such as output location
// and encryption configuration for the query results.
type ResultConfigurationUpdates struct {
	_ struct{} `type:"structure"`

	// The encryption configuration for the query results.
	EncryptionConfiguration *EncryptionConfiguration `type:"structure"`

	// The location in Amazon S3 where your query results are stored, such as s3://path/to/query/bucket/.
	// For more information, see Query Results (https://docs.aws.amazon.com/athena/latest/ug/querying.html)
	// If workgroup settings override client-side settings, then the query uses
	// the location for the query results and the encryption configuration that
	// are specified for the workgroup. The "workgroup settings override" is specified
	// in EnforceWorkGroupConfiguration (true/false) in the WorkGroupConfiguration.
	// See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	OutputLocation *string `type:"string"`

	// If set to "true", indicates that the previously-specified encryption configuration
	// (also known as the client-side setting) for queries in this workgroup should
	// be ignored and set to null. If set to "false" or not set, and a value is
	// present in the EncryptionConfiguration in ResultConfigurationUpdates (the
	// client-side setting), the EncryptionConfiguration in the workgroup's ResultConfiguration
	// will be updated with the new value. For more information, see Workgroup Settings
	// Override Client-Side Settings (https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html).
	RemoveEncryptionConfiguration *bool `type:"boolean"`

	// If set to "true", indicates that the previously-specified query results location
	// (also known as a client-side setting) for queries in this workgroup should
	// be ignored and set to null. If set to "false" or not set, and a value is
	// present in the OutputLocation in ResultConfigurationUpdates (the client-side
	// setting), the OutputLocation in the workgroup's ResultConfiguration will
	// be updated with the new value. For more information, see Workgroup Settings
	// Override Client-Side Settings (https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html).
	RemoveOutputLocation *bool `type:"boolean"`
}

// String returns the string representation
func (s ResultConfigurationUpdates) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResultConfigurationUpdates) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResultConfigurationUpdates"}
	if s.EncryptionConfiguration != nil {
		if err := s.EncryptionConfiguration.Validate(); err != nil {
			invalidParams.AddNested("EncryptionConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The metadata and rows that comprise a query result set. The metadata describes
// the column structure and data types.
type ResultSet struct {
	_ struct{} `type:"structure"`

	// The metadata that describes the column structure and data types of a table
	// of query results.
	ResultSetMetadata *ResultSetMetadata `type:"structure"`

	// The rows in the table.
	Rows []Row `type:"list"`
}

// String returns the string representation
func (s ResultSet) String() string {
	return awsutil.Prettify(s)
}

// The metadata that describes the column structure and data types of a table
// of query results.
type ResultSetMetadata struct {
	_ struct{} `type:"structure"`

	// Information about the columns returned in a query result metadata.
	ColumnInfo []ColumnInfo `type:"list"`
}

// String returns the string representation
func (s ResultSetMetadata) String() string {
	return awsutil.Prettify(s)
}

// The rows that comprise a query result table.
type Row struct {
	_ struct{} `type:"structure"`

	// The data that populates a row in a query result table.
	Data []Datum `type:"list"`
}

// String returns the string representation
func (s Row) String() string {
	return awsutil.Prettify(s)
}

// A tag that you can add to a resource. A tag is a label that you assign to
// an AWS Athena resource (a workgroup). Each tag consists of a key and an optional
// value, both of which you define. Tags enable you to categorize workgroups
// in Athena, for example, by purpose, owner, or environment. Use a consistent
// set of tag keys to make it easier to search and filter workgroups in your
// account. The maximum tag key length is 128 Unicode characters in UTF-8. The
// maximum tag value length is 256 Unicode characters in UTF-8. You can use
// letters and numbers representable in UTF-8, and the following characters:
// + - = . _ : / @. Tag keys and values are case-sensitive. Tag keys must be
// unique per resource.
type Tag struct {
	_ struct{} `type:"structure"`

	// A tag key. The tag key length is from 1 to 128 Unicode characters in UTF-8.
	// You can use letters and numbers representable in UTF-8, and the following
	// characters: + - = . _ : / @. Tag keys are case-sensitive and must be unique
	// per resource.
	Key *string `min:"1" type:"string"`

	// A tag value. The tag value length is from 0 to 256 Unicode characters in
	// UTF-8. You can use letters and numbers representable in UTF-8, and the following
	// characters: + - = . _ : / @. Tag values are case-sensitive.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a named query ID that could not be processed.
type UnprocessedNamedQueryId struct {
	_ struct{} `type:"structure"`

	// The error code returned when the processing request for the named query failed,
	// if applicable.
	ErrorCode *string `min:"1" type:"string"`

	// The error message returned when the processing request for the named query
	// failed, if applicable.
	ErrorMessage *string `type:"string"`

	// The unique identifier of the named query.
	NamedQueryId *string `type:"string"`
}

// String returns the string representation
func (s UnprocessedNamedQueryId) String() string {
	return awsutil.Prettify(s)
}

// Describes a query execution that failed to process.
type UnprocessedQueryExecutionId struct {
	_ struct{} `type:"structure"`

	// The error code returned when the query execution failed to process, if applicable.
	ErrorCode *string `min:"1" type:"string"`

	// The error message returned when the query execution failed to process, if
	// applicable.
	ErrorMessage *string `type:"string"`

	// The unique identifier of the query execution.
	QueryExecutionId *string `type:"string"`
}

// String returns the string representation
func (s UnprocessedQueryExecutionId) String() string {
	return awsutil.Prettify(s)
}

// A workgroup, which contains a name, description, creation time, state, and
// other configuration, listed under WorkGroup$Configuration. Each workgroup
// enables you to isolate queries for you or your group of users from other
// queries in the same account, to configure the query results location and
// the encryption configuration (known as workgroup settings), to enable sending
// query metrics to Amazon CloudWatch, and to establish per-query data usage
// control limits for all queries in a workgroup. The workgroup settings override
// is specified in EnforceWorkGroupConfiguration (true/false) in the WorkGroupConfiguration.
// See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
type WorkGroup struct {
	_ struct{} `type:"structure"`

	// The configuration of the workgroup, which includes the location in Amazon
	// S3 where query results are stored, the encryption configuration, if any,
	// used for query results; whether the Amazon CloudWatch Metrics are enabled
	// for the workgroup; whether workgroup settings override client-side settings;
	// and the data usage limits for the amount of data scanned per query or per
	// workgroup. The workgroup settings override is specified in EnforceWorkGroupConfiguration
	// (true/false) in the WorkGroupConfiguration. See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
	Configuration *WorkGroupConfiguration `type:"structure"`

	// The date and time the workgroup was created.
	CreationTime *time.Time `type:"timestamp"`

	// The workgroup description.
	Description *string `type:"string"`

	// The workgroup name.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The state of the workgroup: ENABLED or DISABLED.
	State enums.WorkGroupState `type:"string" enum:"true"`
}

// String returns the string representation
func (s WorkGroup) String() string {
	return awsutil.Prettify(s)
}

// The configuration of the workgroup, which includes the location in Amazon
// S3 where query results are stored, the encryption option, if any, used for
// query results, whether the Amazon CloudWatch Metrics are enabled for the
// workgroup and whether workgroup settings override query settings, and the
// data usage limits for the amount of data scanned per query or per workgroup.
// The workgroup settings override is specified in EnforceWorkGroupConfiguration
// (true/false) in the WorkGroupConfiguration. See WorkGroupConfiguration$EnforceWorkGroupConfiguration.
type WorkGroupConfiguration struct {
	_ struct{} `type:"structure"`

	// The upper data usage limit (cutoff) for the amount of bytes a single query
	// in a workgroup is allowed to scan.
	BytesScannedCutoffPerQuery *int64 `min:"1e+07" type:"long"`

	// If set to "true", the settings for the workgroup override client-side settings.
	// If set to "false", client-side settings are used. For more information, see
	// Workgroup Settings Override Client-Side Settings (https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html).
	EnforceWorkGroupConfiguration *bool `type:"boolean"`

	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	PublishCloudWatchMetricsEnabled *bool `type:"boolean"`

	// If set to true, allows members assigned to a workgroup to reference Amazon
	// S3 Requester Pays buckets in queries. If set to false, workgroup members
	// cannot query data from Requester Pays buckets, and queries that retrieve
	// data from Requester Pays buckets cause an error. The default is false. For
	// more information about Requester Pays buckets, see Requester Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// in the Amazon Simple Storage Service Developer Guide.
	RequesterPaysEnabled *bool `type:"boolean"`

	// The configuration for the workgroup, which includes the location in Amazon
	// S3 where query results are stored and the encryption option, if any, used
	// for query results. To run the query, you must specify the query results location
	// using one of the ways: either in the workgroup using this setting, or for
	// individual queries (client-side), using ResultConfiguration$OutputLocation.
	// If none of them is set, Athena issues an error that no output location is
	// provided. For more information, see Query Results (https://docs.aws.amazon.com/athena/latest/ug/querying.html).
	ResultConfiguration *ResultConfiguration `type:"structure"`
}

// String returns the string representation
func (s WorkGroupConfiguration) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *WorkGroupConfiguration) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "WorkGroupConfiguration"}
	if s.BytesScannedCutoffPerQuery != nil && *s.BytesScannedCutoffPerQuery < 1e+07 {
		invalidParams.Add(aws.NewErrParamMinValue("BytesScannedCutoffPerQuery", 1e+07))
	}
	if s.ResultConfiguration != nil {
		if err := s.ResultConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ResultConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The configuration information that will be updated for this workgroup, which
// includes the location in Amazon S3 where query results are stored, the encryption
// option, if any, used for query results, whether the Amazon CloudWatch Metrics
// are enabled for the workgroup, whether the workgroup settings override the
// client-side settings, and the data usage limit for the amount of bytes scanned
// per query, if it is specified.
type WorkGroupConfigurationUpdates struct {
	_ struct{} `type:"structure"`

	// The upper limit (cutoff) for the amount of bytes a single query in a workgroup
	// is allowed to scan.
	BytesScannedCutoffPerQuery *int64 `min:"1e+07" type:"long"`

	// If set to "true", the settings for the workgroup override client-side settings.
	// If set to "false" client-side settings are used. For more information, see
	// Workgroup Settings Override Client-Side Settings (https://docs.aws.amazon.com/athena/latest/ug/workgroups-settings-override.html).
	EnforceWorkGroupConfiguration *bool `type:"boolean"`

	// Indicates whether this workgroup enables publishing metrics to Amazon CloudWatch.
	PublishCloudWatchMetricsEnabled *bool `type:"boolean"`

	// Indicates that the data usage control limit per query is removed. WorkGroupConfiguration$BytesScannedCutoffPerQuery
	RemoveBytesScannedCutoffPerQuery *bool `type:"boolean"`

	// If set to true, allows members assigned to a workgroup to specify Amazon
	// S3 Requester Pays buckets in queries. If set to false, workgroup members
	// cannot query data from Requester Pays buckets, and queries that retrieve
	// data from Requester Pays buckets cause an error. The default is false. For
	// more information about Requester Pays buckets, see Requester Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html)
	// in the Amazon Simple Storage Service Developer Guide.
	RequesterPaysEnabled *bool `type:"boolean"`

	// The result configuration information about the queries in this workgroup
	// that will be updated. Includes the updated results location and an updated
	// option for encrypting query results.
	ResultConfigurationUpdates *ResultConfigurationUpdates `type:"structure"`
}

// String returns the string representation
func (s WorkGroupConfigurationUpdates) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *WorkGroupConfigurationUpdates) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "WorkGroupConfigurationUpdates"}
	if s.BytesScannedCutoffPerQuery != nil && *s.BytesScannedCutoffPerQuery < 1e+07 {
		invalidParams.Add(aws.NewErrParamMinValue("BytesScannedCutoffPerQuery", 1e+07))
	}
	if s.ResultConfigurationUpdates != nil {
		if err := s.ResultConfigurationUpdates.Validate(); err != nil {
			invalidParams.AddNested("ResultConfigurationUpdates", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The summary information for the workgroup, which includes its name, state,
// description, and the date and time it was created.
type WorkGroupSummary struct {
	_ struct{} `type:"structure"`

	// The workgroup creation date and time.
	CreationTime *time.Time `type:"timestamp"`

	// The workgroup description.
	Description *string `type:"string"`

	// The name of the workgroup.
	Name *string `type:"string"`

	// The state of the workgroup.
	State enums.WorkGroupState `type:"string" enum:"true"`
}

// String returns the string representation
func (s WorkGroupSummary) String() string {
	return awsutil.Prettify(s)
}
