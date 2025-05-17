package book

type Book struct {
	Author   string
	BookName string
	Pages    int
}

func CreateBook(author, bookName string, pages int) Book {
	return Book{
		Author:   author,
		BookName: bookName,
		Pages:    pages,
	}

}

func GetBook(b *Book) Book {
	return *b
}
func (b *Book) UpdateTitle(newTitle string) {
	b.BookName = newTitle
}
