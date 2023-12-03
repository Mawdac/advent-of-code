package main

import (
	"reflect"
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		games []game
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test data",
			args: args{
				[]game{
					{
						sets: []set{
							{
								red:  4,
								blue: 3,
							},
							{
								red:   1,
								green: 2,
								blue:  6,
							},
							{
								green: 2,
							},
						},
					},
					{
						sets: []set{
							{
								green: 2,
								blue:  1,
							},
							{
								red:   1,
								green: 3,
								blue:  4,
							},
							{
								green: 1,
								blue:  1,
							},
						},
					},
					{
						sets: []set{
							{
								red:   20,
								green: 8,
								blue:  6,
							},
							{
								red:   3,
								green: 13,
								blue:  5,
							},
							{
								green: 5,
								red:   1,
							},
						},
					},
					{
						sets: []set{
							{
								red:   3,
								green: 1,
								blue:  6,
							},
							{
								red:   6,
								green: 3,
							},
							{
								green: 3,
								red:   14,
								blue:  15,
							},
						},
					},
					{
						sets: []set{
							{
								red:   6,
								green: 3,
								blue:  1,
							},
							{
								red:   1,
								blue:  2,
								green: 2,
							},
						},
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.games); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		games []game
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test data",
			args: args{
				[]game{
					{
						sets: []set{
							{
								red:  4,
								blue: 3,
							},
							{
								red:   1,
								green: 2,
								blue:  6,
							},
							{
								green: 2,
							},
						},
					},
					{
						sets: []set{
							{
								green: 2,
								blue:  1,
							},
							{
								red:   1,
								green: 3,
								blue:  4,
							},
							{
								green: 1,
								blue:  1,
							},
						},
					},
					{
						sets: []set{
							{
								red:   20,
								green: 8,
								blue:  6,
							},
							{
								red:   3,
								green: 13,
								blue:  5,
							},
							{
								green: 5,
								red:   1,
							},
						},
					},
					{
						sets: []set{
							{
								red:   3,
								green: 1,
								blue:  6,
							},
							{
								red:   6,
								green: 3,
							},
							{
								green: 3,
								red:   14,
								blue:  15,
							},
						},
					},
					{
						sets: []set{
							{
								red:   6,
								green: 3,
								blue:  1,
							},
							{
								red:   1,
								blue:  2,
								green: 2,
							},
						},
					},
				},
			},
			want: 2286,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.games); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readGames(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readGames(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseGame(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want game
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseGame(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addColorCount(t *testing.T) {
	type args struct {
		countAndColor []string
		s             *set
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addColorCount(tt.args.countAndColor, tt.args.s)
		})
	}
}
