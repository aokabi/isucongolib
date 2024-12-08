package isucongolib

import (
	"math/rand"
	"sync"
	"testing"
)

func TestRandMap_Random(t *testing.T) {
	type fields struct {
		m map[string]string
		s []string
	}
	type expected struct {
		key   string
		value string
		ok    bool
	}
	tests := []struct {
		name   string
		fields fields
		want   expected
	}{
		{
			name: "ok",
			fields: fields{
				m: map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
				s: []string{"a", "b", "c"},
			},
			want: expected{
				key:   "a",
				value: "A",
				ok:    true,
			},
		},
		{
			name: "empty",
			fields: fields{
				m: map[string]string{},
				s: []string{},
			},
			want: expected{
				key:   "",
				value: "",
				ok:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fixed seed
			rnd := rand.New(rand.NewSource(0))
			m := &randMap{
				mu:       sync.Mutex{},
				m:        tt.fields.m,
				keys:     tt.fields.s,
				randFunc: rnd.Intn,
			}
			if gotkey, gotval, gotOk := m.Random(); gotkey != tt.want.key || gotval != tt.want.value || gotOk != tt.want.ok {
				t.Errorf("Random() = (%v, %v, %v), want (%v, %v, %v)", gotkey, gotval, gotOk, tt.want.key, tt.want.value, tt.want.ok)
			}
		})
	}
}

func Test_randMap_RandomPop(t *testing.T) {
	type fields struct {
		m map[string]string
		s []string
	}
	type expected struct {
		key   string
		value string
		ok    bool
	}

	tests := []struct {
		name   string
		fields fields
		want   expected
	}{
		{
			name: "ok",
			fields: fields{
				m: map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
				s: []string{"a", "b", "c"},
			},
			want: expected{
				key:   "a",
				value: "A",
				ok:    true,
			},
		},
		{
			name: "empty",
			fields: fields{
				m: map[string]string{},
				s: []string{},
			},
			want: expected{
				key:   "",
				value: "",
				ok:    false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := rand.New(rand.NewSource(0))
			m := &randMap{
				mu:       sync.Mutex{},
				m:        tt.fields.m,
				keys:     tt.fields.s,
				randFunc: rnd.Intn,
			}
			if gotKey, gotValue, gotOk := m.PopRandom(); gotKey != tt.want.key || gotValue != tt.want.value || gotOk != tt.want.ok {
				t.Errorf("RandomPop() = (%v, %v, %v), want (%v, %v, %v)", gotKey, gotValue, gotOk, tt.want.key, tt.want.value, tt.want.ok)
			}

			if _, ok := m.m[tt.want.key]; ok {
				t.Errorf("RandomPop() = (%v, %v, %v), not remove value", tt.want.key, tt.want.value, tt.want.ok)
			}
		})
	}
}

func Test_randMap_Set(t *testing.T) {
	type fields struct {
		m map[string]string
		s []string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "ok",
			fields: fields{
				m: map[string]string{
					"a": "A",
					"b": "B",
				},
				s: []string{"a", "b"},
			},
			args: args{
				key:   "c",
				value: "C",
			},
		},
		{
			name: "overwrite",
			fields: fields{
				m: map[string]string{
					"a": "A",
					"b": "B",
				},
				s: []string{"a", "b"},
			},
			args: args{
				key:   "a",
				value: "AA",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rnd := rand.New(rand.NewSource(0))
			m := &randMap{
				m:        tt.fields.m,
				keys:     tt.fields.s,
				randFunc: rnd.Intn,
			}
			m.Set(tt.args.key, tt.args.value)
			if _, ok := m.m[tt.args.key]; !ok {
				t.Errorf("Set() = (%v, %v), not set value", tt.args.key, tt.args.value)
			}
			if m.m[tt.args.key] != tt.args.value {
				t.Errorf("Set() = (%v, %v), want (%v, %v)", tt.args.key, m.m[tt.args.key], tt.args.key, tt.args.value)
			}
		})
	}
}