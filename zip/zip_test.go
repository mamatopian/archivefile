package zip

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func ExampleArchiveFile() {
	tmpDir, err := ioutil.TempDir("", "test_zip")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	outFilePath := filepath.Join(tmpDir, "foo.zip")

	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}

	err = ArchiveFile("testdata/foo", outFilePath, progress)
	if err != nil {
		panic(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}

func ExampleUnarchiveFile() {
	tmpDir, err := ioutil.TempDir("", "test_zip")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}

	err = UnarchiveFile("testdata/foo.zip", tmpDir, progress)
	if err != nil {
		panic(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}

func TestArchiveList(t *testing.T) {
	type args struct {
		fileList []string
		progress ProgressFunc
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    error
	}{
		{"test_wildcard", args{fileList: []string{
			"testdata/foo.zip",
			"testdata/foo2.zip",
		}, progress: func(archivePath string) {

		}}, "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			writer := bufio.NewWriter(&b)

			err := ArchiveList(tt.args.fileList, writer, tt.args.progress)

			ioutil.WriteFile("./test.zip", b.Bytes(), 0644)

			if err != nil {
				t.Errorf("ArchiveList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(b.Bytes()) == 0 {
				t.Errorf("ArchiveList() gotWriter = %v, want %v",
					b.Bytes(), tt.wantWriter)
			}
		})
	}
}
