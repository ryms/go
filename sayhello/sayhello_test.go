package sayhello


import "testing"

func TestSayHello(t *testing.T) {
  greeting := SayHello()
  if greeting != "Hello World\n" {
    t.Error("Test failed!")
  }
}
