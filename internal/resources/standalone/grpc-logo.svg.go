package standalone

import (
	"bytes"
	"compress/gzip"
	"io"
)

// grpc_logo_svg returns raw, uncompressed file data.
func grpc_logo_svg() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x9c, 0x59,
		0xcb, 0x6e, 0x5c, 0x49, 0x72, 0xdd, 0xcf, 0x57, 0xa4, 0xcb, 0x1b, 0x1b,
		0xb8, 0x0a, 0xc6, 0xfb, 0xd1, 0x10, 0x35, 0x80, 0x0b, 0x98, 0xd9, 0xa8,
		0x81, 0x81, 0xdb, 0x86, 0x81, 0xde, 0x18, 0x72, 0xab, 0x9a, 0x24, 0x5a,
		0x4d, 0x0a, 0x22, 0x47, 0x92, 0xff, 0xde, 0x88, 0x2c, 0x52, 0xcc, 0xdb,
		0x16, 0x01, 0xc3, 0x5a, 0x89, 0xe7, 0x66, 0xc5, 0x8d, 0xe7, 0x89, 0x93,
		0x55, 0xaf, 0xff, 0xfc, 0xf5, 0xf7, 0x0f, 0xe3, 0xf3, 0xe9, 0xd3, 0xfd,
		0xcd, 0xdd, 0xed, 0xe5, 0x81, 0x00, 0x0f, 0xe3, 0x74, 0xfb, 0xcb, 0xdd,
		0xfb, 0x9b, 0xdb, 0xab, 0xcb, 0xc3, 0xbf, 0xff, 0xdb, 0x5f, 0x5e, 0xe5,
		0x61, 0xdc, 0x3f, 0xbc, 0xbb, 0x7d, 0xff, 0xee, 0xc3, 0xdd, 0xed, 0xe9,
		0xf2, 0x70, 0x7b, 0x77, 0xf8, 0xf3, 0x9b, 0x3f, 0xbd, 0xbe, 0xff, 0x7c,
		0x35, 0xbe, 0xdc, 0xbc, 0x7f, 0xb8, 0xbe, 0x3c, 0x30, 0xe2, 0xc7, 0xaf,
		0x87, 0x71, 0x7d, 0xba, 0xb9, 0xba, 0x7e, 0xb8, 0x3c, 0x04, 0xf5, 0x5f,
		0x9f, 0x6f, 0x4e, 0x5f, 0xfe, 0xe5, 0xee, 0xeb, 0xe5, 0x01, 0x07, 0x0e,
		0x46, 0x1c, 0x41, 0x87, 0xf5, 0x35, 0x74, 0x18, 0x5f, 0x7f, 0xff, 0x70,
		0x7b, 0x7f, 0x79, 0xb8, 0x7e, 0x78, 0xf8, 0xf8, 0xc3, 0xc5, 0xc5, 0x97,
		0x2f, 0x5f, 0xe0, 0x8b, 0xc0, 0xdd, 0xa7, 0xab, 0x0b, 0x46, 0xc4, 0x8b,
		0xfb, 0xcf, 0x57, 0x8f, 0x47, 0x7e, 0xf8, 0xfa, 0xe1, 0xe6, 0xf6, 0xb7,
		0xef, 0x1d, 0xa4, 0xaa, 0xba, 0x98, 0x4f, 0x9f, 0x8e, 0xde, 0xff, 0x76,
		0x7a, 0xf8, 0xe5, 0x7a, 0x77, 0xf6, 0xbf, 0xee, 0xae, 0x4f, 0xbf, 0xdf,
		0xbc, 0x7b, 0x8c, 0x09, 0x7e, 0xb9, 0xfb, 0xfd, 0xe2, 0x7c, 0xea, 0xe2,
		0xf6, 0xfe, 0xf0, 0xe6, 0x4f, 0x63, 0x8c, 0xf1, 0xfa, 0x1f, 0x5e, 0xbd,
		0x1a, 0x7f, 0x3d, 0xdd, 0x9e, 0x3e, 0xbd, 0x7b, 0xb8, 0xfb, 0xf4, 0xc3,
		0xf8, 0x69, 0x3e, 0x1f, 0x02, 0x02, 0x3c, 0xfe, 0x89, 0x18, 0x55, 0xfe,
		0x79, 0xbc, 0x1a, 0xff, 0x17, 0xa3, 0xe3, 0xd5, 0xab, 0x47, 0x9b, 0x0f,
		0x37, 0x0f, 0x1f, 0x4e, 0x6f, 0xae, 0x3e, 0x7d, 0xfc, 0xe5, 0x3f, 0xbf,
		0x5c, 0xdf, 0x3c, 0x9c, 0x5e, 0x5f, 0x9c, 0x91, 0xf3, 0xd3, 0xf7, 0xa7,
		0xfb, 0x5f, 0xde, 0x1c, 0x3f, 0x9d, 0xde, 0x3d, 0x9c, 0xde, 0x8f, 0x2f,
		0x37, 0x0f, 0xd7, 0x8f, 0x2f, 0x85, 0xd7, 0x17, 0xf3, 0xd1, 0xd3, 0xa9,
		0x5f, 0xef, 0xdf, 0x34, 0xf2, 0xeb, 0xfd, 0x23, 0x72, 0x35, 0x6e, 0xde,
		0x5f, 0x1e, 0xae, 0xfe, 0xf5, 0x6f, 0xc7, 0x57, 0x6f, 0xef, 0xae, 0xee,
		0x5e, 0xfd, 0x47, 0x9b, 0xee, 0x0a, 0x7d, 0xba, 0xfb, 0x6d, 0x56, 0xe7,
		0xf6, 0xdb, 0x5f, 0xaf, 0x1e, 0xeb, 0x43, 0x87, 0xf1, 0xeb, 0xcd, 0x87,
		0x0f, 0x4f, 0x0f, 0xfb, 0xff, 0xaf, 0x3e, 0xfd, 0xfd, 0xc3, 0xe9, 0xf2,
		0x70, 0xfa, 0x7c, 0xba, 0xbd, 0x7b, 0xff, 0xfe, 0x30, 0xce, 0xbe, 0xff,
		0xf0, 0xf0, 0xdf, 0x1f, 0x4f, 0x97, 0x87, 0x1f, 0x7f, 0xfa, 0xdb, 0xbb,
		0xab, 0xd3, 0x63, 0x6a, 0xe6, 0x4b, 0x3f, 0xbe, 0x7b, 0xb8, 0x1e, 0xef,
		0x2f, 0x0f, 0x3f, 0x86, 0x82, 0x99, 0x27, 0x57, 0x6e, 0x82, 0x10, 0xc2,
		0x99, 0xaa, 0xe3, 0x6d, 0x22, 0x30, 0x5a, 0x89, 0xf9, 0x0a, 0x1f, 0x93,
		0xa0, 0xac, 0x88, 0x8c, 0x56, 0x38, 0x05, 0xc4, 0x8c, 0x54, 0xad, 0x51,
		0xad, 0xcc, 0x54, 0x1f, 0xa9, 0x20, 0xa5, 0xe8, 0x16, 0x8d, 0xa2, 0x60,
		0x46, 0x8e, 0x63, 0x1a, 0xa8, 0x70, 0xa5, 0xd7, 0xc6, 0x05, 0xe6, 0x9c,
		0x41, 0x36, 0xd2, 0x81, 0x39, 0xa5, 0x6a, 0xe3, 0x84, 0x2c, 0xc6, 0x60,
		0x6a, 0x30, 0x12, 0x85, 0x6d, 0xa2, 0x48, 0xa9, 0x56, 0x6d, 0x21, 0x40,
		0x84, 0xd9, 0x38, 0x37, 0x0e, 0x20, 0xd5, 0xd4, 0xf0, 0x91, 0x01, 0x81,
		0xe6, 0x58, 0xb6, 0xb1, 0x03, 0x7a, 0xaa, 0xe6, 0x44, 0x0b, 0x51, 0x31,
		0x64, 0x63, 0x85, 0xc8, 0x62, 0x41, 0x1f, 0xc7, 0x4c, 0xc0, 0x32, 0x46,
		0xdb, 0x58, 0xc0, 0x08, 0x91, 0x38, 0x46, 0x26, 0x50, 0xb1, 0x23, 0xeb,
		0xc6, 0x0c, 0xc8, 0x89, 0xa4, 0xb4, 0x43, 0x11, 0x44, 0x85, 0x0a, 0x6d,
		0x5a, 0x78, 0x82, 0x29, 0xc1, 0x2d, 0xc5, 0x3d, 0xc7, 0xb3, 0x5d, 0x0a,
		0xa0, 0x70, 0xf1, 0xe4, 0xd5, 0x07, 0x32, 0xc8, 0x0a, 0x32, 0x94, 0x19,
		0xc6, 0x93, 0xc3, 0xa4, 0xe0, 0x14, 0x25, 0x2c, 0x7d, 0x98, 0xab, 0x22,
		0x63, 0x23, 0x01, 0x63, 0x2b, 0x11, 0xed, 0x44, 0x74, 0x85, 0x38, 0x7c,
		0x23, 0x06, 0x67, 0x22, 0xc6, 0x18, 0xc7, 0xec, 0x38, 0x1d, 0x3d, 0x69,
		0x23, 0x82, 0x20, 0x17, 0xcc, 0x91, 0xd6, 0x85, 0x0b, 0x56, 0x6d, 0x10,
		0xd5, 0x0c, 0x73, 0xd6, 0x82, 0x90, 0x32, 0xd0, 0x37, 0x42, 0x70, 0xcc,
		0x40, 0xee, 0x28, 0x18, 0xaa, 0x12, 0xc9, 0xb3, 0x61, 0x0a, 0xca, 0x72,
		0x19, 0x49, 0x60, 0x12, 0x68, 0x96, 0x5b, 0x41, 0x99, 0x68, 0xb9, 0xf9,
		0x88, 0x82, 0x7e, 0xae, 0x12, 0x2b, 0xfa, 0x76, 0xe9, 0x9f, 0x17, 0xe0,
		0xb5, 0xad, 0xbe, 0x0f, 0xff, 0x3c, 0x7e, 0x74, 0x05, 0x67, 0xa6, 0x32,
		0xdd, 0x10, 0x2c, 0x4d, 0xab, 0xa2, 0x62, 0xf6, 0xa1, 0x95, 0x19, 0x17,
		0xed, 0xf0, 0x63, 0x31, 0x70, 0x8a, 0x14, 0xfb, 0x0e, 0xaf, 0x04, 0xe2,
		0x60, 0xf1, 0xd8, 0x02, 0x24, 0x25, 0xc4, 0x90, 0x57, 0x94, 0x11, 0x2a,
		0x51, 0xb1, 0x74, 0x1c, 0x57, 0xd8, 0x00, 0x49, 0x32, 0xd8, 0x47, 0x05,
		0x68, 0x29, 0x99, 0x78, 0x77, 0x9d, 0x46, 0xa2, 0x1a, 0x8d, 0xea, 0x06,
		0x8d, 0xf2, 0xd4, 0x4d, 0x08, 0x24, 0xa4, 0xfb, 0x79, 0x1c, 0x4b, 0xa1,
		0x9c, 0x22, 0x85, 0x37, 0x51, 0x60, 0xef, 0x5c, 0xda, 0x28, 0x86, 0x90,
		0xc8, 0xb0, 0xdc, 0xc4, 0xc1, 0x4a, 0x82, 0xbb, 0x06, 0xd5, 0xb1, 0x53,
		0x94, 0x6f, 0x92, 0xd0, 0x8e, 0x89, 0xc7, 0x78, 0x4b, 0x88, 0xc0, 0x81,
		0x88, 0xb4, 0x05, 0x02, 0xba, 0x26, 0x16, 0x8d, 0xb7, 0x59, 0x10, 0x86,
		0x15, 0xae, 0x7b, 0xb8, 0x7b, 0x90, 0xd0, 0xb9, 0x36, 0x29, 0xc8, 0x54,
		0x93, 0xf8, 0x43, 0x5e, 0xbf, 0x0f, 0xaf, 0x46, 0x96, 0x64, 0xbf, 0x00,
		0xef, 0x6a, 0xf0, 0x02, 0xfe, 0xf3, 0xf8, 0x91, 0xd0, 0x81, 0x48, 0x0a,
		0xff, 0xf0, 0x09, 0x62, 0x02, 0x24, 0x4b, 0xb3, 0x7d, 0xd5, 0x88, 0x05,
		0x42, 0x05, 0x59, 0x76, 0x38, 0xb1, 0x03, 0xa9, 0x69, 0xf1, 0x86, 0x50,
		0x2a, 0xd2, 0x59, 0xa4, 0x41, 0x9c, 0xc0, 0x2c, 0x22, 0xb6, 0x51, 0x8f,
		0x58, 0x91, 0x63, 0x8d, 0x23, 0x09, 0x82, 0x20, 0x51, 0xc4, 0xc6, 0x20,
		0xa1, 0x26, 0x68, 0x32, 0x48, 0x18, 0x28, 0x95, 0x4a, 0x37, 0x01, 0x57,
		0x21, 0xc4, 0xd0, 0x41, 0x22, 0x90, 0x3d, 0x0c, 0xb4, 0x19, 0xa8, 0xab,
		0xbb, 0x36, 0x09, 0x91, 0x18, 0x58, 0x4f, 0x4e, 0x6e, 0x01, 0x9c, 0xce,
		0x9c, 0x49, 0x83, 0xc4, 0x21, 0xb8, 0xd4, 0x7c, 0x2b, 0x50, 0xf2, 0x08,
		0xe7, 0x46, 0x03, 0xba, 0x98, 0xc9, 0x3d, 0x50, 0x69, 0x2e, 0x71, 0xf6,
		0x22, 0x9b, 0x23, 0x08, 0xb3, 0xc7, 0x97, 0xcb, 0x66, 0xdf, 0xd3, 0xac,
		0x2c, 0x87, 0x7b, 0x33, 0x80, 0x55, 0x18, 0x62, 0xad, 0x28, 0x13, 0x84,
		0x13, 0x86, 0xe5, 0xd9, 0xc4, 0x13, 0xac, 0x90, 0x3d, 0x9e, 0x44, 0xf3,
		0x70, 0x1f, 0x70, 0x6a, 0x7a, 0xeb, 0x28, 0x9a, 0x17, 0xda, 0x8b, 0x30,
		0x54, 0xe1, 0xa6, 0x4d, 0x4a, 0x4f, 0x14, 0x6f, 0x13, 0x01, 0x52, 0x22,
		0x28, 0xb3, 0x2d, 0xdd, 0x29, 0x34, 0x67, 0x20, 0xac, 0xe8, 0x59, 0x9b,
		0x18, 0x90, 0x85, 0xa6, 0xf0, 0x20, 0x69, 0x92, 0x60, 0xec, 0x1e, 0x89,
		0x0e, 0xb4, 0x4c, 0xa9, 0x6b, 0x92, 0x90, 0x51, 0x5e, 0xb3, 0xa3, 0x44,
		0xb2, 0x90, 0x72, 0xd0, 0xf4, 0xa8, 0xf3, 0x26, 0x05, 0x45, 0x11, 0x6e,
		0x3c, 0xba, 0xac, 0x44, 0xc2, 0x26, 0x2b, 0xfa, 0x96, 0xc8, 0x01, 0x35,
		0x4d, 0xf2, 0x25, 0x78, 0xed, 0xb2, 0xa5, 0x67, 0x5e, 0x80, 0xf7, 0xad,
		0xf4, 0x7d, 0xbc, 0x9b, 0x6f, 0x79, 0xed, 0x42, 0x31, 0xc4, 0x08, 0xc9,
		0x84, 0xb6, 0x83, 0x8f, 0xc4, 0x0c, 0x4d, 0x18, 0xbe, 0x83, 0x3b, 0x4c,
		0x29, 0x16, 0x9e, 0x3b, 0x4a, 0xa3, 0x89, 0xa6, 0xe3, 0x34, 0xd0, 0x72,
		0xca, 0xce, 0x75, 0x39, 0x62, 0x05, 0xb7, 0x05, 0x07, 0x47, 0x44, 0xd1,
		0x86, 0x55, 0xb5, 0xb0, 0xb3, 0xca, 0x01, 0x4a, 0xe4, 0x31, 0x97, 0x54,
		0xa8, 0x78, 0x74, 0xdf, 0x71, 0x40, 0x09, 0x91, 0x76, 0x0d, 0xbb, 0x67,
		0x2a, 0xeb, 0x9c, 0x6a, 0x35, 0xf4, 0xde, 0x50, 0xe5, 0x45, 0x8c, 0x32,
		0x9b, 0x3c, 0xac, 0x90, 0xad, 0x49, 0x28, 0xcb, 0x92, 0xcc, 0x26, 0x9a,
		0xe6, 0xda, 0x5b, 0x47, 0xc1, 0xc5, 0xd5, 0x58, 0xcf, 0x06, 0xca, 0x24,
		0x59, 0x7a, 0x71, 0x49, 0x04, 0x66, 0xb5, 0x89, 0x02, 0x44, 0x36, 0xe6,
		0x6e, 0xaf, 0x0a, 0x75, 0x89, 0xd8, 0xa1, 0x08, 0xca, 0x85, 0x66, 0x73,
		0x04, 0xbf, 0xc1, 0x94, 0xd0, 0xc9, 0xf4, 0x19, 0xf1, 0x37, 0xc3, 0x14,
		0xe0, 0x92, 0xaa, 0x54, 0xab, 0x17, 0xe4, 0xcd, 0x58, 0x6a, 0xfe, 0x18,
		0xc6, 0xa3, 0xcb, 0x64, 0x80, 0xce, 0x1a, 0xb3, 0x63, 0x12, 0xd4, 0x3d,
		0x45, 0x7a, 0x75, 0x95, 0x3a, 0x8b, 0xd6, 0x39, 0x11, 0x51, 0x29, 0xd5,
		0x28, 0x22, 0x05, 0xf5, 0xe2, 0x99, 0x59, 0x2b, 0x4e, 0xb5, 0xde, 0x68,
		0x68, 0x41, 0x85, 0x3c, 0x79, 0x20, 0xb8, 0xb2, 0xb2, 0x47, 0x4d, 0x48,
		0x50, 0xf8, 0x5c, 0x0d, 0x2f, 0xac, 0x88, 0x5e, 0x52, 0xe1, 0xc5, 0xe4,
		0xd2, 0x26, 0x14, 0xbc, 0x77, 0xed, 0x5c, 0x69, 0xcc, 0xa6, 0x88, 0x9d,
		0x78, 0x01, 0xb2, 0x6c, 0x7e, 0x5c, 0xd6, 0x51, 0x37, 0x2e, 0x13, 0xba,
		0xf1, 0x6e, 0x49, 0x2d, 0x1d, 0xf4, 0x02, 0xbc, 0x6b, 0xac, 0xef, 0xc2,
		0xdd, 0x88, 0x11, 0x20, 0xc4, 0x44, 0xbe, 0x99, 0x82, 0x69, 0xa1, 0x6b,
		0xc7, 0xb8, 0xc0, 0x0e, 0x5e, 0x56, 0x18, 0x39, 0x28, 0x1c, 0xaa, 0xf9,
		0x5b, 0x37, 0xeb, 0x7e, 0x31, 0x12, 0xf7, 0x89, 0x62, 0x75, 0xe1, 0x36,
		0x47, 0xf0, 0xf2, 0x24, 0xee, 0x18, 0xc3, 0x80, 0x93, 0x2b, 0x6a, 0x73,
		0x06, 0xd7, 0xd4, 0xc9, 0x53, 0xa1, 0x40, 0xce, 0x85, 0xb5, 0xb9, 0x82,
		0xf4, 0x4e, 0xc9, 0x36, 0xdc, 0x5b, 0x47, 0xa5, 0x64, 0x73, 0x83, 0xa4,
		0x6a, 0xf1, 0xd6, 0x26, 0x08, 0x04, 0x2d, 0xc3, 0x37, 0x6f, 0xae, 0x13,
		0x2f, 0xd6, 0x41, 0x5e, 0xe0, 0x22, 0x99, 0xb1, 0x79, 0xf7, 0xa4, 0x36,
		0x9f, 0x0c, 0xf2, 0xe8, 0x1d, 0x2f, 0x7d, 0xb6, 0xa9, 0x80, 0x4b, 0xe6,
		0xe8, 0xb4, 0x3d, 0xe4, 0x74, 0xeb, 0x89, 0x65, 0x22, 0xca, 0x6c, 0x13,
		0xcd, 0xe5, 0xc4, 0x19, 0x8d, 0xba, 0xf5, 0x42, 0xac, 0x41, 0xde, 0xe2,
		0x41, 0x2c, 0x6a, 0x45, 0x8f, 0x64, 0xd5, 0x7f, 0x98, 0xd8, 0xee, 0xb0,
		0x05, 0xf4, 0x86, 0x36, 0x6e, 0x54, 0x0a, 0x31, 0x4b, 0x07, 0x99, 0x41,
		0x64, 0x60, 0x58, 0x7b, 0x91, 0x5e, 0xae, 0xb3, 0x65, 0x4c, 0x20, 0xa5,
		0x9a, 0xbe, 0xda, 0x39, 0x2d, 0xaa, 0xae, 0xac, 0x31, 0x60, 0x5a, 0xa9,
		0x3d, 0x05, 0x52, 0x34, 0xc8, 0x10, 0x8c, 0xc3, 0x5c, 0x3a, 0x66, 0x4a,
		0x6f, 0xed, 0x33, 0x8e, 0xa4, 0x39, 0xc7, 0x2e, 0xa9, 0x33, 0x54, 0x14,
		0xd6, 0xbd, 0x48, 0x1a, 0xe0, 0xe9, 0xcd, 0xef, 0xae, 0xc0, 0xda, 0x6d,
		0x2b, 0x83, 0xb4, 0x35, 0x17, 0x36, 0x31, 0x38, 0x03, 0x39, 0x9a, 0x9f,
		0x4d, 0x18, 0x78, 0xb4, 0x50, 0xe8, 0x3a, 0x61, 0x64, 0x44, 0xb3, 0x88,
		0x1a, 0x50, 0xcc, 0xf6, 0xb2, 0x68, 0x66, 0x40, 0x23, 0xdf, 0xa1, 0x0a,
		0x6c, 0x2e, 0xa4, 0xdd, 0x5d, 0xcf, 0x70, 0x37, 0x14, 0x36, 0x69, 0xd3,
		0xd9, 0xf2, 0x13, 0x2c, 0x10, 0xec, 0x81, 0x4e, 0xd3, 0x86, 0xb5, 0xca,
		0x9c, 0xcb, 0xc7, 0x89, 0x95, 0xbb, 0x7a, 0xea, 0x20, 0xd2, 0x8b, 0x72,
		0x2b, 0x70, 0xc7, 0x50, 0xf7, 0x2e, 0xb5, 0x06, 0x10, 0xa5, 0xb3, 0x6e,
		0xad, 0x2d, 0x0b, 0xa3, 0x99, 0xb5, 0xa3, 0x66, 0x66, 0x75, 0xda, 0x1c,
		0x90, 0xc2, 0x55, 0xaa, 0x9d, 0x9b, 0x05, 0xa9, 0xc8, 0x4d, 0xc1, 0xd2,
		0x95, 0xac, 0xf5, 0x0c, 0x19, 0x01, 0x46, 0xe9, 0xdc, 0xa2, 0x64, 0xd6,
		0xd6, 0x66, 0x8e, 0x23, 0xa5, 0xb4, 0x36, 0x06, 0x6c, 0xf1, 0x5e, 0x1a,
		0x83, 0x4c, 0x61, 0x1a, 0xe3, 0xad, 0x87, 0xab, 0xa2, 0x68, 0x52, 0x94,
		0x39, 0x84, 0x6a, 0x4a, 0x0b, 0xb3, 0x56, 0xff, 0xd5, 0xdb, 0x7c, 0x90,
		0x25, 0x94, 0x38, 0x47, 0x6c, 0x38, 0x9b, 0x44, 0x24, 0x23, 0x65, 0xc3,
		0x6e, 0x2f, 0x01, 0x27, 0xd2, 0x16, 0x05, 0xa3, 0x7b, 0x2d, 0x98, 0x62,
		0xff, 0x61, 0x0f, 0x70, 0x2f, 0x6f, 0x06, 0x58, 0x5f, 0xd5, 0x4d, 0x4c,
		0xe1, 0xa6, 0x3b, 0xbf, 0x9e, 0x1b, 0x5e, 0x5a, 0xd2, 0xaa, 0x56, 0xf1,
		0x3a, 0x1d, 0xcd, 0xa7, 0xc6, 0xe8, 0x31, 0x07, 0xec, 0x79, 0x94, 0x00,
		0xab, 0xe5, 0xaf, 0xe9, 0x58, 0xc6, 0xae, 0xd9, 0xdb, 0xb4, 0xe5, 0xd4,
		0x3a, 0xa3, 0x05, 0x85, 0x1a, 0xa8, 0xd5, 0x85, 0x7b, 0x1e, 0x68, 0xea,
		0xdb, 0x91, 0x94, 0x77, 0x4b, 0x3c, 0x0f, 0x3f, 0x29, 0x50, 0x32, 0x96,
		0xf2, 0x0e, 0x75, 0x78, 0x52, 0x34, 0x6f, 0x17, 0xb8, 0xa9, 0xba, 0x58,
		0x8b, 0x1b, 0xf6, 0x56, 0x1e, 0xa1, 0xc9, 0x2f, 0xc1, 0xf3, 0x96, 0xa1,
		0xdd, 0x01, 0x9d, 0x8d, 0x67, 0xd8, 0xba, 0xe2, 0x95, 0x36, 0xc8, 0x1d,
		0x92, 0xad, 0x27, 0xa1, 0xef, 0x13, 0x42, 0x66, 0xf9, 0x2d, 0xc7, 0xd4,
		0xdc, 0x4b, 0xd8, 0x99, 0xeb, 0x38, 0x7c, 0xde, 0x42, 0xf2, 0xcc, 0xa7,
		0xee, 0xd5, 0x4f, 0xe6, 0x94, 0x13, 0xab, 0xfe, 0x81, 0x4f, 0x9d, 0x80,
		0x95, 0x24, 0x75, 0x45, 0x8f, 0x5d, 0xe4, 0x08, 0x29, 0xf5, 0xdd, 0x61,
		0x0b, 0x20, 0xd1, 0xae, 0x6f, 0xf3, 0x37, 0x45, 0x16, 0xf5, 0x94, 0x3b,
		0x08, 0x9f, 0x95, 0x12, 0x03, 0xab, 0x07, 0x4e, 0x35, 0x62, 0x06, 0x46,
		0x44, 0x38, 0xf7, 0x45, 0x84, 0x19, 0x79, 0x4e, 0x4a, 0x20, 0xb4, 0xe6,
		0x1a, 0x32, 0x08, 0xc2, 0xac, 0xb2, 0x1d, 0x9a, 0x80, 0xbd, 0xbb, 0xb1,
		0x79, 0xfa, 0x19, 0x36, 0x01, 0x0b, 0x61, 0x21, 0x3b, 0x5b, 0x7e, 0x82,
		0xad, 0x73, 0x5f, 0x2e, 0x34, 0x6d, 0x98, 0x8a, 0x85, 0xf5, 0xe8, 0x72,
		0x92, 0xd4, 0x74, 0x4d, 0xb9, 0x35, 0x68, 0x53, 0xb4, 0x6b, 0x13, 0x49,
		0x7f, 0xbe, 0xaf, 0x66, 0x1e, 0xa6, 0x73, 0xf4, 0x49, 0x42, 0x7a, 0x40,
		0xad, 0xd5, 0x13, 0x89, 0xfa, 0x4a, 0xdc, 0x9d, 0x1f, 0x52, 0xa9, 0xf4,
		0x3d, 0x9d, 0x3b, 0x41, 0x94, 0x08, 0xf1, 0xfe, 0x30, 0x83, 0x16, 0x09,
		0xcf, 0xc3, 0x96, 0x5c, 0xee, 0xe7, 0xc4, 0xb3, 0x24, 0x1a, 0x35, 0x2a,
		0x66, 0x1c, 0x73, 0x40, 0x5d, 0xa0, 0x52, 0x23, 0xe6, 0x61, 0xe2, 0xb0,
		0x68, 0xf1, 0xdb, 0xb5, 0x73, 0x37, 0x8b, 0xcd, 0xfa, 0xde, 0x20, 0x3d,
		0xd9, 0xb3, 0xce, 0x9c, 0xa2, 0x34, 0x51, 0x16, 0x41, 0xe7, 0x3c, 0xd3,
		0x79, 0x21, 0xf6, 0x6d, 0xb2, 0xeb, 0x45, 0xec, 0x73, 0xc4, 0x1c, 0x14,
		0xbb, 0x62, 0x9d, 0x87, 0xaa, 0xc0, 0xb3, 0x6b, 0xde, 0x0e, 0x63, 0x4d,
		0x62, 0xc3, 0x74, 0x16, 0x94, 0x73, 0xab, 0x51, 0xb2, 0x77, 0x2a, 0x1d,
		0x28, 0x4c, 0x35, 0x79, 0x2c, 0x0d, 0x68, 0x06, 0x28, 0x5e, 0x7d, 0xe1,
		0x5f, 0x51, 0x01, 0x0f, 0x4c, 0xd2, 0xdc, 0x35, 0xb1, 0x21, 0xb0, 0x99,
		0x54, 0xc9, 0x6e, 0x12, 0x5e, 0x82, 0x9f, 0x37, 0xec, 0x0b, 0x70, 0x6f,
		0xe4, 0x72, 0x88, 0xf0, 0xc8, 0x6c, 0x01, 0xcc, 0x7d, 0x0d, 0xaf, 0xf6,
		0x7b, 0x81, 0x19, 0x02, 0xd1, 0xb2, 0x67, 0xa4, 0xb4, 0xf9, 0x56, 0x6a,
		0x6a, 0x40, 0x57, 0x35, 0x6b, 0xc2, 0x28, 0x06, 0xaa, 0x20, 0xe2, 0x15,
		0x3d, 0xfb, 0x9d, 0xae, 0x2e, 0xdf, 0xe0, 0xf4, 0x3d, 0x5c, 0x90, 0x4d,
		0x28, 0xd9, 0x51, 0x66, 0xf4, 0x4c, 0xa9, 0xf8, 0x37, 0x78, 0x86, 0xf3,
		0x0c, 0x5b, 0x6f, 0x91, 0xe6, 0xf2, 0x96, 0xc0, 0xa9, 0xa0, 0x55, 0x7d,
		0xef, 0xde, 0xc1, 0x7d, 0xd7, 0xd5, 0x8a, 0x16, 0x01, 0x04, 0xc8, 0x4a,
		0xed, 0x08, 0x23, 0xfe, 0xe1, 0xd4, 0xb7, 0xd8, 0x5e, 0x80, 0x97, 0x4c,
		0xbc, 0x00, 0xff, 0x3c, 0x7e, 0x64, 0x06, 0x53, 0xc7, 0x8a, 0x27, 0xe1,
		0xaf, 0xdc, 0x99, 0x5e, 0xe0, 0xa5, 0x8a, 0xc7, 0x15, 0xf6, 0xde, 0xfb,
		0x1c, 0x6c, 0x83, 0x05, 0x50, 0x04, 0xb3, 0x2f, 0x5d, 0x05, 0xc8, 0xe4,
		0x4c, 0x39, 0x58, 0xfb, 0x32, 0x83, 0xd1, 0x1b, 0x95, 0x80, 0x64, 0x5e,
		0x8b, 0xc7, 0x91, 0x15, 0x2a, 0x09, 0x9d, 0x64, 0xeb, 0x8e, 0xb7, 0x96,
		0xa6, 0x3e, 0xd8, 0x81, 0x5d, 0xd4, 0xd1, 0x7a, 0x55, 0x63, 0x93, 0xb5,
		0xea, 0x98, 0x5a, 0x5a, 0x05, 0x7b, 0xac, 0xba, 0x5d, 0x7b, 0x70, 0x75,
		0x1c, 0xa7, 0x1a, 0x37, 0x32, 0x8f, 0x5e, 0xf7, 0x89, 0x4a, 0xe2, 0x32,
		0x84, 0x80, 0x05, 0x23, 0x6c, 0x4a, 0x9c, 0x74, 0xe2, 0x48, 0x1e, 0x22,
		0x2d, 0x60, 0x98, 0x8a, 0x5b, 0x45, 0x58, 0x78, 0x16, 0xfb, 0x38, 0xf6,
		0x35, 0xa9, 0xc8, 0x9d, 0xe7, 0x5d, 0x85, 0x8b, 0x0d, 0x63, 0x48, 0x77,
		0x79, 0x0a, 0xd6, 0x4e, 0xb4, 0x48, 0x01, 0x63, 0x5b, 0x90, 0x9d, 0xc2,
		0xd1, 0xde, 0x4c, 0x4a, 0x5c, 0xbc, 0xc2, 0xda, 0xaf, 0x6b, 0x27, 0x6a,
		0xb1, 0xab, 0xf3, 0x65, 0x7d, 0xa1, 0xde, 0xf9, 0xd0, 0xdb, 0x3b, 0xc8,
		0x5b, 0xb4, 0x3c, 0xfb, 0xdb, 0x92, 0xc5, 0x82, 0x5b, 0x9c, 0x3c, 0x87,
		0xd6, 0xf2, 0x46, 0x93, 0x04, 0x69, 0x97, 0x07, 0x63, 0x20, 0xc9, 0x9a,
		0x4a, 0xe6, 0x39, 0x69, 0x26, 0xa0, 0x4c, 0x12, 0xb1, 0x4b, 0xb0, 0xf5,
		0xbd, 0xc7, 0xc4, 0x7d, 0x57, 0x0c, 0x33, 0x10, 0x2f, 0x31, 0xf5, 0xb5,
		0x72, 0x66, 0xf3, 0x0a, 0xa3, 0x67, 0x21, 0xfb, 0x54, 0xe5, 0x15, 0x5d,
		0x06, 0x7b, 0x81, 0x5b, 0x25, 0x1b, 0x2b, 0xd2, 0x78, 0xdb, 0x97, 0x65,
		0xd3, 0xea, 0x0b, 0x29, 0x42, 0xb0, 0xb1, 0x51, 0x9d, 0x51, 0x61, 0x95,
		0x7a, 0x6c, 0x36, 0x43, 0xeb, 0xc3, 0x6a, 0x50, 0x4c, 0x81, 0xe1, 0xdf,
		0x87, 0x65, 0xd7, 0x83, 0x2b, 0xec, 0x80, 0x11, 0x99, 0x64, 0x9d, 0x63,
		0x36, 0x27, 0x4b, 0x6d, 0xe6, 0x4a, 0x43, 0xe7, 0xa0, 0x2e, 0x47, 0x97,
		0x40, 0x71, 0x52, 0x5f, 0x65, 0x2f, 0x8c, 0xce, 0x3c, 0x83, 0x95, 0x14,
		0x73, 0x3e, 0x91, 0x6a, 0x32, 0x0d, 0x25, 0x40, 0x6a, 0xb6, 0xb6, 0x95,
		0xad, 0x97, 0xf2, 0xaf, 0xd4, 0x2e, 0xf3, 0xeb, 0x26, 0xc2, 0xd8, 0x51,
		0xbb, 0xb4, 0x70, 0x4e, 0x23, 0xab, 0xd5, 0xb0, 0x28, 0x68, 0xaf, 0xc9,
		0xa4, 0x9d, 0x17, 0x22, 0x40, 0xea, 0x9c, 0x28, 0xab, 0xcb, 0xc2, 0xa0,
		0x89, 0x21, 0xa4, 0x6b, 0x78, 0x2b, 0xba, 0x24, 0x7f, 0x81, 0xc9, 0xfb,
		0xde, 0xa5, 0x4d, 0x02, 0xc7, 0x15, 0x56, 0xb0, 0x60, 0xe5, 0xd0, 0xb1,
		0xbc, 0x8f, 0x18, 0xa2, 0xda, 0xe7, 0x5c, 0x9d, 0x6b, 0xc5, 0xe9, 0xf3,
		0xab, 0xf2, 0x39, 0x20, 0x4f, 0x91, 0x50, 0x8b, 0x6b, 0x0e, 0xf6, 0x1c,
		0x4b, 0xd4, 0xcb, 0xc2, 0x5f, 0x52, 0xb4, 0xaa, 0x83, 0x25, 0x9f, 0x0b,
		0xbc, 0xe4, 0x7e, 0x31, 0xbc, 0x14, 0x6a, 0xf5, 0x62, 0xa9, 0xea, 0xe2,
		0xf2, 0xd2, 0x01, 0x4b, 0x78, 0x2b, 0xfa, 0x9c, 0x8b, 0xb5, 0x8b, 0x56,
		0x51, 0xb5, 0xb4, 0xed, 0x0b, 0xf0, 0x9a, 0xd0, 0x15, 0x7e, 0x54, 0x77,
		0xcc, 0x63, 0x99, 0x9e, 0xf9, 0x7d, 0x2c, 0xf6, 0x08, 0x2e, 0x93, 0x56,
		0xd0, 0x77, 0x00, 0x17, 0xf1, 0x71, 0x5c, 0xa6, 0x32, 0x40, 0xc9, 0x85,
		0x5b, 0x6d, 0x2d, 0x13, 0x6c, 0xe0, 0x2d, 0x4b, 0x54, 0x6b, 0x1d, 0xf7,
		0xbe, 0xa0, 0x68, 0x86, 0x58, 0x27, 0xe3, 0x1b, 0x35, 0x30, 0xa4, 0x3a,
		0x05, 0x5b, 0x8e, 0x67, 0x12, 0x21, 0x88, 0x2c, 0x14, 0x3c, 0xa7, 0xe2,
		0x89, 0x70, 0x5a, 0xde, 0xb7, 0xdf, 0xfd, 0xf9, 0x67, 0x6e, 0x6a, 0x81,
		0x11, 0x89, 0x94, 0xa6, 0x63, 0x21, 0x32, 0x5c, 0x4b, 0x89, 0xb3, 0xc9,
		0x9f, 0x08, 0x11, 0xc7, 0x42, 0x99, 0xeb, 0xc7, 0x17, 0x82, 0x5d, 0xde,
		0xb5, 0x90, 0xf1, 0xe2, 0xd7, 0x42, 0xdc, 0x4b, 0x08, 0x0b, 0xcb, 0xaf,
		0xe1, 0x2e, 0x2b, 0x61, 0xc9, 0xcd, 0xb2, 0x3e, 0x96, 0x3c, 0x2e, 0xbb,
		0x66, 0x4d, 0xfa, 0xb2, 0x98, 0x96, 0x02, 0x2d, 0x4b, 0x6c, 0x29, 0xe6,
		0x8a, 0x2e, 0xed, 0xb3, 0xee, 0xc7, 0xdd, 0x97, 0x4e, 0x20, 0xc9, 0xe6,
		0xfa, 0x22, 0xdc, 0x5d, 0x95, 0x69, 0x59, 0xad, 0x16, 0x0c, 0x0c, 0x51,
		0x5b, 0x98, 0xad, 0xf0, 0xfc, 0x86, 0x8f, 0x19, 0x73, 0x7a, 0x2d, 0x11,
		0x45, 0x31, 0xde, 0xe2, 0xee, 0x8c, 0x00, 0xb3, 0x30, 0x31, 0xd3, 0x4b,
		0xb0, 0x18, 0x08, 0x1b, 0x59, 0xdf, 0xdc, 0x57, 0x38, 0x7a, 0xa7, 0x48,
		0x19, 0x0f, 0x03, 0x0e, 0x71, 0x7c, 0xfc, 0xb2, 0xf8, 0x69, 0xc7, 0xc7,
		0xbc, 0xe3, 0x67, 0x52, 0xbc, 0xb4, 0xf9, 0xff, 0xdf, 0xf0, 0xcf, 0x87,
		0xf3, 0x4f, 0x5c, 0xdf, 0x7e, 0x38, 0x7b, 0xfa, 0xf9, 0xea, 0x1f, 0xff,
		0x32, 0xff, 0xfd, 0xaf, 0x5f, 0xab, 0x7e, 0xba, 0x7e, 0xf7, 0xf1, 0xf4,
		0xd7, 0x4f, 0x77, 0x7f, 0xff, 0x78, 0x78, 0xf3, 0xfa, 0xe2, 0xe3, 0xbb,
		0x87, 0xeb, 0xc7, 0xdf, 0xca, 0x2e, 0xae, 0xde, 0xfc, 0xe9, 0xf5, 0xc5,
		0xfd, 0xe7, 0xab, 0x37, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x51, 0x67,
		0x9c, 0x18, 0xe6, 0x1c, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}

func init() {
	go_bindata["/grpc-logo.svg"] = grpc_logo_svg
}
