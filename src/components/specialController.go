package components

import (
	"os"
	"path/filepath"
)

func CancelTypingModal(m model) model {
	m.typingModal.textInput.Blur()
	m.typingModal.open = false
	return m
}

func CancelWarnModal(m model) model {
	m.warnModal.open = false
	return m
}

func CreateItem(m model) model {
	if m.typingModal.itemType == newFile {
		path := m.typingModal.location + "/" + m.typingModal.textInput.Value()
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			OutPutLog("Create item function error", err)
		}
		f, err := os.Create(path)
		if err != nil {
			OutPutLog("Create item function create file error", err)
		}
		defer f.Close()
	} else {
		path := m.typingModal.location + "/" + m.typingModal.textInput.Value()
		err := os.MkdirAll(path, 0755)
		if err != nil {
			OutPutLog("Create item function create folder error", err)
		}
	}
	m.typingModal.open = false
	m.typingModal.textInput.Blur()
	return m
}

func CancelReanem(m model) model {
	panel := m.fileModel.filePanels[m.filePanelFocusIndex]
	panel.rename.Blur()
	panel.renaming = false
	m.fileModel.renaming = false
	m.fileModel.filePanels[m.filePanelFocusIndex] = panel
	return m
}

func ConfirmRename(m model) model {
	panel := m.fileModel.filePanels[m.filePanelFocusIndex]
	oldPath := panel.element[panel.cursor].location
	newPath := panel.location + "/" + panel.rename.Value()

	// Rename the file
	err := os.Rename(oldPath, newPath)
	if err != nil {
		OutPutLog("Confirm function rename error", err)
	}

	m.fileModel.renaming = false
	panel.rename.Blur()
	panel.renaming = false
	m.fileModel.filePanels[m.filePanelFocusIndex] = panel
	return m
}
