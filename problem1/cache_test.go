package main

import "testing"

func TestInMemoryCache_Get(t *testing.T) {
	node1 := Node{
		key:   "1",
		value: "satu",
	}
	node2 := Node{
		prev:  &node1,
		key:   "2",
		value: "dua",
	}
	node1.next = &node2

	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success get value of key",
			args: args{
				key: "1",
			},
			want: "satu",
		},
		{
			name: "get empty value of key",
			args: args{
				key: "3",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NoneEvictionManager{}
			m := New(3, &e)
			m.head = &node1
			m.tail = &node2
			e.InMemoryCache = &m
			if got := m.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryCache_Add(t *testing.T) {
	node1 := Node{
		key:   "1",
		value: "satu",
	}
	node2 := Node{
		prev:  &node1,
		key:   "2",
		value: "dua",
	}
	node1.next = &node2
	type args struct {
		key   string
		value string
	}
	type field struct {
		limit int
	}
	tests := []struct {
		name    string
		fields  field
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "success add new key value",
			fields: field{
				limit: 3,
			},
			args: args{
				key:   "3",
				value: "tiga",
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "success update key value",
			fields: field{
				limit: 3,
			},
			args: args{
				key:   "1",
				value: "satu.satu",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "failed when add more than 2 key",
			fields: field{
				limit: 2,
			},
			args: args{
				key:   "3",
				value: "tiga",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := NoneEvictionManager{}
			m := New(tt.fields.limit, &e)
			m.head = &node1
			m.tail = &node2
			e.InMemoryCache = &m
			got, err := m.Add(tt.args.key, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
		})
	}
}
