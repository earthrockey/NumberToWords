package numbertowords

import (
	"fmt"
	"testing"
)

func TestNumToWords(t *testing.T) {
	var tests = []struct {
		input int
		want  string
	}{
		{input: 1, want: "หนึ่ง"},
		{input: 12, want: "สิบสอง"},
		{input: 123, want: "หนึ่งร้อยยี่สิบสาม"},
		{input: 1234, want: "หนึ่งพันสองร้อยสามสิบสี่"},
		{input: 12345, want: "หนึ่งหมื่นสองพันสามร้อยสี่สิบห้า"},
		{input: 123456, want: "หนึ่งแสนสองหมื่นสามพันสี่ร้อยห้าสิบหก"},
		{input: 1234567, want: "หนึ่งล้านสองแสนสามหมื่นสี่พันห้าร้อยหกสิบเจ็ด"},
		{input: 10, want: "สิบ"},
		{input: 101, want: "หนึ่งร้อยเอ็ด"},
		{input: 1011, want: "หนึ่งพันสิบเอ็ด"},
		{input: 10111, want: "หนึ่งหมื่นหนึ่งร้อยสิบเอ็ด"},
		{input: 101111, want: "หนึ่งแสนหนึ่งพันหนึ่งร้อยสิบเอ็ด"},
		{input: 1011111, want: "หนึ่งล้านหนึ่งหมื่นหนึ่งพันหนึ่งร้อยสิบเอ็ด"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.want)
		t.Run(testName, func(t *testing.T) {
			ans := NumToWords(tt.input)

			if ans != tt.want {
				t.Errorf("got : %s, want : %s", ans, tt.want)
			}
		})
	}
}
