package main

import (
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {

	type args struct {
		config *ServerConfig
	}
	tests := []struct {
		name          string
		args          args
		want          *Server
		wantErr       bool
		expectedError error
	}{
		{
			name: "WtiteTimeOut overflows int64",
			args: args{
				config: &ServerConfig{
					Address:      "127.0.0.1:8080",
					WriteTimeout: 15 * time.Second,
					ReadTimeout:  15,
				},
			},
			want:          nil,
			wantErr:       true,
			expectedError: ErrWriteTimeoutOverFlows,
		},
		{
			name: "ReadTimeOut Overflows Int64",
			args: args{
				config: &ServerConfig{
					Address:      "127.0.0.1:8080",
					WriteTimeout: 15,
					ReadTimeout:  15 * time.Second,
				},
			},
			want:          nil,
			wantErr:       true,
			expectedError: ErrReadTimeoutOverFlows,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewServer(tt.args.config)

			if got != tt.want && got.DeepEqual(tt.want) {
				t.Errorf("NewServer() returned wrong server instance config %v %v", *got.config, *tt.want.config)
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("NewServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil && err != tt.expectedError {
				t.Errorf("UnExpected error Type = %v, Expected Error %v", err, tt.expectedError)
				return
			}
		})
	}
}
