package service

import (
	"context"
	"log"
	"time"

	"github.com/sony/gobreaker"
	// articlev1 "github.com/course/tasks/api/article/v1"
)

type SearchClient struct {
	cb *gobreaker.CircuitBreaker
	// searchServiceClient articlev1.SearchServiceClient
}

func NewSearchClient() *SearchClient {
	var st gobreaker.Settings
	st.Name = "SearchService"
	st.MaxRequests = 3
	st.Interval = 5 * time.Second
	st.Timeout = 10 * time.Second
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}
	st.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		log.Printf("Circuit Breaker '%s' changed state from %v to %v\n", name, from, to)
	}

	return &SearchClient{
		cb: gobreaker.NewCircuitBreaker(st),
	}
}

// Emulate calling remote gRPC SearchService
func (s *SearchClient) IndexArticle(ctx context.Context, id int, title string, content string) error {
	_, err := s.cb.Execute(func() (interface{}, error) {
		log.Printf("Calling remote SearchService to index article %d...\n", id)
		// res, err := s.searchServiceClient.IndexArticle(ctx, &articlev1.IndexArticleRequest{...})
		// return res, err
		
		// Simulate network call
		time.Sleep(100 * time.Millisecond)
		return nil, nil // Return error here to test circuit breaker
	})

	return err
}
