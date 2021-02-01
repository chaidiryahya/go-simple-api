package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSampleGetWithRouteAndQueryParam(t *testing.T) {
	type args struct {
		UserID int64
	}

	const (
		SuccessWithUserID100 = "SuccessSuccessWithUserID100"
		SuccessWithUserID101 = "SuccessSuccessWithUserID101"
		FailUserIDNotFound   = "FailUserIDNotFound"
	)

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr error
	}{
		{
			SuccessWithUserID100, args{UserID: 100}, SampleGetResponse{
				UserID:      100,
				Name:        "John Doe",
				Age:         "16 Years Old",
				PhoneNumber: "12345",
			}, nil,
		},
		{
			SuccessWithUserID101, args{UserID: 101}, SampleGetResponse{
				UserID:      101,
				Name:        "Maria",
				Age:         "17 Years Old",
				PhoneNumber: "54321",
			}, nil,
		},
		{
			FailUserIDNotFound, args{UserID: 102}, SampleGetResponse{}, errors.New("UserID not found"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SampleGetWithRouteAndQueryParam(tt.args.UserID)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
