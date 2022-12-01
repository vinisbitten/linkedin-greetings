package main

import "testing"

func TestWelcome(t *testing.T) {			            //█───█─▄▀▀─█───▄▀▀─▄▀▀▄─█▄─▄█─▄▀▀
	want := "Welcome!"						            //█───█─█───█───█───█──█─█▀▄▀█─█──
	got := Hi()							                //█───█─█▀▀─█───█───█──█─█─▀─█─█▀▀
	if want != got {                                    //█▄█▄█─█───█───█───█──█─█───█─█──
		t.Fatal("want:", want, "got:", got)             //─▀─▀───▀▀──▀▀──▀▀──▀▀──▀───▀──▀▀
	}
}
