// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_query service, marshal.go contains the generated operation marshalers and methods on it.

package aws_query

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
)

var _ bytes.Buffer
var _ awserr.Error

// AddListenerCertificatesMarshaler defines marshaler for AddListenerCertificates operation
type AddListenerCertificatesMarshaler struct {
	Input *types.AddListenerCertificatesInput
}

// AddTagsMarshaler defines marshaler for AddTags operation
type AddTagsMarshaler struct {
	Input *types.AddTagsInput
}

// CreateListenerMarshaler defines marshaler for CreateListener operation
type CreateListenerMarshaler struct {
	Input *types.CreateListenerInput
}

// CreateLoadBalancerMarshaler defines marshaler for CreateLoadBalancer operation
type CreateLoadBalancerMarshaler struct {
	Input *types.CreateLoadBalancerInput
}

// CreateRuleMarshaler defines marshaler for CreateRule operation
type CreateRuleMarshaler struct {
	Input *types.CreateRuleInput
}

// CreateTargetGroupMarshaler defines marshaler for CreateTargetGroup operation
type CreateTargetGroupMarshaler struct {
	Input *types.CreateTargetGroupInput
}

// DeleteListenerMarshaler defines marshaler for DeleteListener operation
type DeleteListenerMarshaler struct {
	Input *types.DeleteListenerInput
}

// DeleteLoadBalancerMarshaler defines marshaler for DeleteLoadBalancer operation
type DeleteLoadBalancerMarshaler struct {
	Input *types.DeleteLoadBalancerInput
}

// DeleteRuleMarshaler defines marshaler for DeleteRule operation
type DeleteRuleMarshaler struct {
	Input *types.DeleteRuleInput
}

// DeleteTargetGroupMarshaler defines marshaler for DeleteTargetGroup operation
type DeleteTargetGroupMarshaler struct {
	Input *types.DeleteTargetGroupInput
}

// DeregisterTargetsMarshaler defines marshaler for DeregisterTargets operation
type DeregisterTargetsMarshaler struct {
	Input *types.DeregisterTargetsInput
}

// DescribeAccountLimitsMarshaler defines marshaler for DescribeAccountLimits operation
type DescribeAccountLimitsMarshaler struct {
	Input *types.DescribeAccountLimitsInput
}

// DescribeListenerCertificatesMarshaler defines marshaler for DescribeListenerCertificates operation
type DescribeListenerCertificatesMarshaler struct {
	Input *types.DescribeListenerCertificatesInput
}

// DescribeListenersMarshaler defines marshaler for DescribeListeners operation
type DescribeListenersMarshaler struct {
	Input *types.DescribeListenersInput
}

// DescribeLoadBalancerAttributesMarshaler defines marshaler for DescribeLoadBalancerAttributes operation
type DescribeLoadBalancerAttributesMarshaler struct {
	Input *types.DescribeLoadBalancerAttributesInput
}

// DescribeLoadBalancersMarshaler defines marshaler for DescribeLoadBalancers operation
type DescribeLoadBalancersMarshaler struct {
	Input *types.DescribeLoadBalancersInput
}

// DescribeRulesMarshaler defines marshaler for DescribeRules operation
type DescribeRulesMarshaler struct {
	Input *types.DescribeRulesInput
}

// DescribeSSLPoliciesMarshaler defines marshaler for DescribeSSLPolicies operation
type DescribeSSLPoliciesMarshaler struct {
	Input *types.DescribeSSLPoliciesInput
}

// DescribeTagsMarshaler defines marshaler for DescribeTags operation
type DescribeTagsMarshaler struct {
	Input *types.DescribeTagsInput
}

// DescribeTargetGroupAttributesMarshaler defines marshaler for DescribeTargetGroupAttributes operation
type DescribeTargetGroupAttributesMarshaler struct {
	Input *types.DescribeTargetGroupAttributesInput
}

// DescribeTargetGroupsMarshaler defines marshaler for DescribeTargetGroups operation
type DescribeTargetGroupsMarshaler struct {
	Input *types.DescribeTargetGroupsInput
}

// DescribeTargetHealthMarshaler defines marshaler for DescribeTargetHealth operation
type DescribeTargetHealthMarshaler struct {
	Input *types.DescribeTargetHealthInput
}

// ModifyListenerMarshaler defines marshaler for ModifyListener operation
type ModifyListenerMarshaler struct {
	Input *types.ModifyListenerInput
}

// ModifyLoadBalancerAttributesMarshaler defines marshaler for ModifyLoadBalancerAttributes operation
type ModifyLoadBalancerAttributesMarshaler struct {
	Input *types.ModifyLoadBalancerAttributesInput
}

// ModifyRuleMarshaler defines marshaler for ModifyRule operation
type ModifyRuleMarshaler struct {
	Input *types.ModifyRuleInput
}

// ModifyTargetGroupMarshaler defines marshaler for ModifyTargetGroup operation
type ModifyTargetGroupMarshaler struct {
	Input *types.ModifyTargetGroupInput
}

// ModifyTargetGroupAttributesMarshaler defines marshaler for ModifyTargetGroupAttributes operation
type ModifyTargetGroupAttributesMarshaler struct {
	Input *types.ModifyTargetGroupAttributesInput
}

// RegisterTargetsMarshaler defines marshaler for RegisterTargets operation
type RegisterTargetsMarshaler struct {
	Input *types.RegisterTargetsInput
}

// RemoveListenerCertificatesMarshaler defines marshaler for RemoveListenerCertificates operation
type RemoveListenerCertificatesMarshaler struct {
	Input *types.RemoveListenerCertificatesInput
}

// RemoveTagsMarshaler defines marshaler for RemoveTags operation
type RemoveTagsMarshaler struct {
	Input *types.RemoveTagsInput
}

// SetIpAddressTypeMarshaler defines marshaler for SetIpAddressType operation
type SetIpAddressTypeMarshaler struct {
	Input *types.SetIpAddressTypeInput
}

// SetRulePrioritiesMarshaler defines marshaler for SetRulePriorities operation
type SetRulePrioritiesMarshaler struct {
	Input *types.SetRulePrioritiesInput
}

// SetSecurityGroupsMarshaler defines marshaler for SetSecurityGroups operation
type SetSecurityGroupsMarshaler struct {
	Input *types.SetSecurityGroupsInput
}

// SetSubnetsMarshaler defines marshaler for SetSubnets operation
type SetSubnetsMarshaler struct {
	Input *types.SetSubnetsInput
}

func MarshalAddListenerCertificatesInputShapeAWSQuery(v *types.AddListenerCertificatesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalAddTagsInputShapeAWSQuery(v *types.AddTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreateListenerInputShapeAWSQuery(v *types.CreateListenerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreateLoadBalancerInputShapeAWSQuery(v *types.CreateLoadBalancerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreateRuleInputShapeAWSQuery(v *types.CreateRuleInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreateTargetGroupInputShapeAWSQuery(v *types.CreateTargetGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeleteListenerInputShapeAWSQuery(v *types.DeleteListenerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeleteLoadBalancerInputShapeAWSQuery(v *types.DeleteLoadBalancerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeleteRuleInputShapeAWSQuery(v *types.DeleteRuleInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeleteTargetGroupInputShapeAWSQuery(v *types.DeleteTargetGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeregisterTargetsInputShapeAWSQuery(v *types.DeregisterTargetsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeAccountLimitsInputShapeAWSQuery(v *types.DescribeAccountLimitsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeListenerCertificatesInputShapeAWSQuery(v *types.DescribeListenerCertificatesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeListenersInputShapeAWSQuery(v *types.DescribeListenersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeLoadBalancerAttributesInputShapeAWSQuery(v *types.DescribeLoadBalancerAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeLoadBalancersInputShapeAWSQuery(v *types.DescribeLoadBalancersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeRulesInputShapeAWSQuery(v *types.DescribeRulesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeSSLPoliciesInputShapeAWSQuery(v *types.DescribeSSLPoliciesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeTagsInputShapeAWSQuery(v *types.DescribeTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeTargetGroupAttributesInputShapeAWSQuery(v *types.DescribeTargetGroupAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeTargetGroupsInputShapeAWSQuery(v *types.DescribeTargetGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDescribeTargetHealthInputShapeAWSQuery(v *types.DescribeTargetHealthInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalModifyListenerInputShapeAWSQuery(v *types.ModifyListenerInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalModifyLoadBalancerAttributesInputShapeAWSQuery(v *types.ModifyLoadBalancerAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalModifyRuleInputShapeAWSQuery(v *types.ModifyRuleInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalModifyTargetGroupInputShapeAWSQuery(v *types.ModifyTargetGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalModifyTargetGroupAttributesInputShapeAWSQuery(v *types.ModifyTargetGroupAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalRegisterTargetsInputShapeAWSQuery(v *types.RegisterTargetsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalRemoveListenerCertificatesInputShapeAWSQuery(v *types.RemoveListenerCertificatesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalRemoveTagsInputShapeAWSQuery(v *types.RemoveTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetIpAddressTypeInputShapeAWSQuery(v *types.SetIpAddressTypeInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetRulePrioritiesInputShapeAWSQuery(v *types.SetRulePrioritiesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetSecurityGroupsInputShapeAWSQuery(v *types.SetSecurityGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetSubnetsInputShapeAWSQuery(v *types.SetSubnetsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}
func (m AddListenerCertificatesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAddListenerCertificatesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AddListenerCertificatesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AddListenerCertificates.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m AddTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAddTagsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AddTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AddTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateListenerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateListenerInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateListenerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateListener.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateLoadBalancerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateLoadBalancerInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateLoadBalancerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateLoadBalancer.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateRuleMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateRuleInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateRuleMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateRule.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateTargetGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateTargetGroupInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateTargetGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateTargetGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteListenerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteListenerInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteListenerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteListener.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteLoadBalancerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteLoadBalancerInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteLoadBalancerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteLoadBalancer.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteRuleMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteRuleInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteRuleMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteRule.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteTargetGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteTargetGroupInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteTargetGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteTargetGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeregisterTargetsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeregisterTargetsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeregisterTargetsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeregisterTargets.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeAccountLimitsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeAccountLimitsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeAccountLimitsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeAccountLimits.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeListenerCertificatesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeListenerCertificatesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeListenerCertificatesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeListenerCertificates.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeListenersMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeListenersInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeListenersMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeListeners.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeLoadBalancerAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeLoadBalancerAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeLoadBalancerAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeLoadBalancerAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeLoadBalancersMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeLoadBalancersInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeLoadBalancersMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeLoadBalancers.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeRulesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeRulesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeRulesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeRules.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeSSLPoliciesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeSSLPoliciesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeSSLPoliciesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeSSLPolicies.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeTagsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeTargetGroupAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeTargetGroupAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeTargetGroupAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeTargetGroupAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeTargetGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeTargetGroupsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeTargetGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeTargetGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeTargetHealthMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeTargetHealthInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeTargetHealthMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeTargetHealth.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyListenerMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyListenerInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyListenerMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyListener.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyLoadBalancerAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyLoadBalancerAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyLoadBalancerAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyLoadBalancerAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyRuleMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyRuleInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyRuleMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyRule.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyTargetGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyTargetGroupInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyTargetGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyTargetGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyTargetGroupAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyTargetGroupAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyTargetGroupAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyTargetGroupAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RegisterTargetsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRegisterTargetsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RegisterTargetsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RegisterTargets.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RemoveListenerCertificatesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRemoveListenerCertificatesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RemoveListenerCertificatesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RemoveListenerCertificates.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RemoveTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRemoveTagsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RemoveTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RemoveTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetIpAddressTypeMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetIpAddressTypeInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetIpAddressTypeMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetIpAddressType.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetRulePrioritiesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetRulePrioritiesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetRulePrioritiesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetRulePriorities.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetSecurityGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetSecurityGroupsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetSecurityGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetSecurityGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetSubnetsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetSubnetsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetSubnetsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetSubnets.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
