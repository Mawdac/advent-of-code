package main

import (
	"reflect"
	"testing"
)

func Test_readLines(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test file",
			args: args{file: "test.txt"},
			want: []string{
				"1abc2",
				"pqr3stu8vwx",
				"a1b2c3d4e5f",
				"treb7uchet",
				"two1nine",
				"eightwothree",
				"abcone2threexyz",
				"xtwone3four",
				"4nineeightseven2",
				"zoneight234",
				"7pqrstsixteen",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readLines(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test file",
			args: args{
				lines: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
			},
			want: 351,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test file",
			args: args{
				lines: []string{
					"1abc2",
					"pqr3stu8vwx",
					"a1b2c3d4e5f",
					"treb7uchet",
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
			},
			want: 423,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseOnlyNumbersFromCalibration(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "just numbers",
			args: args{line: "12345"},
			want: 15,
		},
		{
			name: "numbers and letters",
			args: args{line: "2aajc3o4159abc"},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOnlyNumbersFromCalibration(tt.args.line); got != tt.want {
				t.Errorf("parseOnlyNumbersFromCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isRuneHot(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should be hot",
			args: args{r: 'o'},
			want: true,
		},
		{
			name: "should not be hot",
			args: args{r: 'u'},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRuneHot(tt.args.r); got != tt.want {
				t.Errorf("isRuneHot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseNumbersAndStringsFromCalibration(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "just numbers",
			args: args{line: "12345"},
			want: 15,
		},
		{
			name: "numbers and letters",
			args: args{line: "2aajc3o4159abc"},
			want: 29,
		},
		{
			name: "numbers as strings",
			args: args{line: "threetwoeight"},
			want: 38,
		},
		{
			name: "numbers as strings and numbers",
			args: args{line: "9three1kzzi2twonine"},
			want: 99,
		},
		{
			name: "words with shared letters",
			args: args{line: "twoneighthreeightwo"},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseNumbersAndStringsFromCalibration(tt.args.line); got != tt.want {
				t.Errorf("parseNumbersAndStringsFromCalibration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitByNumbersAndStrings(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "just numbers",
			args: args{line: "12345"},
			want: []string{"1", "2", "3", "4", "5"},
		},
		{
			name: "numbers and letters",
			args: args{line: "1abc2"},
			want: []string{"1", "abc", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitByNumbersAndStrings(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitByNumbersAndStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseNumbersFromStrings(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "one number",
			args: args{text: "one"},
			want: []string{"1"},
		},
		{
			name: "multiple numbers",
			args: args{text: "twoneighthreeightwo"},
			want: []string{"2", "1", "8", "3", "8", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseNumbersFromStrings(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseNumbersFromStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
