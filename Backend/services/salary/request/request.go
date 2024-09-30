package request

import "time"

type Request struct {
	EmployeeId             int64     `json:"employeeId"`
	EmployeeName           string    `json:"employeeName"`
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
	TaxEndPay              time.Time `json:"taxEndPay"`
	TaxStartPay            time.Time `json:"taxStartPay"`
}
