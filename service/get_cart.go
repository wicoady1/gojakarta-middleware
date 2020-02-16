package service

import "net/http"

func GetCart(r *http.Request) string {
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
