package vision

import (
	"log"
	"sync"
	"time"

	"github.com/VSSSLeague/vsss-vision-client/pkg/net"
	"google.golang.org/protobuf/proto"
)

type Receiver struct {
	FieldFrames     map[int]*Frame
	FieldGeometry   *Field
	receivedTimes   map[int]time.Time
	MulticastServer *net.MulticastServer
	mutex           sync.Mutex
	ConsumeFrames   func(frame *Frame)
	ConsumeGeometry func(frame *Field)
}

func NewReceiver() (r *Receiver) {
	r = new(Receiver)
	r.FieldFrames = map[int]*Frame{}
	r.receivedTimes = map[int]time.Time{}
	r.FieldGeometry = new(Field)
	r.MulticastServer = net.NewMulticastServer(r.consumeMessage)
	r.ConsumeFrames = func(*Frame) {}
	r.ConsumeGeometry = func(*Field) {}
	return
}

func (r *Receiver) Start(multicastAddress string) {
	r.MulticastServer.Start(multicastAddress)
}

func (r *Receiver) consumeMessage(data []byte) {
	message, err := parseEnvironmentPacket(data)

	if err != nil {
		log.Print("Could not parse message: ", err)
		return
	}

	r.mutex.Lock()

	if message.Frame != nil {
		/// TODO: check this later
		// Simulator doesn't send camera data, so, by now, it will be always 0
		camId := 0

		// Set frame, set received time and cast consumeFrame to behave it
		r.FieldFrames[camId] = message.Frame
		r.receivedTimes[camId] = time.Now()
		r.ConsumeFrames(message.Frame)
	}

	if message.Field != nil {
		r.FieldGeometry = message.Field
		r.ConsumeGeometry(message.Field)
	}

	r.mutex.Unlock()
}

func (r *Receiver) CombinedDetectionFrames() (f *Frame) {
	r.mutex.Lock()

	// Create new frame
	f = new(Frame)
	f.Ball = new(Ball) // Simulator send only one ball... it will change in the vision unified software?
	f.RobotsYellow = make([]*Robot, 0)
	f.RobotsBlue = make([]*Robot, 0)

	// Cleanup old detections (use timestamp for this)
	r.cleanupDetections()

	// For each frame that survived the cleanup, append into frame list
	for _, b := range r.FieldFrames {
		f.Ball = b.Ball // Simulator send only one ball... it will change in the vision unified software?
		f.RobotsYellow = append(f.RobotsYellow, b.RobotsYellow...)
		f.RobotsBlue = append(f.RobotsBlue, b.RobotsBlue...)
	}

	r.mutex.Unlock()

	return
}

func parseEnvironmentPacket(data []byte) (message *Environment, err error) {
	message = new(Environment)
	err = proto.Unmarshal(data, message)
	return
}

func (r *Receiver) cleanupDetections() {
	for camId, t := range r.receivedTimes {
		if time.Now().Sub(t) > time.Second {
			delete(r.receivedTimes, camId)
			delete(r.FieldFrames, camId)
		}
	}
}

func (r *Receiver) CurrentDetections() (result map[int]Frame) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	result = map[int]Frame{}
	for id, frame := range r.FieldFrames {
		result[id] = *frame
	}
	return
}

func (r *Receiver) CurrentGeometry() (geometry *Field) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if r.FieldGeometry != nil {
		geometry = new(Field)
		*geometry = *r.FieldGeometry
	}
	return
}
