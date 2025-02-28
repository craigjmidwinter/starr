package sonarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/craigjmidwinter/starr"
)

/* Custom Formats do not exist in Sonarr v3; this is v4 only. */

const bpCustomFormat = APIver + "/customFormat"

// CustomFormat is the api/customformat endpoint payload.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
type CustomFormat struct {
	ID                    int                 `json:"id"`
	Name                  string              `json:"name"`
	IncludeCFWhenRenaming bool                `json:"includeCustomFormatWhenRenaming"`
	Specifications        []*CustomFormatSpec `json:"specifications"`
}

// CustomFormatSpec is part of a CustomFormat.
type CustomFormatSpec struct {
	Name               string               `json:"name"`
	Implementation     string               `json:"implementation"`
	Implementationname string               `json:"implementationName"`
	Infolink           string               `json:"infoLink"`
	Negate             bool                 `json:"negate"`
	Required           bool                 `json:"required"`
	Fields             []*CustomFormatField `json:"fields"`
}

// CustomFormatField is part of a CustomFormat Specification.
type CustomFormatField struct {
	Order    int         `json:"order"`
	Name     string      `json:"name"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value"` // should be a string, but sometimes it's a number.
	Type     string      `json:"type"`
	Advanced bool        `json:"advanced"`
}

// GetCustomFormats returns all configured Custom Formats.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) GetCustomFormats() ([]*CustomFormat, error) {
	return s.GetCustomFormatsContext(context.Background())
}

// GetCustomFormatsContext returns all configured Custom Formats.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) GetCustomFormatsContext(ctx context.Context) ([]*CustomFormat, error) {
	var output []*CustomFormat

	req := starr.Request{URI: bpCustomFormat}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddCustomFormat creates a new custom format and returns the response (with ID).
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) AddCustomFormat(format *CustomFormat) (*CustomFormat, error) {
	return s.AddCustomFormatContext(context.Background(), format)
}

// AddCustomFormatContext creates a new custom format and returns the response (with ID).
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) AddCustomFormatContext(ctx context.Context, format *CustomFormat) (*CustomFormat, error) {
	var output CustomFormat

	if format == nil {
		return &output, nil
	}

	format.ID = 0 // ID must be zero when adding.

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFormat, err)
	}

	req := starr.Request{URI: bpCustomFormat, Body: &body}
	if err := s.PostInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return &output, nil
}

// UpdateCustomFormat updates an existing custom format and returns the response.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) UpdateCustomFormat(cf *CustomFormat, cfID int) (*CustomFormat, error) {
	return s.UpdateCustomFormatContext(context.Background(), cf, cfID)
}

// UpdateCustomFormatContext updates an existing custom format and returns the response.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) UpdateCustomFormatContext(ctx context.Context, format *CustomFormat, cfID int) (*CustomFormat, error) {
	if cfID == 0 {
		cfID = format.ID
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(format); err != nil {
		return nil, fmt.Errorf("json.Marshal(%s): %w", bpCustomFormat, err)
	}

	var output CustomFormat

	req := starr.Request{URI: path.Join(bpCustomFormat, fmt.Sprint(cfID)), Body: &body}
	if err := s.PutInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return &output, nil
}

// DeleteCustomFormat deletes a custom format.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) DeleteCustomFormat(cfID int) error {
	return s.DeleteCustomFormatContext(context.Background(), cfID)
}

// DeleteCustomFormatContext deletes a custom format.
// This data and these endpoints do not exist in Sonarr v3; this is v4 only.
func (s *Sonarr) DeleteCustomFormatContext(ctx context.Context, cfID int) error {
	req := starr.Request{URI: path.Join(bpCustomFormat, fmt.Sprint(cfID))}
	if err := s.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
