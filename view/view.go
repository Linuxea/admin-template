package view

import "html/template"

type FieldType int

const (
	Number FieldType = iota + 1
	String
	Obj
	List
)

type Column struct {
	Prop        template.JS // properties
	Label       string
	Width       uint
	Type        FieldType
	Candidate   map[interface{}]interface{}
	ShowInTable bool
	Queryable   bool
	Addable     bool
	ShowImage   bool
	Required    bool
	Placeholder string
	Color       string
	AutoColor   bool
}

type Record struct {
	Cols       []Column
	Editable   bool
	Deleteable bool
}

type View struct {
	Record           Record
	Height           uint
	QueryPage        uint
	QueryPageSize    uint
	BucketName       string
	AliOssPath       string
	AddMethodName    string
	EditMethodName   string
	DeleteMethodName string
	ListMethodName   string
	ControllerName   string
	ModelName        string
}
