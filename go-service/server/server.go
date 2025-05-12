package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pb "grpcservice/pb"
)

const (
	djangoAPIBaseURL = "http://localhost:8000/api/v1"
)

// now this server implements the interface defined in our proto file
type Server struct {
	pb.UnimplementedModuleServiceServer
	httpClient *http.Client
}

// creates a new server..used when we star
func NewServer() *Server {
	return &Server{
		httpClient: &http.Client{},
	}
}

//makes HHTP request to API and gets JSON response and this convert it back to gRPC to handle gRPC re
func (s *Server) ListModules(ctx context.Context, _ *pb.Empty) (*pb.ModuleList, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/modules/", djangoAPIBaseURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned error: %s - %s", resp.Status, string(body))
	}

	var modules []struct {
		ID         int32           `json:"id"`
		Subject    string          `json:"subject"`
		Grade      string          `json:"grade"`
		Chapter    string          `json:"chapter"`
		ContentJSON json.RawMessage `json:"content"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&modules); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	result := &pb.ModuleList{
		Modules: make([]*pb.Module, len(modules)),
	}

	for i, m := range modules {
		result.Modules[i] = &pb.Module{
			Id:          m.ID,
			Subject:     m.Subject,
			Grade:       m.Grade,
			Chapter:     m.Chapter,
			ContentJson: string(m.ContentJSON),
		}
	}

	return result, nil
}

// make request to API and convert
func (s *Server) GetModule(ctx context.Context, req *pb.ModuleRequest) (*pb.Module, error) {
	url := fmt.Sprintf("%s/modules/%d/", djangoAPIBaseURL, req.Id)
	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := s.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("module with ID %d not found", req.Id)
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned error: %s - %s", resp.Status, string(body))
	}

	var module struct {
		ID         int32           `json:"id"`
		Subject    string          `json:"subject"`
		Grade      string          `json:"grade"`
		Chapter    string          `json:"chapter"`
		ContentJSON json.RawMessage `json:"content"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&module); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &pb.Module{
		Id:          module.ID,
		Subject:     module.Subject,
		Grade:       module.Grade,
		Chapter:     module.Chapter,
		ContentJson: string(module.ContentJSON),
	}, nil
} 