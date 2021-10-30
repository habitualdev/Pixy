package embed

import (
	"embed"
	_ "embed"
)

//go:embed ldlinux.c32
//go:embed ldlinux.e64
//go:embed libcom32.c32
//go:embed libutil.c32
//go:embed  lpxelinux.0
//go:embed menu.c32
//go:embed pxeboot.0
//go:embed syslinux.efi
//go:embed vesamenu.c32

var F embed.FS
