package main

// Metadata structure for Kubernetes resources
type Metadata struct {
	Name         string                 `yaml:"name,omitempty"`
	GenerateName string                 `yaml:"generateName,omitempty"`
	Annotations  map[string]string      `yaml:"annotations,omitempty"`
	OtherFields  map[string]interface{} `yaml:",inline"` // Captures all other metadata fields
}

// GenericResource represents a Kubernetes resource, like a Job.
type GenericResource struct {
	APIVersion  string                 `yaml:"apiVersion"`
	Kind        string                 `yaml:"kind"`
	Metadata    Metadata               `yaml:"metadata"`
	Spec        map[string]interface{} `yaml:"spec,omitempty"`
	Data        map[string]interface{} `yaml:"data,omitempty"` // For ConfigMap/Secret
	OtherFields map[string]interface{} `yaml:",inline"`        // Captures other top-level fields
}

// FunctionConfigSpec is part of the function's configuration.
type FunctionConfigSpec struct {
	ResourcePath string `yaml:"resourcePath"`
}

// KRMFunctionConfig represents the 'functionConfig' field in the KRM Function input.
// It's the custom resource that configures this function.
type KRMFunctionConfig struct {
	APIVersion string             `yaml:"apiVersion"`
	Kind       string             `yaml:"kind"`
	Metadata   Metadata           `yaml:"metadata,omitempty"`
	Spec       FunctionConfigSpec `yaml:"spec"`
}

// InputResourceList is the structure expected from stdin, conforming to KRM Function spec.
type InputResourceList struct {
	Kind           string            `yaml:"kind"` // Expected to be "ResourceList"
	FunctionConfig KRMFunctionConfig `yaml:"functionConfig"`
	// Items []*GenericResource `yaml:"items,omitempty"` // Other items in the list, not used by this script's logic
}

// OutputResourceList is the structure to be marshalled to YAML and printed to stdout.
type OutputResourceList struct {
	Kind  string             `yaml:"kind"`
	Items []*GenericResource `yaml:"items"`
}
