package rbox

import (
	"io/ioutil"
	"testing"
)

func TestNewRBox(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	var b interface{} = NewRBox(boxname)
	if _, ok := b.(RemoteBox); !ok {
		t.Fatal("Error: not a RemoteBox type")
	}
}

func TestRBoxSetSSID(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)

	tests := []struct {
		given string
		want  []string
	}{
		{"z3n z3n", []string{"z3n z3n", "z3n z3n"}},
		{"z3n", []string{"z3n", "z3n"}},
	}
	for _, test := range tests {
		err := b.SetSSIDs(test.given)
		if err != nil {
			t.Fatalf("Error setting SSIDs (%v): %v.", test.given, err)
		}

		SSIDs, err := b.GetSSIDs()
		for i, SSID := range SSIDs {
			if SSID != test.want[i] {
				t.Fatalf("Error setting SSID. Want %v, got %v.", test.want[i], SSID)
			}
		}
	}
}

func TestRBoxGetSSIDs(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)

	SSIDs, err := b.GetSSIDs()
	if err != nil {
		t.Fatalf("Error getting SSIDs: %v.", err)
	}

	for _, SSID := range SSIDs {
		if SSID != "z3n" {
			t.Fatalf("Error getting SSID. Want: z3n, got: '%v'.", SSID)
		}
	}
}

func TestRBoxSetMACs(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)

	tests := []struct {
		given string
		want  []string
	}{
		{"11:11:11:11:11:11", []string{"11:11:11:11:11:11"}},
		{"11:11:11:11:11:11 22:22:22:22:22:22", []string{"11:11:11:11:11:11", "22:22:22:22:22:22"}},
		{"11:11:11:11:11:11 22:22:22:22:22:22 33:33:33:33:33:33", []string{"11:11:11:11:11:11", "22:22:22:22:22:22", "33:33:33:33:33:33"}},
	}
	for _, test := range tests {
		if err := b.SetMACs(test.given); err != nil {
			t.Fatalf("Error setting MACs: %v.", err)
		}
		MACs, _ := b.GetMACs()
		for i, MAC := range MACs {
			if test.want[i] != MAC {
				t.Fatalf("Error setting MACs. Want %v, got %v.", test.want[i], MAC)
			}
		}
	}
}

func TestRBoxGetMACs(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)

	MACs, err := b.GetMACs()
	if err != nil {
		t.Fatalf("Error getting MACs: %v.", err)
	}

	want := []string{
		"11:11:11:11:11:11",
		"22:22:22:22:22:22",
		"33:33:33:33:33:33",
	}
	for i, MAC := range MACs {
		if want[i] != MAC {
			t.Fatalf("Error getting MAC. Want: %v, got: %v.", want[i], MAC)
		}
	}
}

func TestRBoxGetConfig(t *testing.T) {
	boxname := "788a20298f81.z3n.com.br"
	b := NewRBox(boxname)
	err := b.GetConfig()
	if err != nil {
		t.Fatalf("Error getting config for %v.", boxname)
	}

	f1, err := ioutil.ReadFile("./uci_show")
	if err != nil {
		t.Fatalf("Error reading uci_show: %v.", err)
	}

	f2, err := ioutil.ReadFile(boxname)
	if err != nil {
		t.Fatalf("Error reading %v: %v.", boxname, err)
	}

	if len(f1) != len(f2) {
		t.Fatal("Error config files are not the same size.")
	}

	for index, data := range f1 {
		if data != f2[index] {
			t.Fatal("Error: files are different.")
		}
	}
}
