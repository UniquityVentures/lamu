package p_llm_assistant

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/UniquityVentures/lamu/getters"
)

type skillExportJSON struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Content     string   `json:"content"`
	Files       []string `json:"files"`
}

func sanitizeFilename(s string) string {
	var res []rune
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			res = append(res, r)
		} else {
			res = append(res, '_')
		}
	}
	out := string(res)
	if out == "" {
		return "skill"
	}
	return out
}

func handleSkillExport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid skill ID", http.StatusBadRequest)
		return
	}

	db, err := getters.DBFromContext(ctx)
	if err != nil {
		http.Error(w, "No database connection", http.StatusInternalServerError)
		return
	}

	var skill Skill
	if err := db.WithContext(ctx).Preload("Files").First(&skill, uint(id)).Error; err != nil {
		http.Error(w, "Skill not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", sanitizeFilename(skill.Name)))

	zw := zip.NewWriter(w)
	defer zw.Close()

	exportData := skillExportJSON{
		Name:        skill.Name,
		Description: skill.Description,
		Content:     skill.Content,
		Files:       make([]string, 0, len(skill.Files)),
	}

	for _, file := range skill.Files {
		exportData.Files = append(exportData.Files, file.Name)
	}

	indexBytes, err := json.MarshalIndent(exportData, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal index.json", http.StatusInternalServerError)
		return
	}

	indexFile, err := zw.Create("index.json")
	if err != nil {
		http.Error(w, "Failed to create index.json in zip", http.StatusInternalServerError)
		return
	}
	if _, err := indexFile.Write(indexBytes); err != nil {
		http.Error(w, "Failed to write index.json to zip", http.StatusInternalServerError)
		return
	}

	for _, file := range skill.Files {
		dl, err := file.OpenDownload()
		if err != nil {
			continue
		}
		
		fWriter, err := zw.Create(file.Name)
		if err != nil {
			dl.Reader.Close()
			continue
		}
		
		_, _ = io.Copy(fWriter, dl.Reader)
		dl.Reader.Close()
	}
}
