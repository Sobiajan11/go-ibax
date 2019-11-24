/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
package model

type VDEDestHashTime struct {
	ID         int64 `gorm:"primary_key; not null" json:"id"`
	UpdateTime int64 `gorm:"not null" json:"update_time"`
	CreateTime int64 `gorm:"not null" json:"create_time"`
}

func (VDEDestHashTime) TableName() string {
	return "vde_dest_hash_time"
	return DBConn.Delete(m).Error
}

func (m *VDEDestHashTime) Get() (*VDEDestHashTime, error) {
	err := DBConn.First(&m).Error
	return m, err
}

func (m *VDEDestHashTime) GetAll() ([]VDEDestHashTime, error) {
	var result []VDEDestHashTime
	err := DBConn.Find(&result).Error
	return result, err
}
func (m *VDEDestHashTime) GetOneByID() (*VDEDestHashTime, error) {
	err := DBConn.Where("id=?", m.ID).First(&m).Error
	return m, err
}