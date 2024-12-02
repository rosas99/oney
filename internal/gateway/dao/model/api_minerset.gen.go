// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMinerSetM = "api_minerset"

// MinerSetM mapped from table <api_minerset>
type MinerSetM struct {
	ID                   int64     `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`                             // 主键 ID
	Namespace            string    `gorm:"column:namespace;type:varchar(253);not null;uniqueIndex:uniq_namespace_name,priority:1;comment:命名空间" json:"namespace"` // 命名空间
	Name                 string    `gorm:"column:name;type:varchar(253);not null;uniqueIndex:uniq_namespace_name,priority:2;comment:矿机池名" json:"name"`           // 矿机池名
	Replicas             int32     `gorm:"column:replicas;type:int(8);not null;comment:矿机副本数" json:"replicas"`                                                   // 矿机副本数
	DisplayName          string    `gorm:"column:display_name;type:varchar(253);not null;comment:矿机池展示名" json:"display_name"`                                    // 矿机池展示名
	DeletePolicy         string    `gorm:"column:delete_policy;type:varchar(32);not null;comment:矿机池缩容策略" json:"delete_policy"`                                  // 矿机池缩容策略
	MinReadySeconds      int32     `gorm:"column:min_ready_seconds;type:int(8);not null;comment:矿机 Ready 最小等待时间" json:"min_ready_seconds"`                       // 矿机 Ready 最小等待时间
	FullyLabeledReplicas int32     `gorm:"column:fully_labeled_replicas;type:int(8);not null;comment:所有标签匹配的副本数" json:"fully_labeled_replicas"`                  // 所有标签匹配的副本数
	ReadyReplicas        int32     `gorm:"column:ready_replicas;type:int(8);not null;comment:Ready 副本数" json:"ready_replicas"`                                   // Ready 副本数
	AvailableReplicas    int32     `gorm:"column:available_replicas;type:int(8);not null;comment:可用副本数" json:"available_replicas"`                               // 可用副本数
	FailureReason        *string   `gorm:"column:failure_reason;type:longtext;comment:失败原因" json:"failure_reason"`                                               // 失败原因
	FailureMessage       *string   `gorm:"column:failure_message;type:longtext;comment:失败信息" json:"failure_message"`                                             // 失败信息
	Conditions           *string   `gorm:"column:conditions;type:longtext;comment:矿机池状态" json:"conditions"`                                                      // 矿机池状态
	CreatedAt            time.Time `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp();comment:创建时间" json:"created_at"`                 // 创建时间
	UpdatedAt            time.Time `gorm:"column:updated_at;type:timestamp;not null;default:current_timestamp();comment:最后修改时间" json:"updated_at"`               // 最后修改时间
}

// TableName MinerSetM's table name
func (*MinerSetM) TableName() string {
	return TableNameMinerSetM
}