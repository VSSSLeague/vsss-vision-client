package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/VSSSLeague/vsss-vision-client/pkg/client"
	"github.com/VSSSLeague/vsss-vision-client/pkg/vision"
	"github.com/gobuffalo/packr"
)

var address = flag.String("address", ":8082", "The address on which the UI and API is served, default: :8082")
var visionAddress = flag.String("visionAddress", "224.0.0.1:10002", "The multicast address of ssl-vision, default: 224.0.0.1:10002")
var skipInterfaces = flag.String("skipInterfaces", "", "Comma separated list of interface names to ignore when receiving multicast packets")
var verbose = flag.Bool("verbose", false, "Verbose output")

func main() {
	flag.Parse()

	// Setup vision
	setupVisionClient()
	setupUi()

	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setupVisionClient() {
	receiver := vision.NewReceiver()
	skipIfis := parseSkipInterfaces()

	// Setup publisher
	publisher := client.NewPublisher()
	publisher.DetectionProvider = receiver.CombinedDetectionFrames
	publisher.GeometryProvider = geometryProvider(receiver)
	http.HandleFunc("/api/vision", publisher.Handler)

	// Set configs to MulticastServer
	receiver.MulticastServer.SkipInterfaces = skipIfis
	receiver.MulticastServer.Verbose = *verbose

	// Start MulticastServer
	receiver.Start(*visionAddress)
}

func setupUi() {
	box := packr.NewBox("../dist")

	withResponseHeaders := func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Set some header.
			w.Header().Add("Access-Control-Allow-Origin", "*")
			// Serve with the actual handler.
			h.ServeHTTP(w, r)
		}
	}

	http.Handle("/", withResponseHeaders(http.FileServer(box)))
	if box.Has("index.html") {
		log.Printf("UI is available at http://%v", *address)
	} else {
		log.Print("Backend-only version started. Run the UI separately or get a binary that has the UI included")
	}
}

func geometryProvider(receiver *vision.Receiver) func() *vision.Field {
	return func() *vision.Field {
		geometry := receiver.CurrentGeometry()
		if geometry == nil {
			return defaultGeometry()
		}
		return geometry
	}
}

func defaultGeometry() (g *vision.Field) {
	// Create new field
	g = new(vision.Field)

	// Set default values (VSS 3v3 field)
	g.Width = 1300
	g.Length = 1500
	g.GoalDepth = 100
	g.GoalWidth = 400

	return
}

func parseSkipInterfaces() []string {
	return strings.Split(*skipInterfaces, ",")
}
