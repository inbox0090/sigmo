package notify

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func parseEndpoint(name string, raw string, fallback string) (*url.URL, error) {
	endpoint := strings.TrimSpace(raw)
	if endpoint == "" {
		endpoint = strings.TrimSpace(fallback)
	}
	if endpoint == "" {
		return nil, fmt.Errorf("%s endpoint is required", name)
	}
	parsed, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing %s endpoint: %w", name, err)
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return nil, fmt.Errorf("%s endpoint must include scheme and host", name)
	}
	return parsed, nil
}

func ensureEndpointPath(parsed *url.URL, segment string) {
	trimmed := strings.TrimRight(parsed.Path, "/")
	if trimmed == "" {
		parsed.Path = "/" + segment
		return
	}
	if strings.HasSuffix(trimmed, "/"+segment) {
		parsed.Path = trimmed
		return
	}
	parsed.Path = path.Join(trimmed, segment)
}
