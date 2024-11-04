package utils

import (
	"mime"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kkdai/youtube/v2"
)

const defaultExtension = ".mov"

var canonicals = map[string]string{
	"video/quicktime":  ".mov",
	"video/x-msvideo":  ".avi",
	"video/x-matroska": ".mkv",
	"video/mpeg":       ".mpeg",
	"video/webm":       ".webm",
	"video/3gpp2":      ".3g2",
	"video/x-flv":      ".flv",
	"video/3gpp":       ".3gp",
	"video/mp4":        ".mp4",
	"video/ogg":        ".ogv",
	"video/mp2t":       ".ts",
}

func pickIdealFileExtension(mediaType string) string {
	mediaType, _, err := mime.ParseMediaType(mediaType)
	if err != nil {
		return defaultExtension
	}

	if extension, ok := canonicals[mediaType]; ok {
		return extension
	}

	extensions, err := mime.ExtensionsByType(mediaType)
	if err != nil || extensions == nil {
		return defaultExtension
	}

	return extensions[0]
}

func SanitizeFilename(fileName string) string {
	// Characters not allowed on mac
	//	:/
	// Characters not allowed on linux
	//	/
	// Characters not allowed on windows
	//	<>:"/\|?*

	// Ref https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file#naming-conventions

	fileName = regexp.MustCompile(`[:/<>\:"\\|?*]`).ReplaceAllString(fileName, "")
	fileName = regexp.MustCompile(`\s+`).ReplaceAllString(fileName, " ")

	return fileName
}

func GetOutputFile(
	format youtube.Format,
	outputDir string,
	outputFile string,
) (string, error) {
	outputFile = SanitizeFilename(outputFile)
	outputFile += pickIdealFileExtension(format.MimeType)

	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0o755); err != nil {
			return "", err
		}
		outputFile = filepath.Join(outputDir, outputFile)
	}

	return outputFile, nil
}
