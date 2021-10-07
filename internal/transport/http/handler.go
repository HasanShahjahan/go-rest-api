package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HasanShahjahan/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//Response - stores response from API
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := sendOkResponse(w, Response{Message: "I am alive"}); err != nil {
			panic(err)
		}
	})
}

//GetComment - retrieve a comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error Retrieving Comment By ID", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

//GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all comments", err)
		return
	}

	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

//PostComment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}
	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

//UpdateComment - updates a new comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse unit from ID", err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to update new comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

//DeleteComment - deletes comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	commentID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse unit from ID", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by ID", err)
		return
	}

	if err := sendOkResponse(w, Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
