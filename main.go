package main

import (
	"html/template"
	"os"
	"strings"

	"template.com/linuxea/view"
)

func main() {

	// create table person (
	// 	id bigint(20) primary key auto_increment,
	// 	name varchar(100) ,
	//  pic varchar(100)
	// )

	cols := []view.Column{

		{
			Prop:        "ID",
			Label:       "id",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Queryable:   false,
			Addable:     false,
			ShowImage:   false,
			Color:       "blue",
		},
		{
			Prop:        "type",
			Label:       "类型",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Queryable:   true,
			Addable:     true,
			ShowImage:   false,
			Required:    true,
			Candidate: map[interface{}]interface{}{
				1: "铭牌",
				2: "道具",
				3: "菜币",
			},
		},
		{
			Prop:        "type_id",
			Label:       "类型ID",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Addable:     true,
			ShowImage:   false,
		},
		{
			Prop:        "name",
			Label:       "名称",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Queryable:   true,
			Addable:     true,
			Required:    true,
		},
		{
			Prop:        "color",
			Label:       "颜色",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Queryable:   false,
			Addable:     true,
			Candidate: map[interface{}]interface{}{
				"BLUE":   "蓝色",
				"RED":    "红色",
				"PURPLE": "紫色",
				"YELLOW": "黄色",
			},
			Required:  true,
			AutoColor: true,
		},
		{
			Prop:        "unit",
			Label:       "单位",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Addable:     true,
		},
		{
			Prop:        "pic",
			Label:       "图片",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Queryable:   false,
			Addable:     true,
			ShowImage:   true,
			Placeholder: "请上传图片",
		},
		{
			Prop:        "CreatedAt",
			Label:       "创建时间",
			Width:       150,
			Type:        view.String,
			ShowInTable: true,
			Queryable:   false,
		},
		{
			Prop:        "UpdatedAt",
			Label:       "更新时间",
			Width:       150,
			Type:        view.String,
			ShowInTable: true,
			Queryable:   false,
		},
	}

	record := view.Record{Cols: cols, Deleteable: true}
	view := view.View{
		Record:         record,
		Height:         650,
		QueryPage:      1,
		QueryPageSize:  30,
		ControllerName: "ActivityTwentyOneYearEndPool",
	}

	t := template.Must(template.New("admin.txt").Funcs(template.FuncMap{
		"Cap": func(in template.JS) template.JS {
			return template.JS(strings.ToUpper(string(in[:1])) + strings.ToLower(string(in[1:])))
		},
	}).ParseFiles("view/admin.txt"))
	if err := t.Execute(os.Stdout, view); err != nil {
		panic(err)
	}

}
