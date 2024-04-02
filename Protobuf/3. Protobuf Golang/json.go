package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Function1: create a JSON String from a message
// Docs: https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#MarshalOptions
func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true, // Pretty print

	}
	out, err := option.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return string(out)
}

// Function2: from JSON which will take a string from object
// Docs: https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson#UnmarshalOptions
func fromJSON(in string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true, // Unknown fields and enum values will be ignored
	}
	if err := option.Unmarshal([]byte(in), pb); err != nil {
		log.Fatalln("Could not unmarshal from JSON")
	}
}
