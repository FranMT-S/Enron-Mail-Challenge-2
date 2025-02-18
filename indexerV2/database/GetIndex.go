package db

import (
	"encoding/json"
	"time"
)

func GetIndexString(indexName string) string {
	index := ZincSearchIndex{
		Name:        indexName,
		StorageType: "disk",
		Mappings: Mappings{
			Properties: Properties{
				Date: Property{
					Type:          "date",
					Format:        time.RFC3339,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Bcc: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Cc: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Body: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				ContentTransferEncoding: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				ContentType: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				From: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				MessageID: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				MimeVersion: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Subject: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				To: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XFileName: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XFolder: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XFrom: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XOrigin: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XTo: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XBcc: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XCc: Property{
					Type:          "text",
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
			},
		},
	}

	// jsonData, err := json.MarshalIndent(index, "", "  ")
	jsonData, err := json.Marshal(index)
	if err != nil {
		panic("cannot unmarshal index")
	}

	return string(jsonData)
}
