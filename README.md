# Boilerplate for Golang service and Hasura GraphQL Engine's Event Triggers

## Events Trigger and Serverless functions architecture

![Architecture diagram](assets/basic-event-triggers-arch-diagram.png)

## Setup Postgres + Hasura GraphQL engine
This boilerplate code assumes that you already have a HGE instance running.

If not you can visit the [docs](https://docs.hasura.io/1.0/graphql/manual/getting-started/index.html) and setup Postgres + HGE.

## Documented examples

* soon :)

### Test the trigger

Goto `Data` tab on Hasura console, browse to `note` table and insert a new row.
Once a new row is inserted, goto `Events` table and `note_trigger`. Checkout the
request and response body for the processed events.

Trigger payload (request):
```json
{
    "event": {
        "op": "INSERT",
        "data": {
            "old": null,
            "new": {
                "text": "new-entry",
                "id": 1
            }
        }
    },
    "created_at": "2018-10-01T17:21:03.76895Z",
    "id": "b30cc7e6-9f3b-48ee-9a10-16cce333df40",
    "trigger": {
        "name": "note_trigger",
        "id": "551bd6a9-6f8b-4644-ba7f-80c08eb9227b"
    },
    "table": {
        "schema": "public",
        "name": "note"
    }
}
```

Webhook response:
```json
{
    "message": "got 'b30cc7e6-9f3b-48ee-9a10-16cce333df40' for 'INSERT' operation on 'note' 
}
```

## Directory Structure

The following is a representative tree of the folder structure:

    .
    ├── aws-lambda
    |   |── README.md
    |   |── routes
    |   |   |── routes.go
    |   |── handlers
    |   |   |── article.go
    |   |── main.go