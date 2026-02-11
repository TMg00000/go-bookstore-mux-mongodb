package domain

type MongoDBRepository interface {
	Add(Book) error // Insere um novo livro no banco de dados.
}
