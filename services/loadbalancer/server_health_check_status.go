/*
 * loadbalancer
 *
 * <br/>https://ncloud.apigw.ntruss.com/loadbalancer/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package loadbalancer

type ServerHealthCheckStatus struct {

ProtocolType *CommonCode `json:"protocolType,omitempty"`

LoadBalancerPort *int32 `json:"loadBalancerPort,omitempty"`

	// 서버포트
ServerPort *int32 `json:"serverPort,omitempty"`

L7HealthCheckPath *string `json:"l7HealthCheckPath,omitempty"`

ProxyProtocolUseYn *string `json:"proxyProtocolUseYn,omitempty"`

	// 서버상태
ServerStatus *bool `json:"serverStatus,omitempty"`
}
