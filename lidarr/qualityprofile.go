package lidarr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"

	"github.com/craigjmidwinter/starr"
)

const bpQualityProfile = APIver + "/qualityProfile"

// QualityProfile is the /api/v1/qualityprofile endpoint.
type QualityProfile struct {
	ID             int64            `json:"id"`
	Name           string           `json:"name"`
	UpgradeAllowed bool             `json:"upgradeAllowed"`
	Cutoff         int64            `json:"cutoff"`
	Qualities      []*starr.Quality `json:"items"`
}

// GetQualityProfiles returns the quality profiles.
func (l *Lidarr) GetQualityProfiles() ([]*QualityProfile, error) {
	return l.GetQualityProfilesContext(context.Background())
}

// GetQualityProfilesContext returns the quality profiles.
func (l *Lidarr) GetQualityProfilesContext(ctx context.Context) ([]*QualityProfile, error) {
	var output []*QualityProfile

	req := starr.Request{URI: bpQualityProfile}
	if err := l.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// AddQualityProfile updates a quality profile in place.
func (l *Lidarr) AddQualityProfile(profile *QualityProfile) (int64, error) {
	return l.AddQualityProfileContext(context.Background(), profile)
}

// AddQualityProfileContext updates a quality profile in place.
func (l *Lidarr) AddQualityProfileContext(ctx context.Context, profile *QualityProfile) (int64, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return 0, fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	var output QualityProfile

	req := starr.Request{URI: bpQualityProfile, Body: &body}
	if err := l.PostInto(ctx, req, &output); err != nil {
		return 0, fmt.Errorf("api.Post(%s): %w", &req, err)
	}

	return output.ID, nil
}

// UpdateQualityProfile updates a quality profile in place.
func (l *Lidarr) UpdateQualityProfile(profile *QualityProfile) error {
	return l.UpdateQualityProfileContext(context.Background(), profile)
}

// UpdateQualityProfileContext updates a quality profile in place.
func (l *Lidarr) UpdateQualityProfileContext(ctx context.Context, profile *QualityProfile) error {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(profile); err != nil {
		return fmt.Errorf("json.Marshal(%s): %w", bpQualityProfile, err)
	}

	var output interface{}

	req := starr.Request{URI: path.Join(bpQualityProfile, fmt.Sprint(profile.ID)), Body: &body}
	if err := l.PutInto(ctx, req, &output); err != nil {
		return fmt.Errorf("api.Put(%s): %w", &req, err)
	}

	return nil
}

// DeleteQualityProfile deletes a quality profile.
func (l *Lidarr) DeleteQualityProfile(profileID int64) error {
	return l.DeleteQualityProfileContext(context.Background(), profileID)
}

// DeleteQualityProfileContext deletes a quality profile.
func (l *Lidarr) DeleteQualityProfileContext(ctx context.Context, profileID int64) error {
	req := starr.Request{URI: path.Join(bpQualityProfile, fmt.Sprint(profileID))}
	if err := l.DeleteAny(ctx, req); err != nil {
		return fmt.Errorf("api.Delete(%s): %w", &req, err)
	}

	return nil
}
