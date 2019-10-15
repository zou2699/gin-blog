package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Tag struct {
	Model
	Name       string `form:"name" json:"name" bind:"required" valid:required;maxsize(100)`
	CreatedBy  string `form:"created_by" json:"created_by" bind:"required"`
	ModifiedBy string `form:"modified_by" json:"modified_by" bind:"required"`
	State      int    `form:"state" json:"state" bind:"required"`
}

// 不存在就创建Tag
/*
func init() {
	db.AutoMigrate(&Tag{})
}
*/

// 获取tags
func GetTags(pageNum int, PageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(PageSize).Find(&tags)
	return
}

// 获取tag总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 判断是否存在tag
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}
	return false, nil
}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// 添加tag
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// 编辑tag
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id=?", id).Update(data).Error; err != nil {
		log.Println("###", err)
		return err
	}

	return nil
}

// 删除tag
func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}
	return nil
}
