package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"quiz3/forms"
	"quiz3/middlewares"
	"quiz3/repositories"
	"quiz3/utils"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type BookAPI interface {
	GetAllBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	GetBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetBookByCategoryID(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	InsertBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
	UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type bookAPI struct {
	bookRepo repositories.BookRepo
}

func NewBookAPI(bookRepo repositories.BookRepo) *bookAPI {
	return &bookAPI{bookRepo}
}

func (c *bookAPI) GetAllBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	queryParams := r.URL.Query()
	filter := forms.BookFilter{
		Title:       queryParams.Get("title"),
		MinYear:     queryParams.Get("minYear"),
		MaxYear:     queryParams.Get("maxYear"),
		MinPage:     queryParams.Get("minPage"),
		MaxPage:     queryParams.Get("maxPage"),
		SortByTitle: queryParams.Get("sortByTitle"),
	}
	
	books, err := c.bookRepo.GetAll(ctx, filter)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Success get all books",
		Data: books,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bookAPI) InsertBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if res := middlewares.BasicAuth(r); res != "" {
		utils.ErrorJSON(w, errors.New(res), "UNAUTHORIZED",http.StatusUnauthorized)
		return
	}
	
	if r.Header.Get("Content-Type") != "application/json" {
        utils.ErrorJSON(w, errors.New("Gunakan content type application / json"), "BAD REQUEST",http.StatusBadRequest)
		return
    }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var newBook forms.InsertBook
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
	}
	isDuplicate, err := c.bookRepo.IsDuplicate(ctx, newBook.Title)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	var validationErrors []string

	if isDuplicate {
		validationErrors = append(validationErrors, "Book with this title already exists")
	}

	if !utils.IsValidURL(newBook.ImageURL) {
		validationErrors = append(validationErrors, "Invalid image URL format")
	}

	if newBook.ReleaseYear < 1980 || newBook.ReleaseYear > 2021 {
		validationErrors = append(validationErrors, "Invalid release year")
	}

	if len(validationErrors) > 0 {
		errorMessage := "Validation failed: " + strings.Join(validationErrors, ", ")
		utils.ErrorJSON(w, errors.New(errorMessage), "BAD REQUEST", http.StatusBadRequest)
		return
	}

	newBook.Thickness = utils.CheckThickness(newBook.TotalPage)

	if err := c.bookRepo.Insert(ctx, newBook); err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success insert new category %v", newBook.Title),
		Data: newBook,
	}
	
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (c *bookAPI) GetBookById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idBook = utils.ParseInt(ps.ByName("id"))
	
	book, err := c.bookRepo.GetById(ctx, idBook)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success get book with id %v", idBook),
		Data: book,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bookAPI) GetBookByCategoryID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategory = utils.ParseInt(ps.ByName("id"))

	queryParams := r.URL.Query()
	filter := forms.BookFilter{
		Title:       queryParams.Get("title"),
		MinYear:     queryParams.Get("minYear"),
		MaxYear:     queryParams.Get("maxYear"),
		MinPage:     queryParams.Get("minPage"),
		MaxPage:     queryParams.Get("maxPage"),
		SortByTitle: queryParams.Get("sortByTitle"),
	}
	
	book, err := c.bookRepo.GetByCategoryID(ctx, idCategory, filter)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success get book with category id %v", idCategory),
		Data: book,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (c *bookAPI) UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if res := middlewares.BasicAuth(r); res != "" {
		utils.ErrorJSON(w, errors.New(res), "UNAUTHORIZED",http.StatusUnauthorized)
		return
	}
	
	if r.Header.Get("Content-Type") != "application/json" {
        utils.ErrorJSON(w, errors.New("Gunakan content type application / json"), "BAD REQUEST",http.StatusBadRequest)
		return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var newBook forms.UpdateBook

	var idBook = utils.ParseInt(ps.ByName("id"))

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
	}

	isDuplicate, err := c.bookRepo.IsDuplicate(ctx, newBook.Title)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	isExist, err := c.bookRepo.IsExist(ctx, idBook)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if isExist {
		utils.ErrorJSON(w, errors.New("Data category with this ID doesn't exist"), "CONFLICT", http.StatusConflict)
		return
	}

	var validationErrors []string

	if isDuplicate {
		validationErrors = append(validationErrors, "Book with this title already exists")
	}

	if !utils.IsValidURL(newBook.ImageURL) {
		validationErrors = append(validationErrors, "Invalid image URL format")
	}

	if newBook.ReleaseYear < 1980 || newBook.ReleaseYear > 2021 {
		validationErrors = append(validationErrors, "Invalid release year")
	}

	if len(validationErrors) > 0 {
		errorMessage := "Validation failed: " + strings.Join(validationErrors, ", ")
		utils.ErrorJSON(w, errors.New(errorMessage), "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	newBook.Thickness = utils.CheckThickness(newBook.TotalPage)

	if err := c.bookRepo.UpdateByID(ctx, newBook, idBook); err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

    payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success update book %v", newBook.Title),
		Data: newBook,
	}  
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (c *bookAPI) DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if res := middlewares.BasicAuth(r); res != "" {
		utils.ErrorJSON(w, errors.New(res), "UNAUTHORIZED",http.StatusUnauthorized)
		return
	}
	
	ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idBook = utils.ParseInt(ps.ByName("id"))

    if err := c.bookRepo.DeleteByID(ctx, idBook); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success delete book %v", idBook),
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}