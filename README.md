## Getting started

### Run frontends

```bash
cd react-v
npm i
npm run dev

cd vue-v
npm i
npm run dev
```

### Run the backend

```bash
go run main.go
```

#### Available Routes

| Path      | Description                                                                                 |
| --------- | ------------------------------------------------------------------------------------------- |
| GET /     | Get all new pastes                                                                          |
| GET /{id} | Get a specific paste                                                                        |
| POST /    | Create a new paste, expects a request: `{ "body": "sample text" }`                          |
| GET /new  | GET version of creating new pastes, body is passed in as a query string: `/new?body=sample` |
