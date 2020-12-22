/*
 * autoscaling
 *
 * <br/>https://ncloud.apigw.ntruss.com/autoscaling/v2
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package autoscaling

type GetAutoScalingActivityLogListRequest struct {

	// 액티비티번호리스트
ActivityNoList []*string `json:"activityNoList,omitempty"`

	// 오토스케일링그룹명
AutoScalingGroupName *string `json:"autoScalingGroupName,omitempty"`

	// 페이지번호
PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
PageSize *int32 `json:"pageSize,omitempty"`
}
