package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"github.com/stretchr/testify/require"
)

func TestMocktail(t *testing.T) {
	testRoot := "./testdata/src"

	dir, err := os.ReadDir(testRoot)
	require.NoError(t, err)

	for _, entry := range dir {
		if !entry.IsDir() {
			continue
		}

		t.Setenv("MOCKTAIL_TEST_PATH", filepath.Join(testRoot, entry.Name()))

		output, err := exec.Command("go", "run", ".").CombinedOutput()
		t.Log(string(output))

		require.NoError(t, err)
	}

	require.NoError(t, filepath.WalkDir(testRoot, func(path string, d fs.DirEntry, errW error) error {
		if errW != nil {
			return errW
		}

		if d.IsDir() || d.Name() != outputMockFile {
			return nil
		}

		genBytes, err := os.ReadFile(path)
		require.NoError(t, err)

		goldenBytes, err := os.ReadFile(path + ".golden")
		require.NoError(t, err)

		edits := myers.ComputeEdits(span.URIFromPath(d.Name()), string(genBytes), string(goldenBytes))

		if len(edits) > 0 {
			diff := fmt.Sprint(gotextdiff.ToUnified(d.Name(), d.Name()+".golden", string(genBytes), edits))
			t.Error(diff)
		}

		return nil
	}))
}
