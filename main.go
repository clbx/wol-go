package main

import (
	"log"
	"net/http"

	"github.com/linde12/gowol"
)

type MACAddress [6]byte

func SendPacket(mac string) error {

	if packet, err := gowol.NewMagicPacket(mac); err == nil {
		packet.Send("255.255.255.255")          // send to broadcast
		packet.SendPort("255.255.255.255", "7") // specify receiving port
	}

	return nil
}

func WakeHandler(w http.ResponseWriter, r *http.Request) {
	mac := r.URL.Query().Get("mac")

	log.Printf("Got request: " + mac)

	if mac == "" {
		http.Error(w, "Mac required", http.StatusBadRequest)
		return
	}

	err := SendPacket(mac)
	if err != nil {
		http.Error(w, "failed to send magic packet", http.StatusInternalServerError)
		return

	}

}

func main() {
	http.HandleFunc("/wake", WakeHandler)
	port := "8080"
	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
