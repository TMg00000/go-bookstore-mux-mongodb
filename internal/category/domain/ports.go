package domain

type Book struct {
	Title       string
	Description string
	ReleaseDate string
	Author      string
	Value       float32
}

type MongoDBRepository interface {
	Search(title string) (Book, error) // busca se o livro ja existe no banco de dados para evitar duplicação.
	Add(Book) error                    // Insere um novo livro no banco de dados.
}
