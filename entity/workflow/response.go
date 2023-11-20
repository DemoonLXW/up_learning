package workflow

import (
	"time"
)

type Variables struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
	Type  string      `json:"type,omitempty"`
	Scope string      `json:"scope,omitempty"`
}

type HistoricProcessInstances struct {
	ID        string      `json:"id,omitempty"`
	StartTime *time.Time  `json:"startTime,omitempty"`
	EndTime   *time.Time  `json:"endTime,omitempty"`
	Variables []Variables `json:"variables,omitempty"`
}

type HistoricProcessInstancesPageList struct {
	Data  []HistoricProcessInstances `json:"data,omitempty"`
	Total int                        `json:"total,omitempty"`
	Size  int                        `json:"size,omitempty"`
	Start int                        `json:"start,omitempty"`
	Order string                     `json:"order,omitempty"`
	Sort  string                     `json:"sort,omitempty"`
}

type HistoricTaskInstances struct {
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Assignee  string      `json:"assignee,omitempty"`
	StartTime *time.Time  `json:"startTime,omitempty"`
	EndTime   *time.Time  `json:"endTime,omitempty"`
	Variables []Variables `json:"variables,omitempty"`
}

type HistoricTaskInstancesPageList struct {
	Data  []HistoricTaskInstances `json:"data,omitempty"`
	Total int                     `json:"total,omitempty"`
	Size  int                     `json:"size,omitempty"`
	Start int                     `json:"start,omitempty"`
	Order string                  `json:"order,omitempty"`
	Sort  string                  `json:"sort,omitempty"`
}
