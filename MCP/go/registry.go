package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_policy "github.com/apimanagementclient/mcp-server/tools/policy"
	tools_policysnippets "github.com/apimanagementclient/mcp-server/tools/policysnippets"
	tools_regions "github.com/apimanagementclient/mcp-server/tools/regions"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_policy.CreatePolicy_deleteTool(cfg),
		tools_policy.CreatePolicy_getTool(cfg),
		tools_policy.CreatePolicy_createorupdateTool(cfg),
		tools_policysnippets.CreatePolicysnippets_listbyserviceTool(cfg),
		tools_regions.CreateRegions_listbyserviceTool(cfg),
		tools_policy.CreatePolicy_listbyserviceTool(cfg),
	}
}
