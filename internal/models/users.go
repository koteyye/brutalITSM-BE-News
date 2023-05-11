package models

type UserSingle struct {
	Id          string   `json:"id"`
	Login       string   `json:"login"`
	LastName    string   `json:"lastName"`
	FirstName   string   `json:"firstName"`
	SurName     string   `json:"surName"`
	Job         string   `json:"job"`
	Org         string   `json:"org"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type UserList struct {
	Id         string  `json:"id"`
	LastName   string  `json:"lastName"`
	FirstName  string  `json:"firstName"`
	SurName    string  `json:"surName"`
	MimeType   *string `json:"mimeType"`
	BucketName *string `json:"bucketName"`
	FileName   *string `json:"fileName"`
}
