package some_data

type Test1 interface {
	Test1()
}
type Test2 interface {
	Test2()
}
type Test3 interface {
	Test3()
}
type Test4 interface {
	Test4()
}

type ConcreteTest1 struct {
	t1 Test2
}

func NewConcreteTest1(t1 Test2) *ConcreteTest1 {
	return &ConcreteTest1{t1: t1}
}

func (receiver ConcreteTest1) Test1() {

}

type ConcreteTest2 struct {
	t2 Test3
}

func NewConcreteTest2(t2 Test3) *ConcreteTest2 {
	return &ConcreteTest2{t2: t2}
}

func (receiver ConcreteTest2) Test2() {

}

type ConcreteTest3 struct {
	t3 Test4
}

func NewConcreteTest3(t3 Test4) *ConcreteTest3 {
	return &ConcreteTest3{t3: t3}
}

func (receiver ConcreteTest3) Test3() {

}

type ConcreteTest4 struct {
}

func NewConcreteTest4() *ConcreteTest4 {
	return &ConcreteTest4{}
}

func (receiver ConcreteTest4) Test4() {

}

type Test5 interface {
	Test5()
}

type Test6 interface {
	Test6()
}
type ConcreteTest5 struct {
	t6 Test6
}

func NewConcreteTest5(t6 Test6) *ConcreteTest5 {
	return &ConcreteTest5{t6: t6}
}

func (receiver ConcreteTest5) Test5() {

}

type ConcreteTest6 struct {
	t5 Test6
}

func NewConcreteTest6(t5 Test6) *ConcreteTest6 {
	return &ConcreteTest6{t5: t5}
}

func (receiver ConcreteTest6) Test6() {

}
