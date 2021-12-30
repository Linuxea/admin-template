package main

import (
	"html/template"
	"os"
	"strings"

	"template.com/linuxea/view"
)

func main() {

	yesNo := map[interface{}]interface{}{
		"1": "是",
		"2": "否",
	}

	colType := map[interface{}]interface{}{
		1: "number",
		2: "string",
		3: "obj",
		4: "List",
	}

	cols := []view.Column{
		{
			Prop:        "ID",
			Label:       "id",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Color:       "blue",
		},

		{
			Prop:        "tb_name",
			Label:       "表名",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Color:       "blue",
			Queryable:   true,
			Addable:     true,
		},
		{
			Prop:        "col_name",
			Label:       "列名",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
		},
		{
			Prop:        "label",
			Label:       "标签名",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
		},
		{
			Prop:        "type",
			Label:       "类型",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
			Candidate:   colType,
		},
		{
			Prop:        "show_in_table",
			Label:       "列表展示",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
			Candidate:   yesNo,
		},
		{
			Prop:        "query_able",
			Label:       "条件查询",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
			Candidate:   yesNo,
		},
		{
			Prop:        "add_able",
			Label:       "支持添加/编辑",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
			Candidate:   yesNo,
		},
		{
			Prop:        "show_image",
			Label:       "图片展示",
			Width:       100,
			Type:        view.Number,
			ShowInTable: true,
			Color:       "blue",
			Addable:     true,
			Candidate:   yesNo,
		},
		{
			Prop:        "CreatedAt",
			Label:       "创建时间",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Color:       "blue",
		},
		{
			Prop:        "UpdatedAt",
			Label:       "更新时间",
			Width:       100,
			Type:        view.String,
			ShowInTable: true,
			Color:       "blue",
		},
	}

	record := view.Record{Cols: cols, Deleteable: true, Editable: true}
	View := view.View{
		Record:         record,
		Height:         650,
		QueryPage:      1,
		QueryPageSize:  30,
		ControllerName: cap(sneak2Camcel("tbName")),
		ModelName:      cap(sneak2Camcel("tbName")) + "Model",
	}

	fm := template.FuncMap{
		"Cap": func(in template.JS) template.JS {
			return template.JS(cap(string(in)))
		},
		"Sneak2Camcel": func(in template.JS) template.JS {
			return template.JS(sneak2Camcel(string(in)))
		},
	}

	t := template.Must(template.New("admin.txt").Funcs(fm).ParseFiles("view/admin.txt"))
	if err := t.Execute(os.Stdout, View); err != nil {
		panic(err)
	}

	t = template.Must(template.New("ctrl.txt").Funcs(fm).ParseFiles("view/ctrl.txt"))
	if err := t.Execute(os.Stdout, View); err != nil {
		panic(err)
	}

}

func cap(in string) string {
	return strings.ToUpper(string(in[:1])) + strings.ToLower(string(in[1:]))
}

func sneak2Camcel(in string) string {
	first := true
	up := false
	newIn := ""
	for _, i := range in {

		if first || up {
			newIn += strings.ToUpper(string(i))
			first = false
			up = false
		} else if string(i) == "_" {

		} else {
			newIn += string(i)
		}

		if string(i) == "_" {
			up = true
		}

	}

	return newIn
}
