# Learning Protocol Buffer and gRPC

---

- Author: Hoplin

---

## Protocol Buffer

- [Protocol Buffer Document](https://protobuf.dev)

1. [Protocol Buffer Theory](./Protobuf/1.%20Protobuf%20Theory/)

   1. [Protocol Buffer Basic](./Protobuf/1.%20Protobuf%20Theory/1.%20Protocol%20Buffer%20Basic.md)

   2. [Protocol Buffer Scalar Type](./Protobuf/1.%20Protobuf%20Theory/2.%20Protocol%20Buffer%20Data%20Type%20-%20Scalar%20Type.md)

   3. [Protocol Buffer Repeated, Oneof, Map Type](./Protobuf/1.%20Protobuf%20Theory/3.%20Protocol%20Buffer%20Data%20Type%20-%20Repeated,Oneof,Map%20Type.md)

   4. [Protocol Buffer Enumeration & Comment](./Protobuf/1.%20Protobuf%20Theory/4.%20Protocol%20Buffer%20Enumeration%20&%20Comment.md)

   5. [Protocol Buffer Nested Message](./Protobuf/1.%20Protobuf%20Theory/5.%20Protocol%20Buffer%20Nested%20Message.md)

   6. [Protocol Buffer Import](./Protobuf/1.%20Protobuf%20Theory/6.%20Protocol%20Buffer%20Import.md)

   7. [Advanced - What if schema changes?](./Protobuf/1.%20Protobuf%20Theory/7.%20Advanced%20-%20What%20if%20schema%20changes.md)

   8. [Protocol Buffer Service](./Protobuf/1.%20Protobuf%20Theory/8.%20Protocol%20Buffer%20Service.md)

2. [Protocol Buffer Excersice](./Protobuf/2.%20Protobuf%20Excercise/)

3. [Protocol Buffer Golang](./Protobuf/3.%20Protobuf%20Golang/)

## gRPC Excersice

### Change client commands

Remove/Add annotation in `excersice/client/main.go`, which service you want to try with.

```go
func main() {

   ...


	// Invoke Unary Sum

	//InvokeUnarySum(c)
	//InvokeClientStream(c)
	//InvokeServerStream(c)
	InvokeBidrectionalStream(c)
}
```

**You need to execute after compile each time you modify main.go**

### Compile and execute

1. Compile proto and rpc services

```bash
cd excersice

make excersice
```

2. Prepare 2 terminal session and execute server, client each

```bash
# session 1
./bin/excersice/server

# session 2
./bin/excersice/client
```

- Unary Communication
- Client streaming
- Server streaming
- Bi-Directional stream
