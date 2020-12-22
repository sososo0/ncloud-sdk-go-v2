/*
 * monitoring
 *
 * <br/>https://ncloud.apigw.ntruss.com/monitoring/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package monitoring

type MetricData struct {

Label *string `json:"label,omitempty"`

Average *float64 `json:"average,omitempty"`

Maximum *float64 `json:"maximum,omitempty"`

Minimum *float64 `json:"minimum,omitempty"`

Sum *float64 `json:"sum,omitempty"`

DataPointList []*DataPoint `json:"dataPointList,omitempty"`
}
