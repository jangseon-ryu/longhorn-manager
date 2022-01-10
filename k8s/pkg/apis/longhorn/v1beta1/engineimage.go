package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type EngineImageState string

const (
	EngineImageStateDeploying    = EngineImageState("deploying")
	EngineImageStateDeployed     = EngineImageState("deployed")
	EngineImageStateIncompatible = EngineImageState("incompatible")
	EngineImageStateError        = EngineImageState("error")
)

const (
	EngineImageConditionTypeReady = "ready"

	EngineImageConditionTypeReadyReasonDaemonSet = "daemonSet"
	EngineImageConditionTypeReadyReasonBinary    = "binary"
)

type EngineVersionDetails struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	BuildDate string `json:"buildDate"`

	CLIAPIVersion           int `json:"cliAPIVersion"`
	CLIAPIMinVersion        int `json:"cliAPIMinVersion"`
	ControllerAPIVersion    int `json:"controllerAPIVersion"`
	ControllerAPIMinVersion int `json:"controllerAPIMinVersion"`
	DataFormatVersion       int `json:"dataFormatVersion"`
	DataFormatMinVersion    int `json:"dataFormatMinVersion"`
}

// EngineImageSpec defines the desired state of the Longhorn engine image
type EngineImageSpec struct {
	Image string `json:"image"`
}

// EngineImageStatus defines the observed state of the Longhorn engine image
type EngineImageStatus struct {
	OwnerID           string               `json:"ownerID"`
	State             EngineImageState     `json:"state"`
	RefCount          int                  `json:"refCount"`
	NoRefSince        string               `json:"noRefSince"`
	Conditions        map[string]Condition `json:"conditions"`
	NodeDeploymentMap map[string]bool      `json:"nodeDeploymentMap"`

	EngineVersionDetails `json:""`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:shortName=lhei
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`,description="State of the engine image"
// +kubebuilder:printcolumn:name="Image",type=string,JSONPath=`.spec.image`,description="The Longhorn engine image"
// +kubebuilder:printcolumn:name="RefCount",type=integer,JSONPath=`.status.refCount`,description="Number of volumes are using the engine image"
// +kubebuilder:printcolumn:name="BuildDate",type=date,JSONPath=`.status.buildDate`,description="The build date of the engine image"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// EngineImage is where Longhorn stores engine image object.
type EngineImage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Spec EngineImageSpec `json:"spec,omitempty"`
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Status EngineImageStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EngineImageList is a list of EngineImages.
type EngineImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EngineImage `json:"items"`
}
