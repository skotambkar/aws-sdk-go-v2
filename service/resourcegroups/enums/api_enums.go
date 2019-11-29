// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type GroupFilterName string

// Enum values for GroupFilterName
const (
	GroupFilterNameResourceType GroupFilterName = "resource-type"
)

func (enum GroupFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GroupFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QueryErrorCode string

// Enum values for QueryErrorCode
const (
	QueryErrorCodeCloudformationStackInactive    QueryErrorCode = "CLOUDFORMATION_STACK_INACTIVE"
	QueryErrorCodeCloudformationStackNotExisting QueryErrorCode = "CLOUDFORMATION_STACK_NOT_EXISTING"
)

func (enum QueryErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QueryErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type QueryType string

// Enum values for QueryType
const (
	QueryTypeTagFilters10          QueryType = "TAG_FILTERS_1_0"
	QueryTypeCloudformationStack10 QueryType = "CLOUDFORMATION_STACK_1_0"
)

func (enum QueryType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum QueryType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceFilterName string

// Enum values for ResourceFilterName
const (
	ResourceFilterNameResourceType ResourceFilterName = "resource-type"
)

func (enum ResourceFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}