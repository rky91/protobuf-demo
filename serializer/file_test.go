package serializer

import (
	"github.com/stretchr/testify/require"
	pb "protobuf-demo/pb/proto"
	"protobuf-demo/sample"
	"testing"
)

func TestProtoBufSerialize(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	jsonFile := "../temp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := WriteProtoBufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.Equal(t, laptop1.Name, laptop2.Name)

	err = ProtoButToJson(laptop1, jsonFile)
	require.NoError(t, err)

}
