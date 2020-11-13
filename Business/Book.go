package Business

import (
	"Book/Database"
	"Book/Database/DAL"
	"Book/Database/Models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type bookBusiness struct {
	BookRepository Models.BookRepository
}

func (b bookBusiness) Create(book Models.Book) (*Models.Book, error) {
	result, err := b.BookRepository.Create(book)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return result, err
}

func (b bookBusiness) GetById(id string) (*Models.Book, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}

	result, err := b.BookRepository.GetById(ID)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return result, err
}

func (b bookBusiness) Filter(text string, author string, status bool, category []string, page int, pageSize int) ([]*Models.Book,int64 , error) {
	results, total, err := b.BookRepository.Filter(text, author, status, nil, page, pageSize)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, 0, err
	}
	return results, total, err

}

func (b bookBusiness) Update(id string, book Models.Book) (*Models.Book, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	result, err := b.BookRepository.Update(ID, book)
	if err != nil {
		Database.ErrLog.Print(err)
		return nil, err
	}
	return result, err

}

func (b bookBusiness) Active(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Database.ErrLog.Print(err)
		return  err
	}
	err = b.BookRepository.UpdateStatus(true, ID)
	if err != nil {
		return err
	}
	return err
}

func (b bookBusiness) DeActive(id string) error {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		Database.ErrLog.Print(err)
		return  err
	}
	err = b.BookRepository.UpdateStatus(false, ID)
	if err != nil {
		return err
	}
	return err
}

type BookBusiness interface {
	Create(book Models.Book) (*Models.Book, error)
	GetById(id string) (*Models.Book, error)
	Filter(id string, author string, status bool, category []string, page int, pageSize int) ([]*Models.Book, int64,  error)
	Update(id string, book Models.Book) (*Models.Book, error)
	Active(id string) error
	DeActive(id string)error
}

func NewBookBusiness() BookBusiness {
	bookRepo := DAL.NewBookRepository(Database.Db)
	return &bookBusiness{BookRepository: bookRepo}
}

