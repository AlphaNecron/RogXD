package dmi

const (
	ProductFamily DmiProp = "product_family"
	BoardName             = "board_name"
	BiosDate              = "bios_date"
	BiosRelease           = "bios_release"
	BiosVendor            = "bios_vendor"
)

type DmiProp string
