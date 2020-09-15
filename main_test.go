package main

import "testing"

func TestLine(t *testing.T) {
    tests := []struct {
        line           *line
        expectedResult bool
    }{
        {&line{empty, empty, empty}, false},
        {&line{o, x, x}, false},
        {&line{x, x, x}, true},
    }

    for _, tt := range tests {
        if tt.line.check() != tt.expectedResult {
            t.Errorf("failed")
            return
        }
    }
}

func TestInput(t *testing.T) {
    test := struct {
        g * grid
        expected bool
    }{
        &grid{
            {x, empty, x},
            {o, x, o},
            {x, empty, o},
            {x, o, x},
            {empty, x, empty},
            {x, o, o},
            {x, x, o},
            {x, x, x},
        }, true}

    if winner := test.g.check(); winner != test.expected {
        t.Errorf{"input not correct"}
    }
}

func assertEquals(t *testing.T, value, expected bool) {
    if value != expected {
        t.Errorf{"Values did not match"}
    }
}