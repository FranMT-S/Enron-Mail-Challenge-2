package db

import "fmt"

func GetIndexString(indexName string) string {

	return fmt.Sprintf(`{
		"name": "%v",
		"storage_type": "disk",
		"mappings": {
		"properties": {
		  "Date": {
			"type": "date",
			"format":"2006-01-02T15:04:05Z",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Bcc": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Cc": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Content": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Content_Transfer_Encoding": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Content_Type": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "From": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Message_ID": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Mime_Version": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "Subject": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "To": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_FileName": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_Folder": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_From": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_Origin": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_To": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_bcc": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "X_cc": {
			"type": "text",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  },
		  "_id": {
			"type": "keyword",
			"index": true,
			"store": false,
			"sortable": true,
			"aggregatable": true,
			"highlightable": true
		  }
		}
		}
	}`, indexName)
}
