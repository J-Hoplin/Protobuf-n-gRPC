package main

/*
Author: Hoplin
*/

import (
	"errors"
	"fmt"
	"log"
	pb "protoexample/proto"
	"reflect"

	"google.golang.org/protobuf/proto"
)

// Basic
func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A Name",
		SampleLists: []int32{1, 2, 3, 4, 5},
	}
}

// Complex
func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy2{Id: 42, Name: "My Name"},
		MultipleDummies: []*pb.Dummy2{
			{Id: 43, Name: "Name2"},
			{Id: 44, Name: "Name3"},
			{Id: 45, Name: "Name4"},
		},
	}
}

// Enumeration
func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

// Handle oneof
func doOneOf(message interface{}) (err error) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		err = errors.New(fmt.Sprintf("Invalid type: %v", x))
	}
	return
}

// Handle Map
func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"test":  {Id: 42},
			"test2": {Id: 43},
		},
	}
}

// Read, Write from dis
func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := new(pb.Simple)
	readFromFile(path, message)
}

// Handle with JSON
// 1. Call the to JSON Function
func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)
	fmt.Println(jsonString)
	return jsonString
}

// 2. Call the from JSON Function
// Why using reflect.Type? jsonstring을 받았을때 어떤 타입인지 유추할 수 없기때문
func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	// Golang interface conversion: https://go.dev/doc/effective_go
	message, ok := reflect.New(t).Interface().(proto.Message)
	if !ok {
		log.Fatalln("Invalid type")
		return nil
	}
	fromJSON(jsonString, message)
	return message
}

func main() {
	// Simple protobuf
	fmt.Println(doSimple())

	// Multiple message and reference message
	fmt.Println(doComplex())

	// Protobuf Enum
	fmt.Println(doEnum())

	// Handling Oneof Type
	doOneOf(&pb.Result_Id{Id: 42})
	doOneOf(&pb.Result_Message{Message: "a message"})

	// Handling Map Type
	fmt.Println(doMap())

	// Marshal protobuf to binary, save to file as '.bin' and unmarshal binary array
	doFile(doSimple())

	// Golang protobuf struct to JSON string
	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	jsonString = doToJSON(doComplex())
	message = doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))
	fmt.Println(jsonString)
	fmt.Println(message)

	// JSON string to proto.Message
	simple := doFromJSON(`{"id":42,"unknown": "test"}`, reflect.TypeOf(pb.Simple{}))

	// protoreflect.Protomessage to proto.Message implemented object
	simple2, ok := simple.ProtoReflect().Interface().(*pb.Simple)
	if ok {
		fmt.Println(simple2.SampleLists)
	} else {
		fmt.Println("Fail to unmarshal")
	}
}
