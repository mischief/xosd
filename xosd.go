package xosd

// #cgo LDFLAGS: -lxosd
// #include <stdlib.h>
// #include <xosd.h>
// int
// xosd_display_string(xosd *xosd, int line, char *str)
// {
//	return xosd_display(xosd, line, XOSD_string, str);
// }
import "C"
import "unsafe"
import "errors"
import "sync"

type Align C.xosd_align

// Horizontal alignment types for libxosd
const (
	Left   Align = C.XOSD_left
	Center Align = C.XOSD_center
	Right  Align = C.XOSD_right
)

// XOSD is a wrapper around libxosd.
type XOSD struct {
	osd *C.xosd
	mu  sync.Mutex
}

// NewXOSD created a new XOSD instance. 'lines' is the maximum number of lines
// that can be drawn.
func NewXOSD(lines int) (*XOSD, error) {
	osd := C.xosd_create(C.int(lines))
	if osd == nil {
		return nil, errors.New(C.GoString(C.xosd_error))
	}

	x := &XOSD{
		osd: osd,
	}

	return x, nil
}

// SetFont sets the font used to render displayed text.
// The input must be a X logical font description name.
func (x *XOSD) SetFont(xfont string) error {
	cs := C.CString(xfont)
	defer C.free(unsafe.Pointer(cs))

	x.mu.Lock()
	defer x.mu.Unlock()

	r := C.xosd_set_font(x.osd, cs)
	if r == -1 {
		return errors.New(C.GoString(C.xosd_error))
	}

	return nil
}

// SetTimeout sets the number of seconds text will be displayed on screen.
func (x *XOSD) SetTimeout(seconds int) error {
	x.mu.Lock()
	defer x.mu.Unlock()

	r := C.xosd_set_timeout(x.osd, C.int(seconds))
	if r == -1 {
		return errors.New(C.GoString(C.xosd_error))
	}

	return nil
}

// SetAlign sets the alignment of the text along the horizontal axis.
func (x *XOSD) SetAlign(alignment Align) error {
	x.mu.Lock()
	defer x.mu.Unlock()

	r := C.xosd_set_align(x.osd, C.xosd_align(alignment))
	if r == -1 {
		return errors.New(C.GoString(C.xosd_error))
	}

	return nil
}

// DisplayString draws a string 'str' at line index 'line'.
func (x *XOSD) DisplayString(line int, str string) error {
	cs := C.CString(str)
	defer C.free(unsafe.Pointer(cs))

	x.mu.Lock()
	defer x.mu.Unlock()

	r := C.xosd_display_string(x.osd, C.int(line), cs)
	if r == -1 {
		return errors.New(C.GoString(C.xosd_error))
	}

	return nil
}

// Wait blocks until no message is displayed.
func (x *XOSD) Wait() error {
	x.mu.Lock()
	defer x.mu.Unlock()

	r := C.xosd_wait_until_no_display(x.osd)
	if r == -1 {
		return errors.New(C.GoString(C.xosd_error))
	}

	return nil
}
