package config

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
	"github.com/k0kubun/pp"
	"github.com/spf13/viper"
)

type fakeFlag struct {
	Key   string
	Value string
}
type fakeENVLoader struct {
	flags []fakeFlag
}

func (l *fakeENVLoader) Load(v viper.Viper) (*viper.Viper, error) {
	for idx := range l.flags {
		f := l.flags[idx]
		v.Set(f.Key, f.Value)
	}
	return &v, nil
}

func TestLoadConfig(t *testing.T) {
	fileLoader := NewFileLoader(".env.sample", "../..")
	errFileLoader := NewFileLoader(".env.err", "../..")
	type args struct {
		loaders []Loader
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Unable to load env",
			args: args{
				loaders: []Loader{
					errFileLoader,
					&fakeENVLoader{flags: []fakeFlag{}},
				},
			},
			want: Config{
				ServiceName:     "",
				BaseURL:         "",
				Port:            "8080",
				Env:             "dev",
				AllowedOrigins:  "",
				AccessTokenTTL:  0,
				RefreshTokenTTL: 0,
			},
		},
		{
			name: "Load from env",
			args: args{
				loaders: []Loader{
					fileLoader,
					&fakeENVLoader{flags: []fakeFlag{}},
				},
			},
			want: Config{
				ServiceName:     "tiny",
				BaseURL:         "http://localhost:8020",
				Port:            "8020",
				Env:             "dev",
				AllowedOrigins:  "",
				AccessTokenTTL:  600,
				RefreshTokenTTL: 7776000,
				DBHost:          "127.0.0.1",
				DBPort:          "5432",
				DBUser:          "user",
				DBName:          "dbname",
				DBPass:          "pass",
				DBSSLMode:       "disable",
			},
		},
		{
			name: "Load from flag env",
			args: args{
				loaders: []Loader{
					fileLoader,
					&fakeENVLoader{flags: []fakeFlag{}},
					NewENVLoader(),
				},
			},
			want: Config{
				ServiceName:     "tiny",
				BaseURL:         "http://localhost:8020",
				Port:            "8020",
				Env:             "dev",
				AllowedOrigins:  "",
				AccessTokenTTL:  600,
				RefreshTokenTTL: 7776000,
				DBHost:          "127.0.0.1",
				DBPort:          "5432",
				DBUser:          "user",
				DBName:          "dbname",
				DBPass:          "pass",
				DBSSLMode:       "disable",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadConfig(tt.args.loaders)
			if diff := deep.Equal(got, tt.want); diff != nil {
				pp.Println(got)
				t.Error(diff)
			}
		})
	}
}

func TestConfig_GetCORS(t *testing.T) {
	type fields struct {
		AllowedOrigins string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Get cors list",
			fields: fields{
				AllowedOrigins: "localhost:8020;",
			},
			want: []string{
				"localhost:8020",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				AllowedOrigins: tt.fields.AllowedOrigins,
			}
			if got := c.GetCORS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.GetCORS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultConfigLoaders(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "Get default config with 2 item",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultConfigLoaders(); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("DefaultConfigLoaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
