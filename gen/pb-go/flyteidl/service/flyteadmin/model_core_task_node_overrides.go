/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Optional task node overrides that will be applied at task execution time.
type CoreTaskNodeOverrides struct {
	// A customizable interface to convey resources requested for a task container.
	Resources *CoreResources `json:"resources,omitempty"`
	ResourceMetadata *CoreResourceMetadata `json:"resource_metadata,omitempty"`
}
