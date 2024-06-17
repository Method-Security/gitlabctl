package config

import (
	"strings"
)

// NormalizeGitlabURL normalizes the Gitlab URL to ensure it is in the correct format for the Gitlab API, which requires
// the URL to end with "/api/v4" and to start with "https".
func NormalizeGitlabURL(url string) string {
	if !strings.HasSuffix(url, "/api/v4") {
		url = strings.TrimSuffix(url, "/") + "/api/v4"
	}
	if !strings.HasPrefix(url, "https") {
		url = "https://" + url
	}
	return url
}
