module github.com/fabioaraujopt/provider-coreweave

go 1.26.8

tool golang.org/x/tools/cmd/goimports

require (
	dario.cat/mergo v1.0.2
	github.com/alecthomas/kingpin/v2 v2.4.0
	github.com/coreweave/terraform-provider-coreweave/xpprovider v0.0.0-00010101000000-000000000000
	github.com/crossplane/crossplane-runtime/v2 v2.5.0-rc.0.0.20260908074656-9b2fb6b1d1ff
	github.com/crossplane/crossplane-tools v0.0.0-20260719180100-659f1dc036c5
	github.com/crossplane/crossplane/apis/v2 v2.4.1
	github.com/crossplane/upjet/v2 v2.5.0
	github.com/pkg/errors v0.9.1
	google.golang.org/grpc v1.83.2
	k8s.io/api v0.36.3
	k8s.io/apiextensions-apiserver v0.36.3
	k8s.io/apimachinery v0.36.3
	k8s.io/client-go v0.36.3
	sigs.k8s.io/controller-runtime v0.24.1
	sigs.k8s.io/controller-tools v0.20.1
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.12-20260415201107-50325440f8f2.1 // indirect
	buf.build/gen/go/coreweave/cks/connectrpc/go v1.19.1-20260326215135-be22b90e1aa6.2 // indirect
	buf.build/gen/go/coreweave/cks/protocolbuffers/go v1.36.11-20260326215135-be22b90e1aa6.1 // indirect
	buf.build/gen/go/coreweave/cwobject/connectrpc/go v1.19.1-20260729060712-bb1877226449.2 // indirect
	buf.build/gen/go/coreweave/cwobject/protocolbuffers/go v1.36.11-20260729060712-bb1877226449.1 // indirect
	buf.build/gen/go/coreweave/inference/connectrpc/go v1.20.0-20260825010541-7cc01a738b29.1 // indirect
	buf.build/gen/go/coreweave/inference/protocolbuffers/go v1.36.12-20260825010541-7cc01a738b29.1 // indirect
	buf.build/gen/go/coreweave/networking/connectrpc/go v1.19.1-20260121155637-a637e7777165.2 // indirect
	buf.build/gen/go/coreweave/networking/protocolbuffers/go v1.36.11-20260121155637-a637e7777165.1 // indirect
	buf.build/gen/go/coreweave/sandbox/connectrpc/go v1.20.0-20260909020241-374666cd3a25.1 // indirect
	buf.build/gen/go/coreweave/sandbox/protocolbuffers/go v1.36.11-20260909020241-374666cd3a25.1 // indirect
	buf.build/gen/go/coreweave/workload-federation/connectrpc/go v1.20.0-20260825153113-5bfae6057a98.1 // indirect
	buf.build/gen/go/coreweave/workload-federation/protocolbuffers/go v1.36.12-20260825153113-5bfae6057a98.1 // indirect
	buf.build/gen/go/gnostic/gnostic/protocolbuffers/go v1.36.11-20230414000709-087bc8072ce4.1 // indirect
	buf.build/gen/go/grpc-ecosystem/grpc-gateway/protocolbuffers/go v1.36.11-20241220201140-4c5ba75caaf8.1 // indirect
	cloud.google.com/go v0.123.0 // indirect
	connectrpc.com/connect v1.20.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/units v0.0.0-20240927000941-0f3dac36c52b // indirect
	github.com/antchfx/htmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.3.6 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/aws/aws-sdk-go-v2 v1.41.7 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.6 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.16 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v1.97.0 // indirect
	github.com/aws/smithy-go v1.25.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coreweave/terraform-provider-coreweave v0.24.0 // indirect
	github.com/dave/jennifer v1.7.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emicklei/go-restful/v3 v3.13.0 // indirect
	github.com/evanphx/json-patch v5.9.11+incompatible // indirect
	github.com/evanphx/json-patch/v5 v5.9.11 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/fxamacker/cbor/v2 v2.9.3 // indirect
	github.com/getkin/kin-openapi v0.144.0 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/zapr v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v1.0.0 // indirect
	github.com/go-openapi/jsonreference v1.0.0 // indirect
	github.com/go-openapi/swag v0.29.0 // indirect
	github.com/go-openapi/swag/cmdutils v0.29.0 // indirect
	github.com/go-openapi/swag/conv v0.29.0 // indirect
	github.com/go-openapi/swag/fileutils v0.29.0 // indirect
	github.com/go-openapi/swag/jsonutils v0.29.0 // indirect
	github.com/go-openapi/swag/loading v0.29.0 // indirect
	github.com/go-openapi/swag/mangling v0.29.0 // indirect
	github.com/go-openapi/swag/netutils v0.29.0 // indirect
	github.com/go-openapi/swag/pools v0.29.0 // indirect
	github.com/go-openapi/swag/stringutils v0.29.0 // indirect
	github.com/go-openapi/swag/typeutils v0.29.0 // indirect
	github.com/go-openapi/swag/yamlutils v0.29.0 // indirect
	github.com/gobuffalo/flect v1.0.3 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.7.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.5.0 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-plugin v1.7.0 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.8 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/hcl/v2 v2.24.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-json v0.27.1 // indirect
	github.com/hashicorp/terraform-plugin-framework v1.15.1 // indirect
	github.com/hashicorp/terraform-plugin-framework-jsontypes v0.2.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-nettypes v0.3.0 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.18.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.29.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.38.1 // indirect
	github.com/hashicorp/terraform-registry-address v0.4.0 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.2 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-ps v1.0.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.3-0.20250322232337-35a7c28c31ee // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/muvaf/typewriter v0.0.0-20240614220100-70f9d4a54ea0 // indirect
	github.com/oasdiff/oasdiff v1.26.1 // indirect
	github.com/oasdiff/yaml v0.1.1 // indirect
	github.com/oasdiff/yaml3 v0.0.14 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.23.2 // indirect
	github.com/prometheus/client_model v0.6.2 // indirect
	github.com/prometheus/common v0.67.5 // indirect
	github.com/prometheus/procfs v0.20.1 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 // indirect
	github.com/spf13/afero v1.15.0 // indirect
	github.com/spf13/cobra v1.10.2 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
	github.com/stretchr/testify v1.11.1 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tmccombs/hcl2json v0.3.3 // indirect
	github.com/upbound/uptest v0.12.1-0.20260728095952-c230a8044006 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/wI2L/jsondiff v0.7.1 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	github.com/yargevad/filepathx v1.0.0 // indirect
	github.com/yuin/goldmark v1.8.4 // indirect
	github.com/zclconf/go-cty v1.17.0 // indirect
	github.com/zclconf/go-cty-yaml v1.0.3 // indirect
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/trace v1.44.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.28.0 // indirect
	go.yaml.in/yaml/v2 v2.4.4 // indirect
	go.yaml.in/yaml/v3 v3.0.5 // indirect
	golang.org/x/mod v0.40.0 // indirect
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/oauth2 v0.36.0 // indirect
	golang.org/x/sync v0.22.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/telemetry v0.0.0-20260819180204-f79da969f323 // indirect
	golang.org/x/term v0.45.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	golang.org/x/tools v0.49.0 // indirect
	gomodules.xyz/jsonpatch/v2 v2.5.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20260810153831-ec0a7760b754 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260810153831-ec0a7760b754 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.13.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/code-generator v0.36.3 // indirect
	k8s.io/component-base v0.36.3 // indirect
	k8s.io/gengo/v2 v2.0.0-20251215205346-5ee0d033ba5b // indirect
	k8s.io/klog/v2 v2.140.0 // indirect
	k8s.io/kube-openapi v0.0.0-20260721132016-d427ff9ee9ad // indirect
	k8s.io/utils v0.0.0-20260707023825-cf1189d6abe3 // indirect
	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v6 v6.4.2 // indirect
	sigs.k8s.io/yaml v1.6.0 // indirect
)

// CoreWeave's Terraform provider keeps its provider in internal/, which only
// packages under its module path may import. ./xpprovider is a tiny module
// under that path that re-exports it from the unmodified upstream release.
replace github.com/coreweave/terraform-provider-coreweave/xpprovider => ./xpprovider
