# Middleware - GoJakarta (20200220)

## Introduction

This repository is made by [wicoady1](https://github.com/wicoady1), for GoJakarta workshop purpose on February 20th, 2020.  
Presentation File will be included soon.

## Repository Structure
```
|-controller
|-middleware
|-service
|-validator
.gitignore
main.go
```

**Controller**  
Contains functions, which directly serve received network HTTP Request contents and write the response in HTTP Response.

**Service**  
Functions that server business logics to process request, which is received from `controller` and return proper response to `controller`.

**Middleware**  
Intermediary functions, bridge between router in `main.go` and `controller`

**Validator**  
Contains functions for business logic validation purpose. In this case, for `middleware` folder purpose.

## How to Run
- Make sure your Golang version is 1.13.4 or above
- `go get && go build && ./demo`
- Use postman or cURL (in your local) with this URL `localhost:6969/v1/cart/1`
- If it returns following response, then it is working properly
```
{
	"id": 1,
	"items": [
		{
			"name": "masker",
			"qty": 1
		},
		{
			"name": "corona",
			"qty": 2
		}
	]
}
```