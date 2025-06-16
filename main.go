package main

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	stdinData, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}

	var inputList InputResourceList
	err = yaml.Unmarshal(stdinData, &inputList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling stdin YAML: %v\n", err)
		os.Exit(1)
	}

	resourcePath := inputList.FunctionConfig.Spec.ResourcePath
	if resourcePath == "" {
		fmt.Fprintln(os.Stderr, "Error: functionConfig.spec.resourcePath is empty")
		os.Exit(1)
	}

	resourceFileContent, err := os.ReadFile(resourcePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading resource file %s: %v\n", resourcePath, err)
		os.Exit(1)
	}

	var resource GenericResource
	err = yaml.Unmarshal(resourceFileContent, &resource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling resource file YAML from %s: %v\n", resourcePath, err)
		os.Exit(1)
	}

	if resource.Metadata.Annotations == nil {
		resource.Metadata.Annotations = make(map[string]string)
	}

	resource.Metadata.Annotations["kustomize.config.k8s.io/needs-hash"] = "true"

	outputList := OutputResourceList{
		Kind:  "ResourceList",
		Items: []*GenericResource{&resource},
	}

	outputYAML, err := yaml.Marshal(&outputList)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshalling output YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(outputYAML))
}
