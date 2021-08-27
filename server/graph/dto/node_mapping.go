package dto

import "github.com/aligator/godrop/server/repository/model"

func FileNodeFromModel(fileNode model.FileNode) FileNode {
	result := FileNode{
		ID:          fileNode.ID,
		Name:        fileNode.Name,
		Description: fileNode.Description,
		IsFolder:    fileNode.IsFolder,
		MimeType:    &fileNode.MimeType,
	}

	if fileNode.MimeType == "" {
		result.MimeType = nil
	}

	if fileNode.Children != nil {
		for _, child := range fileNode.Children {
			result.Children = append(result.Children, FileNodeFromModel(child))
		}
	}

	return result
}

func CreateFileNodeFromDTO(FileNode CreateFileNode) model.CreateFileNode {
	result := model.CreateFileNode{
		Name:        FileNode.Name,
		Description: FileNode.Description,
		IsFolder:    FileNode.IsFolder,
	}

	if FileNode.MimeType == nil {
		result.MimeType = ""
	} else {
		result.MimeType = *FileNode.MimeType
	}

	if FileNode.File == nil {
		result.File = []byte{}
	} else {
		result.File = []byte(*FileNode.File)
	}

	return result
}
