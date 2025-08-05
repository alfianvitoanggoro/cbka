package kafka

import "github.com/hamba/avro"

// SchemaStr is the compiled Avro schema
var SchemaStr = avro.MustParse(`
{
	"type": "record",
	"name": "UserReconcile",
	"fields": [
		{"name": "user_id", "type": "string"},
		{"name": "action", "type": "string"},
		{"name": "timestamp", "type": "string"}
	]
}
`)
