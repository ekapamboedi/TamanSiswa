package salary

import (
	"database/sql"
	"errors"
	"time"

	"TamanSiswa/helper"
	"TamanSiswa/model"
	"TamanSiswa/services/salary/request"
)

func GetAllRecord(identity helper.TokenPayload) (*[]model.Salary, error) {
	var resultData []model.Salary
	result := model.DB.Table("salaries").Where("company_id = ?", identity.CompanyId).Find(&resultData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &resultData, nil
}

func GetWithId(identity helper.TokenPayload, id string) (*model.Salary, error) {
	var resultData model.Salary
	result := model.DB.Table("salaries").Where("employee_id = ? AND company_id = ?", id, identity.CompanyId).First(&resultData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &resultData, nil
}

func UpdateOne(identity helper.TokenPayload, req request.Request, id string) error {
	var data model.Salary
	var resultEmployee model.Employee

	result := model.DB.Table("employees").Where("id = ? AND company_id = ?", req.EmployeeId, identity.CompanyId).First(&resultEmployee)
	if result.Error != nil {
		return result.Error
	}

	data.EmployeeId = *helper.ValidateSqlInt64(req.EmployeeId)
	data.EmployeeName = *helper.ValidateSqlString(req.EmployeeName)
	data.CompanyId = sql.NullString{String: identity.CompanyId, Valid: true}
	data.DepartmentId = resultEmployee.DepartmentId
	data.DepartmentName = resultEmployee.DepartmentName
	data.PositionId = resultEmployee.PositionId
	data.PositionName = resultEmployee.PositionName
	data.Currency = *helper.ValidateSqlString(req.Currency)
	data.Earning = *helper.ValidateSqlInt64(req.Earning)
	data.MartialAndDependents = *helper.ValidateSqlString(req.MartialAndDependents)
	data.BankName = *helper.ValidateSqlString(req.BankName)
	data.BankAccount = *helper.ValidateSqlString(req.BankAccount)
	data.PaymentMethod = *helper.ValidateSqlString(req.PaymentMethod)
	data.BpjsEmployeeNo = *helper.ValidateSqlString(req.BpjsHealthCareNo)
	data.BpjsEmployeeStartPay = *helper.ValidateSqlTime(req.BpjsEmployeeStartPay)
	data.BpjsEmployeeEndPay = *helper.ValidateSqlTime(req.BpjsEmployeeEndPay)
	data.BpjsHealthCareNo = *helper.ValidateSqlString(req.BpjsHealthCareNo)
	data.BpjsHealthCareStartPay = *helper.ValidateSqlTime(req.BpjsHealthCareStartPay)
	data.BpjsHealthCareEndPay = *helper.ValidateSqlTime(req.BpjsHealthCareEndPay)
	data.NpwpNo = *helper.ValidateSqlString(req.NpwpNo)
	data.TaxStartPay = *helper.ValidateSqlTime(req.TaxStartPay)
	data.TaxEndPay = *helper.ValidateSqlTime(req.TaxEndPay)
	data.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	result = model.DB.Table("salaries").Where("employee_id = ? AND company_id = ?", id, identity.CompanyId).Updates(&data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func CreateOne(identity helper.TokenPayload, req request.Request) error {
	var data model.Salary
	var resultEmployee model.Employee

	result := model.DB.Table("employees").Where("id = ? AND company_id = ?", req.EmployeeId, identity.CompanyId).First(&resultEmployee)
	if result.Error != nil {
		return result.Error
	}

	data.EmployeeId = *helper.ValidateSqlInt64(req.EmployeeId)
	data.EmployeeName = *helper.ValidateSqlString(req.EmployeeName)
	data.CompanyId = *helper.ValidateSqlString(identity.CompanyId)
	data.DepartmentId = resultEmployee.DepartmentId
	data.DepartmentName = resultEmployee.DepartmentName
	data.PositionId = resultEmployee.PositionId
	data.PositionName = resultEmployee.PositionName
	data.Currency = *helper.ValidateSqlString(req.Currency)
	data.Earning = *helper.ValidateSqlInt64(req.Earning)
	data.MartialAndDependents = *helper.ValidateSqlString(req.MartialAndDependents)
	data.BankName = *helper.ValidateSqlString(req.BankName)
	data.BankAccount = *helper.ValidateSqlString(req.BankAccount)
	data.PaymentMethod = *helper.ValidateSqlString(req.PaymentMethod)
	data.BpjsEmployeeNo = *helper.ValidateSqlString(req.BpjsHealthCareNo)
	data.BpjsEmployeeStartPay = *helper.ValidateSqlTime(req.BpjsEmployeeStartPay)
	data.BpjsEmployeeEndPay = *helper.ValidateSqlTime(req.BpjsEmployeeEndPay)
	data.BpjsHealthCareNo = *helper.ValidateSqlString(req.BpjsHealthCareNo)
	data.BpjsHealthCareStartPay = *helper.ValidateSqlTime(req.BpjsHealthCareStartPay)
	data.BpjsHealthCareEndPay = *helper.ValidateSqlTime(req.BpjsHealthCareEndPay)
	data.NpwpNo = *helper.ValidateSqlString(req.NpwpNo)
	data.TaxStartPay = *helper.ValidateSqlTime(req.TaxStartPay)
	data.TaxEndPay = *helper.ValidateSqlTime(req.TaxEndPay)
	data.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	result = model.DB.Table("salaries").Create(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
