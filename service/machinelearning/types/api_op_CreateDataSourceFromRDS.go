// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateDataSourceFromRDSInput struct {
	_ struct{} `type:"structure"`

	// The compute statistics for a DataSource. The statistics are generated from
	// the observation data referenced by a DataSource. Amazon ML uses the statistics
	// internally during MLModel training. This parameter must be set to true if
	// the DataSource needs to be used for MLModel training.
	ComputeStatistics *bool `type:"boolean"`

	// A user-supplied ID that uniquely identifies the DataSource. Typically, an
	// Amazon Resource Number (ARN) becomes the ID for a DataSource.
	//
	// DataSourceId is a required field
	DataSourceId *string `min:"1" type:"string" required:"true"`

	// A user-supplied name or description of the DataSource.
	DataSourceName *string `type:"string"`

	// The data specification of an Amazon RDS DataSource:
	//
	//    * DatabaseInformation -
	//    * DatabaseName - The name of the Amazon RDS database.
	//
	//    * InstanceIdentifier - A unique identifier for the Amazon RDS database
	//    instance.
	//
	//    * DatabaseCredentials - AWS Identity and Access Management (IAM) credentials
	//    that are used to connect to the Amazon RDS database.
	//
	//    * ResourceRole - A role (DataPipelineDefaultResourceRole) assumed by an
	//    EC2 instance to carry out the copy task from Amazon RDS to Amazon Simple
	//    Storage Service (Amazon S3). For more information, see Role templates
	//    (http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-iam-roles.html)
	//    for data pipelines.
	//
	//    * ServiceRole - A role (DataPipelineDefaultRole) assumed by the AWS Data
	//    Pipeline service to monitor the progress of the copy task from Amazon
	//    RDS to Amazon S3. For more information, see Role templates (http://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-iam-roles.html)
	//    for data pipelines.
	//
	//    * SecurityInfo - The security information to use to access an RDS DB instance.
	//    You need to set up appropriate ingress rules for the security entity IDs
	//    provided to allow access to the Amazon RDS instance. Specify a [SubnetId,
	//    SecurityGroupIds] pair for a VPC-based RDS DB instance.
	//
	//    * SelectSqlQuery - A query that is used to retrieve the observation data
	//    for the Datasource.
	//
	//    * S3StagingLocation - The Amazon S3 location for staging Amazon RDS data.
	//    The data retrieved from Amazon RDS using SelectSqlQuery is stored in this
	//    location.
	//
	//    * DataSchemaUri - The Amazon S3 location of the DataSchema.
	//
	//    * DataSchema - A JSON string representing the schema. This is not required
	//    if DataSchemaUri is specified.
	//
	//    * DataRearrangement - A JSON string that represents the splitting and
	//    rearrangement requirements for the Datasource. Sample - "{\"splitting\":{\"percentBegin\":10,\"percentEnd\":60}}"
	//
	// RDSData is a required field
	RDSData *RDSDataSpec `type:"structure" required:"true"`

	// The role that Amazon ML assumes on behalf of the user to create and activate
	// a data pipeline in the user's account and copy data using the SelectSqlQuery
	// query from Amazon RDS to Amazon S3.
	//
	// RoleARN is a required field
	RoleARN *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDataSourceFromRDSInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSourceFromRDSInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSourceFromRDSInput"}

	if s.DataSourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSourceId"))
	}
	if s.DataSourceId != nil && len(*s.DataSourceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DataSourceId", 1))
	}

	if s.RDSData == nil {
		invalidParams.Add(aws.NewErrParamRequired("RDSData"))
	}

	if s.RoleARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoleARN"))
	}
	if s.RoleARN != nil && len(*s.RoleARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleARN", 1))
	}
	if s.RDSData != nil {
		if err := s.RDSData.Validate(); err != nil {
			invalidParams.AddNested("RDSData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a CreateDataSourceFromRDS operation, and is an acknowledgement
// that Amazon ML received the request.
//
// The CreateDataSourceFromRDS> operation is asynchronous. You can poll for
// updates by using the GetBatchPrediction operation and checking the Status
// parameter. You can inspect the Message when Status shows up as FAILED. You
// can also check the progress of the copy operation by going to the DataPipeline
// console and looking up the pipeline using the pipelineId from the describe
// call.
type CreateDataSourceFromRDSOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the datasource. This value should
	// be identical to the value of the DataSourceID in the request.
	DataSourceId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDataSourceFromRDSOutput) String() string {
	return awsutil.Prettify(s)
}