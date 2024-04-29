# Concurrent Prime Checker API

The Concurrent Prime Checker API is a GoLang project aimed at providing a practical learning experience in concurrent and asynchronous programming using Goroutines and channels. The project implements a worker pool layered with a REST API for checking prime numbers concurrently. Users can submit tasks, monitor task statuses, and retrieve results through the API.

## Getting Started

### Prerequisites
1. GoLang installed on your machine
2. Basic understanding of GoLang programming concepts

### Installation

1. Clone the repository:

```bash
git clone https://github.com/hasssanezzz/rest-workers.git
```

2. Navigate to the project directory:

```bash
cd rest-workers
```

3. Build and run the project:

```bash
go build
./rest-workers -a :3030 -w 5
```

The API server will start listening on the specified address (-a) with the specified number of workers (-w).

### API Endpoints

* List Tasks: GET /api/v0/task
    * Retrieve a list of all tasks.
* Get Task Details: GET /api/v0/task/{id}
    * Retrieve details of a specific task by ID.
* Create Task: POST /api/v0/task
    * Create a new task for prime number checking.

### Request Payload

To create a new task, send a POST request to /api/v0/task with the following JSON payload:

```json
{
  "value": "2024444666688888688681"
}
```
### Response Format

List Tasks Response:
```json
[
	{
		"id": 1,
		"payload": {
			"number": 7999909
		},
		"result": {
			"result": true
		},
		"status": "Finished",
		"placedAt": "2024-04-29T16:30:10.337569715+03:00",
		"startedAt": "2024-04-29T16:30:10.337688913+03:00",
		"finishedAt": "2024-04-29T16:30:10.338323058+03:00"
	},
	{
		"id": 0,
		"payload": {
			"number": 1235607889460606009419
		},
		"result": {
			"result": false
		},
		"status": "Working",
		"placedAt": "2024-04-29T16:29:42.962283309+03:00",
		"startedAt": "2024-04-29T16:29:42.962331781+03:00",
		"finishedAt": "0001-01-01T00:00:00Z"
	}
]
```

Get Task Details Response:

```json
{
	"id": 0,
	"payload": {
		"number": 1235607889460606009419
	},
	"result": {
		"result": true
	},
	"status": "Finished",
	"placedAt": "2024-04-29T16:29:42.962283309+03:00",
	"startedAt": "2024-04-29T16:29:42.962331781+03:00",
	"finishedAt": "0001-01-01T00:00:00Z"
}
```

Create Task Response:

```json
{
	"id": 0,
	"payload": {
		"number": 1235607889460606009419
	},
	"result": {
		"result": false
	},
	"status": "Waiting",
	"placedAt": "2024-04-29T16:29:42.962283309+03:00",
	"startedAt": "2024-04-29T16:29:42.962331781+03:00",
	"finishedAt": "0001-01-01T00:00:00Z"
}
```
Configuration

* `-a`: Specify the listen address for the API server (default: :3030).
* `-w`: Specify the number of workers in the worker pool (default: 5).

## Learning Objectives

- Understand Goroutines and channels for concurrent programming.
- Implement a scalable worker pool for parallel task processing.
- Develop a REST API for asynchronous task submission and status tracking.
- Explore error handling, graceful shutdown, and best practices in concurrent programming.

## Contributing

Contributions to the project are welcome! Please fork the repository, make your changes, and submit a pull request.