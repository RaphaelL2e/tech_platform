package jwtauth

import (
	"testing"
	"time"
)

func TestJWTHelper_GenToken(t *testing.T) {
	type fields struct {
		Conf JWTConf
	}
	tests := []struct {
		name            string
		fields          fields
		wantTokenString string
		wantErr         bool
	}{
		{name: "simple gen", fields: fields{Conf: JWTConf{
			Key:      "hello",
			Duration: 10 * 365 * 24 * 60 * 60, // 10 year
		}}, wantTokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE4OTMxOTY4MDAsImlzcyI6InNlbnNlYmVlZHNlIn0.CL6TkbYTPy6VwW5ek7nlMqq7b8nDSnh7NWbgfn9Cv6w", wantErr: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := &JWTHelper{
				Conf: tt.fields.Conf,
			}
			gotTokenString, err := h.GenToken(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).Unix())
			if (err != nil) != tt.wantErr {
				t.Errorf("GenToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTokenString != tt.wantTokenString {
				t.Errorf("GenToken() gotTokenString = %v, want %v", gotTokenString, tt.wantTokenString)
			}
		})
	}
}

func TestJWTHelper_VerifyToken(t *testing.T) {
	type fields struct {
		Conf JWTConf
	}
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "not token", fields: fields{Conf: JWTConf{
			Key: "hello",
		}}, args: args{tokenString: "123123"}, wantErr: true},
		{name: "out date token", fields: fields{Conf: JWTConf{
			Key: "hello",
		}}, args: args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1Nzc4MDgwMDAsImlzcyI6InNlbnNlYmVlZHNlIn0.cjjp-eCYXM0_riPhzzVlZF3sCFU-KU0l0CHSbGsUyj4"}, wantErr: true},
		{name: "other sign method token", fields: fields{Conf: JWTConf{
			Key: "hello",
		}}, args: args{tokenString: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5MjQ1MTAwMjgsImlzcyI6InNlbnNlYmVlZHNlIn0.FJPR9SkXYohTSZkPanjZY5MG26ElLonIj_zHAnrAn6WtcYNeeErii5Bh_a8YFU0swGbfYBHDk4VRG6wxsANMrfO76viHPU1qYSe5avKj6Z9o5L0jXnD_oIBK23DAI7m4P2psacBWYjwvgfN4v62QcpCdwqX5Tk2d9F2oYDznsEyG0TNmymrTGVVnrFT0ora2VUBmZ8VD4TcvHRGGb1-MWtBf3MPxfbBbE-QjSnASjrrSB0g_PbDsXlrEiEw5crwsD0WoITiq023jjSagqeFzBIen4zkV7rj4MvF6ufMow6Ffew5jr65bsC9qMFTTZAxidCDopRsX76-DVUjPf6xEKlEEQSWcC3KIvGE_SLMFgx4EnF75AvkeRwVwr9ExOWm71NA85JpmP-M0bRMAY7qj-SAn_mkNb-2_6q4jXYJLf8eNtgvH_b47lo-Bfx6A-K5aVVQLzO4l2W5uCL60w9cKQDYk_2BNk3J4ViHV8eq3oF85AMPfPPnARoVq6r9AWVe7q09TEZvWpt-bWDosScP36y54BDaxtx0JXMRlGzdOQXDXsMX84lilwLKSjk0wfJ-tcXTObXU818EpN6wG2y0asQOYfhMk5IqHoWCoZ1_aOZQxQeQmshOqg4LB6SXUezjga5N6pkpYXduf2VT-iS6P1ynbRjfom9hrm9gkvCWvDyE"}, wantErr: true},
		{name: "other key token", fields: fields{Conf: JWTConf{
			Key: "hello",
		}}, args: args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5MjQ1MTAwMjgsImlzcyI6InNlbnNlYmVlZHNlIn0.ZNUsTllCNB60wyDekGR9i-8SxQpC04aacWO63YwXsZg"}, wantErr: true},
		{name: "valid token", fields: fields{Conf: JWTConf{
			Key: "hello",
		}}, args: args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE5MjQ1MTAwMjgsImlzcyI6InNlbnNlYmVlZHNlIn0.95j3-iLkZyzgpuZMrH30QencVAVfLV3epWf2oAlVPbI"}, wantErr: false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := &JWTHelper{
				Conf: tt.fields.Conf,
			}
			if err := h.VerifyToken(tt.args.tokenString); (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
