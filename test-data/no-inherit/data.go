package no_inherit

import "context"

type SomeInterfaceCase1 interface {
	HelloWorld(str string)
	HelloWorld2(str string)
	HelloWorld1(str string)
	HelloWorld3(str string)
}

type SomeStructCase1 struct {
}

func NewSomeStructCase1() *SomeStructCase1 {
	return &SomeStructCase1{}
}

func (s SomeStructCase1) HelloWorld(ctx context.Context, str string)  {}
func (s SomeStructCase1) HelloWorld1(ctx context.Context, str string) {}
func (s SomeStructCase1) HelloWorld2(ctx context.Context, str string) {}
func (s SomeStructCase1) HelloWorld3(ctx context.Context, str string) {}

type SomeInterfaceCase2 interface {
	HelloWorld01(ctx context.Context, str string)
	HelloWorld11(ctx context.Context, str string)
	HelloWorld21(ctx context.Context, str string)
	HelloWorld31(ctx context.Context, str string)
}

type SomeStructCase2 struct {
}

func (s *SomeStructCase2) HelloWorld01(ctx context.Context, str string) {}
func (s *SomeStructCase2) HelloWorld11(ctx context.Context, str string) {}
func (s *SomeStructCase2) HelloWorld21(ctx context.Context, str string) {}
func (s *SomeStructCase2) HelloWorld31(ctx context.Context, str string) {}
