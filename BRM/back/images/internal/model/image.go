package model

type Image []byte

const ImageMaxSize = 5 * 1024 * 1024 // 5 MB

var PermittedImageTypes = map[string]struct{}{
	"image/jpeg": {},
	"image/png":  {},
}
