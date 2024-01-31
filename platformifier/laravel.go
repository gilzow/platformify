package platformifier

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/platformsh/platformify/internal/colors"
	"github.com/platformsh/platformify/internal/utils"
)

const (
	composerJSONFile = "composer.json"
)

func newLaravelPlatformifier(templates fs.FS, fileSystem FS) *laravelPlatformifier {
	return &laravelPlatformifier{
		templates:  templates,
		fileSystem: fileSystem,
	}
}

type laravelPlatformifier struct {
	templates  fs.FS
	fileSystem FS
}

func (p *laravelPlatformifier) Platformify(ctx context.Context, input *UserInput) error {
	// Check for the Laravel Bridge.
	out, _, _ := colors.FromContext(ctx)
	fmt.Fprintln(out, colors.Colorize(colors.AccentCode, "yo yo! you be runnin' Laravel!"))
	appRoot := filepath.Join(input.Root, input.ApplicationRoot)
	composerJSONPaths := p.fileSystem.Find(appRoot, composerJSONFile, false)
	for _, composerJSONPath := range composerJSONPaths {
		_, required := utils.GetJSONValue([]string{"require", "platformsh/laravel-bridge"}, composerJSONPath, true)
		if !required {
			out, _, ok := colors.FromContext(ctx)
			if !ok {
				return fmt.Errorf("output context failed")
			}

			var suggest = "\nPlease use composer to add the Laravel Bridge to your project:\n"
			var composerRequire = "\n    composer require platformsh/laravel-bridge\n"
			fmt.Fprintln(out, colors.Colorize(colors.WarningCode, suggest+composerRequire))
		}
	}

	return nil
}
