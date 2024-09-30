package response

import (
	"time"

	"TamanSiswa/model"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseDetail struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    *Salary `json:"data,omitempty"`
}

type Salary struct {
	Id                     string    `gorm:"primaryKey" json:"id"`
	EmployeeId             int64     `json:"employeeId"`
	EmployeeName           string    `json:"employeeName"`
	CompanyId              string    `json:"companyId"`
	DepartmentId           string    `json:"departmentId"`
	DepartmentName         string    `json:"departmentName"`
	PositionId             string    `json:"positionId"`
	PositionName           string    `json:"positionName"`
	Earning                int64     `json:"earning"`
	MartialAndDependents   string    `json:"martialAndDependents"`
	Currency               string    `json:"currency"`
	BankName               string    `json:"bankName"`
	BankAccount            string    `json:"bankAccount"`
	PaymentMethod          string    `json:"paymentMethod"`
	BpjsEmployeeNo         string    `json:"bpjsEmployeeNo"`
	BpjsEmployeeStartPay   time.Time `json:"bpjsEmployeeStartPay"`
	BpjsEmployeeEndPay     time.Time `json:"bpjsEmployeeEndPay"`
	BpjsHealthCareNo       string    `json:"bpjsHealthCareNo"`
	BpjsHealthCareStartPay time.Time `json:"bpjsHealthCareStartPay"`
	BpjsHealthCareEndPay   time.Time `json:"bpjsHealthCareEndPay"`
	NpwpNo                 string    `json:"npwpNo"`
	TaxStartPay            time.Time `json:"taxStartPay"`
	TaxEndPay              time.Time `json:"taxEndPay"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}

func (res *ResponseDetail) Set(data model.Salary) {
	res.Data = &Salary{
		Id:                     data.Id,
		EmployeeId:             data.EmployeeId.Int64,
		EmployeeName:           data.EmployeeName.String,
		CompanyId:              data.CompanyId.String,
		DepartmentId:           data.DepartmentId.String,
		DepartmentName:         data.DepartmentName.String,
		PositionId:             data.PositionId.String,
		PositionName:           data.PositionName.String,
		Earning:                data.Earning.Int64,
		Currency:               data.Currency.String,
		MartialAndDependents:   data.MartialAndDependents.String,
		BankName:               data.BankName.String,
		BankAccount:            data.BankAccount.String,
		PaymentMethod:          data.PaymentMethod.String,
		BpjsEmployeeNo:         data.BpjsEmployeeNo.String,
		BpjsEmployeeStartPay:   data.BpjsEmployeeStartPay.Time,
		BpjsEmployeeEndPay:     data.BpjsEmployeeEndPay.Time,
		BpjsHealthCareNo:       data.BpjsHealthCareNo.String,
		BpjsHealthCareStartPay: data.BpjsHealthCareStartPay.Time,
		BpjsHealthCareEndPay:   data.BpjsHealthCareEndPay.Time,
		NpwpNo:                 data.NpwpNo.String,
		TaxStartPay:            data.TaxStartPay.Time,
		TaxEndPay:              data.TaxEndPay.Time,
		CreatedAt:              data.CreatedAt.Time,
		UpdatedAt:              data.UpdatedAt.Time,
	}
}

type ResponseList struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    *[]Salary `json:"data,omitempty"`
}

func (res *ResponseList) Set(data []model.Salary) {
	var allData []Salary

	for _, item := range data {
		var singleData = Salary{
			Id:                     item.Id,
			EmployeeId:             item.EmployeeId.Int64,
			EmployeeName:           item.EmployeeName.String,
			CompanyId:              item.CompanyId.String,
			DepartmentId:           item.DepartmentId.String,
			DepartmentName:         item.DepartmentName.String,
			PositionId:             item.PositionId.String,
			PositionName:           item.PositionName.String,
			Earning:                item.Earning.Int64,
			Currency:               item.Currency.String,
			MartialAndDependents:   item.MartialAndDependents.String,
			BankName:               item.BankName.String,
			BankAccount:            item.BankAccount.String,
			PaymentMethod:          item.PaymentMethod.String,
			BpjsEmployeeNo:         item.BpjsEmployeeNo.String,
			BpjsEmployeeStartPay:   item.BpjsEmployeeStartPay.Time,
			BpjsEmployeeEndPay:     item.BpjsEmployeeEndPay.Time,
			BpjsHealthCareNo:       item.BpjsHealthCareNo.String,
			BpjsHealthCareStartPay: item.BpjsHealthCareStartPay.Time,
			BpjsHealthCareEndPay:   item.BpjsHealthCareEndPay.Time,
			NpwpNo:                 item.NpwpNo.String,
			TaxStartPay:            item.TaxStartPay.Time,
			TaxEndPay:              item.TaxEndPay.Time,
			CreatedAt:              item.CreatedAt.Time,
			UpdatedAt:              item.UpdatedAt.Time,
		}

		allData = append(allData, singleData)
	}

	res.Data = &allData
}
