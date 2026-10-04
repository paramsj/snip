package main

import (
	"net/url"
	"regexp"
	"strings"
)

var aliasPattern = regexp.MustCompile(`^[A-Za-z0-9_-]{3,20}$`)

func normalizeURL(raw string) string {
	raw = strings.TrimSpace(raw)
	u, err := url.Parse(raw)

	if err != nil {
		return ""
	}

	if u.Scheme == "http" || u.Scheme == "https" {
		return raw
	}

	if u.Scheme == "" {
		return "https://" + raw
	}

	return raw
}

func validURL(raw string) bool {
	if raw == "" || len(raw) > 2048 {
		return false
	}
	u, err := url.Parse(raw)
	return err == nil && (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
}

func validAlias(alias string) bool {
	return aliasPattern.MatchString(alias)
}
