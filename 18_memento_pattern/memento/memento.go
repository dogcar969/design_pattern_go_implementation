package memento

import "fmt"

type File struct {
	Content   FileContent
	saveStack memento
}

type FileContent struct {
	content string
}

type BackUp struct {
	content string
}

type memento struct {
	versionCount uint // 现有版本的数量
	versionIndex uint // 版本现在的位置
	saves        []BackUp
}

func (file *FileContent) GetContent() string {
	return file.content
}

func (file *FileContent) SetContent(newContent string) {
	file.content = newContent
}

func (file *FileContent) save() BackUp {
	// 生成一个备份
	return BackUp{content: file.content}
}

func (file *FileContent) load(backUp BackUp) {
	// 从备份中加载内容
	file.content = backUp.content
}

func (file *File) Save() {
	// 将内容转化为BackUp然后保存到saves[versionIndex]中
	file.saveStack.versionIndex++
	save := file.Content.save()
	if uint(len(file.saveStack.saves)) <= file.saveStack.versionIndex {
		file.saveStack.saves = append(file.saveStack.saves, save)
		file.saveStack.versionCount = file.saveStack.versionIndex + 1
		return
	}
	file.saveStack.saves[file.saveStack.versionIndex] = save
	file.saveStack.versionCount = file.saveStack.versionIndex + 1
}

func (file *File) Load() {
	// 加载最近的版本
	file.Content.load(file.saveStack.saves[len(file.saveStack.saves)-1])
}

func (file *File) Undo() {
	// 加载比现在版本更靠前一个的版本
	if file.saveStack.versionIndex > 1 {
		file.saveStack.versionIndex--
		file.Content.load(file.saveStack.saves[file.saveStack.versionIndex - 1])
		
	}
}

func (file *File) Redo() {
	// 加载比现在版本靠后一个的版本
	if file.saveStack.versionIndex+1 < file.saveStack.versionCount {
		file.Content.load(file.saveStack.saves[file.saveStack.versionIndex])
		file.saveStack.versionIndex++
	}
}


func (file *File) ShowRecord() {
	for _,save := range file.saveStack.saves {
		fmt.Println(save.content)
	}
}