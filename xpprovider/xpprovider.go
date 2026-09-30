// Package xpprovider exposes the CoreWeave Terraform Plugin Framework provider
// to provider-coreweave.
//
// The upstream provider constructor lives in internal/provider, which Go only
// lets packages under github.com/coreweave/terraform-provider-coreweave/
// import. This package is its own module whose path sits under that prefix,
// so it can import internal/provider from the unmodified upstream release,
// and provider-coreweave points at it with a replace directive. No fork of
// the Terraform provider is needed: bumping CoreWeave is a version change in
// both go.mod files.
package xpprovider

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/coreweave/terraform-provider-coreweave/internal/provider"
)

// New returns a new instance of the CoreWeave Terraform Plugin Framework
// provider for the given version string.
func New(version string) fwprovider.Provider {
	return provider.New(version)()
}
