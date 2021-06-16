package main

func protocolBuffer() {
	//To Receive
	dataRec := make([]byte, 4096)
	lengthRec, err := connection.Read(dataRec)
	checkerror(err)

	messagePB := pb.Message()
	err = proto.Unmarshall(dataRec[:lengthRec])
}
