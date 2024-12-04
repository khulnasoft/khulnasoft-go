package khulnasoft

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// WAFOverridesResponse represents the response from the WAF overrides endpoint.
type WAFOverridesResponse struct {
	Response
	Result     []WAFOverride `json:"result"`
	ResultInfo ResultInfo    `json:"result_info"`
}

// WAFOverrideResponse represents the response from a single WAF override endpoint.
type WAFOverrideResponse struct {
	Response
	Result     WAFOverride `json:"result"`
	ResultInfo ResultInfo  `json:"result_info"`
}

// WAFOverride represents a WAF override configuration.
type WAFOverride struct {
	ID            string            `json:"id,omitempty"`
	Description   string            `json:"description"`
	URLs          []string          `json:"urls"`
	Priority      int               `json:"priority"`
	Groups        map[string]string `json:"groups"`
	RewriteAction map[string]string `json:"rewrite_action"`
	Rules         map[string]string `json:"rules"`
	Paused        bool              `json:"paused"`
}

// Utility function to unmarshal JSON responses.
func unmarshalResponse(res []byte, out interface{}) error {
	if err := json.Unmarshal(res, out); err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return nil
}

// ListWAFOverrides returns a slice of the WAF overrides for a given zone ID.
//
// API Reference: https://api.khulnasoft.com/#waf-overrides-list-uri-controlled-waf-configurations
func (api *API) ListWAFOverrides(ctx context.Context, zoneID string) ([]WAFOverride, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/waf/overrides", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch WAF overrides: %w", err)
	}

	var r WAFOverridesResponse
	if err := unmarshalResponse(res, &r); err != nil {
		return nil, fmt.Errorf("failed to parse WAF overrides response: %w", err)
	}

	if !r.Success {
		return nil, fmt.Errorf("API error: %v", r.Errors)
	}

	return r.Result, nil
}

// WAFOverride fetches a specific WAF override by its ID.
//
// API Reference: https://api.khulnasoft.com/#waf-overrides-uri-controlled-waf-configuration-details
func (api *API) WAFOverride(ctx context.Context, zoneID, overrideID string) (WAFOverride, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/waf/overrides/%s", zoneID, overrideID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WAFOverride{}, fmt.Errorf("failed to fetch WAF override: %w", err)
	}

	var r WAFOverrideResponse
	if err := unmarshalResponse(res, &r); err != nil {
		return WAFOverride{}, fmt.Errorf("failed to parse WAF override response: %w", err)
	}

	if !r.Success {
		return WAFOverride{}, fmt.Errorf("API error: %v", r.Errors)
	}

	return r.Result, nil
}

// CreateWAFOverride creates a new WAF override.
//
// API Reference: https://api.khulnasoft.com/#waf-overrides-create-a-uri-controlled-waf-configuration
func (api *API) CreateWAFOverride(ctx context.Context, zoneID string, override WAFOverride) (WAFOverride, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/waf/overrides", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, override)
	if err != nil {
		return WAFOverride{}, fmt.Errorf("failed to create WAF override: %w", err)
	}

	var r WAFOverrideResponse
	if err := unmarshalResponse(res, &r); err != nil {
		return WAFOverride{}, fmt.Errorf("failed to parse WAF override creation response: %w", err)
	}

	if !r.Success {
		return WAFOverride{}, fmt.Errorf("API error: %v", r.Errors)
	}

	return r.Result, nil
}

// UpdateWAFOverride updates an existing WAF override.
//
// API Reference: https://api.khulnasoft.com/#waf-overrides-update-uri-controlled-waf-configuration
func (api *API) UpdateWAFOverride(ctx context.Context, zoneID, overrideID string, override WAFOverride) (WAFOverride, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/waf/overrides/%s", zoneID, overrideID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, override)
	if err != nil {
		return WAFOverride{}, fmt.Errorf("failed to update WAF override: %w", err)
	}

	var r WAFOverrideResponse
	if err := unmarshalResponse(res, &r); err != nil {
		return WAFOverride{}, fmt.Errorf("failed to parse WAF override update response: %w", err)
	}

	if !r.Success {
		return WAFOverride{}, fmt.Errorf("API error: %v", r.Errors)
	}

	return r.Result, nil
}

// DeleteWAFOverride deletes a WAF override by its ID.
//
// API Reference: https://api.khulnasoft.com/#waf-overrides-delete-lockdown-rule
func (api *API) DeleteWAFOverride(ctx context.Context, zoneID, overrideID string) error {
	uri := fmt.Sprintf("/zones/%s/firewall/waf/overrides/%s", zoneID, overrideID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("failed to delete WAF override: %w", err)
	}

	var r WAFOverrideResponse
	if err := unmarshalResponse(res, &r); err != nil {
		return fmt.Errorf("failed to parse WAF override deletion response: %w", err)
	}

	if !r.Success {
		return fmt.Errorf("API error: %v", r.Errors)
	}

	return nil
}
