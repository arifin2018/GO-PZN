package gopzn

type MyPage struct {
	Name string
}

func (MyPage MyPage) SayHello(name string) string {
	return "hello " + name + "my name is " + MyPage.Name
}
