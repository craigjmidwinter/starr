package sonarr

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"time"

	"github.com/craigjmidwinter/starr"
)

// Define Base Path for Calendar queries.
const bpCalendar = APIver + "/calendar"

// Calendar defines the filters for fetching calendar items.
type Calendar struct {
	Start                time.Time
	End                  time.Time
	Unmonitored          bool
	IncludeSeries        bool
	IncludeEpisodeFile   bool
	IncludeEpisodeImages bool
}

// GetCalendar returns calendars based on filters.
func (s *Sonarr) GetCalendar(filter Calendar) ([]*Episode, error) {
	return s.GetCalendarContext(context.Background(), filter)
}

// GetCalendarContext returns calendars based on filters.
func (s *Sonarr) GetCalendarContext(ctx context.Context, filter Calendar) ([]*Episode, error) {
	var output []*Episode

	req := starr.Request{URI: bpCalendar, Query: make(url.Values)}
	req.Query.Add("unmonitored", fmt.Sprint(filter.Unmonitored))
	req.Query.Add("includeSeries", fmt.Sprint(filter.IncludeSeries))
	req.Query.Add("includeEpisodeFile", fmt.Sprint(filter.IncludeEpisodeFile))
	req.Query.Add("includeEpisodeImages", fmt.Sprint(filter.IncludeEpisodeImages))

	if !filter.Start.IsZero() {
		req.Query.Add("start", filter.Start.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if !filter.End.IsZero() {
		req.Query.Add("end", filter.End.UTC().Format(starr.CalendarTimeFilterFormat))
	}

	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}

// GetCalendarID returns a single calendar by ID.
func (s *Sonarr) GetCalendarID(calendarID int64) (*Episode, error) {
	return s.GetCalendarIDContext(context.Background(), calendarID)
}

// GetCalendarIDContext returns a single calendar by ID.
func (s *Sonarr) GetCalendarIDContext(ctx context.Context, calendarID int64) (*Episode, error) {
	var output *Episode

	req := starr.Request{URI: path.Join(bpCalendar, fmt.Sprint(calendarID))}
	if err := s.GetInto(ctx, req, &output); err != nil {
		return nil, fmt.Errorf("api.Get(%s): %w", &req, err)
	}

	return output, nil
}
