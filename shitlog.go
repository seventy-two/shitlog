package shitlog

import (
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
)

// LogProto dumps json proto representations to standard out
func LogProto(p proto.Message) {
	m := jsonpb.Marshaler{}
	s, _ := m.MarshalToString(p)
	fmt.Println(s)
}
