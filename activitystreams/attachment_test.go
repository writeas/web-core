package activitystreams

import (
	"reflect"
	"testing"
)

func TestNewImageAttachment(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *Attachment
	}{
		{name: "good svg", args: args{"https://writefreely.org/img/writefreely.svg"}, want: &Attachment{
			Type:      "Image",
			URL:       "https://writefreely.org/img/writefreely.svg",
			MediaType: "image/svg+xml",
		}},
		{name: "good png", args: args{"https://i.snap.as/12345678.png"}, want: &Attachment{
			Type:      "Image",
			URL:       "https://i.snap.as/12345678.png",
			MediaType: "image/png",
		}},
		{name: "no extension", args: args{"https://i.snap.as/12345678"}, want: &Attachment{
			Type:      "Image",
			URL:       "https://i.snap.as/12345678",
			MediaType: "",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageAttachment(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
