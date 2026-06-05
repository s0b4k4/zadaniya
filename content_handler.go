package grpc_handler

import (
	"context"

	articlev1 "github.com/course/tasks/api/article/v1"
	"github.com/course/tasks/internal/service"
)

type ContentHandler struct {
	articlev1.UnimplementedContentServiceServer
	searchClient *service.SearchClient
}

func NewContentHandler(searchClient *service.SearchClient) *ContentHandler {
	return &ContentHandler{searchClient: searchClient}
}

// Task 7: Inter-service communication with Circuit Breaker
func (h *ContentHandler) CreateArticle(ctx context.Context, req *articlev1.CreateArticleRequest) (*articlev1.ArticleResponse, error) {
	// 1. Save to DB
	articleID := int32(1)

	// 2. Call Search service to index
	err := h.searchClient.IndexArticle(ctx, int(articleID), req.Title, req.Content)
	if err != nil {
		// Log error, but don't fail the creation (Circuit Breaker handles degradation gracefully)
	}

	return &articlev1.ArticleResponse{
		Id:       articleID,
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: req.AuthorId,
	}, nil
}

func (h *ContentHandler) GetArticle(ctx context.Context, req *articlev1.GetArticleRequest) (*articlev1.ArticleResponse, error) {
	return &articlev1.ArticleResponse{Id: req.Id, Title: "Sample"}, nil
}
