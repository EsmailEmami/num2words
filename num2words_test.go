package num2words_test

import (
	"testing"

	"github.com/esmailemami/num2words"
)

func TestConvert(t *testing.T) {
	Convert := num2words.Convert

	tests := []struct {
		number        int
		expectedWords string
	}{
		{0, "صفر"},
		{1, "یک"},
		{5, "پنج"},
		{9, "نه"},
		{10, "ده"},
		{12, "دوازده"},
		{17, "هفده"},
		{20, "بیست"},
		{30, "سی"},
		{40, "چهل"},
		{50, "پنجاه"},
		{60, "شصت"},
		{90, "نود"},
		{21, "بیست و یک"},
		{49, "چهل و نه"},
		{53, "پنجاه و سه"},
		{68, "شصت و هشت"},
		{99, "نود و نه"},
		{100, "یکصد"},
		{200, "دویست"},
		{500, "پانصد"},
		{123, "یکصد و بیست و سه"},
		{666, "ششصد و شصت و شش"},
		{666, "ششصد و شصت و شش"},
		{666, "ششصد و شصت و شش"},
		{666, "ششصد و شصت و شش"},
		{666, "ششصد و شصت و شش"},
		{1024, "یک هزار و بیست و چهار"},
		{1_000_000, "یک میلیون"},
		{10_000_000, "ده میلیون"},
		{100_000_000, "یکصد میلیون"},
		{100_000_000_000, "یکصد میلیارد"},
		{365_657_764_000, "سیصد و شصت و پنج میلیارد و ششصد و پنجاه و هفت میلیون و هفتصد و شصت و چهار هزار"},
	}

	for _, tt := range tests {
		words := Convert(tt.number)
		if words != tt.expectedWords {
			t.Fatalf("tests[%d] - words wrong. expected=%q, got=%q", tt.number, tt.expectedWords, words)
		}
	}
}