package controllers

type RegisterUser struct {
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	Password    string `json:"password"`
}

type TokenRequest struct {
	PhoneNumber int    `json:"phone_number"`
	Password    string `json:"password"`
}

type AddBlogRequest struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
