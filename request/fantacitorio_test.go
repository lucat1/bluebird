package request

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSearchNameBySurname(t *testing.T) {
	name, err := searchNameBySurname("BERLUSCONI")
	assert.Nil(t, err, "Politician with Berlusconi as surname exists")
	assert.Equal(t, "SILVIO", name, "Name of Berlusconi should be Silvio")
	name, err = searchNameBySurname("CATALINI")
	assert.NotNil(t, err, "Politician with this surname doesn't exist")
	assert.Equal(t, "", name, "Politician with this surname doesn't exist")
}

func TestCheckNameSurname(t *testing.T) {
	name, surname, err := checkNameSurname("SILVIO BERLUSCONI")
	assert.Equal(t, "SILVIO", name, "Expected first name to be Silvio")
	assert.Equal(t, "BERLUSCONI", surname, "Expected first name to be Berlusconi")
	assert.Nil(t, err, "Expected no error with Silvio Berlusconi")

	name, surname, err = checkNameSurname("SILVIA BERLUSCONI")
	assert.NotNil(t, err, "Silvia Berlusconi doesn't exist")

	name, surname, err = checkNameSurname("BERLUSCONI")
	assert.Equal(t, "SILVIO", name, "Expected first name to be Silvio")
	assert.Equal(t, "BERLUSCONI", surname, "Expected first name to be Berlusconi")
	assert.Nil(t, err, "Expected no error with Silvio Berlusconi")

	name, surname, err = checkNameSurname("CATALINI")
	assert.NotNil(t, err, "Politician with Catalini as surname exists")

	name, surname, err = checkNameSurname("")
	assert.Nil(t, err, "Empty strings should be not allowed")
}

func TestRemoveEmptyStrings(t *testing.T) {
	before := []string{"", "s1", "", "s2"}
	expected := []string{"s1", "s2"}
	actual := removeEmptyStrings(before)
	assert.Equal(t, expected, actual, "Expected no empty strings")
}

func TestParsePolitician(t *testing.T) {
	politician, err := parsePolitician("600 PUNTI - SIMONE PILLON")
	assert.Nil(t, err, "Error should be nil, input is correctly formed")
	assert.Equal(t, "SIMONE", politician.Name, "Politician's name should be SIMONE")
	assert.Equal(t, "PILLON", politician.Surname, "Politician's surname should be SIMONE")

	politician, err = parsePolitician("SIMONE PILLON 600 punti")
	assert.Nil(t, err, "Error should be nil, input is correctly formed")
	assert.Equal(t, "SIMONE", politician.Name, "Politician's name should be SIMONE")
	assert.Equal(t, "PILLON", politician.Surname, "Politician's surname should be SIMONE")

	politician, err = parsePolitician("PUNTI - SIMON PILLON")
	assert.NotNil(t, err, "Expected error, points field is missing")

	politician, err = parsePolitician("SIMONE 600 punti")
	assert.NotNil(t, err, "Expected error, SIMON PILLON ERROR is not a valid politician")
}

func TestPoliticianScore(t *testing.T) {
	PoliticiansScore(100, time.Now(), time.Now().Add(64))
}
