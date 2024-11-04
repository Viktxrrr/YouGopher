package utils

import "regexp"

func IsValidYouTubeURL(url string) bool {
	re := regexp.MustCompile(
		`^(https?://)?(www\.)?(youtube\.com/watch\?v=|youtu\.be/)([a-zA-Z0-9_-]{11})$`,
	)
	return re.MatchString(url)
}
