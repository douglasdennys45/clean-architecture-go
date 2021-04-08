package protocols

type Validation interface {
	Validate(inputs interface{}) error
}
