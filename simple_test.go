package simple_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jutkko/mockComp/mockCompfakes"
	"github.com/jutkko/mockComp/mockgen"
	"github.com/jutkko/mockComp/mocks"
	"github.com/stretchr/testify/mock"
)

func TestSimpleCounterfeiterPerson(t *testing.T) {
	var fakePerson = new(mockCompfakes.FakePerson)

	// The return values are all type checked
	fakePerson.DoThingsReturns(42, errors.New("An error"))

	res, err := fakePerson.DoThings("Hi it's 24", 24)
	// Do you assertion on res and err
	if res != 42 || err == nil {
		t.Fatal("Wrong result")
	}

	// Get the args from the call, this is an array
	firstArg, secondArg := fakePerson.DoThingsArgsForCall(0)
	if firstArg != "Hi it's 24" || secondArg != 24 {
		t.Fatal("Wrong result")
	}
}

func TestSimpleMockeryPerson(t *testing.T) {
	var fakePerson = &mocks.Person{}

	// The function name is a string, the arguments are not type checked,
	// including the return value
	fakePerson.On("DoThings", "Hi it's 24", 24).Return(42, errors.New("An error"))

	res, err := fakePerson.DoThings("Hi it's 24", 24)
	// Do you assertion on res and err
	if res != 42 || err == nil {
		t.Fatal("Wrong result")
	}
}

func TestSimpleGogenPerson(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var fakePerson = mockgen.NewMockPerson(ctrl)

	// The inputs are not type checked
	fakePerson.EXPECT().DoThings("Hi it's 24", 24).Return(42, nil)

	// The actual call
	fakePerson.DoThings("Hi it's 24", 24)
}

func TestAdvancedCounterfeiterPerson(t *testing.T) {
	var fakePerson = new(mockCompfakes.FakePerson)

	// The function is type checked
	fakePerson.DoThingsStub = func(inputString string, inputInt int) (int, error) {
		if inputString == "Hi it's 24" {
			return inputInt + 1, nil
		}

		return 0, errors.New("An error")
	}

	res, err := fakePerson.DoThings("Hi it's 24", 24)
	// Do you assertion on res and err
	if res != 25 || err != nil {
		t.Fatal("Wrong result")
	}

	// Get the args from the call, this is an array
	firstArg, secondArg := fakePerson.DoThingsArgsForCall(0)
	if firstArg != "Hi it's 24" || secondArg != 24 {
		t.Fatal("Wrong result")
	}
}

func TestAdvancedMockeryPerson(t *testing.T) {
	var fakePerson = &mocks.Person{}

	// The function is not type checked
	doThingsStub := func(inputString string, inputInt int) (int, error) {
		if inputString == "Hi it's 24" {
			return inputInt + 1, nil
		}

		return 0, errors.New("An error")
	}

	fakePerson.On("DoThings", mock.Anything, mock.Anything).Return(doThingsStub)

	res, err := fakePerson.DoThings("Hi it's 24", 24)
	// Do you assertion on res and err
	if res != 42 || err == nil {
		t.Fatal("Wrong result")
	}
}
