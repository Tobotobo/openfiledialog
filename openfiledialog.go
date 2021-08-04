// openfiledialog.go
// Copyright (c) 2021 Tobotobo
// This software is released under the MIT License.
// http://opensource.org/licenses/mit-license.php

package openfiledialog

import (
	"github.com/Tobotobo/commondialogs"
	"github.com/lxn/win"
)

type openFileDialog struct {
	Owner          win.HWND
	Title          string
	Filter         string
	FilterIndex    int
	FilePath       string
	InitialDirPath string
}

type OpenFileDialog struct {
	InnerValue openFileDialog
}

type multOpenFileDialog struct {
	Owner          win.HWND
	Title          string
	Filter         string
	FilterIndex    int
	FilePaths      []string
	InitialDirPath string
}

type MultOpenFileDialog struct {
	InnerValue multOpenFileDialog
}

func New() *OpenFileDialog {
	return &OpenFileDialog{
		InnerValue: openFileDialog{
			Owner:          0,
			Title:          "ファイルを開く",
			Filter:         "すべてのファイル(*.*)|*.*",
			FilterIndex:    1,
			FilePath:       "",
			InitialDirPath: "",
		},
	}
}

func NewMult() *MultOpenFileDialog {
	return &MultOpenFileDialog{
		InnerValue: multOpenFileDialog{
			Owner:          0,
			Title:          "ファイルを開く",
			Filter:         "すべてのファイル(*.*)|*.*",
			FilterIndex:    1,
			FilePaths:      []string{},
			InitialDirPath: "",
		},
	}
}

func (dlg *OpenFileDialog) convertToMult() *MultOpenFileDialog {
	multDlg := NewMult()
	multDlg.InnerValue.Owner = dlg.InnerValue.Owner
	multDlg.InnerValue.Title = dlg.InnerValue.Title
	multDlg.InnerValue.Filter = dlg.InnerValue.Filter
	multDlg.InnerValue.FilterIndex = dlg.InnerValue.FilterIndex
	multDlg.InnerValue.InitialDirPath = dlg.InnerValue.InitialDirPath
	return multDlg
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Show() (filePath string, accepted bool) {
	wdlg := new(commondialogs.FileDialog)
	wdlg.Title = dlg.InnerValue.Title
	wdlg.Filter = dlg.InnerValue.Filter
	wdlg.FilterIndex = dlg.InnerValue.FilterIndex
	// wdlg.ShowReadOnlyCB ※読み取り専用で開くが選ばれたかどうかが渡ってこないため未対応
	wdlg.FilePath = dlg.InnerValue.FilePath
	wdlg.InitialDirPath = dlg.InnerValue.InitialDirPath
	// wdlg.Flags

	ok, err := wdlg.ShowOpen(dlg.InnerValue.Owner)
	if err != nil {
		panic(err)
	}
	dlg.InnerValue.FilePath = wdlg.FilePath

	return wdlg.FilePath, ok
}

func (dlg *MultOpenFileDialog) Show() (filePaths []string, accepted bool) {
	wdlg := new(commondialogs.FileDialog)
	wdlg.Title = dlg.InnerValue.Title
	wdlg.Filter = dlg.InnerValue.Filter
	wdlg.FilterIndex = dlg.InnerValue.FilterIndex
	// wdlg.ShowReadOnlyCB ※読み取り専用で開くが選ばれたかどうかが渡ってこないため未対応
	wdlg.FilePaths = dlg.InnerValue.FilePaths
	wdlg.InitialDirPath = dlg.InnerValue.InitialDirPath
	// wdlg.Flags

	ok, err := wdlg.ShowOpenMultiple(dlg.InnerValue.Owner)
	if err != nil {
		panic(err)
	}
	dlg.InnerValue.FilePaths = wdlg.FilePaths

	return wdlg.FilePaths, ok
}

func Show() (filePath string, accepted bool) {
	dlg := New()
	return dlg.Show()
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Mult() *MultOpenFileDialog {
	return dlg.convertToMult()
}

func Mult() *MultOpenFileDialog {
	return NewMult()
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Owner(owner win.HWND) *OpenFileDialog {
	dlg.InnerValue.Owner = owner
	return dlg
}

func (dlg *MultOpenFileDialog) Owner(owner win.HWND) *MultOpenFileDialog {
	dlg.InnerValue.Owner = owner
	return dlg
}

func Owner(owner win.HWND) *OpenFileDialog {
	dlg := New()
	return dlg.Owner(owner)
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Title(title string) *OpenFileDialog {
	dlg.InnerValue.Title = title
	return dlg
}

func (dlg *MultOpenFileDialog) Title(title string) *MultOpenFileDialog {
	dlg.InnerValue.Title = title
	return dlg
}

func Title(title string) *OpenFileDialog {
	dlg := New()
	return dlg.Title(title)
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Filter(filter string, index ...int) *OpenFileDialog {
	dlg.InnerValue.Filter = filter
	if len(index) > 0 {
		dlg.InnerValue.FilterIndex = index[0]
	}
	return dlg
}

func (dlg *MultOpenFileDialog) Filter(filter string, index ...int) *MultOpenFileDialog {
	dlg.InnerValue.Filter = filter
	if len(index) > 0 {
		dlg.InnerValue.FilterIndex = index[0]
	}
	return dlg
}

func Filter(filter string, index ...int) *OpenFileDialog {
	dlg := New()
	return dlg.Filter(filter, index...)
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) InitFilePath(path string) *OpenFileDialog {
	dlg.InnerValue.FilePath = path
	return dlg
}

func InitFilePath(path string) *OpenFileDialog {
	dlg := New()
	return dlg.InitFilePath(path)
}

// ----------------------------------------------------------------

// 予め格納してもFilePathと違って初期表示に反映されず
// func (dlg *MultOpenFileDialog) InitFilePaths(paths []string) *MultOpenFileDialog {
// 	dlg.FilePaths = paths
// 	return dlg
// }

// func InitFilePaths(paths []string) *MultOpenFileDialog {
// 	dlg := NewMult()
// 	return dlg.InitFilePaths(paths)
// }

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) InitDirPath(path string) *OpenFileDialog {
	dlg.InnerValue.InitialDirPath = path
	return dlg
}

func (dlg *MultOpenFileDialog) InitDirPath(path string) *MultOpenFileDialog {
	dlg.InnerValue.InitialDirPath = path
	return dlg
}

func InitDirPath(path string) *OpenFileDialog {
	dlg := New()
	return dlg.InitDirPath(path)
}
