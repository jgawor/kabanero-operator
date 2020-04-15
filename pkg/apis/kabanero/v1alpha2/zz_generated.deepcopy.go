// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerWebhookCustomizationSpec) DeepCopyInto(out *AdmissionControllerWebhookCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerWebhookCustomizationSpec.
func (in *AdmissionControllerWebhookCustomizationSpec) DeepCopy() *AdmissionControllerWebhookCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerWebhookCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerWebhookStatus) DeepCopyInto(out *AdmissionControllerWebhookStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerWebhookStatus.
func (in *AdmissionControllerWebhookStatus) DeepCopy() *AdmissionControllerWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerWebhookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppsodyStatus) DeepCopyInto(out *AppsodyStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppsodyStatus.
func (in *AppsodyStatus) DeepCopy() *AppsodyStatus {
	if in == nil {
		return nil
	}
	out := new(AppsodyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRWCustomizationSpec) DeepCopyInto(out *CRWCustomizationSpec) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	in.Operator.DeepCopyInto(&out.Operator)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRWCustomizationSpec.
func (in *CRWCustomizationSpec) DeepCopy() *CRWCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(CRWCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRWInstanceStatus) DeepCopyInto(out *CRWInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRWInstanceStatus.
func (in *CRWInstanceStatus) DeepCopy() *CRWInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CRWInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRWOperatorCRInstanceSpec) DeepCopyInto(out *CRWOperatorCRInstanceSpec) {
	*out = *in
	out.DevFileRegistryImage = in.DevFileRegistryImage
	if in.OpenShiftOAuth != nil {
		in, out := &in.OpenShiftOAuth, &out.OpenShiftOAuth
		*out = new(bool)
		**out = **in
	}
	if in.SelfSignedCert != nil {
		in, out := &in.SelfSignedCert, &out.SelfSignedCert
		*out = new(bool)
		**out = **in
	}
	if in.TLSSupport != nil {
		in, out := &in.TLSSupport, &out.TLSSupport
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRWOperatorCRInstanceSpec.
func (in *CRWOperatorCRInstanceSpec) DeepCopy() *CRWOperatorCRInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CRWOperatorCRInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRWOperatorSpec) DeepCopyInto(out *CRWOperatorSpec) {
	*out = *in
	in.CustomResourceInstance.DeepCopyInto(&out.CustomResourceInstance)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRWOperatorSpec.
func (in *CRWOperatorSpec) DeepCopy() *CRWOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(CRWOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRWOperatorStatus) DeepCopyInto(out *CRWOperatorStatus) {
	*out = *in
	out.Instance = in.Instance
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRWOperatorStatus.
func (in *CRWOperatorStatus) DeepCopy() *CRWOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(CRWOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRWStatus) DeepCopyInto(out *CRWStatus) {
	*out = *in
	out.Operator = in.Operator
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRWStatus.
func (in *CRWStatus) DeepCopy() *CRWStatus {
	if in == nil {
		return nil
	}
	out := new(CRWStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CWRCustomResourceDevFileRegImage) DeepCopyInto(out *CWRCustomResourceDevFileRegImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CWRCustomResourceDevFileRegImage.
func (in *CWRCustomResourceDevFileRegImage) DeepCopy() *CWRCustomResourceDevFileRegImage {
	if in == nil {
		return nil
	}
	out := new(CWRCustomResourceDevFileRegImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CliStatus) DeepCopyInto(out *CliStatus) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CliStatus.
func (in *CliStatus) DeepCopy() *CliStatus {
	if in == nil {
		return nil
	}
	out := new(CliStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionControllerSpec) DeepCopyInto(out *CollectionControllerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionControllerSpec.
func (in *CollectionControllerSpec) DeepCopy() *CollectionControllerSpec {
	if in == nil {
		return nil
	}
	out := new(CollectionControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionControllerStatus) DeepCopyInto(out *CollectionControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionControllerStatus.
func (in *CollectionControllerStatus) DeepCopy() *CollectionControllerStatus {
	if in == nil {
		return nil
	}
	out := new(CollectionControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsCustomizationSpec) DeepCopyInto(out *EventsCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsCustomizationSpec.
func (in *EventsCustomizationSpec) DeepCopy() *EventsCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(EventsCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventsStatus) DeepCopyInto(out *EventsStatus) {
	*out = *in
	if in.Hostnames != nil {
		in, out := &in.Hostnames, &out.Hostnames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventsStatus.
func (in *EventsStatus) DeepCopy() *EventsStatus {
	if in == nil {
		return nil
	}
	out := new(EventsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitReleaseSpec) DeepCopyInto(out *GitReleaseSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitReleaseSpec.
func (in *GitReleaseSpec) DeepCopy() *GitReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(GitReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GithubConfig) DeepCopyInto(out *GithubConfig) {
	*out = *in
	if in.Teams != nil {
		in, out := &in.Teams, &out.Teams
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GithubConfig.
func (in *GithubConfig) DeepCopy() *GithubConfig {
	if in == nil {
		return nil
	}
	out := new(GithubConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitopsSpec) DeepCopyInto(out *GitopsSpec) {
	*out = *in
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitopsSpec.
func (in *GitopsSpec) DeepCopy() *GitopsSpec {
	if in == nil {
		return nil
	}
	out := new(GitopsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitopsStatus) DeepCopyInto(out *GitopsStatus) {
	*out = *in
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitopsStatus.
func (in *GitopsStatus) DeepCopy() *GitopsStatus {
	if in == nil {
		return nil
	}
	out := new(GitopsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GovernancePolicyConfig) DeepCopyInto(out *GovernancePolicyConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GovernancePolicyConfig.
func (in *GovernancePolicyConfig) DeepCopy() *GovernancePolicyConfig {
	if in == nil {
		return nil
	}
	out := new(GovernancePolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpsProtocolFile) DeepCopyInto(out *HttpsProtocolFile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpsProtocolFile.
func (in *HttpsProtocolFile) DeepCopy() *HttpsProtocolFile {
	if in == nil {
		return nil
	}
	out := new(HttpsProtocolFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDigest) DeepCopyInto(out *ImageDigest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDigest.
func (in *ImageDigest) DeepCopy() *ImageDigest {
	if in == nil {
		return nil
	}
	out := new(ImageDigest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	out.Digest = in.Digest
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStackConfig) DeepCopyInto(out *InstanceStackConfig) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]RepositoryConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStackConfig.
func (in *InstanceStackConfig) DeepCopy() *InstanceStackConfig {
	if in == nil {
		return nil
	}
	out := new(InstanceStackConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kabanero) DeepCopyInto(out *Kabanero) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kabanero.
func (in *Kabanero) DeepCopy() *Kabanero {
	if in == nil {
		return nil
	}
	out := new(Kabanero)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kabanero) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroCliServicesCustomizationSpec) DeepCopyInto(out *KabaneroCliServicesCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroCliServicesCustomizationSpec.
func (in *KabaneroCliServicesCustomizationSpec) DeepCopy() *KabaneroCliServicesCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroCliServicesCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroInstanceStatus) DeepCopyInto(out *KabaneroInstanceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroInstanceStatus.
func (in *KabaneroInstanceStatus) DeepCopy() *KabaneroInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroLandingCustomizationSpec) DeepCopyInto(out *KabaneroLandingCustomizationSpec) {
	*out = *in
	if in.Enable != nil {
		in, out := &in.Enable, &out.Enable
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroLandingCustomizationSpec.
func (in *KabaneroLandingCustomizationSpec) DeepCopy() *KabaneroLandingCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroLandingCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroLandingPageStatus) DeepCopyInto(out *KabaneroLandingPageStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroLandingPageStatus.
func (in *KabaneroLandingPageStatus) DeepCopy() *KabaneroLandingPageStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroLandingPageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroList) DeepCopyInto(out *KabaneroList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Kabanero, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroList.
func (in *KabaneroList) DeepCopy() *KabaneroList {
	if in == nil {
		return nil
	}
	out := new(KabaneroList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KabaneroList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroSpec) DeepCopyInto(out *KabaneroSpec) {
	*out = *in
	if in.TargetNamespaces != nil {
		in, out := &in.TargetNamespaces, &out.TargetNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Github.DeepCopyInto(&out.Github)
	out.GovernancePolicy = in.GovernancePolicy
	in.Stacks.DeepCopyInto(&out.Stacks)
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]TriggerSpec, len(*in))
		copy(*out, *in)
	}
	out.CliServices = in.CliServices
	in.Landing.DeepCopyInto(&out.Landing)
	in.CodereadyWorkspaces.DeepCopyInto(&out.CodereadyWorkspaces)
	out.Events = in.Events
	out.CollectionController = in.CollectionController
	out.StackController = in.StackController
	out.AdmissionControllerWebhook = in.AdmissionControllerWebhook
	out.Sso = in.Sso
	in.Gitops.DeepCopyInto(&out.Gitops)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroSpec.
func (in *KabaneroSpec) DeepCopy() *KabaneroSpec {
	if in == nil {
		return nil
	}
	out := new(KabaneroSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KabaneroStatus) DeepCopyInto(out *KabaneroStatus) {
	*out = *in
	out.KabaneroInstance = in.KabaneroInstance
	out.Serverless = in.Serverless
	out.Tekton = in.Tekton
	in.Cli.DeepCopyInto(&out.Cli)
	if in.Landing != nil {
		in, out := &in.Landing, &out.Landing
		*out = new(KabaneroLandingPageStatus)
		**out = **in
	}
	out.Appsody = in.Appsody
	if in.Kappnav != nil {
		in, out := &in.Kappnav, &out.Kappnav
		*out = new(KappnavStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.CodereadyWorkspaces != nil {
		in, out := &in.CodereadyWorkspaces, &out.CodereadyWorkspaces
		*out = new(CRWStatus)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = new(EventsStatus)
		(*in).DeepCopyInto(*out)
	}
	out.CollectionController = in.CollectionController
	out.StackController = in.StackController
	out.AdmissionControllerWebhook = in.AdmissionControllerWebhook
	out.Sso = in.Sso
	in.Gitops.DeepCopyInto(&out.Gitops)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KabaneroStatus.
func (in *KabaneroStatus) DeepCopy() *KabaneroStatus {
	if in == nil {
		return nil
	}
	out := new(KabaneroStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KappnavStatus) DeepCopyInto(out *KappnavStatus) {
	*out = *in
	if in.UiLocations != nil {
		in, out := &in.UiLocations, &out.UiLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ApiLocations != nil {
		in, out := &in.ApiLocations, &out.ApiLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KappnavStatus.
func (in *KappnavStatus) DeepCopy() *KappnavStatus {
	if in == nil {
		return nil
	}
	out := new(KappnavStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnativeServingStatus) DeepCopyInto(out *KnativeServingStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnativeServingStatus.
func (in *KnativeServingStatus) DeepCopy() *KnativeServingStatus {
	if in == nil {
		return nil
	}
	out := new(KnativeServingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	out.Https = in.Https
	out.GitRelease = in.GitRelease
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	out.GitRelease = in.GitRelease
	if in.ActiveAssets != nil {
		in, out := &in.ActiveAssets, &out.ActiveAssets
		*out = make([]RepositoryAssetStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryAssetStatus) DeepCopyInto(out *RepositoryAssetStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryAssetStatus.
func (in *RepositoryAssetStatus) DeepCopy() *RepositoryAssetStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryAssetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryConfig) DeepCopyInto(out *RepositoryConfig) {
	*out = *in
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineSpec, len(*in))
		copy(*out, *in)
	}
	out.Https = in.Https
	out.GitRelease = in.GitRelease
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryConfig.
func (in *RepositoryConfig) DeepCopy() *RepositoryConfig {
	if in == nil {
		return nil
	}
	out := new(RepositoryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerlessStatus) DeepCopyInto(out *ServerlessStatus) {
	*out = *in
	out.KnativeServing = in.KnativeServing
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerlessStatus.
func (in *ServerlessStatus) DeepCopy() *ServerlessStatus {
	if in == nil {
		return nil
	}
	out := new(ServerlessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoCustomizationSpec) DeepCopyInto(out *SsoCustomizationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoCustomizationSpec.
func (in *SsoCustomizationSpec) DeepCopy() *SsoCustomizationSpec {
	if in == nil {
		return nil
	}
	out := new(SsoCustomizationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SsoStatus) DeepCopyInto(out *SsoStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SsoStatus.
func (in *SsoStatus) DeepCopy() *SsoStatus {
	if in == nil {
		return nil
	}
	out := new(SsoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Stack) DeepCopyInto(out *Stack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stack.
func (in *Stack) DeepCopy() *Stack {
	if in == nil {
		return nil
	}
	out := new(Stack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackControllerSpec) DeepCopyInto(out *StackControllerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackControllerSpec.
func (in *StackControllerSpec) DeepCopy() *StackControllerSpec {
	if in == nil {
		return nil
	}
	out := new(StackControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackControllerStatus) DeepCopyInto(out *StackControllerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackControllerStatus.
func (in *StackControllerStatus) DeepCopy() *StackControllerStatus {
	if in == nil {
		return nil
	}
	out := new(StackControllerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackList) DeepCopyInto(out *StackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackList.
func (in *StackList) DeepCopy() *StackList {
	if in == nil {
		return nil
	}
	out := new(StackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackSpec) DeepCopyInto(out *StackSpec) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]StackVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackSpec.
func (in *StackSpec) DeepCopy() *StackSpec {
	if in == nil {
		return nil
	}
	out := new(StackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackStatus) DeepCopyInto(out *StackStatus) {
	*out = *in
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]StackVersionStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackStatus.
func (in *StackStatus) DeepCopy() *StackStatus {
	if in == nil {
		return nil
	}
	out := new(StackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackVersion) DeepCopyInto(out *StackVersion) {
	*out = *in
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineSpec, len(*in))
		copy(*out, *in)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]Image, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackVersion.
func (in *StackVersion) DeepCopy() *StackVersion {
	if in == nil {
		return nil
	}
	out := new(StackVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StackVersionStatus) DeepCopyInto(out *StackVersionStatus) {
	*out = *in
	if in.Pipelines != nil {
		in, out := &in.Pipelines, &out.Pipelines
		*out = make([]PipelineStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StackVersionStatus.
func (in *StackVersionStatus) DeepCopy() *StackVersionStatus {
	if in == nil {
		return nil
	}
	out := new(StackVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TektonStatus) DeepCopyInto(out *TektonStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TektonStatus.
func (in *TektonStatus) DeepCopy() *TektonStatus {
	if in == nil {
		return nil
	}
	out := new(TektonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	out.Https = in.Https
	out.GitRelease = in.GitRelease
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}
