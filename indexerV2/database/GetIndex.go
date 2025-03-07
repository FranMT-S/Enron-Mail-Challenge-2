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
				MessageID: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				From: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				From_Sort: Property{
					Type:          Keyword,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				To: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				To_Sort: Property{
					Type:          Keyword,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Subject: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Subject_Sort: Property{
					Type:          Keyword,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Date: Property{
					Type:          Date,
					Format:        time.DateOnly,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				DateTime: Property{
					Type:          Date,
					Format:        time.RFC3339,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Bcc: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Cc: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				Body: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				ContentTransferEncoding: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				ContentType: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				MimeVersion: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XFileName: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XFolder: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XFrom: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XOrigin: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XTo: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XBcc: Property{
					Type:          Text,
					Index:         true,
					Store:         true,
					Sortable:      true,
					Aggregratable: true,
					Highlightable: true,
				},
				XCc: Property{
					Type:          Text,
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
