package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	color     Color
	columnNum int
}

func TestNewBoard(t *testing.T) {
	cases := []struct {
		desc   string
		output [][]Color
	}{
		{
			"simple test",
			[][]Color{
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
			},
		},
	}

	for _, test := range cases {
		result := NewBoard()
		assert.Equal(t, test.output, result.State)
	}
}

func TestInsertColor(t *testing.T) {
	cases := []struct {
		desc          string
		inputs        []Input
		output        [][]Color
		expectedError error
	}{
		{
			"insert red color",
			[]Input{Input{Red, 0}},
			[][]Color{
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
			},
			nil,
		},
		{
			"insert red color swice",
			[]Input{Input{Red, 0}, Input{Red, 0}},
			[][]Color{
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
			},
			nil,
		},
		{
			"insert red color twice",
			[]Input{Input{Red, 0}, Input{Red, 0}},
			[][]Color{
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
			},
			nil,
		},
		{
			"insert red and yellow colors",
			[]Input{Input{Red, 0}, Input{Yellow, 1}, Input{Yellow, 0}},
			[][]Color{
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Blank, Blank, Blank, Blank, Blank, Blank, Blank},
				{Yellow, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Yellow, Blank, Blank, Blank, Blank, Blank},
			},
			nil,
		},
		{
			"full inserting",
			[]Input{Input{Red, 0}, Input{Red, 0}, Input{Yellow, 0}, Input{Yellow, 0}, Input{Yellow, 0}, Input{Red, 0}, Input{Red, 0}},
			[][]Color{
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
				{Yellow, Blank, Blank, Blank, Blank, Blank, Blank},
				{Yellow, Blank, Blank, Blank, Blank, Blank, Blank},
				{Yellow, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
				{Red, Blank, Blank, Blank, Blank, Blank, Blank},
			},
			ErrColumnFull,
		},
	}

	for _, test := range cases {
		b := NewBoard()
		for _, input := range test.inputs {
			if err := b.InsertColor(input.color, input.columnNum); err != nil {
				if err != test.expectedError {
					t.Fatalf("%s: %s", test.desc, err)
				}
			}
		}
		assert.Equal(t, test.output, b.State)
	}
}

type Output struct {
	found bool
	color Color
}

func TestHasWinner(t *testing.T) {
	cases := []struct {
		desc   string
		inputs []Input
		expect Output
	}{
		{
			desc:   "init status",
			inputs: []Input{},
			expect: Output{
				found: false,
				color: 0,
			},
		},
		{
			desc:   "vertical winner check",
			inputs: []Input{Input{Red, 0}, Input{Red, 0}, Input{Red, 0}, Input{Red, 0}},
			expect: Output{
				found: true,
				color: Red,
			},
		},
		{
			desc:   "horizontal winner check",
			inputs: []Input{Input{Red, 0}, Input{Red, 1}, Input{Red, 2}, Input{Red, 3}},
			expect: Output{
				found: true,
				color: Red,
			},
		},
		{
			desc:   "diagonal winner check left-top to right-down",
			inputs: []Input{Input{Red, 0}, Input{Yellow, 0}, Input{Red, 0}, Input{Yellow, 0}, Input{Red, 1}, Input{Yellow, 1}, Input{Red, 1}, Input{Red, 2}, Input{Yellow, 2}, Input{Red, 3}},
			expect: Output{
				found: true,
				color: Red,
			},
		},
		{
			desc:   "diagonal winner check left-down to right-top",
			inputs: []Input{Input{Red, 0}, Input{Yellow, 1}, Input{Red, 1}, Input{Yellow, 2}, Input{Red, 2}, Input{Red, 2}, Input{Red, 3}, Input{Red, 3}, Input{Yellow, 3}, Input{Red, 3}},
			expect: Output{
				found: true,
				color: Red,
			},
		},
	}

	for _, test := range cases {
		b := NewBoard()

		for _, input := range test.inputs {
			if err := b.InsertColor(input.color, input.columnNum); err != nil {
				t.Fatalf("%s: %s", test.desc, err)
			}
		}

		found, color := b.HasWinner()
		assert.Equal(t, test.expect.found, found)
		assert.Equal(t, test.expect.color, color)

	}

}
