package db_gorm

import (
	"errors"
	dbstruct "./db/struct"

	"gorm.io/gorm"
)

func (x *DbGorm) GetBlackDeviceByYostarUid(deviceId string) *dbstruct.BlackDevice {
	var data *dbstruct.BlackDevice
	err := x.sql.Model(&dbstruct.BlackDevice{}).Where("device_id = ?", deviceId).Select("*").First(&data).Error
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil
	}
	return data
}
