package internal

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"time"

	"lecture9.demo/grpc/stream"
)

type StreamingService struct {
	rnd *rand.Rand
	stream.UnimplementedIntStreamServer
}

func NewStreamingService() *StreamingService {
	return &StreamingService{
		rnd: rand.New(rand.NewSource(time.Now().Unix())),
	}
}

func (s *StreamingService) ServerSideStream(req *stream.StreamRequest, str stream.IntStream_ServerSideStreamServer) error {
	for i := 0; i < int(req.GetCount()); i++ {
		time.Sleep(200 * time.Millisecond)
		if err := str.Send(&stream.Value{Value: s.rnd.Int63()}); err != nil {
			return err
		}
	}
	return nil
}

func (s *StreamingService) ClientSideStream(str stream.IntStream_ClientSideStreamServer) error {
	defer fmt.Println("stream finished")
	for {
		val, err := str.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		fmt.Println(val)
	}
}
func (s *StreamingService) BidirectionalStream(str stream.IntStream_BidirectionalStreamServer) error {
	ch := make(chan int64)
	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		defer close(ch)
		for {
			val, err := str.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				fmt.Println("FAILED TO READ")
			}
			ch <- val.GetValue()
		}
	}()
	go func() {
		defer wg.Done()
		for val := range ch {
			_ = str.Send(&stream.Value{Value: val})
		}
	}()
	return nil
}
