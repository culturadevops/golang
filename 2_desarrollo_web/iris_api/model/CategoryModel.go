package model

import (
	"cmsx/libs"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

var ListTree []Category

type Category struct {
	gorm.Model
	Name     string `gorm:"not null;VARCHAR(100);"validate:"required"`
	ParentId int    `gorm:"default:'0';not null;"validate:"required,number,gte=1"`
	Sort     int    `gorm:"default:'0';not null;"validate:"number,min=0"`
	Level    int    `gorm:"-"`
}

func (this *Category) List() []Category {
	var data = []Category{}
	db := libs.DB

	err := db.Order("sort desc").Find(&data).Error
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}

func (this *Category) CategoryInfo(id uint) (Category, error) {
	var category Category
	db := libs.DB

	if db.Where("id = ?", id).First(&category).RecordNotFound() {
		return Category{}, errors.New("分类未找到")
	}
	return category, nil
}

func (this *Category) CategoryMoreInfo(ids string) ([]Category, error) {
	var data = []Category{}
	db := libs.DB

	if db.Where("id in (?)", strings.Split(ids, ",")).Find(&data).RecordNotFound() {
		return []Category{}, errors.New("分类未找到")
	}
	return data, nil
}

func (this *Category) CategoryAdd(postValues map[string][]string) error {
	var category Category
	db := libs.DB

	if err := libs.FormDecode(&category, postValues); err != nil {
		return err
	}
	if err := libs.Validate(category); err != nil {
		return err
	}
	if !db.Where("name = ? ", category.Name).First(&Category{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if err := db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

func (this *Category) CategoryUpdate(postValues map[string][]string) error {
	var category Category
	db := libs.DB

	if err := libs.FormDecode(&category, postValues); err != nil {
		return err
	}
	if err := libs.Validate(category); err != nil {
		return err
	}
	if !db.Where("name = ? and id != ?", category.Name, category.ID).Find(&Category{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if db.Where("id = ? ", category.ID).Find(&Category{}).RecordNotFound() {
		return errors.New("未查询到分类id")
	}
	if err := db.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func (this *Category) CategoryDel(id uint) error {
	var category Category
	db := libs.DB

	if !db.Where("parent_id = ?", id).Find(&category).RecordNotFound() {
		return errors.New("该分类下存在子级分类，请先删除子级分类")
	}
	if err := db.Where("id = ?", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (this *Category) GetTree(list []Category, pid int, level int) []Category {
	for _, v := range list {
		if v.ParentId == pid {
			v.Level = level
			v.Name = strings.Repeat("————", v.Level) + v.Name
			ListTree = append(ListTree, v)
			/*if len(list) == 0 {
				list = []Category{}
			} else {
				list = append(list[:index], list[index+1:]...)
			}*/
			this.GetTree(list, int(v.Model.ID), v.Level+1)
		}
	}
	return ListTree
}
