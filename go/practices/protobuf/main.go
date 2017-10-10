package main

import (
	"log"

	"github.com/golang/protobuf/proto"
	"practices/protobuf/example"
)

func main() {
	test := &example.Test {
		Label: proto.String("hello"),
		Type: proto.Int32(17),
		Reps: []int64{1, 2, 3},
		Optionalgroup: &example.Test_OptionalGroup {
			RequiredField: proto.String("good byte"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	log.Println(test)
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatal("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
}
