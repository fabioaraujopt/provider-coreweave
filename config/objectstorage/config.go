// Package objectstorage configures CoreWeave AI Object Storage resources.
package objectstorage

import "github.com/crossplane/upjet/v2/pkg/config"

const group = "objectstorage"

// bucketSubresources are keyed by the bucket they configure; each gets a
// bucketRef/bucketSelector so it can point at a Bucket managed resource.
var bucketSubresources = map[string]string{
	"coreweave_object_storage_bucket_inventory":               "BucketInventory",
	"coreweave_object_storage_bucket_lifecycle_configuration": "BucketLifecycleConfiguration",
	"coreweave_object_storage_bucket_policy":                  "BucketPolicy",
	"coreweave_object_storage_bucket_settings":                "BucketSettings",
	"coreweave_object_storage_bucket_versioning":              "BucketVersioning",
}

// Configure configures the object storage resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("coreweave_object_storage_bucket", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "Bucket"
	})
	p.AddResourceConfigurator("coreweave_object_storage_access_key", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "AccessKey"
		// secret_key is Sensitive in the Terraform schema, so Upjet already
		// publishes it to the connection secret as attribute.secret_key. Also
		// expose friendly keys ready for S3 clients.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if v, ok := attr["id"].(string); ok && v != "" {
				conn["access_key_id"] = []byte(v)
			}
			if v, ok := attr["secret_key"].(string); ok && v != "" {
				conn["secret_access_key"] = []byte(v)
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("coreweave_object_storage_organization_access_policy", func(r *config.Resource) {
		r.ShortGroup = group
		r.Kind = "OrganizationAccessPolicy"
	})
	for name, kind := range bucketSubresources {
		p.AddResourceConfigurator(name, func(r *config.Resource) {
			r.ShortGroup = group
			r.Kind = kind
			r.References["bucket"] = config.Reference{
				TerraformName: "coreweave_object_storage_bucket",
			}
		})
	}
}
