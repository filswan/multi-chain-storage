package models

type SourceFileDealFileMap struct {
	SourceFileId int64 `json:"source_file_id"`
	FileIndex    int   `json:"file_index"`
	DealFileId   int64 `json:"source_file_id"`
}
