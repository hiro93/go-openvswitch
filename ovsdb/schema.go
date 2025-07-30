package ovsdb

// DatabaseSchema
//
// An Open vSwitch configuration database consists of a set of tables,
// each of which has a number of columns and zero or more rows.  A
// schema for the database is represented by <database-schema>, as
// described below.
//
//	<database-schema>
//	   A JSON object with the following members:
//
//	        "name": <id>                            required
//	        "version": <version>                    required
//	        "cksum": <string>                       optional
//	        "tables": {<id>: <table-schema>, ...}   required
//
// The "name" identifies the database as a whole.  It must be
// provided to most JSON-RPC requests to identify the database being
// operated on.
//
// The "version" reports the version of the database schema.  It is
// REQUIRED to be present.  Open vSwitch semantics for "version" are
// described in [DB-SCHEMA].  Other schemas may use it differently.
// The "cksum" optionally reports an implementation-defined checksum
// for the database schema.  Its use is primarily as a tool for
// schema developers, and clients SHOULD ignore it.
type DatabaseSchema struct {
	Name     string                 `json:"name"`
	Version  string                 `json:"version"`
	Checksum string                 `json:"cksum"`
	Tables   map[string]TableSchema `json:"tables"`
}

// TableSchema
//
//	A JSON object with the following members:
//
//	   "columns": {<id>: <column-schema>, ...}   required
//	   "maxRows": <integer>                      optional
//	   "isRoot": <boolean>                       optional
//	   "indexes": [<column-set>*]                optional
type TableSchema struct {
	Columns map[string]ColumnSchema `json:"columns"`
	MaxRows int                     `json:"maxRows"`
	IsRoot  bool                    `json:"isRoot"`
	Indexes []ColumnSet             `json:"indexes"`
}

// ColumnSchema
//
//	<column-schema>
//	   A JSON object with the following members:
//
//	      "type": <type>                            required
//	      "ephemeral": <boolean>                    optional
//	      "mutable": <boolean>                      optional
type ColumnSchema struct {
	Type      interface{} `json:"type"`
	Ephemeral bool        `json:"ephemeral"`
	Mutable   bool        `json:"mutable"`
}

type ColumnSet []string
