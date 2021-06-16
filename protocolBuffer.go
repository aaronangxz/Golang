package main

type Item struct {
	itemid          string
	shopid          string
	price           int
	discountedPrice int
}

func protocolBuffer() {
	//Require protobuf installation,
	// *.proto, *.pb.go files with struct inside

	//Assuming TCP connection is established

	//To Receive
	dataRec := make([]byte, 4096)
	lengthRec, err := connection.Read(dataRec)
	checkerror(err)

	messagePB := pb.Message()
	err = proto.Unmarshall(dataRec[:lengthRec], &messagePB)
	checkerror(err)

	//To Send
	itemProto := pb.Item{Itemid: "11111", Shopid: "22222", price: 5, discountedPrice: 4}
	dataSend, err := proto.Marshal(&itemProto)
	checkerror(err)
	lengthSend, err := connection.Write(dataSend)
	checkerror(err)
}
