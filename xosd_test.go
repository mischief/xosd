package xosd

import (
	"testing"
)

func TestDisplayString(t *testing.T) {
	osd, err := NewXOSD(1)
	if err != nil {
		t.Errorf("Failed to create XOSD: %v", err)
		return
	}

	if err := osd.SetTimeout(3); err != nil {
		t.Errorf("Failed to set XOSD timeout: %v", err)
		return
	}

	if err := osd.SetFont("-misc-fixed-*-*-*-*-20-*-*-*-*-*-*-*"); err != nil {
		t.Errorf("Failed to set XOSD font: %v", err)
		return
	}

	if err := osd.SetAlign(Center); err != nil {
		t.Errorf("Failed to set XOSD alignment: %v", err)
		return
	}

	if err := osd.DisplayString(0, "Hello, World!"); err != nil {
		t.Errorf("Failed to display XOSD string: %v", err)
		return
	}

	if err := osd.Wait(); err != nil {
		t.Errorf("Failed to display XOSD string: %v", err)
		return
	}
}
