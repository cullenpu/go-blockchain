# Go Blockchain

A project I built to demonstrate how a blockchain works. However, it does not demonstrate the concept of decentralization because blocks are created and stored on a web server.

## Installation

Clone the repo
```
git clone https://github.com/cullenpu/go-blockchain.git
```

Build and start the server
```
go build
./go-blockchain
```

## Usage

The blockchain is created and accessed through the web server. There are a total of 3 endpoints.

`POST /mine` creates a new block and adds it to the chain

```
$ curl -X POST \
    --url http://localhost:8080/mine \
    --header 'content-type: application/json' \
    --data '{"data": "sample data 1"}'

{
  "Index": 1,
  "Timestamp": "2022-03-06T19:58:25.30689-05:00",
  "Hash": "00474c0711cd7ed1c2a423d11c3bec5002a9738a77cf296f86943dd4bff71795",
  "PrevHash": "0",
  "Data": "sample data 1",
  "Pow": 34
}
```

`GET /` returns the blockchain

```
$ curl -X GET --url http://localhost:8080/

[
  {
    "Index": 0,
    "Timestamp": "2022-03-06T19:56:50.772698-05:00",
    "Hash": "0",
    "PrevHash": "",
    "Data": "",
    "Pow": 0
  },
  {
    "Index": 1,
    "Timestamp": "2022-03-06T19:58:25.30689-05:00",
    "Hash": "00474c0711cd7ed1c2a423d11c3bec5002a9738a77cf296f86943dd4bff71795",
    "PrevHash": "0",
    "Data": "sample data 1",
    "Pow": 34
  }
]
```

`GET /:index` returns the block at index `index`

```
$ curl -X GET --url http://localhost:8080/1

{
  "Index": 1,
  "Timestamp": "2022-03-06T19:58:25.30689-05:00",
  "Hash": "00474c0711cd7ed1c2a423d11c3bec5002a9738a77cf296f86943dd4bff71795",
  "PrevHash": "0",
  "Data": "sample data 1",
  "Pow": 34
}
```