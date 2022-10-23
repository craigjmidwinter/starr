package radarr

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/craigjmidwinter/starr"
)

const bpSystem = APIver + "/system"

// SystemStatus is the /api/v3/system/status endpoint.
type SystemStatus struct {
	AppData                string    `json:"appData"`
	AppName                string    `json:"appName"`
	Authentication         string    `json:"authentication"`
	Branch                 string    `json:"branch"`
	BuildTime              time.Time `json:"buildTime"`
	DatabaseType           string    `json:"databaseType"`
	DatabaseVersion        string    `json:"databaseVersion"`
	InstanceName           string    `json:"instanceName"`
	IsAdmin                bool      `json:"isAdmin"`
	IsDebug                bool      `json:"isDebug"`
	IsDocker               bool      `json:"isDocker"`
	IsLinux                bool      `json:"isLinux"`
	IsNetCore              bool      `json:"isNetCore"`
	IsOsx                  bool      `json:"isOsx"`
	IsProduction           bool      `json:"isProduction"`
	IsUserInteractive      bool      `json:"isUserInteractive"`
	IsWindows              bool      `json:"isWindows"`
	MigrationVersion       int64     `json:"migrationVersion"`
	Mode                   string    `json:"mode"`
	OsName                 string    `json:"osName"`
	PackageAuthor          string    `json:"packageAuthor"`
	PackageUpdateMechanism string    `json:"packageUpdateMechanism"`
	PackageVersion         string    `json:"packageVersion"`
	RuntimeName            string    `json:"runtimeName"`
	RuntimeVersion         string    `json:"runtimeVersion"`
	StartTime              time.Time `json:"startTime"`
	StartupPath            string    `json:"startupPath"`
	URLBase                string    `json:"urlBase"`
	Version                string    `json:"version"`
}

// GetSystemStatus returns system status.
func (r *Radarr) GetSystemStatus() (*SystemStatus, error) {
	return r.GetSystemStatusContext(context.Background())
}

// GetSystemStatusContext returns system status.
func (r *Radarr) GetSystemStatusContext(ctx context.Context) (*SystemStatus, error) {
	var output SystemStatus

	req := starr.Request{URI: path.Join(bpSystem, "status")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%v): %w", &req, err)
	}

	return &output, nil
}

// GetBackupFiles returns all available Radarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Radarr) GetBackupFiles() ([]*starr.BackupFile, error) {
	return r.GetBackupFilesContext(context.Background())
}

// GetBackupFilesContext returns all available Radarr backup files.
// Use GetBody to download a file using BackupFile.Path.
func (r *Radarr) GetBackupFilesContext(ctx context.Context) ([]*starr.BackupFile, error) {
	var output []*starr.BackupFile

	req := starr.Request{URI: path.Join(bpSystem, "backup")}
	if err := r.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
