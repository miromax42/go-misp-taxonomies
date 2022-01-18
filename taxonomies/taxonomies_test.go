package taxonomies

import (
	"reflect"
	"testing"
)

func Test_fromString1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantT   Taxonomy
		wantErr bool
	}{
		{
			"good with value",
			args{s: `test:pred="228"`},
			Taxonomy{
				Namespace: "test",
				Predicate: "pred",
				Value:     "228",
			},
			false,
		},
		{
			"good without value",
			args{s: `test:pred`},
			Taxonomy{
				Namespace: "test",
				Predicate: "pred",
				Value:     "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, err := fromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("fromString() gotT = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestTaxonomy_String(t1 *testing.T) {
	type fields struct {
		Namespace string
		Predicate string
		Value     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"good with 2",
			fields{
				Namespace: "testinn",
				Predicate: "pred",
				Value:     "",
			},
			"testinn:pred",
		},
		{
			"good with 3",
			fields{
				Namespace: "testinn",
				Predicate: "pred",
				Value:     "228",
			},
			`testinn:pred="228"`,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Taxonomy{
				Namespace: tt.fields.Namespace,
				Predicate: tt.fields.Predicate,
				Value:     tt.fields.Value,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Circle(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			"good with value",
			`test:pred="228"`,
			false,
		},
		{
			"good without value",
			`test:pred`,
			false,
		},
		{
			"bad without value",
			`:te==st:pred=213`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, err := fromString(tt.s)
			if (err != nil) == tt.wantErr {
				return
			}

			gotS := gotT.String()
			if !reflect.DeepEqual(gotS, tt.s) {
				t.Errorf("fromString() gotT = %v, want %v", gotS, tt.s)
			}
		})
	}
}
