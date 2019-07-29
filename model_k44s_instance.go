/*
 * Oktawave KaaS API
 *
 * Oktawave KaaS API
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type K44sInstance struct {
	Id float32 `json:"Id"`
	Name string `json:"Name"`
	CreationDate string `json:"CreationDate"`
	Subregion *K44sInstanceSubregion `json:"Subregion"`
	Type_ *K44sInstanceType `json:"Type"`
	Status *K44sInstanceStatus `json:"Status"`
	IpAddress string `json:"IpAddress"`
	TotalDisksCapacity float32 `json:"TotalDisksCapacity"`
	CpuNumber float32 `json:"CpuNumber"`
	RamMb float32 `json:"RamMb"`
}
