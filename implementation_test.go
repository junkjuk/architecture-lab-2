package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix1(t *testing.T) {
	res, err := PrefixToPostfix("+ 7 9")
	if assert.Nil(t, err) {
		assert.Equal(t, "7 9 +", res)
	}
}

func TestPrefixToPostfix2(t *testing.T) {
	res, err := PrefixToPostfix("* - 2 4 1")
	if assert.Nil(t, err) {
		assert.Equal(t, "1 2 4 - *", res)
	}
}

func TestPrefixToPostfix3(t *testing.T) {
	res, err := PrefixToPostfix("/ * + 2 5 + - 2 5 6 - 5 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "2 5 + 2 5 - 6 + * 5 3 - /", res)
	}
}

func TestPrefixToPostfix4(t *testing.T) {
	res, err := PrefixToPostfix("/ * + - - 3 5 1 9 + + + 3 4 9 2 8")
	if assert.Nil(t, err) {
		assert.Equal(t, "3 4 9 2 + + + + 3 5 - 1 - 9 + * 8 /", res)
	}
}

func TestPrefixToPostfix5(t *testing.T) {
	_, err := PrefixToPostfix("")

	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Inappropriate input"), err)
	}
}
func TestPrefixToPostfix6(t *testing.T) {
	_, err := PrefixToPostfix("plus three  eight")

	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Inappropriate input"), err)
	}
}
func TestPrefixToPostfix7(t *testing.T) {
	_, err := PrefixToPostfix("-....- ..... ...--")

	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("Inappropriate input"), err)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- 5 8")
	fmt.Println(res)

	// Output:
	// 5 8 -
}
