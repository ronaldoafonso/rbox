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
	bnameOK := "788a20298f81.z3n.com.br"
	b1 := NewRBox(bnameOK)
	if SSIDs, err := b1.GetSSIDs(); err != nil {
		t.Fatalf("Error getting SSIDs: %v.", err)
	} else {
		for _, SSID := range SSIDs {
			if SSID != "z3n" {
				t.Fatalf("Error GetSSIDs. Want: z3n, got: '%v'.", SSID)
			}
		}
	}

	bnameNotOK := "notok.z3n.com.br"
	b2 := NewRBox(bnameNotOK)
	if SSIDs, err := b2.GetSSIDs(); SSIDs != nil && err == nil {
		t.Fatal("GetSSIDs should return 'nil' and an error.")
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
	bnameOK := "788a20298f81.z3n.com.br"
	b1 := NewRBox(bnameOK)

	MACs, err := b1.GetMACs()
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

	bnameNotOK := "notok.z3n.com.br"
	b2 := NewRBox(bnameNotOK)
	if MACs, err := b2.GetMACs(); MACs != nil && err == nil {
		t.Fatal("GetMACs should return 'nil' and an error.")
	}
}

func TestRBoxGetConfig(t *testing.T) {
	wantConfig, err := ioutil.ReadFile("./uci_show")
	if err != nil {
		t.Fatalf("Error reading uci_show: %v.", err)
	}

	bnameOK := "788a20298f81.z3n.com.br"
	b1 := NewRBox(bnameOK)
	if config, err := b1.GetConfig(); err != nil {
		t.Fatalf("Error getting config for %v: %v.", bnameOK, err)
	} else {
		for i, data := range config {
			if data != wantConfig[i] {
				t.Fatalf("Content error. Want %v, got %v.", wantConfig[i], data)
			}
		}
	}

	bnameNotOK := "notok.z3n.com.br"
	b2 := NewRBox(bnameNotOK)
	if config, err := b2.GetConfig(); config != nil && err == nil {
		t.Fatal("GetConfig should return 'nil' and an error.")
	}
}
