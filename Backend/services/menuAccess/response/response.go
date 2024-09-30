package response

import (
	"encoding/json"
	"time"

	"TamanSiswa/model"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseDetail struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    *MenuAccess `json:"data,omitempty"`
}

type ResponseList struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    *[]MenuAccess `json:"data,omitempty"`
}

type MenuAccess struct {
	Id             string           `json:"id"`
	PositionId     string           `json:"positionId"`
	PositionName   string           `json:"positionName"`
	DepartmentName string           `json:"departmentName"`
	Access         []map[string]any `json:"access"`
	CreatedAt      time.Time        `json:"createdAt"`
	UpdatedAt      time.Time        `json:"updatedAt"`
}

func (r *ResponseDetail) Set(data model.MenuAccess) {
	var tmp MenuAccess
	var jsonData []map[string]any

	json.Unmarshal([]byte(data.Access.String), &jsonData)

	tmp.Id = data.Id
	tmp.PositionId = data.PositionId.String
	tmp.Access = jsonData
	tmp.CreatedAt = data.CreatedAt.Time
	tmp.UpdatedAt = data.UpdatedAt.Time

	r.Data = &tmp
}

func (r *ResponseList) Set(menuAccessData []model.MenuAccess, positionData []model.Position) {
	var tmp []MenuAccess

	for _, item := range menuAccessData {
		var jsonData []map[string]any
		var singleData MenuAccess
		json.Unmarshal([]byte(item.Access.String), &jsonData)

		singleData.Id = item.Id
		singleData.PositionId = item.PositionId.String
		singleData.Access = jsonData
		singleData.CreatedAt = item.CreatedAt.Time
		singleData.UpdatedAt = item.UpdatedAt.Time

		for _, item2 := range positionData {
			if item.PositionId.String == item2.Id {
				singleData.PositionName = item2.PositionName.String
				singleData.DepartmentName = item2.DepartmentName.String
				break
			}
		}

		tmp = append(tmp, singleData)
	}

	r.Data = &tmp
}
