// Evermind serves as the core for the home infrastructure I'm designing with my roommate for our IoT devices and homelab services.
package main

import (
	"fmt"
	"math/rand"
)

// Begun 11/22/2026 by guac, with the hopes of building something from scratch.

var startupMessages = [...]string{
	"oh boy, time to serve again",
	"evermind whirring, vibrating, foaming into being",
};

func main() {
	fmt.Println(startupMessages[rand.Intn(len(startupMessages))]);
}