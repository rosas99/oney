// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameJobM = "nw_job"

// JobM mapped from table <nw_job>
type JobM struct {
	ID          int64          `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`                        // 主键 ID
	JobID       string         `gorm:"column:job_id;type:varchar(100);not null;comment:Job ID" json:"job_id"`                                           // Job ID
	UserID      string         `gorm:"column:user_id;type:varchar(100);not null;index:idx_user_id,priority:1;comment:创建人" json:"user_id"`               // 创建人
	Scope       string         `gorm:"column:scope;type:varchar(256);not null;index:idx_scope,priority:1;default:default;comment:Job 作用域" json:"scope"` // Job 作用域
	Name        string         `gorm:"column:name;type:varchar(255);not null;comment:Job 名称" json:"name"`                                               // Job 名称
	Description string         `gorm:"column:description;type:varchar(256);not null;comment:Job 描述" json:"description"`                                 // Job 描述
	CronJobID   *string        `gorm:"column:cronjob_id;type:varchar(100);index:idx_task_id,priority:1;comment:CronJob ID，可选" json:"cronjob_id"`        // CronJob ID，可选
	Watcher     string         `gorm:"column:watcher;type:varchar(255);not null;comment:eam-nightwatch watcher 名字" json:"watcher"`                      // eam-nightwatch watcher 名字
	Suspend     int32          `gorm:"column:suspend;type:tinyint(4);not null;comment:是否挂起（1 挂起，0 不挂起）" json:"suspend"`                                 // 是否挂起（1 挂起，0 不挂起）
	Params      *JobParams     `gorm:"column:params;type:longtext;comment:Job 参数" json:"params"`                                                        // Job 参数
	Results     *JobResults    `gorm:"column:results;type:longtext;comment:Job 执行结果" json:"results"`                                                    // Job 执行结果
	Status      string         `gorm:"column:status;type:varchar(32);not null;default:Pending;comment:Job 状态" json:"status"`                            // Job 状态
	Conditions  *JobConditions `gorm:"column:conditions;type:longtext;comment:Job 运行状态" json:"conditions"`                                              // Job 运行状态
	StartedAt   time.Time      `gorm:"column:started_at;type:datetime;not null;comment:Job 开始时间" json:"started_at"`                                     // Job 开始时间
	EndedAt     time.Time      `gorm:"column:ended_at;type:datetime;not null;comment:Job 结束时间" json:"ended_at"`                                         // Job 结束时间
	CreatedAt   time.Time      `gorm:"column:created_at;type:timestamp;not null;default:current_timestamp();comment:创建时间" json:"created_at"`            // 创建时间
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:timestamp;not null;default:current_timestamp();comment:更新时间" json:"updated_at"`            // 更新时间
}

// TableName JobM's table name
func (*JobM) TableName() string {
	return TableNameJobM
}
