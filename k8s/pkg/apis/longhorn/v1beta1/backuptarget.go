package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	BackupTargetConditionTypeUnavailable = "Unavailable"

	BackupTargetConditionReasonUnavailable = "Unavailable"
)

// BackupTargetSpec defines the desired state of the Longhorn backup target
type BackupTargetSpec struct {
	// The backup target URL.
	BackupTargetURL string `json:"backupTargetURL"`
	// The backup target credential secret.
	CredentialSecret string `json:"credentialSecret"`
	// The interval that the cluster needs to run sync with the backup target.
	PollInterval metav1.Duration `json:"pollInterval"`
	// The time to request run sync the remote backup target.
	SyncRequestedAt metav1.Time `json:"syncRequestedAt"`
}

// BackupTargetStatus defines the observed state of the Longhorn backup target
type BackupTargetStatus struct {
	// The node ID on which the controller is responsible to reconcile this backup target CR.
	OwnerID string `json:"ownerID"`
	// Available indicates if the remote backup target is available or not.
	Available bool `json:"available"`
	// Records the reason on why the backup target is unavailable.
	Conditions map[string]Condition `json:"conditions"`
	// The last time that the controller synced with the remote backup target.
	LastSyncedAt metav1.Time `json:"lastSyncedAt"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=lhbt
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="URL",type=string,JSONPath=`.spec.backupTargetURL`,description="The backup target URL"
// +kubebuilder:printcolumn:name="Credential",type=string,JSONPath=`.spec.credentialSecret`,description="The backup target credential secret"
// +kubebuilder:printcolumn:name="LastBackupAt",type=string,JSONPath=`.spec.pollInterval`,description="The backup target poll interval"
// +kubebuilder:printcolumn:name="Available",type=boolean,JSONPath=`.status.available`,description="Indicate whether the backup target is available or not"
// +kubebuilder:printcolumn:name="LastSyncedAt",type=string,JSONPath=`.status.lastSyncedAt`,description="The backup target last synced time"

// BackupTarget is where Longhorn stores backup target object.
type BackupTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Spec BackupTargetSpec `json:"spec,omitempty"`
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Status BackupTargetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BackupTargetList is a list of BackupTargets.
type BackupTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupTarget `json:"items"`
}
