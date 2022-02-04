package main

// to create new entity, 
// 1. create new type, which implements Animal interface
// 2. add new _Type const
// 3. modify Create function: switch has to match new _Type case with new type

type AnimalType int

const (
	CatType AnimalType = iota
	DogType
)

// Create builds animal according to type received;
// returns pointer to requested animal if everything went fine,
// nil otherwise
func Create(at AnimalType) *Animal {
	switch at {
	case CatType:
		var a Animal = Cat{}
		return &a
	case DogType:
		var a Animal = Dog{}
		return &a
	default:
		return nil
	}
}

type Cat struct{}

func (c Cat) Sound() string {
	return `"Мяу"`
}

type Dog struct{}

func (c Dog) Sound() string {
	return `"Гав"`
}

