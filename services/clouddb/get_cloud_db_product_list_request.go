/*
 * clouddb
 *
 * Cloud DB<br/>https://ncloud.apigw.ntruss.com/clouddb/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package clouddb

type GetCloudDbProductListRequest struct {

	// 제외할상품코드
ExclusionProductCode *string `json:"exclusionProductCode,omitempty"`

	// CloudDB이미지상품코드
CloudDBImageProductCode *string `json:"cloudDBImageProductCode"`

	// 조회할상품코드
ProductCode *string `json:"productCode,omitempty"`

	// 리전번호
RegionNo *string `json:"regionNo,omitempty"`

	// zone번호
ZoneNo *string `json:"zoneNo,omitempty"`
}
