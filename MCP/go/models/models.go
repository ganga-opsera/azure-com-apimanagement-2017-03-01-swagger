package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// PolicyContractProperties represents the PolicyContractProperties schema from the OpenAPI specification
type PolicyContractProperties struct {
	Policycontent string `json:"policyContent"` // Json escaped Xml Encoded contents of the Policy.
}

// RegionListResult represents the RegionListResult schema from the OpenAPI specification
type RegionListResult struct {
	Value []RegionContract `json:"value,omitempty"` // Lists of Regions.
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
}

// ErrorFieldContract represents the ErrorFieldContract schema from the OpenAPI specification
type ErrorFieldContract struct {
	Message string `json:"message,omitempty"` // Human-readable representation of property-level error.
	Target string `json:"target,omitempty"` // Property name.
	Code string `json:"code,omitempty"` // Property level error code.
}

// PolicySnippetsCollection represents the PolicySnippetsCollection schema from the OpenAPI specification
type PolicySnippetsCollection struct {
	Value []PolicySnippetContract `json:"value,omitempty"` // Policy snippet value.
}

// PolicyContract represents the PolicyContract schema from the OpenAPI specification
type PolicyContract struct {
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
	Id string `json:"id,omitempty"` // Resource ID.
}

// RegionContract represents the RegionContract schema from the OpenAPI specification
type RegionContract struct {
	Name string `json:"name,omitempty"` // Region name.
	Isdeleted bool `json:"isDeleted,omitempty"` // whether Region is deleted.
	Ismasterregion bool `json:"isMasterRegion,omitempty"` // whether Region is the master region.
}

// PolicyCollection represents the PolicyCollection schema from the OpenAPI specification
type PolicyCollection struct {
	Value []PolicyContract `json:"value,omitempty"` // Policy Contract value.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
}

// PolicySnippetContract represents the PolicySnippetContract schema from the OpenAPI specification
type PolicySnippetContract struct {
	Name string `json:"name,omitempty"` // Snippet name.
	Scope int `json:"scope,omitempty"` // Binary OR value of the Snippet scope.
	Tooltip string `json:"toolTip,omitempty"` // Snippet toolTip.
	Content string `json:"content,omitempty"` // Snippet content.
}

// Resource represents the Resource schema from the OpenAPI specification
type Resource struct {
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Code string `json:"code,omitempty"` // Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
	Details []ErrorFieldContract `json:"details,omitempty"` // The list of invalid fields send in request, in case of validation error.
	Message string `json:"message,omitempty"` // Human-readable representation of the error.
}
