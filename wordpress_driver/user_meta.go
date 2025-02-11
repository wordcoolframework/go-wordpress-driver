package wordpressdriver

import "gorm.io/gorm"

type UserMeta struct {
	ID        int    `gorm:"column:umeta_id"`
	UserID    int    `gorm:"column:user_id"`
	MetaKey   string `gorm:"column:meta_key"`
	MetaValue string `gorm:"column:meta_value"`
}

func (w *WpFacade) UserMeta() *UserMeta {
	return &UserMeta{}
}

func (u *UserMeta) GetAllUserMeta(prefixTable string) ([]UserMeta, error) {
	var userMeta []UserMeta
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Find(&userMeta)
	return userMeta, result.Error
}

func (u *UserMeta) GetUserMetaByID(id int, prefixTable string) (*UserMeta, error) {
	var userMeta UserMeta
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Where("umeta_id = ?", id).First(&userMeta)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &userMeta, result.Error
}

func (u *UserMeta) GetUserMetaByUserID(userID int, prefixTable string) ([]UserMeta, error) {
	var userMeta []UserMeta
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Where("user_id = ?", userID).Find(&userMeta)
	return userMeta, result.Error
}

func (u *UserMeta) GetUserMetaByKey(userID int, key, prefixTable string) (*UserMeta, error) {
	var userMeta UserMeta
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Where("user_id = ? AND meta_key = ?", userID, key).First(&userMeta)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &userMeta, result.Error
}

func (u *UserMeta) SearchUserMeta(keyword, prefixTable string) ([]UserMeta, error) {
	var userMeta []UserMeta
	pTable := prefixTable + "_usermeta"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("meta_value LIKE ?", query).Find(&userMeta)
	return userMeta, result.Error
}

func (u *UserMeta) AddUserMeta(userID int, key, value, prefixTable string) error {
	pTable := prefixTable + "_usermeta"
	userMeta := UserMeta{UserID: userID, MetaKey: key, MetaValue: value}

	result := DB.Table(pTable).Create(&userMeta)
	return result.Error
}

func (u *UserMeta) UpdateUserMeta(userID int, key, value, prefixTable string) error {
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Where("user_id = ? AND meta_key = ?", userID, key).Update("meta_value", value)
	return result.Error
}

func (u *UserMeta) DeleteUserMeta(userID int, key, prefixTable string) error {
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Where("user_id = ? AND meta_key = ?", userID, key).Delete(&UserMeta{})
	return result.Error
}

func (u *UserMeta) GetUsersByMetaKeyValue(key, value, prefixTable string) ([]int, error) {
	var userIDs []int
	pTable := prefixTable + "_usermeta"

	result := DB.Table(pTable).Select("user_id").Where("meta_key = ? AND meta_value = ?", key, value).Find(&userIDs)
	return userIDs, result.Error
}
