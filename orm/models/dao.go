package models

type DataSet struct {
	tableName      struct{}  `sql:"data_set,alias:dataset"`
	Id             string
	MajorLabel     string `sql:"major_label"`
	MajorVersion   int `sql:"major_version"`
	Metadata       string `sql:"metadata"`
	MinorVersion   int `sql:"minor_version"`
	RevisionNotes  string `sql:"revision_notes"`
	RevisionReason string `sql:"revision_reason"`
	S3URL          string `sql:"s3_url"`
	Status         string `sql:"status"`
	Title          string `sql:"title"`
	TotalRowCount  int `sql:"total_row_count"`
	DataResource   string `pg:"data_resource, fk:data_resource"`
}

type DataResource struct {
	tableName      struct{} `sql:"data_resource,alias:dataResource"`
	DataResource  string `sql:"data_resource,pk"`
	ColumnConcept string `sql:"column_concept"`
	Metadata      string `sql:"metadata"`
	RowConcept    string `sql:"row_concept"`
	Title         string `sql:"title"`
}
