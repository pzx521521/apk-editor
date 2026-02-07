package editor

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

func TestDecompressXML(t *testing.T) {
	data, err := os.ReadFile("/Users/parapeng/Downloads/app-release/AndroidManifest.xml")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	m1 := ModifyInfo[string]{DefaultManifest.VersionName, "6.6.6"}
	m2 := ModifyInfo[uint32]{DefaultManifest.VersionCode, 666}
	m3 := ModifyInfo[string]{DefaultManifest.Label, "TestDemo"}
	m4 := ModifyInfo[string]{DefaultManifest.Package, "com.test.webvievwtr"}
	result, err := ModifyAll(data, m1, m2, m3, m4)
	if err != nil {
		t.Fatalf("Failed to modify manifest: %v", err)
	}
	os.WriteFile("/Users/parapeng/Downloads/app-release/AndroidManifest.new.xml", result, 0644)
}

func TestAdjustStringLength(t *testing.T) {
	_, err := os.Stat("http://www.baidu.com")
	errors.Is(err, fs.ErrNotExist)
	if _, ok := err.(*fs.PathError); ok {
		fmt.Printf("file not exist\n")
	}
}
