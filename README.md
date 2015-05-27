# arraycoder
The encoder and decoder of string array for shopify's sarama.  
Simple, easy and fast to use.

####Install:  
`go get github.com/LeeQY/arraycoder`

####Encode
```Go
s := arraycoder.ArrayEncoder{"hello", "你好"}

//sarama producer
msg := &sarama.ProducerMessage{Topic: "test", Value: s}
```

####Decode
```Go
arr := arraycoder.Decode(b)
```