package main

import "testing"

// The WhatsApp media CDN rejects unsigned paths with 403, so the direct path
// must keep the signed query string (oh/oe) from the stored URL.
func TestExtractDirectPathKeepsSignature(t *testing.T) {
	cases := map[string]string{
		"https://mmg.whatsapp.net/o1/v/t24/f2/m239/AQN?ccb=9-4&oh=01_abc&oe=6ACBB0FC&_nc_sid=e6ed6c&mms3=true": "/o1/v/t24/f2/m239/AQN?ccb=9-4&oh=01_abc&oe=6ACBB0FC&_nc_sid=e6ed6c",
		"https://mmg.whatsapp.net/v/t62.7118-24/123_n.enc?ccb=11-4&oh=01_def&oe=6ACBB0FC&_nc_sid=5e03e0&mms3=true": "/v/t62.7118-24/123_n.enc?ccb=11-4&oh=01_def&oe=6ACBB0FC&_nc_sid=5e03e0",
	}
	for in, want := range cases {
		if got := extractDirectPathFromURL(in); got != want {
			t.Errorf("extractDirectPathFromURL(%q)\n got  %q\n want %q", in, got, want)
		}
	}
}
