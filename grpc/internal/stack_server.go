package internal

import (
	"context"
	"errors"

	"gitlab.tcsbank.ru/dealer/toolbox/edu9/grpc/unary"
	"google.golang.org/protobuf/types/known/emptypb"
)

type StackServer struct {
	stack *Stack
	unary.UnimplementedIntStackServer
}

func NewStackServer(stack *Stack) *StackServer {
	return &StackServer{stack: stack}
}

func (s *StackServer) Push(_ context.Context, value *unary.Value) (*emptypb.Empty, error) {
	s.stack.Push(value.GetValue())
	return &emptypb.Empty{}, nil
}

func (s *StackServer) Pop(context.Context, *emptypb.Empty) (*unary.Value, error) {
	value, err := s.stack.Pop()
	if err != nil {
		return nil, err
	}
	return &unary.Value{Value: value}, nil
}

func (s *StackServer) Peek(context.Context, *emptypb.Empty) (*unary.Value, error) {
	value, err := s.stack.Peek()
	if err != nil {
		return nil, err
	}
	return &unary.Value{Value: value}, nil
}

var ErrStackEmpty = errors.New("stack is empty")

type stackValue struct {
	value  int64
	parent *stackValue
}

type Stack struct {
	top *stackValue
}

func NewStack() *Stack {
	return &Stack{top: nil}
}

func (s *Stack) Push(value int64) {
	current := s.top
	s.top = &stackValue{
		value:  value,
		parent: current,
	}
}

func (s *Stack) Pop() (int64, error) {
	if s.top == nil {
		return 0, ErrStackEmpty
	}
	value := s.top.value
	s.top = s.top.parent
	return value, nil
}

func (s *Stack) Peek() (int64, error) {
	if s.top == nil {
		return 0, ErrStackEmpty
	}
	return s.top.value, nil
}
