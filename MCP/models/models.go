package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SentenceDependenciesOut represents the SentenceDependenciesOut schema from the OpenAPI specification
type SentenceDependenciesOut struct {
	Sentence_dependencies []SentenceDependencyOut `json:"sentence_dependencies"`
}

// Word represents the Word schema from the OpenAPI specification
type Word struct {
	Tag string `json:"tag"`
	Text string `json:"text"`
}

// DependenciesOut represents the DependenciesOut schema from the OpenAPI specification
type DependenciesOut struct {
	Arcs []Arc `json:"arcs"`
	Words []Word `json:"words"`
}

// EntityOut represents the EntityOut schema from the OpenAPI specification
type EntityOut struct {
	End int `json:"end"`
	Start int `json:"start"`
	Text string `json:"text"`
	TypeField string `json:"type"`
}

// VersionOut represents the VersionOut schema from the OpenAPI specification
type VersionOut struct {
	Spacy string `json:"spacy"`
}

// EntitiesOut represents the EntitiesOut schema from the OpenAPI specification
type EntitiesOut struct {
	Entities []EntityOut `json:"entities"`
}

// UserRequestIn represents the UserRequestIn schema from the OpenAPI specification
type UserRequestIn struct {
	Text string `json:"text"`
}

// ValidationError represents the ValidationError schema from the OpenAPI specification
type ValidationError struct {
	TypeField string `json:"type"`
	Loc []string `json:"loc"`
	Msg string `json:"msg"`
}

// Arc represents the Arc schema from the OpenAPI specification
type Arc struct {
	Dir string `json:"dir"`
	End int `json:"end"`
	Label string `json:"label"`
	Start int `json:"start"`
	Text string `json:"text"`
}

// HTTPValidationError represents the HTTPValidationError schema from the OpenAPI specification
type HTTPValidationError struct {
	Detail []ValidationError `json:"detail,omitempty"`
}

// SentenceDependencyOut represents the SentenceDependencyOut schema from the OpenAPI specification
type SentenceDependencyOut struct {
	Dependencies DependenciesOut `json:"dependencies"`
	Sentence string `json:"sentence"`
}
