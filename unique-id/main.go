package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	var num int64 = 0 // Use int64 for atomic operations
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			log.Println("JSON Unmarshal error:", err)
			return err
		}
		// Here we set a unique ID
		nodeId := n.ID()
		body["type"] = "generate_ok"
		num := atomic.AddInt64(&num, 1)
		id := fmt.Sprintf("%s-%d", nodeId, num)
		body["id"] = id
		log.Println("Sending response:", body) // Debugging log
		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
