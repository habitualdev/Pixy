package embed

import (
	"embed"
)

//go:embed  ldlinux.e64 ldlinux.c32 libcom32.c32 libutil.c32 lpxelinux.0 menu.c32 pxeboot.0 pxelinux.0 syslinux.efi vesamenu.c32

var F embed.FS

