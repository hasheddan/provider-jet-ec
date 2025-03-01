/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ElasticsearchKeystoreObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ElasticsearchKeystoreParameters struct {

	// Optionally stores the remote keystore setting as a file. The default is false, which stores the keystore setting as string when value is a plain string
	// +kubebuilder:validation:Optional
	AsFile *bool `json:"asFile,omitempty" tf:"as_file,omitempty"`

	// Required deployment ID of the Deployment that holds the Elasticsearch cluster where the keystore setting will be written to
	// +kubebuilder:validation:Required
	DeploymentID *string `json:"deploymentId" tf:"deployment_id,omitempty"`

	// Required name for the keystore setting, if the setting already exists in the Elasticsearch cluster, it will be overridden
	// +kubebuilder:validation:Required
	SettingName *string `json:"settingName" tf:"setting_name,omitempty"`

	// Required value of this setting. This can either be a string or a JSON object that is stored as a JSON string in the keystore.
	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

// ElasticsearchKeystoreSpec defines the desired state of ElasticsearchKeystore
type ElasticsearchKeystoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ElasticsearchKeystoreParameters `json:"forProvider"`
}

// ElasticsearchKeystoreStatus defines the observed state of ElasticsearchKeystore.
type ElasticsearchKeystoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ElasticsearchKeystoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchKeystore is the Schema for the ElasticsearchKeystores API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ecjet}
type ElasticsearchKeystore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchKeystoreSpec   `json:"spec"`
	Status            ElasticsearchKeystoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchKeystoreList contains a list of ElasticsearchKeystores
type ElasticsearchKeystoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchKeystore `json:"items"`
}

// Repository type metadata.
var (
	ElasticsearchKeystore_Kind             = "ElasticsearchKeystore"
	ElasticsearchKeystore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ElasticsearchKeystore_Kind}.String()
	ElasticsearchKeystore_KindAPIVersion   = ElasticsearchKeystore_Kind + "." + CRDGroupVersion.String()
	ElasticsearchKeystore_GroupVersionKind = CRDGroupVersion.WithKind(ElasticsearchKeystore_Kind)
)

func init() {
	SchemeBuilder.Register(&ElasticsearchKeystore{}, &ElasticsearchKeystoreList{})
}
