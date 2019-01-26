## Setup

Start kafka container at *localhost:9092*

`docker-compose up -d kafka`

If this address is already in use, edit the host and port into *docker-compose.yml* file. Also, edit consumer and producer vars for broker list.

__Producer__

`go run cmd/producer/main.go`

Available flags:
  * topic (default: the_topic)
  * msg (default: an important message)
  * key (default: a_key)

Output:
  `Send a message in topic the_topic - partition 0  - offset 1`

__Consumer__

`go run cmd/producer/main.go`

Available flags:
  * topic (default: the_topic)

Output:
  `Got key "a_key" with message "an important message"`

To stop consumer's execution send an interruption singal `Ctrl+C`
