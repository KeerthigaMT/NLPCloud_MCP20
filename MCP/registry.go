package main

import (
	"github.com/nlpcloud/mcp-server/config"
	"github.com/nlpcloud/mcp-server/models"
	tools_v1 "github.com/nlpcloud/mcp-server/tools/v1"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v1.CreateRead_version_v1_en_core_web_sm_version_getTool(cfg),
		tools_v1.CreateRead_root_v1_en_core_web_sm_getTool(cfg),
		tools_v1.CreateRead_dependencies_v1_en_core_web_sm_dependencies_postTool(cfg),
		tools_v1.CreateRead_entities_v1_en_core_web_sm_entities_postTool(cfg),
		tools_v1.CreateRead_sentence_dependencies_v1_en_core_web_sm_sentence_dependencies_postTool(cfg),
	}
}
