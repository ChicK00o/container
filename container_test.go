package container_test

import (
	"github.com/golobby/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Shape interface {
	SetArea(int)
	GetArea() int
}

type Circle struct {
	a int
}

func (c *Circle) SetArea(a int) {
	c.a = a
}

func (c Circle) GetArea() int {
	return c.a
}

type Database interface {
	Connect() bool
}

type MySQL struct{}

func (m MySQL) Connect() bool {
	return true
}

type Combined struct {
	Cir *Circle
	Db *MySQL
}

type Combined2 struct {
	Cir Circle
	Db *MySQL
}

func TestCombinedObject1(t *testing.T) {
	container.Reset()

	container.Singleton(func() *Circle {
		return &Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir *Circle, db *MySQL) *Combined{
		return &Combined{
			Cir: cir,
			Db:  db,
		}
	})

	var combined *Combined
	container.Make(&combined)
	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 5)
}

func TestCombinedObject2(t *testing.T) {
	container.Reset()

	container.Singleton(func() *Circle {
		return &Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir *Circle, db *MySQL) *Combined{
		return &Combined{
			Cir: cir,
			Db:  db,
		}
	})

	var combined Combined
	container.Make(&combined)
	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 5)
}

func TestCombinedObject3(t *testing.T) {
	container.Reset()

	container.Singleton(func() *Circle {
		return &Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir Circle, db *MySQL) *Combined2{
		return &Combined2{
			Cir: cir,
			Db:  db,
		}
	})

	var combined Combined2
	container.Make(&combined)
	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 5)
}

func TestCombinedObject4(t *testing.T) {
	container.Reset()

	container.Singleton(func() Circle {
		return Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir *Circle, db *MySQL) *Combined{
		return &Combined{
			Cir: cir,
			Db:  db,
		}
	})

	var combined Combined
	container.Make(&combined)
	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 5)
}

func TestCombinedObject5(t *testing.T) {
	container.Reset()

	container.Singleton(func() *Circle {
		return &Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir Circle, db *MySQL) Combined2{
		return Combined2{
			Cir: cir,
			Db:  db,
		}
	})

	var combined *Combined2
	container.Make(&combined)
	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 5)
}

func TestCombinedObject6(t *testing.T) {
	container.Reset()

	container.Singleton(func() Circle {
		return Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir *Circle, db *MySQL) Combined{
		return Combined{
			Cir: cir,
			Db:  db,
		}
	})

	var combined *Combined
	container.Make(&combined)
	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 5)
}

func TestCombinedObject7(t *testing.T) {
	container.Reset()

	container.Singleton(func() Circle {
		return Circle{a:5}
	})

	container.Singleton(func() *MySQL {
		return &MySQL{}
	})

	container.Singleton(func(cir *Circle, db *MySQL) Combined{
		return Combined{
			Cir: cir,
			Db:  db,
		}
	})

	var circle Circle
	container.Make(&circle)

	var combined *Combined
	container.Make(&combined)
	combined.Cir.a = 6

	assert.NotNil(t, combined)
	assert.NotNil(t, combined.Cir)
	assert.Equal(t, combined.Cir.a, 6)
	assert.Equal(t, circle.a, 6)
}

func TestSingletonItShouldMakeAnInstanceOfTheAbstraction(t *testing.T) {
	area := 5

	container.Singleton(func() Shape {
		return &Circle{a: area}
	})

	container.Make(func(s Shape) {
		a := s.GetArea()
		assert.Equal(t, area, a)
	})
}

func TestSingletonItShouldMakeSameObjectEachMake(t *testing.T) {
	container.Singleton(func() Shape {
		return &Circle{a: 5}
	})

	area := 6

	container.Make(func(s1 Shape) {
		s1.SetArea(area)
	})

	container.Make(func(s2 Shape) {
		a := s2.GetArea()
		assert.Equal(t, a, area)
	})
}

func TestSingletonWithNonFunctionResolverItShouldPanic(t *testing.T) {
	value := "the resolver must be a function"
	assert.PanicsWithValue(t, value, func() {
		container.Singleton("STRING!")
	}, "Expected panic")
}

func TestSingletonItShouldResolveResolverArguments(t *testing.T) {
	area := 5
	container.Singleton(func() Shape {
		return &Circle{a: area}
	})

	container.Singleton(func(s Shape) Database {
		assert.Equal(t, s.GetArea(), area)
		return &MySQL{}
	})
}

func TestTransientItShouldMakeDifferentObjectsOnMake(t *testing.T) {
	area := 5

	container.Transient(func() Shape {
		return &Circle{a: area}
	})

	container.Make(func(s1 Shape) {
		s1.SetArea(6)
	})

	container.Make(func(s2 Shape) {
		a := s2.GetArea()
		assert.Equal(t, a, area)
	})
}

func TestTransientItShouldMakeAnInstanceOfTheAbstraction(t *testing.T) {
	area := 5

	container.Transient(func() Shape {
		return &Circle{a: area}
	})

	container.Make(func(s Shape) {
		a := s.GetArea()
		assert.Equal(t, a, area)
	})
}

func TestMakeWithSingleInputAndCallback(t *testing.T) {
	container.Singleton(func() Shape {
		return &Circle{a: 5}
	})

	container.Make(func(s Shape) {
		if _, ok := s.(*Circle); !ok {
			t.Error("Expected Circle")
		}
	})
}

func TestMakeWithMultipleInputsAndCallback(t *testing.T) {
	container.Singleton(func() Shape {
		return &Circle{a: 5}
	})

	container.Singleton(func() Database {
		return &MySQL{}
	})

	container.Make(func(s Shape, m Database) {
		if _, ok := s.(*Circle); !ok {
			t.Error("Expected Circle")
		}

		if _, ok := m.(*MySQL); !ok {
			t.Error("Expected MySQL")
		}
	})
}

func TestMakeWithSingleInputAndReference(t *testing.T) {
	container.Singleton(func() Shape {
		return &Circle{a: 5}
	})

	var s Shape

	container.Make(&s)

	if _, ok := s.(*Circle); !ok {
		t.Error("Expected Circle")
	}
}

func TestMakeWithMultipleInputsAndReference(t *testing.T) {
	container.Singleton(func() Shape {
		return &Circle{a: 5}
	})

	container.Singleton(func() Database {
		return &MySQL{}
	})

	var (
		s Shape
		d Database
	)

	container.Make(&s)
	container.Make(&d)

	if _, ok := s.(*Circle); !ok {
		t.Error("Expected Circle")
	}

	if _, ok := d.(*MySQL); !ok {
		t.Error("Expected MySQL")
	}
}

func TestMakeWithUnsupportedReceiver(t *testing.T) {
	value := "the receiver must be either a reference or a callback"
	assert.PanicsWithValue(t, value, func() {
		container.Reset()
		container.Make("STRING!")
	}, "Expected panic")
}

func TestMakeWithNonReference(t *testing.T) {
	value := "cannot detect type of the receiver, make sure your are passing reference of the object"
	assert.PanicsWithValue(t, value, func() {
		var s Shape
		container.Reset()
		container.Make(s)
	}, "Expected panic")
}

func TestMakeWithUnboundedAbstraction(t *testing.T) {
	value := "no concrete found for the abstraction container_test.Shape"
	assert.PanicsWithValue(t, value, func() {
		var s Shape
		container.Reset()
		container.Make(&s)
	}, "Expected panic")
}
func TestMakeWithUnboundedAbstractionPtr(t *testing.T) {
	value := "no concrete found for the abstraction *container_test.Shape"
	assert.PanicsWithValue(t, value, func() {
		var s *Shape
		container.Reset()
		container.Make(&s)
	}, "Expected panic")
}

func TestMakeWithCallbackThatHasAUnboundedAbstraction(t *testing.T) {
	value := "no concrete found for the abstraction container_test.Database"
	assert.PanicsWithValue(t, value, func() {
		container.Reset()
		container.Singleton(func() Shape {
			return &Circle{}
		})
		container.Make(func(s Shape, d Database) {})
	}, "Expected panic")
}
