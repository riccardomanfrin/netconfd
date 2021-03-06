/*
 * netConfD API
 *
 * Network Configurator service
 *
 * API version: 0.3.0
 * Contact: support@athonet.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"

)

// NewSystemApiService creates a default api service
func NewSystemApiService() SystemApiServicer {
	return &SystemApiService{}
}

// ConfigGet - Get current live configuration 
func (s *SystemApiService) ConfigGet(ctx context.Context) (ImplResponse, error) {
	return s.ConfigGet_Impl(ctx)
}

// ConfigPatch - Patch existing configuration with new one 
func (s *SystemApiService) ConfigPatch(ctx context.Context, config Config) (ImplResponse, error) {
	return s.ConfigPatch_Impl(ctx, config)
}

// ConfigSet - Replace existing configuration with new one 
func (s *SystemApiService) ConfigSet(ctx context.Context, config Config) (ImplResponse, error) {
	return s.ConfigSet_Impl(ctx, config)
}

// PersistConfig - Persist live configuration
func (s *SystemApiService) PersistConfig(ctx context.Context) (ImplResponse, error) {
	return s.PersistConfig_Impl(ctx)
}

// ResetConfig - Reload persisted configuration back
func (s *SystemApiService) ResetConfig(ctx context.Context) (ImplResponse, error) {
	return s.ResetConfig_Impl(ctx)
}

