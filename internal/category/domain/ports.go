package domain

type BookTypePost struct {
	Title       string
	Description string
	ReleaseDate string
	Author      string
	Value       float32
}

type ServicesMongoDBRepository interface {
	Add(BookTypePost) error
}
