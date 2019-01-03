package model

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Article struct {
	Model

	Tag   Tag `form:"-" json:"tag"` //json:"tag"`
	TagID int `form:"tag_id" json:"tag_id" gorm:"index" bind:"required"`

	Title      string `form:"title" json:"title" bind:"required"`
	Desc       string `form:"desc" json:"desc" bind:"required"`
	Content    string `form:"content" json:"content" bind:"required"`
	CreateBy   string `form:"create_by" json:"create_by" bind:"required"`
	ModifiedBy string `form:"modified_by" json:"modified_by" bind:"required"`
	State      int    `form:"state" json:"state" bind:"required"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if article.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetArticleTotal(maps interface{}) (count int, err error) {
	if err := db.Model(&Article{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

/*
能够达到关联，首先是gorm本身做了大量的约定俗成

Article有一个结构体成员是TagID，就是外键。gorm会通过类名+ID的方式去找到这两个类之间的关联关系
Article有一个结构体成员是Tag，就是我们嵌套在Article里的Tag结构体，我们可以通过Related进行关联查询
*/
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []*Article, err error) {
	err = db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return articles, nil
}

func GetArticle(id int) (*Article, error) {
	var article Article
	log.Println(article.Tag)
	db.Where("id = ?", id).First(&article)
	err := db.Model(&article).Related(&article.Tag, "TagID").Error
	log.Println("####", article)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

func AddArticle(article Article) error {
	/*
		article := Article{
			TagID:    data["tag_id"].(int),
			Title:    data["title"].(string),
			Desc:     data["desc"].(string),
			Content:  data["content"].(string),
			CreateBy: data["create_by"].(string),
			State:    data["state"].(int),
		}
	*/

	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func EditArticle(article Article) error {
	if err := db.Model(&Article{}).Where("id = ?", article.ID).Update(article).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) error {
	if err := db.Where("id = ?", id).Delete(&Article{}).Error; err != nil {
		return err
	}
	return nil
}
