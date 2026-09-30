package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/pipeline"

	"github.com/fabioaraujopt/provider-coreweave/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	pc := config.GetProvider()
	dumpGeneratedResourceList(pc, filepath.Join(absRootDir, "config", "generated.lst"))
	pipeline.Run(pc, config.GetProviderNamespaced(), absRootDir)
}

// dumpGeneratedResourceList writes the Terraform names of the generated
// resources as a JSON array. `make schema-version-diff` uses it in CI to
// report native schema version changes when the CoreWeave Terraform provider
// is bumped.
func dumpGeneratedResourceList(p *ujconfig.Provider, targetPath string) {
	generatedResources := make([]string, 0, len(p.Resources))
	for name := range p.Resources {
		generatedResources = append(generatedResources, name)
	}
	sort.Strings(generatedResources)
	// One name per line, so PRs adding resources don't conflict.
	buff, err := json.MarshalIndent(generatedResources, "", "")
	if err != nil {
		panic(fmt.Sprintf("cannot marshal the generated resource list to JSON: %s", err.Error()))
	}
	if err := os.WriteFile(targetPath, buff, 0o600); err != nil {
		panic(fmt.Sprintf("cannot write the generated resource list to file %s: %s", targetPath, err.Error()))
	}
}
