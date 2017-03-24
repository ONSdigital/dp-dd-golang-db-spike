package dataset

import (
	"database/sql"
)

const INSERT_STMT = "INSERT INTO data_set (id, major_version, major_label, metadata, minor_version," +
	" revision_notes, revision_reason, s3_url, status, title, total_row_count) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"

const SELECT_ALL_QUERY = "SELECT * FROM data_set ds"

type DAO struct {
	Id             string `json:"id"`
	MajorLabel     string `json:"majorLabel"`
	MajorVersion   int64 `json:"majorVersion"`
	Metadata       string `json:"metadata"`
	MinorVersion   int64 `json:"minorVersion"`
	RevisionNotes  string `json:"revisionNotes"`
	RevisionReason string `json:"revisionReason"`
	S3URL          string `json:"s3URL"`
	Status         string `json:"status"`
	Title          string `json:"title"`
	TotalRowCount  int64 `json:"totalRowCount"`
	DataResource   string `json:"dataResource"`
}

func MapFromRow(r *sql.Rows) *DAO {
	id := &sql.NullString{}
	majorLabel := &sql.NullString{}
	majorVersion := &sql.NullInt64{}
	metadata := &sql.NullString{}
	minorVersion := &sql.NullInt64{}
	revisionNotes := &sql.NullString{}
	revisionReason := &sql.NullString{}
	s3URL := &sql.NullString{}
	status := &sql.NullString{}
	title := &sql.NullString{}
	totalRowCount := &sql.NullInt64{}
	dataResource := &sql.NullString{}

	if err := r.Scan(id, majorVersion, majorLabel, metadata,
		minorVersion, revisionReason, revisionNotes, s3URL,
		status, title, totalRowCount, dataResource); err != nil {
		panic(err.Error())
	}
	return &DAO{
		Id:             strDefaultIfNull(id),
		MajorLabel:     strDefaultIfNull(majorLabel),
		MajorVersion:   intDefaultIfNull(majorVersion),
		Metadata:       strDefaultIfNull(metadata),
		MinorVersion:   intDefaultIfNull(minorVersion),
		RevisionNotes:  strDefaultIfNull(revisionNotes),
		RevisionReason: strDefaultIfNull(revisionReason),
		S3URL:          strDefaultIfNull(s3URL),
		Status:         strDefaultIfNull(status),
		Title:          strDefaultIfNull(title),
		TotalRowCount:  intDefaultIfNull(totalRowCount),
		DataResource:   strDefaultIfNull(dataResource),
	}
}

func strDefaultIfNull(sqlVal *sql.NullString) string {
	if sqlVal.Valid {
		return sqlVal.String
	}
	return ""
}

func intDefaultIfNull(sqlVal *sql.NullInt64) int64 {
	if sqlVal.Valid {
		return sqlVal.Int64
	}
	return 0
}
