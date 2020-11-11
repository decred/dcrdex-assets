package assets

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	_, goFilename, _, _ := runtime.Caller(0)
	basepath = filepath.Dir(filepath.Dir(goFilename))
}

// Path returns the absolute filepath of a file or directory relative to the
// root of the dcrdex-assets main module.
//
// For example, to access the filepath of the site assets directory, call
// Path("dexc/site").
//
// This function is only usable when built without -trimpath and on the host the
// Go program was compiled on.
func Path(asset string) string {
	return filepath.Join(basepath, asset)
}
