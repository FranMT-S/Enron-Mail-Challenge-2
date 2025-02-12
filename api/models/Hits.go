package models

type Hit[T any] struct {
	ID     string `json:"_id"`
	Source T      `json:"_source"`
}

type Hits[T any] struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`

		Hits []Hit[T] `json:"hits"`
	} `json:"hits"`
}

// type HitsResponse[T any] struct {
// 	Hits Hits[T] `json:"hits"`
// }
