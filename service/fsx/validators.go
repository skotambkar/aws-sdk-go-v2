// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateFileSystemAliases struct {
}

func (*validateOpAssociateFileSystemAliases) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateFileSystemAliases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateFileSystemAliasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateFileSystemAliasesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelDataRepositoryTask struct {
}

func (*validateOpCancelDataRepositoryTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelDataRepositoryTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelDataRepositoryTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelDataRepositoryTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCopyBackup struct {
}

func (*validateOpCopyBackup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCopyBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CopyBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCopyBackupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateBackup struct {
}

func (*validateOpCreateBackup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBackupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataRepositoryTask struct {
}

func (*validateOpCreateDataRepositoryTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataRepositoryTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataRepositoryTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataRepositoryTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateFileSystemFromBackup struct {
}

func (*validateOpCreateFileSystemFromBackup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateFileSystemFromBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateFileSystemFromBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateFileSystemFromBackupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateFileSystem struct {
}

func (*validateOpCreateFileSystem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateFileSystemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBackup struct {
}

func (*validateOpDeleteBackup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBackupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteFileSystem struct {
}

func (*validateOpDeleteFileSystem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteFileSystemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeFileSystemAliases struct {
}

func (*validateOpDescribeFileSystemAliases) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeFileSystemAliases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeFileSystemAliasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeFileSystemAliasesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateFileSystemAliases struct {
}

func (*validateOpDisassociateFileSystemAliases) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateFileSystemAliases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateFileSystemAliasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateFileSystemAliasesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateFileSystem struct {
}

func (*validateOpUpdateFileSystem) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateFileSystemInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateFileSystemAliasesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateFileSystemAliases{}, middleware.After)
}

func addOpCancelDataRepositoryTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelDataRepositoryTask{}, middleware.After)
}

func addOpCopyBackupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCopyBackup{}, middleware.After)
}

func addOpCreateBackupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBackup{}, middleware.After)
}

func addOpCreateDataRepositoryTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataRepositoryTask{}, middleware.After)
}

func addOpCreateFileSystemFromBackupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateFileSystemFromBackup{}, middleware.After)
}

func addOpCreateFileSystemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateFileSystem{}, middleware.After)
}

func addOpDeleteBackupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBackup{}, middleware.After)
}

func addOpDeleteFileSystemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteFileSystem{}, middleware.After)
}

func addOpDescribeFileSystemAliasesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeFileSystemAliases{}, middleware.After)
}

func addOpDisassociateFileSystemAliasesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateFileSystemAliases{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateFileSystemValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateFileSystem{}, middleware.After)
}

func validateCompletionReport(v *types.CompletionReport) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CompletionReport"}
	if v.Enabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Enabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCreateFileSystemWindowsConfiguration(v *types.CreateFileSystemWindowsConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateFileSystemWindowsConfiguration"}
	if v.SelfManagedActiveDirectoryConfiguration != nil {
		if err := validateSelfManagedActiveDirectoryConfiguration(v.SelfManagedActiveDirectoryConfiguration); err != nil {
			invalidParams.AddNested("SelfManagedActiveDirectoryConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.ThroughputCapacity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThroughputCapacity"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeleteFileSystemLustreConfiguration(v *types.DeleteFileSystemLustreConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFileSystemLustreConfiguration"}
	if v.FinalBackupTags != nil {
		if err := validateTags(v.FinalBackupTags); err != nil {
			invalidParams.AddNested("FinalBackupTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeleteFileSystemWindowsConfiguration(v *types.DeleteFileSystemWindowsConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFileSystemWindowsConfiguration"}
	if v.FinalBackupTags != nil {
		if err := validateTags(v.FinalBackupTags); err != nil {
			invalidParams.AddNested("FinalBackupTags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSelfManagedActiveDirectoryConfiguration(v *types.SelfManagedActiveDirectoryConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SelfManagedActiveDirectoryConfiguration"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.UserName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserName"))
	}
	if v.Password == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Password"))
	}
	if v.DnsIps == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DnsIps"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTags(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tags"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateFileSystemAliasesInput(v *AssociateFileSystemAliasesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateFileSystemAliasesInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.Aliases == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Aliases"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelDataRepositoryTaskInput(v *CancelDataRepositoryTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelDataRepositoryTaskInput"}
	if v.TaskId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCopyBackupInput(v *CopyBackupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CopyBackupInput"}
	if v.SourceBackupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceBackupId"))
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBackupInput(v *CreateBackupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBackupInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataRepositoryTaskInput(v *CreateDataRepositoryTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataRepositoryTaskInput"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.Report == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Report"))
	} else if v.Report != nil {
		if err := validateCompletionReport(v.Report); err != nil {
			invalidParams.AddNested("Report", err.(smithy.InvalidParamsError))
		}
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateFileSystemFromBackupInput(v *CreateFileSystemFromBackupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateFileSystemFromBackupInput"}
	if v.BackupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackupId"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.WindowsConfiguration != nil {
		if err := validateCreateFileSystemWindowsConfiguration(v.WindowsConfiguration); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateFileSystemInput(v *CreateFileSystemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateFileSystemInput"}
	if len(v.FileSystemType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemType"))
	}
	if v.StorageCapacity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StorageCapacity"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if v.WindowsConfiguration != nil {
		if err := validateCreateFileSystemWindowsConfiguration(v.WindowsConfiguration); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBackupInput(v *DeleteBackupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBackupInput"}
	if v.BackupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BackupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteFileSystemInput(v *DeleteFileSystemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteFileSystemInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.WindowsConfiguration != nil {
		if err := validateDeleteFileSystemWindowsConfiguration(v.WindowsConfiguration); err != nil {
			invalidParams.AddNested("WindowsConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.LustreConfiguration != nil {
		if err := validateDeleteFileSystemLustreConfiguration(v.LustreConfiguration); err != nil {
			invalidParams.AddNested("LustreConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeFileSystemAliasesInput(v *DescribeFileSystemAliasesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeFileSystemAliasesInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateFileSystemAliasesInput(v *DisassociateFileSystemAliasesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateFileSystemAliasesInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if v.Aliases == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Aliases"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	} else if v.Tags != nil {
		if err := validateTags(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateFileSystemInput(v *UpdateFileSystemInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateFileSystemInput"}
	if v.FileSystemId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FileSystemId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
