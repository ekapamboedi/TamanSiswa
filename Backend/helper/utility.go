package helper

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"TamanSiswa/config"
	viewmodel "TamanSiswa/viewModel"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var Month = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var MartialStatus = map[string]float64{
	"TK/0": 54000000,
	"K/0":  58500000,
	"K/1":  63000000,
	"K/2":  67500000,
	"K/3":  72000000,
}

var PPH21TaxPercentage = []map[string]float64{
	{
		"amount":     60000000,
		"percentage": 0.05,
	},
	{
		"amount":     190000000,
		"percentage": 0.15,
	},
	{
		"amount":     250000000,
		"percentage": 0.25,
	},
	{
		"amount":     4500000000,
		"percentage": 0.3,
	},
	{
		"percentage": 0.35,
	},
}

type TokenPayload struct {
	Username       string `json:"username"`
	Role           string `json:"role"`
	UserId         string `json:"id"`
	EmployeeId     string `json:"employee_id"`
	CompanyId      string `json:"company_id"`
	CompanyName    string `json:"company_name"`
	City           string `json:"city"`
	CompanyLogoUrl string `json:"icon_url"`
}

func FindMonth(month string) string {
	for index, item := range Month {
		if month == item {
			return strconv.Itoa(index + 1)
		}
	}
	return "-1"
}

func FindTax(martialStatus string, yearlyEarning float64) float64 {
	if yearlyEarning > MartialStatus[martialStatus] {
		taxedEarning := yearlyEarning - MartialStatus[martialStatus]
		var pph21Tax float64 = 0
		for _, item := range PPH21TaxPercentage {
			if taxedEarning-item["amount"] < 0 {
				pph21Tax += (taxedEarning * item["percentage"]) / 12
				break
			} else {
				taxedEarning = taxedEarning - item["amount"]
				pph21Tax = (item["amount"] * item["percentage"]) / 12
			}
		}

		return float64(int(pph21Tax * 1))
	}
	return 0
}

func CalculateAttendance(data []viewmodel.ViewAttendances, employeeId int64) map[string]int64 {
	var late int64 = 0
	var notAttend int64 = 0

	for _, item := range data {
		if item.Id == employeeId {
			if item.ClockIn == "" {
				notAttend++
			} else if value, _ := strconv.ParseFloat(item.Late, 64); value > 0 {
				late++
			}
		}
	}

	return map[string]int64{
		"late":      late,
		"notAttend": notAttend,
	}
}

func StructThousandNumber(data int64) string {
	newPrint := message.NewPrinter(language.Indonesian)
	seperatedNumber := newPrint.Sprintf("%d", data)
	return seperatedNumber
}

func CalculateJoinMonths(joinDate time.Time) int64 {
	today := time.Now()

	yearDiff := int64(today.Year()) - int64(joinDate.Year())
	monthDiff := int64(today.Month()) - int64(joinDate.Month())

	totalMonthDiff := yearDiff*12 + monthDiff

	if totalMonthDiff > 12 {
		return 12
	}
	return totalMonthDiff
}

func Rounding(n float64) float64 {
	return math.Round(n*math.Pow(10, 0)) / math.Pow(10, 0)
}

func TodayDate() string {
	today := time.Now()

	return today.Format("02-January-2006")
}

func CalculateWorkdays(start time.Time) float64 {
	var days float64 = 0
	today := time.Now()
	for d := start; !d.After(today); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			days++
		}

		if days >= 22 {
			break
		}
	}

	return days / 22
}

func GetIdentity(ctx *gin.Context) TokenPayload {
	var tokenBody TokenPayload
	header := ctx.GetHeader("authorization")

	headerSplit := strings.Split(header, " ")
	tokenText := headerSplit[1]

	token, _ := jwt.Parse(tokenText, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.JWT_KEY), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// tokenBody.CompanyId = claims["company_id"].(string)
		// tokenBody.CompanyName = claims["company_name"].(string)
		// tokenBody.CompanyLogoUrl = claims["icon_url"].(string)
		// tokenBody.UserId = claims["id"].(string)
		// tokenBody.EmployeeId = claims["employee_id"].(string)
		// tokenBody.City = claims["company_city"].(string)
		tokenBody.Username = claims["username"].(string)
		tokenBody.Role = claims["role"].(string)
	}

	return tokenBody
}
