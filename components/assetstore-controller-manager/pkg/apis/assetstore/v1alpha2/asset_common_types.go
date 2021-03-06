package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// CommonAssetSpec defines the desired state of Asset
type CommonAssetSpec struct {
	Source    AssetSource    `json:"source"`
	BucketRef AssetBucketRef `json:"bucketRef,omitempty"`
}

// CommonAssetStatus defines the observed state of Asset
type CommonAssetStatus struct {
	Phase              AssetPhase     `json:"phase"`
	Message            string         `json:"message,omitempty"`
	Reason             string         `json:"reason,omitempty"`
	AssetRef           AssetStatusRef `json:"assetRef,omitempty"`
	LastHeartbeatTime  metav1.Time    `json:"lastHeartbeatTime"`
	ObservedGeneration int64          `json:"observedGeneration"`
}

type AssetPhase string

const (
	AssetReady   AssetPhase = "Ready"
	AssetPending AssetPhase = "Pending"
	AssetFailed  AssetPhase = "Failed"
)

type AssetStatusRef struct {
	BaseUrl string   `json:"baseUrl"`
	Assets  []string `json:"assets,omitempty"`
}

type AssetWebhookService struct {
	Name      string                `json:"name,omitempty"`
	Namespace string                `json:"namespace,omitempty"`
	Endpoint  string                `json:"endpoint,omitempty"`
	Metadata  *runtime.RawExtension `json:"metadata,omitempty"`
}

type AssetMode string

const (
	AssetSingle  AssetMode = "single"
	AssetPackage AssetMode = "package"
	AssetIndex   AssetMode = "index"
)

type AssetBucketRef struct {
	Name string `json:"name"`
}

type AssetSource struct {
	// +kubebuilder:validation:Enum=single,package,index
	Mode AssetMode `json:"mode"`
	Url  string    `json:"url"`
	// +optional
	Filter string `json:"filter,omitempty"`

	// +optional
	ValidationWebhookService []AssetWebhookService `json:"validationWebhookService,omitempty"`

	// +optional
	MutationWebhookService []AssetWebhookService `json:"mutationWebhookService,omitempty"`
}
