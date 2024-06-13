package config

import (
	"strings"
)

func NormalizeGitlabURL(url string) string {
	if !strings.HasSuffix(url, "/api/v4") {
		url = strings.TrimSuffix(url, "/") + "/api/v4"
	}
	if !strings.HasPrefix(url, "https") {
		url = "https://" + url
	}
	return url
}
