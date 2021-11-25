package models

// CertModel Structure of the certificates table
type Certificate struct {
	ID             string   `json:"id"`
	ServerHostname string   `json:"server_hostname"`
	Origin         string   `json:"origin"`
	tableName      struct{} `pg:"certificates"`
}
