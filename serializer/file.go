package serializer

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"os"
)

func WriteProtoBufToBinaryFile(message proto.Message, fileName string) error {
	marshal, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("unable to serialize the message: %w", err)
	}
	err = os.WriteFile(fileName, marshal, 0644)
	if err != nil {
		return fmt.Errorf("unable to write the message ti the file : %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(fileName string, message proto.Message) error {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("unable to read the message: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("unable to deserialize the message: %w", err)
	}
	return nil
}

func ProtoButToJson(message proto.Message, filName string) error {

	jsonBytes, err := formatToJson(message)
	if err != nil {
		return fmt.Errorf("unable to format to JSON.: %w", err)
	}

	err = os.WriteFile(filName, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to JSON file.: %w", err)
	}

	return nil

}

func formatToJson(message proto.Message) ([]byte, error) {
	marshalOptions := protojson.MarshalOptions{
		Indent:         " ",
		UseEnumNumbers: false,
		UseProtoNames:  true,
	}
	return marshalOptions.Marshal(message)
}
