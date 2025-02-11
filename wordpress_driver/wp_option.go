package wordpressdriver

import "gorm.io/gorm"

type Option struct {
	ID       int    `gorm:"column:option_id"`
	Name     string `gorm:"column:option_name"`
	Value    string `gorm:"column:option_value"`
	Autoload string `gorm:"column:autoload"`
}

func (w *WpFacade) Option() *Option {
	return &Option{}
}

func (o *Option) GetAllOptions(prefixTable string) ([]Option, error) {
	var options []Option
	pTable := prefixTable + "_options"

	result := DB.Table(pTable).Find(&options)
	return options, result.Error
}

func (o *Option) GetOptionByID(id int, prefixTable string) (*Option, error) {
	var option Option
	pTable := prefixTable + "_options"

	result := DB.Table(pTable).Where("option_id = ?", id).First(&option)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &option, result.Error
}

func (o *Option) GetOptionByName(name, prefixTable string) (*Option, error) {
	var option Option
	pTable := prefixTable + "_options"

	result := DB.Table(pTable).Where("option_name = ?", name).First(&option)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &option, result.Error
}

func (o *Option) SearchOptions(keyword, prefixTable string) ([]Option, error) {
	var options []Option
	pTable := prefixTable + "_options"
	query := "%" + keyword + "%"

	result := DB.Table(pTable).Where("option_name LIKE ? OR option_value LIKE ?", query, query).Find(&options)
	return options, result.Error
}

func (o *Option) AddOption(name, value, autoload, prefixTable string) error {
	pTable := prefixTable + "_options"
	option := Option{Name: name, Value: value, Autoload: autoload}

	result := DB.Table(pTable).Create(&option)
	return result.Error
}

func (o *Option) UpdateOption(name, value, prefixTable string) error {
	pTable := prefixTable + "_options"

	result := DB.Table(pTable).Where("option_name = ?", name).Update("option_value", value)
	return result.Error
}

func (o *Option) DeleteOption(name, prefixTable string) error {
	pTable := prefixTable + "_options"

	result := DB.Table(pTable).Where("option_name = ?", name).Delete(&Option{})
	return result.Error
}

func (o *Option) GetAutoloadOptions(prefixTable string) ([]Option, error) {
	var options []Option
	pTable := prefixTable + "_options"

	result := DB.Table(pTable).Where("autoload = ?", "yes").Find(&options)
	return options, result.Error
}
