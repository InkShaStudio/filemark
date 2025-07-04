package files

import (
	"fmt"
	"os"
	"path"

	"github.com/InkShaStudio/filemark/pkg/storage"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const SELECT_COLOR = "#4F8A95"

type FileType int

const (
	FileTypeFile FileType = iota
	FileTypeDir
)

func (t FileType) String() string {
	if t == FileTypeFile {
		return "file"
	}
	return "dir"
}

type FileInfo struct {
	Size int      `json:"size"`
	Name string   `json:"name"`
	Path string   `json:"path"`
	Type FileType `json:"type"`
	Dir  string   `json:"dir"`
	Ext  string   `json:"ext"`
}

type FileInfoFilter struct {
}

type FileInfoList struct {
	Files      *[]FileInfo
	Cursor     int
	Filter     FileInfoFilter
	CurentPath string
	Selected   int
}

func NewFileInfo(file_path string) FileInfo {
	info, err := os.Stat(file_path)
	t := FileTypeFile

	if err != nil {
		fmt.Println("get file info error: ", err)
	}

	if info.IsDir() {
		t = FileTypeDir
	}

	return FileInfo{
		Size: int(info.Size()),
		Name: info.Name(),
		Path: file_path,
		Type: t,
		Dir:  path.Dir(file_path),
		Ext:  path.Ext(file_path),
	}
}

func NewFileInfoFilter() FileInfoFilter {
	return FileInfoFilter{}
}

func NewFileInfoList(filter FileInfoFilter, files *[]FileInfo) FileInfoList {
	return FileInfoList{
		Files:      files,
		Cursor:     0,
		Filter:     filter,
		CurentPath: "",
		Selected:   -1,
	}
}

func ReadPath(p string) []FileInfo {
	list := make([]FileInfo, 0)

	info, err := os.Stat(p)

	if err != nil {
		println("read path error: ", err)
	}

	if info.IsDir() {
		dirs, _ := os.ReadDir(p)
		for _, item := range dirs {
			list = append(list, NewFileInfo(path.Join(p, item.Name())))
		}
	} else {
		list = append(list, NewFileInfo(p))
	}

	for _, item := range list {
		storage.InsertFile(item.Path)
	}

	return list
}

func (f *FileInfoFilter) Filter(item *FileInfo) bool {
	return false
}

func (f *FileInfoList) Init() tea.Cmd {
	return nil
}

func (f *FileInfoList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		files := *f.Files
		switch key {
		case "esc":
			return f, tea.Quit
		case "up", "k":
			if f.Cursor > 0 {
				f.Cursor--
			} else {
				f.Cursor = len(files) - 1
			}
		case "down", "j":
			if f.Cursor < len(files)-1 {
				f.Cursor++
			} else {
				f.Cursor = 0
			}
		case " ":
			f.Selected = f.Cursor
		case "enter", "right":
			item := files[f.Cursor]
			if item.Type == FileTypeDir {
				newFiles := ReadPath(item.Path)
				f.Files = &newFiles
				f.Cursor = 0
				f.CurentPath = item.Path
				f.Selected = -1
			}
		case "backspace", "left":
			father := path.Dir(f.CurentPath)
			newFiles := ReadPath(father)
			f.Files = &newFiles
			f.CurentPath = father
			f.Cursor = 0
			f.Selected = -1
		}
	}

	return f, nil
}

func (f *FileInfoList) View() string {
	view := "\n"
	files := *f.Files
	fc := lipgloss.NewStyle().Foreground(lipgloss.Color(SELECT_COLOR))

	if f.Selected != -1 {
		view += fmt.Sprintf("\n%s", files[f.Cursor].Path)

	} else {
		for i, item := range files {
			if f.Filter.Filter(&item) {
				continue
			}

			cursor := " "

			if f.Cursor == i {
				cursor = ">"
			}

			name := item.Name
			if item.Type == FileTypeDir {
				name += "/"
			}

			line := fmt.Sprintf("\n%s %s", cursor, name)

			if f.Selected == i {
				view += fc.Render(line)
			} else {
				view += line
			}
		}
	}

	return view
}

func (f *FileInfoList) Run() {
	if _, err := tea.NewProgram(f).Run(); err != nil {
		fmt.Println("raw mode error: ", err)
		os.Exit(1)
	}
}
