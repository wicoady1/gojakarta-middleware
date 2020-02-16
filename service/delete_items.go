package service

import "net/http"

func DeleteItems(r *http.Request) string {
	return `{
		"id": 1,
		"items": [{
				"name": "masker",
				"qty": 1
			},
			{
				"name": "corona",
				"qty": 2
			}
		]
	}
	`
}
