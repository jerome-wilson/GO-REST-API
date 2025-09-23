package services

import "github.com/jerome-wilson/GO-REST-API/repository"

type BookService struct {
	Repo *repository.BookRepository
}
