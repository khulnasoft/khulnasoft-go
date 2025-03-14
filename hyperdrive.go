package khulnasoft

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

var (
	ErrMissingHyperdriveConfigID             = errors.New("required hyperdrive config id is missing")
	ErrMissingHyperdriveConfigName           = errors.New("required hyperdrive config name is missing")
	ErrMissingHyperdriveConfigOriginDatabase = errors.New("required hyperdrive config origin database is missing")
	ErrMissingHyperdriveConfigOriginPassword = errors.New("required hyperdrive config origin password is missing")
	ErrMissingHyperdriveConfigOriginHost     = errors.New("required hyperdrive config origin host is missing")
	ErrMissingHyperdriveConfigOriginScheme   = errors.New("required hyperdrive config origin scheme is missing")
	ErrMissingHyperdriveConfigOriginUser     = errors.New("required hyperdrive config origin user is missing")
)

type HyperdriveConfig struct {
	ID      string                  `json:"id,omitempty"`
	Name    string                  `json:"name,omitempty"`
	Origin  HyperdriveConfigOrigin  `json:"origin,omitempty"`
	Caching HyperdriveConfigCaching `json:"caching,omitempty"`
}

type HyperdriveOriginType string

type HyperdriveConfigOrigin struct {
	Database       string `json:"database,omitempty"`
	Host           string `json:"host,omitempty"`
	Port           int    `json:"port,omitempty"`
	Scheme         string `json:"scheme,omitempty"`
	User           string `json:"user,omitempty"`
	AccessClientID string `json:"access_client_id,omitempty"`
}

type HyperdriveConfigOriginWithSecrets struct {
	HyperdriveConfigOrigin
	Password           string `json:"password"`
	AccessClientSecret string `json:"access_client_secret,omitempty"`
}

type HyperdriveConfigCaching struct {
	Disabled             *bool `json:"disabled,omitempty"`
	MaxAge               int   `json:"max_age,omitempty"`
	StaleWhileRevalidate int   `json:"stale_while_revalidate,omitempty"`
}

type HyperdriveConfigListResponse struct {
	Response
	Result []HyperdriveConfig `json:"result"`
}

type CreateHyperdriveConfigParams struct {
	Name    string                            `json:"name"`
	Origin  HyperdriveConfigOriginWithSecrets `json:"origin"`
	Caching HyperdriveConfigCaching           `json:"caching,omitempty"`
}

type HyperdriveConfigResponse struct {
	Response
	Result HyperdriveConfig `json:"result"`
}

type UpdateHyperdriveConfigParams struct {
	HyperdriveID string                            `json:"-"`
	Name         string                            `json:"name"`
	Origin       HyperdriveConfigOriginWithSecrets `json:"origin"`
	Caching      HyperdriveConfigCaching           `json:"caching,omitempty"`
}

type ListHyperdriveConfigParams struct{}

// ListHyperdriveConfigs returns the Hyperdrive configs owned by an account.
//
// API reference: https://developers.khulnasoft.com/api/resources/hyperdrive/subresources/configs/methods/list/
func (api *API) ListHyperdriveConfigs(ctx context.Context, rc *ResourceContainer, params ListHyperdriveConfigParams) ([]HyperdriveConfig, error) {
	if rc.Identifier == "" {
		return []HyperdriveConfig{}, ErrMissingAccountID
	}

	hResponse := HyperdriveConfigListResponse{}
	uri := fmt.Sprintf("/accounts/%s/hyperdrive/configs", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []HyperdriveConfig{}, err
	}

	err = json.Unmarshal(res, &hResponse)
	if err != nil {
		return []HyperdriveConfig{}, fmt.Errorf("failed to unmarshal filters JSON data: %w", err)
	}

	return hResponse.Result, nil
}

// CreateHyperdriveConfig creates a new Hyperdrive config.
//
// API reference: https://developers.khulnasoft.com/api/resources/hyperdrive/subresources/configs/methods/create/
func (api *API) CreateHyperdriveConfig(ctx context.Context, rc *ResourceContainer, params CreateHyperdriveConfigParams) (HyperdriveConfig, error) {
	if rc.Identifier == "" {
		return HyperdriveConfig{}, ErrMissingAccountID
	}

	if params.Name == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigName
	}

	uri := fmt.Sprintf("/accounts/%s/hyperdrive/configs", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return HyperdriveConfig{}, err
	}

	var r HyperdriveConfigResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return HyperdriveConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DeleteHyperdriveConfig deletes a Hyperdrive config.
//
// API reference: https://developers.khulnasoft.com/api/resources/hyperdrive/subresources/configs/methods/delete/
func (api *API) DeleteHyperdriveConfig(ctx context.Context, rc *ResourceContainer, hyperdriveID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if hyperdriveID == "" {
		return ErrMissingHyperdriveConfigID
	}

	uri := fmt.Sprintf("/accounts/%s/hyperdrive/configs/%s", rc.Identifier, hyperdriveID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	return nil
}

// GetHyperdriveConfig returns a single Hyperdrive config based on the ID.
//
// API reference: https://developers.khulnasoft.com/api/resources/hyperdrive/subresources/configs/methods/get/
func (api *API) GetHyperdriveConfig(ctx context.Context, rc *ResourceContainer, hyperdriveID string) (HyperdriveConfig, error) {
	if rc.Identifier == "" {
		return HyperdriveConfig{}, ErrMissingAccountID
	}

	if hyperdriveID == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigID
	}

	uri := fmt.Sprintf("/accounts/%s/hyperdrive/configs/%s", rc.Identifier, hyperdriveID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return HyperdriveConfig{}, err
	}

	var r HyperdriveConfigResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return HyperdriveConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateHyperdriveConfig updates a Hyperdrive config.
//
// API reference: https://developers.khulnasoft.com/api/resources/hyperdrive/subresources/configs/methods/update/
func (api *API) UpdateHyperdriveConfig(ctx context.Context, rc *ResourceContainer, params UpdateHyperdriveConfigParams) (HyperdriveConfig, error) {
	if rc.Identifier == "" {
		return HyperdriveConfig{}, ErrMissingAccountID
	}

	if params.HyperdriveID == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigID
	}

	if params.Origin.Database == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigOriginDatabase
	}

	if params.Origin.Password == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigOriginPassword
	}

	if params.Origin.Host == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigOriginHost
	}

	if params.Origin.Scheme == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigOriginScheme
	}

	if params.Origin.User == "" {
		return HyperdriveConfig{}, ErrMissingHyperdriveConfigOriginUser
	}

	uri := fmt.Sprintf("/accounts/%s/hyperdrive/configs/%s", rc.Identifier, params.HyperdriveID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return HyperdriveConfig{}, err
	}

	var r HyperdriveConfigResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return HyperdriveConfig{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}
