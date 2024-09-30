package menuaccess

import (
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	"TamanSiswa/helper"
	"TamanSiswa/model"
	"TamanSiswa/services/menuAccess/request"
	"golang.org/x/exp/slices"
)

func GetAllRecord(identity helper.TokenPayload) (*[]model.MenuAccess, *[]model.Position, error) {
	var menuAccessData []model.MenuAccess
	var positionData []model.Position
	var positionIds []string

	result := model.DB.Model(&model.Position{}).Where("company_id = ?", identity.CompanyId).Order("id").Find(&positionData)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	for _, item := range positionData {
		positionIds = append(positionIds, item.Id)
	}

	result = model.DB.Table("menu_access").Where("position_id IN ?", positionIds).Order("position_id").Find(&menuAccessData)
	if result.Error != nil {
		return nil, nil, result.Error
	}

	return &menuAccessData, &positionData, nil
}

func GetWithId(identity helper.TokenPayload, id string) (*model.MenuAccess, error) {
	var resultData model.MenuAccess

	result := model.DB.Table("menu_access").Where("id = ? AND company_id = ?", id, identity.CompanyId).First(&resultData)

	if result.Error != nil {
		return nil, result.Error
	}

	return &resultData, nil
}

func DeleteOne(identity helper.TokenPayload, id string) error {
	result := model.DB.Delete(&model.MenuAccess{}, "id = ? AND company_id = ?", id, identity.CompanyId)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func UpdateOne(identity helper.TokenPayload, req request.RequestUpdate, id string) error {
	var data model.MenuAccess

	data.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	access, err := json.Marshal(req.Access)
	if err != nil {
		return err
	}
	data.Access = *helper.ValidateSqlString(string(access))

	err = checkBaseAuthority(req.Access, data.CompanyId.String)
	if err != nil {
		return err
	}

	result := model.DB.Table("menu_access").Where("id = ? AND company_id = ?", id, identity.CompanyId).Updates(&data)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func CreateOne(identity helper.TokenPayload, req request.Request) error {
	var data model.MenuAccess

	data.PositionId = *helper.ValidateSqlString(req.PositionId)
	data.CompanyId = sql.NullString{String: identity.CompanyId, Valid: true}
	data.CreatedAt = sql.NullTime{Time: time.Now(), Valid: true}

	access, err := json.Marshal(req.Access)
	if err != nil {
		return err
	}
	data.Access = *helper.ValidateSqlString(string(access))

	err = checkBaseAuthority(req.Access, data.CompanyId.String)
	if err != nil {
		return err
	}

	result := model.DB.Table("menu_access").Create(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func checkBaseAuthority(access []request.AccessPermission, companyId string) error {
	var subscriptionData model.Subscription
	var availablePackageData model.AvailablePackage
	var availableMenu []request.AccessPermission

	result := model.DB.Model(&model.Subscription{}).Where("company_id = ?", companyId).First(&subscriptionData)
	if result.Error != nil {
		return result.Error
	}

	result = model.DB.Model(&model.AvailablePackage{}).Where("id = ?", subscriptionData.SubscribedPackage.String).First(&availablePackageData)
	if result.Error != nil {
		return result.Error
	}

	err := json.Unmarshal([]byte(availablePackageData.AvailableMenu.String), &availableMenu)
	if result.Error != nil {
		return err
	}

	for _, item := range access {
		for index2, item2 := range availableMenu {
			if item.MainMenu == item2.MainMenu {
				for _, item3 := range item.SubMenu {
					if !slices.Contains(item2.SubMenu, item3) {
						return errors.New("requested access is not included in base authority")
					}
				}
				break
			}

			if index2 == len(availableMenu)-1 {
				return errors.New("requested access is not included in base authority")
			}
		}
	}

	return nil
}
