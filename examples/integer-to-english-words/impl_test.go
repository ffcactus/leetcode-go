package integer_to_english_words

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestConvert00to99(t *testing.T) {
	assert.Equal(t, convert00to99("00"), "")
	assert.Equal(t, convert00to99("12"), "Twelve")
	assert.Equal(t, convert00to99("99"), "Ninety Nine")
}

func TestConvertPart(t *testing.T) {
	assert.Equal(t, convertPart([]byte("1")), "One")
	assert.Equal(t, convertPart([]byte("12")), "Twelve")
	assert.Equal(t, convertPart([]byte("99")), "Ninety Nine")
	assert.Equal(t, convertPart([]byte("999")), "Nine Hundred Ninety Nine")
}

func TestNumberToWords(t *testing.T) {
	assert.Equal(t, numberToWords(0), "Zero")
	assert.Equal(t, numberToWords(1), "One")
	assert.Equal(t, numberToWords(12), "Twelve")
	assert.Equal(t, numberToWords(99), "Ninety Nine")
	assert.Equal(t, numberToWords(999), "Nine Hundred Ninety Nine")
	assert.Equal(t, numberToWords(9999), "Nine Thousand Nine Hundred Ninety Nine")
	assert.Equal(t, numberToWords(99999), "Ninety Nine Thousand Nine Hundred Ninety Nine")
	assert.Equal(t, numberToWords(999999), "Nine Hundred Ninety Nine Thousand Nine Hundred Ninety Nine")
	assert.Equal(t, numberToWords(10), "Ten")
	assert.Equal(t, numberToWords(100), "One Hundred")
	assert.Equal(t, numberToWords(1000), "One Thousand")
	assert.Equal(t, numberToWords(1000000), "One Million")
	assert.Equal(t, numberToWords(1001000), "One Million One Thousand")
	assert.Equal(t, numberToWords(1000000000), "One Billion")
	assert.Equal(t, numberToWords(1001001100), "One Billion One Million One Thousand One Hundred")
	assert.Equal(t, numberToWords(1234567), "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven")
	assert.Equal(t, numberToWords(1234567891), "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One")
}
