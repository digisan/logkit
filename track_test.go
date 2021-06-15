package logkit

import (
	"testing"
)

func Test_trackCaller(t *testing.T) {
	fPln(track(0))
}

func TestTrackCaller(t *testing.T) {
	fPln(TrackCaller(0))
}

func TestCallerSrc(t *testing.T) {
	fPln(CallerSrc())
}

func TestCaller(t *testing.T) {
	fPln(Caller(true))
	fPln(Caller(false))
}
