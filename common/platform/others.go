//go:build !windows
// +build !windows

package platform

import (
	"os"
	"path/filepath"
)

var assetDirPath = ""

func ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func LineSeparator() string {
	return "\n"
}

func SetAssetPath(path string) {
	assetDirPath = path
}

func GetToolLocation(file string) string {
	toolPath := NewEnvFlag(ToolLocation).GetValue(getExecutableDir)
	return filepath.Join(toolPath, file)
}

// GetAssetLocation searches for `file` in certain locations
func GetAssetLocation(file string) string {
	assetPath := assetDirPath
	if assetPath == "" {
		assetPath = NewEnvFlag(AssetLocation).GetValue(getExecutableDir)
	}
	defPath := filepath.Join(assetPath, file)
	for _, p := range []string{
		defPath,
		filepath.Join("/usr/local/share/xray/", file),
		filepath.Join("/usr/share/xray/", file),
		filepath.Join("/opt/share/xray/", file),
	} {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			continue
		}

		// asset found
		return p
	}

	// asset not found, let the caller throw out the error
	return defPath
}
