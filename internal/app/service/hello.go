package service

// IHelloService interface for function on person service
type IHelloService interface {
	Greet() (greeting string)
}

// HelloService struct for person service
type HelloService struct {
	opt Option
}

// NewHelloService initiate new person service
func NewHelloService(opt Option) IHelloService {
	return &HelloService{
		opt: opt,
	}
}

// Greet function to greet person. Introducing with random name
func (hs *HelloService) Greet() (greeting string) {
	name := hs.opt.Repository.Person.RandomPerson()

	greeting = "Hi, my name is " + name + ". Nice to meet you!"

	return
}
