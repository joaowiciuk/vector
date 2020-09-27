package main

import (
	"reflect"
	"testing"
)

func TestVector_Add(t *testing.T) {
	type args struct {
		b Vector
	}
	tests := []struct {
		name string
		a    Vector
		args args
		want Vector
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Add(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Vector.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
