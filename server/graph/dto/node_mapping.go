package dto

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/aligator/godrop/server/repository/model"
)

func FileNodeFromModel(fileNode model.FileNode) FileNode {
	result := FileNode{
		ID:          fileNode.ID,
		Name:        fileNode.Name,
		Description: fileNode.Description,
		IsFolder:    fileNode.IsFolder,
		MimeType:    &fileNode.MimeType,
		Size:        fileNode.Size,
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

func CreateFileNodeFromDTO(meta CreateFileNode, file *graphql.Upload) model.CreateFileNode {
	result := model.CreateFileNode{
		Path:        meta.Path,
		Name:        meta.Name,
		Description: meta.Description,
		IsFolder:    meta.IsFolder,
	}

	if file != nil {
		result.File = file.File
	}

	if meta.MimeType == nil {
		result.MimeType = ""
	} else {
		result.MimeType = *meta.MimeType
	}

	return result
}
