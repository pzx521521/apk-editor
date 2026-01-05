package main

import (
	"embed"
	"encoding/json"
	"errors"
	"github.com/pzx521521/apk-editor/editor"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed release/*
var embedFiles embed.FS

func Html2Apk(w http.ResponseWriter, r *http.Request) {
	err := html2Apk(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func html2Apk(w http.ResponseWriter, r *http.Request) error {
	crt, err := embedFiles.ReadFile("release/signing.crt")
	if err != nil {
		return err
	}
	apk, err := embedFiles.ReadFile("release/html.apk")
	if err != nil {
		return err
	}

	key, err := embedFiles.ReadFile("release/signing.key")
	if err != nil {
		return err
	}
	apkEditor := editor.NewApkEditor(apk, key, crt)
	// 获取manifest信息
	var manifest editor.Manifest
	manifestJson := r.FormValue("manifest")
	if err := json.Unmarshal([]byte(manifestJson), &manifest); err != nil {
		return err
	}
	apkEditor.Manifest = &manifest
	if url := r.FormValue("url"); url != "" {
		if !strings.HasPrefix(url, "http") {
			return errors.New("url must start with http")
		}
		apkEditor.Url = url
	} else if html, ok := r.MultipartForm.File["html_file"]; ok {
		apkEditor.IndexHtml, err = getFileData(html[0])
		if err != nil {
			return err
		}
	} else if zip, ok := r.MultipartForm.File["zip_file"]; ok {
		apkEditor.HtmlZip, err = getFileData(zip[0])
		if err != nil {
			return err
		}
	} else if so, ok := r.MultipartForm.File["so_file"]; ok {
		apkEditor.SoFile, err = getFileData(so[0])
		apkEditor.SoFileName = filepath.Base(so[0].Filename)
		if err != nil {
			return err
		}
	}

	edit, err := apkEditor.Edit()
	if err != nil {
		return err
	}
	// 用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	// 获取桌面路径
	desktopPath := filepath.Join(homeDir, "Desktop", "webview.apk")

	err = os.WriteFile(desktopPath, edit, 0644)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte("success save at: " + desktopPath))
	if err != nil {
		return err
	}
	return nil
}
func getFileData(file *multipart.FileHeader) ([]byte, error) {
	open, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer open.Close()
	return io.ReadAll(open)
}
