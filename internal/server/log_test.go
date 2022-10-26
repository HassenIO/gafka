package server

import (
	"bytes"
	"testing"
)

func TestAppend(t *testing.T) {
	recordValue := []byte("Hello")
	l := NewLog()
	r := Record{Value: recordValue}
	offset, err := l.Append(r)
	if err != nil {
		t.Error("Should not get any error when good append")
	}
	if offset != 0 {
		t.Errorf("Expected offset=0, got offset=%d", offset)
	}
	if bytes.Compare(l.records[0].Value, recordValue) != 0 {
		t.Errorf("Expected record value to be %s, got %s", recordValue, l.records[0].Value)
	}
}

func TestRead(t *testing.T) {
	recordValue := []byte("Hello")
	l := NewLog()
	r := Record{Value: recordValue}
	offset, _ := l.Append(r)
	result, err := l.Read(offset)
	if err != nil {
		t.Errorf("Get an error=%e", err)
	}
	if bytes.Compare(result.Value, recordValue) != 0 {
		t.Errorf("Expected record value to be %s, but got %s", recordValue, result.Value)
	}
}
