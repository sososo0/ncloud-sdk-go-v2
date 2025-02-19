/*
 * vhadoop
 *
 * <br/>https://ncloud.apigw.ntruss.com/vhadoop/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package vhadoop

type CloudHadoopInstance struct {

	// CloudHadoop인스턴스번호
	CloudHadoopInstanceNo *string `json:"cloudHadoopInstanceNo,omitempty"`

	// CloudHadoop클러스터이름
	CloudHadoopClusterName *string `json:"cloudHadoopClusterName,omitempty"`

	// CloudHadoop이미지상품코드
	CloudHadoopImageProductCode *string `json:"cloudHadoopImageProductCode,omitempty"`

	// CloudHadoop인스턴스상태이름
	CloudHadoopInstanceStatusName *string `json:"cloudHadoopInstanceStatusName,omitempty"`

	// CloudHadoop인스턴스상태
	CloudHadoopInstanceStatus *CommonCode `json:"cloudHadoopInstanceStatus,omitempty"`

	// CloudHadoop인스턴스OP
	CloudHadoopInstanceOperation *CommonCode `json:"cloudHadoopInstanceOperation,omitempty"`

	// CloudHadoop클러스터타입
	CloudHadoopClusterType *CloudHadoopClusterType `json:"cloudHadoopClusterType,omitempty"`

	// CloudHadoop클러스터버전
	CloudHadoopVersion *CloudHadoopVersion `json:"cloudHadoopVersion,omitempty"`

	// CloudHadoop인스턴스OP
	CloudHadoopAddOnList []*CloudHadoopAddOn `json:"cloudHadoopAddOnList,omitempty"`

	// 고가용성여부
	IsHa *string `json:"isHa,omitempty"`

	// Ambari host 이름
	AmbariServerHost *string `json:"ambariServerHost,omitempty"`

	// 클러스터 직접 접속 계정
	ClusterDirectAccessAccount *string `json:"clusterDirectAccessAccount,omitempty"`

	// Cloud Hadoop 인증키
	LoginKey *string `json:"loginKey,omitempty"`

	// Object Storage 버킷 이름
	ObjectStorageBucket *string `json:"objectStorageBucket,omitempty"`

	// KDC Realm
	KdcRealm *string `json:"kdcRealm,omitempty"`

	// 도메인
	Domain *string `json:"domain,omitempty"`

	// CloudHadoop Role
	Role *CommonCode `json:"role,omitempty"`

	// 생성일자
	CreateDate *string `json:"createDate,omitempty"`

	// ACG번호리스트
	AccessControlGroupNoList []*string `json:"accessControlGroupNoList,omitempty"`

	// CloudHadoop서버인스턴스리스트
	CloudHadoopServerInstanceList []*CloudHadoopServerInstance `json:"CloudHadoopServerInstanceList,omitempty"`
}
