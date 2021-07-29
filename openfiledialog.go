package main

import (
	"github.com/lxn/walk"
)

type OpenFileDialog struct {
	OwnerForm      *walk.Form
	TitleText      string
	FilterText     string
	FilterIndex    int
	FilePath       string
	InitialDirPath string
}

type MultOpenFileDialog struct {
	OwnerForm      *walk.Form
	TitleText      string
	FilterText     string
	FilterIndex    int
	FilePaths      []string
	InitialDirPath string
}

func New() *OpenFileDialog {
	return &OpenFileDialog{
		OwnerForm:      nil,
		TitleText:      "ファイルを開く",
		FilterText:     "すべてのファイル(*.*)|*.*",
		FilterIndex:    1,
		FilePath:       "",
		InitialDirPath: "",
	}
}

func NewMult() *MultOpenFileDialog {
	return &MultOpenFileDialog{
		OwnerForm:      nil,
		TitleText:      "ファイルを開く",
		FilterText:     "すべてのファイル(*.*)|*.*",
		FilterIndex:    1,
		FilePaths:      []string{},
		InitialDirPath: "",
	}
}

func (dlg *OpenFileDialog) convertToMult() *MultOpenFileDialog {
	multDlg := NewMult()
	multDlg.OwnerForm = dlg.OwnerForm
	multDlg.TitleText = dlg.TitleText
	multDlg.FilterText = dlg.FilterText
	multDlg.FilterIndex = dlg.FilterIndex
	multDlg.InitialDirPath = dlg.InitialDirPath
	return multDlg
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Show() (accepted bool, filePath string) {
	wdlg := new(walk.FileDialog)
	wdlg.Title = dlg.TitleText
	wdlg.Filter = dlg.FilterText
	wdlg.FilterIndex = dlg.FilterIndex
	// wdlg.ShowReadOnlyCB ※読み取り専用で開くが選ばれたかどうかが渡ってこないため未対応
	wdlg.FilePath = dlg.FilePath
	wdlg.InitialDirPath = dlg.InitialDirPath
	// wdlg.Flags

	var owner walk.Form = nil
	if dlg.OwnerForm != nil {
		owner = *dlg.OwnerForm
	}

	ok, err := wdlg.ShowOpen(owner)
	if err != nil {
		panic(err)
	}

	return ok, wdlg.FilePath
}

func (dlg *MultOpenFileDialog) Show() (accepted bool, filePaths []string) {
	wdlg := new(walk.FileDialog)
	wdlg.Title = dlg.TitleText
	wdlg.Filter = dlg.FilterText
	wdlg.FilterIndex = dlg.FilterIndex
	// wdlg.ShowReadOnlyCB ※読み取り専用で開くが選ばれたかどうかが渡ってこないため未対応
	wdlg.FilePaths = dlg.FilePaths
	wdlg.InitialDirPath = dlg.InitialDirPath
	// wdlg.Flags

	var owner walk.Form = nil
	if dlg.OwnerForm != nil {
		owner = *dlg.OwnerForm
	}

	ok, err := wdlg.ShowOpenMultiple(owner)
	if err != nil {
		panic(err)
	}

	return ok, wdlg.FilePaths
}

func Show() (accepted bool, filePath string) {
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

func (dlg *OpenFileDialog) Owner(owner *walk.Form) *OpenFileDialog {
	dlg.OwnerForm = owner
	return dlg
}

func (dlg *MultOpenFileDialog) Owner(owner *walk.Form) *MultOpenFileDialog {
	dlg.OwnerForm = owner
	return dlg
}

func Owner(owner *walk.Form) *OpenFileDialog {
	dlg := New()
	return dlg.Owner(owner)
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Title(title string) *OpenFileDialog {
	dlg.TitleText = title
	return dlg
}

func (dlg *MultOpenFileDialog) Title(title string) *MultOpenFileDialog {
	dlg.TitleText = title
	return dlg
}

func Title(title string) *OpenFileDialog {
	dlg := New()
	return dlg.Title(title)
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) Filter(filter string, index int) *OpenFileDialog {
	dlg.FilterText = filter
	dlg.FilterIndex = index
	return dlg
}

func (dlg *MultOpenFileDialog) Filter(filter string, index int) *MultOpenFileDialog {
	dlg.FilterText = filter
	dlg.FilterIndex = index
	return dlg
}

func Filter(filter string, index int) *OpenFileDialog {
	dlg := New()
	return dlg.Filter(filter, index)
}

// ----------------------------------------------------------------

func (dlg *OpenFileDialog) InitFilePath(path string) *OpenFileDialog {
	dlg.FilePath = path
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
	dlg.InitialDirPath = path
	return dlg
}

func (dlg *MultOpenFileDialog) InitDirPath(path string) *MultOpenFileDialog {
	dlg.InitialDirPath = path
	return dlg
}

func InitDirPath(path string) *OpenFileDialog {
	dlg := New()
	return dlg.InitDirPath(path)
}
