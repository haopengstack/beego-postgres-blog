package common



type PostCreate struct {
	Title 			string 	`form:"title" valid:"Required;MaxSize(33)"`
	Abstract 		string	`form:"abstract" valid:"Required;MaxSize(100)"`
	Category 		int64 	`form:"category" valid:"Required"`
	Tag 			string 	`form:"tag" valid:"Required"`
	Content 		string 	`form:"editormd-html-code" valid:"Required"`
	BodyOriginal 	string 	`form:"content" valid:"Required"`
}

type CateRequest struct {
	ParentId		string 	`form:"parentId" valid:"Required"`
	Name 			string 	`form:"name" valid:"Required;MaxSize(30)"`
	DisplayName 	string 	`form:"displayName" valid:"Required;MaxSize(30)"`
	Description 	string 	`form:"description" valid:"MaxSize(150)"`
}

type LinkCreate struct {
	Name 	string 	`form:"name" valid:"Required;MaxSize(23)"`
	Link 	string 	`form:"link" valid:"Required;MaxSize(100)"`
	Sort 	int64   `form:"ordering" valid:"Required"`
}


type SystemUpdate struct {
	Title			string 		`form:"title" valid:"Required;MaxSize(23)"`
	STitle			string 		`form:"s_title" `
	Description 	string 		`form:"description" `
	SeoKey			string 		`form:"seo_key" `
	SeoDes			string 		`form:"seo_des" `
	RecordNumber	string 		`form:"record_number" `
}

type TagRequest struct {
	Name 	string `form:"name" valid:"Required;MaxSize(30)"`
}