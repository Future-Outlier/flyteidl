/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Encapsulates specifications for routing an execution onto a specific cluster.
type AdminClusterAssignment struct {
	Affinity *AdminAffinity `json:"affinity,omitempty"`
	Toleration *AdminToleration `json:"toleration,omitempty"`
}