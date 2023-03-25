package file_

import "os"

func File_() {
	tmpFile, _ := os.CreateTemp("", "tmp")
	defer tmpFile.Close()
	tmpFile.WriteString("hello")
}
