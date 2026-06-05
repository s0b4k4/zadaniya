package grpc_handler

import (
	"context"

	documentv1 "github.com/course/tasks/api/document/v1"
	// "github.com/course/tasks/internal/repository/db"
)

type DocumentHandler struct {
	documentv1.UnimplementedDocumentServiceServer
	// repo *db.Queries
}

func (h *DocumentHandler) CreateUser(ctx context.Context, req *documentv1.CreateUserRequest) (*documentv1.UserResponse, error) {
	// u, err := h.repo.CreateUser(ctx, req.Name)
	// if err != nil { return nil, err }
	return &documentv1.UserResponse{
		Id:   1, // u.ID
		Name: req.Name,
	}, nil
}

func (h *DocumentHandler) AddDocumentToUser(ctx context.Context, req *documentv1.AddDocumentRequest) (*documentv1.DocumentResponse, error) {
	// doc, err := h.repo.AddDocument(ctx, db.AddDocumentParams{ ... })
	return &documentv1.DocumentResponse{
		Id:      1,
		UserId:  req.UserId,
		Title:   req.Title,
		Content: req.Content,
	}, nil
}

func (h *DocumentHandler) GetUserDocuments(ctx context.Context, req *documentv1.GetUserDocumentsRequest) (*documentv1.GetUserDocumentsResponse, error) {
	// docs, err := h.repo.GetUserDocuments(ctx, req.UserId)
	return &documentv1.GetUserDocumentsResponse{
		Documents: []*documentv1.DocumentResponse{}, // Convert docs to protobuf
	}, nil
}
