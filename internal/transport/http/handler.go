package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/singfield/rest-api-golang/internal/comment"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// return pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comments/{id}", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comments/{id}", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'am alive !")
	})
}

// retrieve a comment by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "unable to parse UINT from ID")
	}

	comment, err := h.Service.GetComment(uint(i))

	if err != nil {
		fmt.Fprintf(w, "Error Retrieving Comment by ID")
	}

	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "Failed to retrieve all comments")
	}
	fmt.Fprintf(w, "%+v", comment)

}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "Failed to post new comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// i, err := strconv.ParseUint(id, 10, 64)

	commentID, err := strconv.ParseUint(id,10, 64)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse uint from ID")
	}
	fmt.Fprintf(w, "%+v", commentID)

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprintf(w, "Failed to Delete comment by comment ID")
	}
	fmt.Fprintf(w, "Successfully deleted commment")
}
