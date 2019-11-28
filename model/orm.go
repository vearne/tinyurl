package models

// 域名的解析规则
type TinyURL struct {
	//gorm.Model
	ID    uint64 `gorm:"column:id" json:"id"`
	URL   string `gorm:"column:url" json:"url"`
}

func (TinyURL) TableName() string {
	return "tinyurl"
}
