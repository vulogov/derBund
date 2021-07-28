// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // Bund

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 44, 303,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 7,
	2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 3, 3, 3, 5, 3, 85, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 118, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 133,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 139, 10, 6, 12, 6, 14, 6, 142, 11,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 148, 10, 7, 12, 7, 14, 7, 151, 11, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 6, 8, 157, 10, 8, 13, 8, 14, 8, 158, 3, 8, 3, 8,
	3, 9, 3, 9, 7, 9, 165, 10, 9, 12, 9, 14, 9, 168, 11, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 7, 10, 174, 10, 10, 12, 10, 14, 10, 177, 11, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 7, 11, 183, 10, 11, 12, 11, 14, 11, 186, 11, 11, 3, 11, 3,
	11, 3, 12, 3, 12, 7, 12, 192, 10, 12, 12, 12, 14, 12, 195, 11, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 7, 13, 201, 10, 13, 12, 13, 14, 13, 204, 11, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 7, 14, 210, 10, 14, 12, 14, 14, 14, 213, 11, 14,
	3, 14, 3, 14, 3, 15, 3, 15, 7, 15, 219, 10, 15, 12, 15, 14, 15, 222, 11,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 230, 10, 16, 12, 16,
	14, 16, 233, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 7, 17, 239, 10, 17, 12,
	17, 14, 17, 242, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18,
	250, 10, 18, 12, 18, 14, 18, 253, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	38, 3, 38, 3, 38, 2, 2, 39, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 2, 4, 3, 2, 26, 27, 4, 2, 36, 36, 38, 38, 2,
	322, 2, 79, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6, 117, 3, 2, 2, 2, 8, 132,
	3, 2, 2, 2, 10, 134, 3, 2, 2, 2, 12, 145, 3, 2, 2, 2, 14, 154, 3, 2, 2,
	2, 16, 162, 3, 2, 2, 2, 18, 171, 3, 2, 2, 2, 20, 180, 3, 2, 2, 2, 22, 189,
	3, 2, 2, 2, 24, 198, 3, 2, 2, 2, 26, 207, 3, 2, 2, 2, 28, 216, 3, 2, 2,
	2, 30, 225, 3, 2, 2, 2, 32, 236, 3, 2, 2, 2, 34, 245, 3, 2, 2, 2, 36, 256,
	3, 2, 2, 2, 38, 258, 3, 2, 2, 2, 40, 260, 3, 2, 2, 2, 42, 262, 3, 2, 2,
	2, 44, 264, 3, 2, 2, 2, 46, 266, 3, 2, 2, 2, 48, 268, 3, 2, 2, 2, 50, 270,
	3, 2, 2, 2, 52, 272, 3, 2, 2, 2, 54, 274, 3, 2, 2, 2, 56, 276, 3, 2, 2,
	2, 58, 278, 3, 2, 2, 2, 60, 283, 3, 2, 2, 2, 62, 285, 3, 2, 2, 2, 64, 290,
	3, 2, 2, 2, 66, 292, 3, 2, 2, 2, 68, 294, 3, 2, 2, 2, 70, 296, 3, 2, 2,
	2, 72, 298, 3, 2, 2, 2, 74, 300, 3, 2, 2, 2, 76, 78, 5, 4, 3, 2, 77, 76,
	3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2,
	80, 3, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 85, 5, 10, 6, 2, 83, 85, 5,
	12, 7, 2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 5, 3, 2, 2, 2, 86,
	118, 5, 10, 6, 2, 87, 118, 5, 12, 7, 2, 88, 118, 5, 30, 16, 2, 89, 118,
	5, 32, 17, 2, 90, 118, 5, 14, 8, 2, 91, 118, 5, 16, 9, 2, 92, 118, 5, 18,
	10, 2, 93, 118, 5, 20, 11, 2, 94, 118, 5, 22, 12, 2, 95, 118, 5, 24, 13,
	2, 96, 118, 5, 26, 14, 2, 97, 118, 5, 28, 15, 2, 98, 118, 5, 36, 19, 2,
	99, 118, 5, 38, 20, 2, 100, 118, 5, 40, 21, 2, 101, 118, 5, 42, 22, 2,
	102, 118, 5, 44, 23, 2, 103, 118, 5, 46, 24, 2, 104, 118, 5, 48, 25, 2,
	105, 118, 5, 50, 26, 2, 106, 118, 5, 54, 28, 2, 107, 118, 5, 56, 29, 2,
	108, 118, 5, 58, 30, 2, 109, 118, 5, 60, 31, 2, 110, 118, 5, 62, 32, 2,
	111, 118, 5, 64, 33, 2, 112, 118, 5, 66, 34, 2, 113, 118, 5, 72, 37, 2,
	114, 118, 5, 74, 38, 2, 115, 118, 5, 68, 35, 2, 116, 118, 5, 70, 36, 2,
	117, 86, 3, 2, 2, 2, 117, 87, 3, 2, 2, 2, 117, 88, 3, 2, 2, 2, 117, 89,
	3, 2, 2, 2, 117, 90, 3, 2, 2, 2, 117, 91, 3, 2, 2, 2, 117, 92, 3, 2, 2,
	2, 117, 93, 3, 2, 2, 2, 117, 94, 3, 2, 2, 2, 117, 95, 3, 2, 2, 2, 117,
	96, 3, 2, 2, 2, 117, 97, 3, 2, 2, 2, 117, 98, 3, 2, 2, 2, 117, 99, 3, 2,
	2, 2, 117, 100, 3, 2, 2, 2, 117, 101, 3, 2, 2, 2, 117, 102, 3, 2, 2, 2,
	117, 103, 3, 2, 2, 2, 117, 104, 3, 2, 2, 2, 117, 105, 3, 2, 2, 2, 117,
	106, 3, 2, 2, 2, 117, 107, 3, 2, 2, 2, 117, 108, 3, 2, 2, 2, 117, 109,
	3, 2, 2, 2, 117, 110, 3, 2, 2, 2, 117, 111, 3, 2, 2, 2, 117, 112, 3, 2,
	2, 2, 117, 113, 3, 2, 2, 2, 117, 114, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2,
	117, 116, 3, 2, 2, 2, 118, 7, 3, 2, 2, 2, 119, 133, 5, 36, 19, 2, 120,
	133, 5, 38, 20, 2, 121, 133, 5, 40, 21, 2, 122, 133, 5, 44, 23, 2, 123,
	133, 5, 46, 24, 2, 124, 133, 5, 48, 25, 2, 125, 133, 5, 50, 26, 2, 126,
	133, 5, 54, 28, 2, 127, 133, 5, 56, 29, 2, 128, 133, 5, 58, 30, 2, 129,
	133, 5, 60, 31, 2, 130, 133, 5, 62, 32, 2, 131, 133, 5, 42, 22, 2, 132,
	119, 3, 2, 2, 2, 132, 120, 3, 2, 2, 2, 132, 121, 3, 2, 2, 2, 132, 122,
	3, 2, 2, 2, 132, 123, 3, 2, 2, 2, 132, 124, 3, 2, 2, 2, 132, 125, 3, 2,
	2, 2, 132, 126, 3, 2, 2, 2, 132, 127, 3, 2, 2, 2, 132, 128, 3, 2, 2, 2,
	132, 129, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133,
	9, 3, 2, 2, 2, 134, 135, 7, 3, 2, 2, 135, 136, 7, 36, 2, 2, 136, 140, 7,
	31, 2, 2, 137, 139, 5, 6, 4, 2, 138, 137, 3, 2, 2, 2, 139, 142, 3, 2, 2,
	2, 140, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 143, 3, 2, 2, 2, 142,
	140, 3, 2, 2, 2, 143, 144, 7, 4, 2, 2, 144, 11, 3, 2, 2, 2, 145, 149, 7,
	5, 2, 2, 146, 148, 5, 6, 4, 2, 147, 146, 3, 2, 2, 2, 148, 151, 3, 2, 2,
	2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151,
	149, 3, 2, 2, 2, 152, 153, 7, 6, 2, 2, 153, 13, 3, 2, 2, 2, 154, 156, 7,
	7, 2, 2, 155, 157, 5, 8, 5, 2, 156, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2,
	2, 158, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	161, 7, 6, 2, 2, 161, 15, 3, 2, 2, 2, 162, 166, 7, 8, 2, 2, 163, 165, 5,
	8, 5, 2, 164, 163, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2,
	2, 166, 167, 3, 2, 2, 2, 167, 169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169,
	170, 7, 6, 2, 2, 170, 17, 3, 2, 2, 2, 171, 175, 7, 9, 2, 2, 172, 174, 5,
	48, 25, 2, 173, 172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2,
	2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2,
	178, 179, 7, 6, 2, 2, 179, 19, 3, 2, 2, 2, 180, 184, 7, 10, 2, 2, 181,
	183, 5, 44, 23, 2, 182, 181, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184, 182,
	3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 184, 3, 2,
	2, 2, 187, 188, 7, 6, 2, 2, 188, 21, 3, 2, 2, 2, 189, 193, 7, 11, 2, 2,
	190, 192, 5, 46, 24, 2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193,
	191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 193,
	3, 2, 2, 2, 196, 197, 7, 6, 2, 2, 197, 23, 3, 2, 2, 2, 198, 202, 7, 12,
	2, 2, 199, 201, 5, 6, 4, 2, 200, 199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2,
	202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204,
	202, 3, 2, 2, 2, 205, 206, 7, 6, 2, 2, 206, 25, 3, 2, 2, 2, 207, 211, 7,
	13, 2, 2, 208, 210, 5, 6, 4, 2, 209, 208, 3, 2, 2, 2, 210, 213, 3, 2, 2,
	2, 211, 209, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 214, 3, 2, 2, 2, 213,
	211, 3, 2, 2, 2, 214, 215, 7, 6, 2, 2, 215, 27, 3, 2, 2, 2, 216, 220, 7,
	14, 2, 2, 217, 219, 5, 6, 4, 2, 218, 217, 3, 2, 2, 2, 219, 222, 3, 2, 2,
	2, 220, 218, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222,
	220, 3, 2, 2, 2, 223, 224, 7, 6, 2, 2, 224, 29, 3, 2, 2, 2, 225, 226, 7,
	3, 2, 2, 226, 227, 7, 36, 2, 2, 227, 231, 7, 15, 2, 2, 228, 230, 5, 6,
	4, 2, 229, 228, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2,
	231, 232, 3, 2, 2, 2, 232, 234, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 234,
	235, 7, 16, 2, 2, 235, 31, 3, 2, 2, 2, 236, 240, 7, 17, 2, 2, 237, 239,
	5, 6, 4, 2, 238, 237, 3, 2, 2, 2, 239, 242, 3, 2, 2, 2, 240, 238, 3, 2,
	2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2,
	243, 244, 7, 16, 2, 2, 244, 33, 3, 2, 2, 2, 245, 246, 7, 18, 2, 2, 246,
	247, 7, 37, 2, 2, 247, 251, 7, 19, 2, 2, 248, 250, 5, 6, 4, 2, 249, 248,
	3, 2, 2, 2, 250, 253, 3, 2, 2, 2, 251, 249, 3, 2, 2, 2, 251, 252, 3, 2,
	2, 2, 252, 254, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 254, 255, 7, 16, 2, 2,
	255, 35, 3, 2, 2, 2, 256, 257, 7, 21, 2, 2, 257, 37, 3, 2, 2, 2, 258, 259,
	7, 22, 2, 2, 259, 39, 3, 2, 2, 2, 260, 261, 7, 29, 2, 2, 261, 41, 3, 2,
	2, 2, 262, 263, 7, 30, 2, 2, 263, 43, 3, 2, 2, 2, 264, 265, 7, 23, 2, 2,
	265, 45, 3, 2, 2, 2, 266, 267, 7, 24, 2, 2, 267, 47, 3, 2, 2, 2, 268, 269,
	7, 26, 2, 2, 269, 49, 3, 2, 2, 2, 270, 271, 7, 27, 2, 2, 271, 51, 3, 2,
	2, 2, 272, 273, 9, 2, 2, 2, 273, 53, 3, 2, 2, 2, 274, 275, 7, 28, 2, 2,
	275, 55, 3, 2, 2, 2, 276, 277, 7, 36, 2, 2, 277, 57, 3, 2, 2, 2, 278, 279,
	7, 36, 2, 2, 279, 280, 7, 20, 2, 2, 280, 281, 9, 3, 2, 2, 281, 282, 7,
	6, 2, 2, 282, 59, 3, 2, 2, 2, 283, 284, 7, 37, 2, 2, 284, 61, 3, 2, 2,
	2, 285, 286, 7, 37, 2, 2, 286, 287, 7, 20, 2, 2, 287, 288, 9, 3, 2, 2,
	288, 289, 7, 6, 2, 2, 289, 63, 3, 2, 2, 2, 290, 291, 7, 31, 2, 2, 291,
	65, 3, 2, 2, 2, 292, 293, 7, 32, 2, 2, 293, 67, 3, 2, 2, 2, 294, 295, 7,
	34, 2, 2, 295, 69, 3, 2, 2, 2, 296, 297, 7, 35, 2, 2, 297, 71, 3, 2, 2,
	2, 298, 299, 7, 39, 2, 2, 299, 73, 3, 2, 2, 2, 300, 301, 7, 40, 2, 2, 301,
	75, 3, 2, 2, 2, 19, 79, 84, 117, 132, 140, 149, 158, 166, 175, 184, 193,
	202, 211, 220, 231, 240, 251,
}
var literalNames = []string{
	"", "'['", "';;'", "'('", "')'", "'(?'", "'(*'", "'(float'", "'(int'",
	"'(uint'", "'(true'", "'(false'", "'(ignore'", "']'", "'.'", "'[]'", "'[['",
	"']]'", "'.('", "", "", "", "", "", "", "", "", "", "", "':'", "';'", "'/'",
	"','", "'^'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "TRUE", "FALSE", "INTEGER", "UINTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER",
	"UFLOAT_NUMBER", "COMPLEX_NUMBER", "STRING", "GLOB", "TOBEGIN", "TOEND",
	"SLASH", "DROP", "DUPLICATE", "NAME", "CMD", "SYS", "EXECUTE", "RETURN",
	"COMMENT", "BLOCK_COMMENT", "WS", "SHEBANG",
}

var ruleNames = []string{
	"expressions", "root_term", "term", "data", "ns", "block", "cblock", "datablock",
	"floatblock", "intblock", "uintblock", "trueblock", "falseblock", "ignoreblock",
	"lambda", "alambda", "lambda_cmd", "true_term", "false_term", "string_term",
	"glob_term", "integer", "uinteger", "float", "ufloat", "allfloat", "complex_term",
	"call_term", "call_sys", "cmd_term", "cmd_sys", "begin", "end", "drop",
	"duplicate", "execute_term", "return_term",
}

type BundParser struct {
	*antlr.BaseParser
}

// NewBundParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *BundParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundParser(input antlr.TokenStream) *BundParser {
	this := new(BundParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Bund.g4"

	return this
}

// BundParser tokens.
const (
	BundParserEOF             = antlr.TokenEOF
	BundParserT__0            = 1
	BundParserT__1            = 2
	BundParserT__2            = 3
	BundParserT__3            = 4
	BundParserT__4            = 5
	BundParserT__5            = 6
	BundParserT__6            = 7
	BundParserT__7            = 8
	BundParserT__8            = 9
	BundParserT__9            = 10
	BundParserT__10           = 11
	BundParserT__11           = 12
	BundParserT__12           = 13
	BundParserT__13           = 14
	BundParserT__14           = 15
	BundParserT__15           = 16
	BundParserT__16           = 17
	BundParserT__17           = 18
	BundParserTRUE            = 19
	BundParserFALSE           = 20
	BundParserINTEGER         = 21
	BundParserUINTEGER        = 22
	BundParserDECIMAL_INTEGER = 23
	BundParserFLOAT_NUMBER    = 24
	BundParserUFLOAT_NUMBER   = 25
	BundParserCOMPLEX_NUMBER  = 26
	BundParserSTRING          = 27
	BundParserGLOB            = 28
	BundParserTOBEGIN         = 29
	BundParserTOEND           = 30
	BundParserSLASH           = 31
	BundParserDROP            = 32
	BundParserDUPLICATE       = 33
	BundParserNAME            = 34
	BundParserCMD             = 35
	BundParserSYS             = 36
	BundParserEXECUTE         = 37
	BundParserRETURN          = 38
	BundParserCOMMENT         = 39
	BundParserBLOCK_COMMENT   = 40
	BundParserWS              = 41
	BundParserSHEBANG         = 42
)

// BundParser rules.
const (
	BundParserRULE_expressions  = 0
	BundParserRULE_root_term    = 1
	BundParserRULE_term         = 2
	BundParserRULE_data         = 3
	BundParserRULE_ns           = 4
	BundParserRULE_block        = 5
	BundParserRULE_cblock       = 6
	BundParserRULE_datablock    = 7
	BundParserRULE_floatblock   = 8
	BundParserRULE_intblock     = 9
	BundParserRULE_uintblock    = 10
	BundParserRULE_trueblock    = 11
	BundParserRULE_falseblock   = 12
	BundParserRULE_ignoreblock  = 13
	BundParserRULE_lambda       = 14
	BundParserRULE_alambda      = 15
	BundParserRULE_lambda_cmd   = 16
	BundParserRULE_true_term    = 17
	BundParserRULE_false_term   = 18
	BundParserRULE_string_term  = 19
	BundParserRULE_glob_term    = 20
	BundParserRULE_integer      = 21
	BundParserRULE_uinteger     = 22
	BundParserRULE_float        = 23
	BundParserRULE_ufloat       = 24
	BundParserRULE_allfloat     = 25
	BundParserRULE_complex_term = 26
	BundParserRULE_call_term    = 27
	BundParserRULE_call_sys     = 28
	BundParserRULE_cmd_term     = 29
	BundParserRULE_cmd_sys      = 30
	BundParserRULE_begin        = 31
	BundParserRULE_end          = 32
	BundParserRULE_drop         = 33
	BundParserRULE_duplicate    = 34
	BundParserRULE_execute_term = 35
	BundParserRULE_return_term  = 36
)

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllRoot_term() []IRoot_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRoot_termContext)(nil)).Elem())
	var tst = make([]IRoot_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRoot_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Root_term(i int) IRoot_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoot_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRoot_termContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *BundParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BundParserRULE_expressions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserT__0 || _la == BundParserT__2 {
		{
			p.SetState(74)
			p.Root_term()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRoot_termContext is an interface to support dynamic dispatch.
type IRoot_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoot_termContext differentiates from other interfaces.
	IsRoot_termContext()
}

type Root_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoot_termContext() *Root_termContext {
	var p = new(Root_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_root_term
	return p
}

func (*Root_termContext) IsRoot_termContext() {}

func NewRoot_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Root_termContext {
	var p = new(Root_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_root_term

	return p
}

func (s *Root_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Root_termContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Root_termContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Root_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Root_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Root_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterRoot_term(s)
	}
}

func (s *Root_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitRoot_term(s)
	}
}

func (p *BundParser) Root_term() (localctx IRoot_termContext) {
	localctx = NewRoot_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BundParserRULE_root_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case BundParserT__0:
		{
			p.SetState(80)
			p.Ns()
		}

	case BundParserT__2:
		{
			p.SetState(81)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *TermContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TermContext) Lambda() ILambdaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *TermContext) Alambda() IAlambdaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlambdaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlambdaContext)
}

func (s *TermContext) Cblock() ICblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICblockContext)
}

func (s *TermContext) Datablock() IDatablockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatablockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatablockContext)
}

func (s *TermContext) Floatblock() IFloatblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatblockContext)
}

func (s *TermContext) Intblock() IIntblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntblockContext)
}

func (s *TermContext) Uintblock() IUintblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUintblockContext)
}

func (s *TermContext) Trueblock() ITrueblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrueblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrueblockContext)
}

func (s *TermContext) Falseblock() IFalseblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalseblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalseblockContext)
}

func (s *TermContext) Ignoreblock() IIgnoreblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIgnoreblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIgnoreblockContext)
}

func (s *TermContext) True_term() ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *TermContext) False_term() IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *TermContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *TermContext) Glob_term() IGlob_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlob_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlob_termContext)
}

func (s *TermContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TermContext) Uinteger() IUintegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUintegerContext)
}

func (s *TermContext) Float() IFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *TermContext) Ufloat() IUfloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUfloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUfloatContext)
}

func (s *TermContext) Complex_term() IComplex_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplex_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplex_termContext)
}

func (s *TermContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *TermContext) Call_sys() ICall_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_sysContext)
}

func (s *TermContext) Cmd_term() ICmd_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_termContext)
}

func (s *TermContext) Cmd_sys() ICmd_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_sysContext)
}

func (s *TermContext) Begin() IBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeginContext)
}

func (s *TermContext) End() IEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEndContext)
}

func (s *TermContext) Execute_term() IExecute_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecute_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecute_termContext)
}

func (s *TermContext) Return_term() IReturn_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_termContext)
}

func (s *TermContext) Drop() IDropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDropContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDropContext)
}

func (s *TermContext) Duplicate() IDuplicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDuplicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDuplicateContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *BundParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BundParserRULE_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(84)
			p.Ns()
		}

	case 2:
		{
			p.SetState(85)
			p.Block()
		}

	case 3:
		{
			p.SetState(86)
			p.Lambda()
		}

	case 4:
		{
			p.SetState(87)
			p.Alambda()
		}

	case 5:
		{
			p.SetState(88)
			p.Cblock()
		}

	case 6:
		{
			p.SetState(89)
			p.Datablock()
		}

	case 7:
		{
			p.SetState(90)
			p.Floatblock()
		}

	case 8:
		{
			p.SetState(91)
			p.Intblock()
		}

	case 9:
		{
			p.SetState(92)
			p.Uintblock()
		}

	case 10:
		{
			p.SetState(93)
			p.Trueblock()
		}

	case 11:
		{
			p.SetState(94)
			p.Falseblock()
		}

	case 12:
		{
			p.SetState(95)
			p.Ignoreblock()
		}

	case 13:
		{
			p.SetState(96)
			p.True_term()
		}

	case 14:
		{
			p.SetState(97)
			p.False_term()
		}

	case 15:
		{
			p.SetState(98)
			p.String_term()
		}

	case 16:
		{
			p.SetState(99)
			p.Glob_term()
		}

	case 17:
		{
			p.SetState(100)
			p.Integer()
		}

	case 18:
		{
			p.SetState(101)
			p.Uinteger()
		}

	case 19:
		{
			p.SetState(102)
			p.Float()
		}

	case 20:
		{
			p.SetState(103)
			p.Ufloat()
		}

	case 21:
		{
			p.SetState(104)
			p.Complex_term()
		}

	case 22:
		{
			p.SetState(105)
			p.Call_term()
		}

	case 23:
		{
			p.SetState(106)
			p.Call_sys()
		}

	case 24:
		{
			p.SetState(107)
			p.Cmd_term()
		}

	case 25:
		{
			p.SetState(108)
			p.Cmd_sys()
		}

	case 26:
		{
			p.SetState(109)
			p.Begin()
		}

	case 27:
		{
			p.SetState(110)
			p.End()
		}

	case 28:
		{
			p.SetState(111)
			p.Execute_term()
		}

	case 29:
		{
			p.SetState(112)
			p.Return_term()
		}

	case 30:
		{
			p.SetState(113)
			p.Drop()
		}

	case 31:
		{
			p.SetState(114)
			p.Duplicate()
		}

	}

	return localctx
}

// IDataContext is an interface to support dynamic dispatch.
type IDataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataContext differentiates from other interfaces.
	IsDataContext()
}

type DataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataContext() *DataContext {
	var p = new(DataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_data
	return p
}

func (*DataContext) IsDataContext() {}

func NewDataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataContext {
	var p = new(DataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_data

	return p
}

func (s *DataContext) GetParser() antlr.Parser { return s.parser }

func (s *DataContext) True_term() ITrue_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_termContext)
}

func (s *DataContext) False_term() IFalse_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_termContext)
}

func (s *DataContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *DataContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *DataContext) Uinteger() IUintegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUintegerContext)
}

func (s *DataContext) Float() IFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *DataContext) Ufloat() IUfloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUfloatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUfloatContext)
}

func (s *DataContext) Complex_term() IComplex_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComplex_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComplex_termContext)
}

func (s *DataContext) Call_term() ICall_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_termContext)
}

func (s *DataContext) Call_sys() ICall_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICall_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICall_sysContext)
}

func (s *DataContext) Cmd_term() ICmd_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_termContext)
}

func (s *DataContext) Cmd_sys() ICmd_sysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmd_sysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmd_sysContext)
}

func (s *DataContext) Glob_term() IGlob_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlob_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGlob_termContext)
}

func (s *DataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterData(s)
	}
}

func (s *DataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitData(s)
	}
}

func (p *BundParser) Data() (localctx IDataContext) {
	localctx = NewDataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BundParserRULE_data)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(117)
			p.True_term()
		}

	case 2:
		{
			p.SetState(118)
			p.False_term()
		}

	case 3:
		{
			p.SetState(119)
			p.String_term()
		}

	case 4:
		{
			p.SetState(120)
			p.Integer()
		}

	case 5:
		{
			p.SetState(121)
			p.Uinteger()
		}

	case 6:
		{
			p.SetState(122)
			p.Float()
		}

	case 7:
		{
			p.SetState(123)
			p.Ufloat()
		}

	case 8:
		{
			p.SetState(124)
			p.Complex_term()
		}

	case 9:
		{
			p.SetState(125)
			p.Call_term()
		}

	case 10:
		{
			p.SetState(126)
			p.Call_sys()
		}

	case 11:
		{
			p.SetState(127)
			p.Cmd_term()
		}

	case 12:
		{
			p.SetState(128)
			p.Cmd_sys()
		}

	case 13:
		{
			p.SetState(129)
			p.Glob_term()
		}

	}

	return localctx
}

// INsContext is an interface to support dynamic dispatch.
type INsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsNsContext differentiates from other interfaces.
	IsNsContext()
}

type NsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyNsContext() *NsContext {
	var p = new(NsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ns
	return p
}

func (*NsContext) IsNsContext() {}

func NewNsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NsContext {
	var p = new(NsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ns

	return p
}

func (s *NsContext) GetParser() antlr.Parser { return s.parser }

func (s *NsContext) GetName() antlr.Token { return s.name }

func (s *NsContext) SetName(v antlr.Token) { s.name = v }

func (s *NsContext) Get_term() ITermContext { return s._term }

func (s *NsContext) Set_term(v ITermContext) { s._term = v }

func (s *NsContext) GetBody() []ITermContext { return s.body }

func (s *NsContext) SetBody(v []ITermContext) { s.body = v }

func (s *NsContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

func (s *NsContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *NsContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *NsContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *NsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterNs(s)
	}
}

func (s *NsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitNs(s)
	}
}

func (p *BundParser) Ns() (localctx INsContext) {
	localctx = NewNsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BundParserRULE_ns)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(133)

		var _m = p.Match(BundParserNAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(134)
		p.Match(BundParserTOBEGIN)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(135)

			var _x = p.Term()

			localctx.(*NsContext)._term = _x
		}
		localctx.(*NsContext).body = append(localctx.(*NsContext).body, localctx.(*NsContext)._term)

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(BundParserT__1)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) Get_term() ITermContext { return s._term }

func (s *BlockContext) Set_term(v ITermContext) { s._term = v }

func (s *BlockContext) GetBody() []ITermContext { return s.body }

func (s *BlockContext) SetBody(v []ITermContext) { s.body = v }

func (s *BlockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *BlockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *BundParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BundParserRULE_block)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(BundParserT__2)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(144)

			var _x = p.Term()

			localctx.(*BlockContext)._term = _x
		}
		localctx.(*BlockContext).body = append(localctx.(*BlockContext).body, localctx.(*BlockContext)._term)

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ICblockContext is an interface to support dynamic dispatch.
type ICblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_data returns the _data rule contexts.
	Get_data() IDataContext

	// Set_data sets the _data rule contexts.
	Set_data(IDataContext)

	// GetBody returns the body rule context list.
	GetBody() []IDataContext

	// SetBody sets the body rule context list.
	SetBody([]IDataContext)

	// IsCblockContext differentiates from other interfaces.
	IsCblockContext()
}

type CblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_data  IDataContext
	body   []IDataContext
}

func NewEmptyCblockContext() *CblockContext {
	var p = new(CblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_cblock
	return p
}

func (*CblockContext) IsCblockContext() {}

func NewCblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CblockContext {
	var p = new(CblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_cblock

	return p
}

func (s *CblockContext) GetParser() antlr.Parser { return s.parser }

func (s *CblockContext) Get_data() IDataContext { return s._data }

func (s *CblockContext) Set_data(v IDataContext) { s._data = v }

func (s *CblockContext) GetBody() []IDataContext { return s.body }

func (s *CblockContext) SetBody(v []IDataContext) { s.body = v }

func (s *CblockContext) AllData() []IDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataContext)(nil)).Elem())
	var tst = make([]IDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataContext)
		}
	}

	return tst
}

func (s *CblockContext) Data(i int) IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *CblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCblock(s)
	}
}

func (s *CblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCblock(s)
	}
}

func (p *BundParser) Cblock() (localctx ICblockContext) {
	localctx = NewCblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BundParserRULE_cblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(152)
		p.Match(BundParserT__4)
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(BundParserTRUE-19))|(1<<(BundParserFALSE-19))|(1<<(BundParserINTEGER-19))|(1<<(BundParserUINTEGER-19))|(1<<(BundParserFLOAT_NUMBER-19))|(1<<(BundParserUFLOAT_NUMBER-19))|(1<<(BundParserCOMPLEX_NUMBER-19))|(1<<(BundParserSTRING-19))|(1<<(BundParserGLOB-19))|(1<<(BundParserNAME-19))|(1<<(BundParserCMD-19)))) != 0) {
		{
			p.SetState(153)

			var _x = p.Data()

			localctx.(*CblockContext)._data = _x
		}
		localctx.(*CblockContext).body = append(localctx.(*CblockContext).body, localctx.(*CblockContext)._data)

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(158)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IDatablockContext is an interface to support dynamic dispatch.
type IDatablockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_data returns the _data rule contexts.
	Get_data() IDataContext

	// Set_data sets the _data rule contexts.
	Set_data(IDataContext)

	// GetBody returns the body rule context list.
	GetBody() []IDataContext

	// SetBody sets the body rule context list.
	SetBody([]IDataContext)

	// IsDatablockContext differentiates from other interfaces.
	IsDatablockContext()
}

type DatablockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_data  IDataContext
	body   []IDataContext
}

func NewEmptyDatablockContext() *DatablockContext {
	var p = new(DatablockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_datablock
	return p
}

func (*DatablockContext) IsDatablockContext() {}

func NewDatablockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatablockContext {
	var p = new(DatablockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_datablock

	return p
}

func (s *DatablockContext) GetParser() antlr.Parser { return s.parser }

func (s *DatablockContext) Get_data() IDataContext { return s._data }

func (s *DatablockContext) Set_data(v IDataContext) { s._data = v }

func (s *DatablockContext) GetBody() []IDataContext { return s.body }

func (s *DatablockContext) SetBody(v []IDataContext) { s.body = v }

func (s *DatablockContext) AllData() []IDataContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataContext)(nil)).Elem())
	var tst = make([]IDataContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataContext)
		}
	}

	return tst
}

func (s *DatablockContext) Data(i int) IDataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataContext)
}

func (s *DatablockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatablockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatablockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDatablock(s)
	}
}

func (s *DatablockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDatablock(s)
	}
}

func (p *BundParser) Datablock() (localctx IDatablockContext) {
	localctx = NewDatablockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BundParserRULE_datablock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(BundParserT__5)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-19)&-(0x1f+1)) == 0 && ((1<<uint((_la-19)))&((1<<(BundParserTRUE-19))|(1<<(BundParserFALSE-19))|(1<<(BundParserINTEGER-19))|(1<<(BundParserUINTEGER-19))|(1<<(BundParserFLOAT_NUMBER-19))|(1<<(BundParserUFLOAT_NUMBER-19))|(1<<(BundParserCOMPLEX_NUMBER-19))|(1<<(BundParserSTRING-19))|(1<<(BundParserGLOB-19))|(1<<(BundParserNAME-19))|(1<<(BundParserCMD-19)))) != 0 {
		{
			p.SetState(161)

			var _x = p.Data()

			localctx.(*DatablockContext)._data = _x
		}
		localctx.(*DatablockContext).body = append(localctx.(*DatablockContext).body, localctx.(*DatablockContext)._data)

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IFloatblockContext is an interface to support dynamic dispatch.
type IFloatblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_float returns the _float rule contexts.
	Get_float() IFloatContext

	// Set_float sets the _float rule contexts.
	Set_float(IFloatContext)

	// GetBody returns the body rule context list.
	GetBody() []IFloatContext

	// SetBody sets the body rule context list.
	SetBody([]IFloatContext)

	// IsFloatblockContext differentiates from other interfaces.
	IsFloatblockContext()
}

type FloatblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_float IFloatContext
	body   []IFloatContext
}

func NewEmptyFloatblockContext() *FloatblockContext {
	var p = new(FloatblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_floatblock
	return p
}

func (*FloatblockContext) IsFloatblockContext() {}

func NewFloatblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatblockContext {
	var p = new(FloatblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_floatblock

	return p
}

func (s *FloatblockContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatblockContext) Get_float() IFloatContext { return s._float }

func (s *FloatblockContext) Set_float(v IFloatContext) { s._float = v }

func (s *FloatblockContext) GetBody() []IFloatContext { return s.body }

func (s *FloatblockContext) SetBody(v []IFloatContext) { s.body = v }

func (s *FloatblockContext) AllFloat() []IFloatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFloatContext)(nil)).Elem())
	var tst = make([]IFloatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFloatContext)
		}
	}

	return tst
}

func (s *FloatblockContext) Float(i int) IFloatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFloatContext)
}

func (s *FloatblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFloatblock(s)
	}
}

func (s *FloatblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFloatblock(s)
	}
}

func (p *BundParser) Floatblock() (localctx IFloatblockContext) {
	localctx = NewFloatblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BundParserRULE_floatblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(BundParserT__6)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserFLOAT_NUMBER {
		{
			p.SetState(170)

			var _x = p.Float()

			localctx.(*FloatblockContext)._float = _x
		}
		localctx.(*FloatblockContext).body = append(localctx.(*FloatblockContext).body, localctx.(*FloatblockContext)._float)

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IIntblockContext is an interface to support dynamic dispatch.
type IIntblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_integer returns the _integer rule contexts.
	Get_integer() IIntegerContext

	// Set_integer sets the _integer rule contexts.
	Set_integer(IIntegerContext)

	// GetBody returns the body rule context list.
	GetBody() []IIntegerContext

	// SetBody sets the body rule context list.
	SetBody([]IIntegerContext)

	// IsIntblockContext differentiates from other interfaces.
	IsIntblockContext()
}

type IntblockContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_integer IIntegerContext
	body     []IIntegerContext
}

func NewEmptyIntblockContext() *IntblockContext {
	var p = new(IntblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_intblock
	return p
}

func (*IntblockContext) IsIntblockContext() {}

func NewIntblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntblockContext {
	var p = new(IntblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_intblock

	return p
}

func (s *IntblockContext) GetParser() antlr.Parser { return s.parser }

func (s *IntblockContext) Get_integer() IIntegerContext { return s._integer }

func (s *IntblockContext) Set_integer(v IIntegerContext) { s._integer = v }

func (s *IntblockContext) GetBody() []IIntegerContext { return s.body }

func (s *IntblockContext) SetBody(v []IIntegerContext) { s.body = v }

func (s *IntblockContext) AllInteger() []IIntegerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntegerContext)(nil)).Elem())
	var tst = make([]IIntegerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntegerContext)
		}
	}

	return tst
}

func (s *IntblockContext) Integer(i int) IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *IntblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterIntblock(s)
	}
}

func (s *IntblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitIntblock(s)
	}
}

func (p *BundParser) Intblock() (localctx IIntblockContext) {
	localctx = NewIntblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BundParserRULE_intblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(BundParserT__7)
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserINTEGER {
		{
			p.SetState(179)

			var _x = p.Integer()

			localctx.(*IntblockContext)._integer = _x
		}
		localctx.(*IntblockContext).body = append(localctx.(*IntblockContext).body, localctx.(*IntblockContext)._integer)

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IUintblockContext is an interface to support dynamic dispatch.
type IUintblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_uinteger returns the _uinteger rule contexts.
	Get_uinteger() IUintegerContext

	// Set_uinteger sets the _uinteger rule contexts.
	Set_uinteger(IUintegerContext)

	// GetBody returns the body rule context list.
	GetBody() []IUintegerContext

	// SetBody sets the body rule context list.
	SetBody([]IUintegerContext)

	// IsUintblockContext differentiates from other interfaces.
	IsUintblockContext()
}

type UintblockContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_uinteger IUintegerContext
	body      []IUintegerContext
}

func NewEmptyUintblockContext() *UintblockContext {
	var p = new(UintblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_uintblock
	return p
}

func (*UintblockContext) IsUintblockContext() {}

func NewUintblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UintblockContext {
	var p = new(UintblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_uintblock

	return p
}

func (s *UintblockContext) GetParser() antlr.Parser { return s.parser }

func (s *UintblockContext) Get_uinteger() IUintegerContext { return s._uinteger }

func (s *UintblockContext) Set_uinteger(v IUintegerContext) { s._uinteger = v }

func (s *UintblockContext) GetBody() []IUintegerContext { return s.body }

func (s *UintblockContext) SetBody(v []IUintegerContext) { s.body = v }

func (s *UintblockContext) AllUinteger() []IUintegerContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUintegerContext)(nil)).Elem())
	var tst = make([]IUintegerContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUintegerContext)
		}
	}

	return tst
}

func (s *UintblockContext) Uinteger(i int) IUintegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUintegerContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUintegerContext)
}

func (s *UintblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UintblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UintblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterUintblock(s)
	}
}

func (s *UintblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitUintblock(s)
	}
}

func (p *BundParser) Uintblock() (localctx IUintblockContext) {
	localctx = NewUintblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BundParserRULE_uintblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(BundParserT__8)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == BundParserUINTEGER {
		{
			p.SetState(188)

			var _x = p.Uinteger()

			localctx.(*UintblockContext)._uinteger = _x
		}
		localctx.(*UintblockContext).body = append(localctx.(*UintblockContext).body, localctx.(*UintblockContext)._uinteger)

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ITrueblockContext is an interface to support dynamic dispatch.
type ITrueblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsTrueblockContext differentiates from other interfaces.
	IsTrueblockContext()
}

type TrueblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyTrueblockContext() *TrueblockContext {
	var p = new(TrueblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_trueblock
	return p
}

func (*TrueblockContext) IsTrueblockContext() {}

func NewTrueblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrueblockContext {
	var p = new(TrueblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_trueblock

	return p
}

func (s *TrueblockContext) GetParser() antlr.Parser { return s.parser }

func (s *TrueblockContext) Get_term() ITermContext { return s._term }

func (s *TrueblockContext) Set_term(v ITermContext) { s._term = v }

func (s *TrueblockContext) GetBody() []ITermContext { return s.body }

func (s *TrueblockContext) SetBody(v []ITermContext) { s.body = v }

func (s *TrueblockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *TrueblockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TrueblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrueblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTrueblock(s)
	}
}

func (s *TrueblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTrueblock(s)
	}
}

func (p *BundParser) Trueblock() (localctx ITrueblockContext) {
	localctx = NewTrueblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BundParserRULE_trueblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(BundParserT__9)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(197)

			var _x = p.Term()

			localctx.(*TrueblockContext)._term = _x
		}
		localctx.(*TrueblockContext).body = append(localctx.(*TrueblockContext).body, localctx.(*TrueblockContext)._term)

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IFalseblockContext is an interface to support dynamic dispatch.
type IFalseblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsFalseblockContext differentiates from other interfaces.
	IsFalseblockContext()
}

type FalseblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyFalseblockContext() *FalseblockContext {
	var p = new(FalseblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_falseblock
	return p
}

func (*FalseblockContext) IsFalseblockContext() {}

func NewFalseblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FalseblockContext {
	var p = new(FalseblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_falseblock

	return p
}

func (s *FalseblockContext) GetParser() antlr.Parser { return s.parser }

func (s *FalseblockContext) Get_term() ITermContext { return s._term }

func (s *FalseblockContext) Set_term(v ITermContext) { s._term = v }

func (s *FalseblockContext) GetBody() []ITermContext { return s.body }

func (s *FalseblockContext) SetBody(v []ITermContext) { s.body = v }

func (s *FalseblockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FalseblockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *FalseblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FalseblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFalseblock(s)
	}
}

func (s *FalseblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFalseblock(s)
	}
}

func (p *BundParser) Falseblock() (localctx IFalseblockContext) {
	localctx = NewFalseblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BundParserRULE_falseblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(BundParserT__10)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(206)

			var _x = p.Term()

			localctx.(*FalseblockContext)._term = _x
		}
		localctx.(*FalseblockContext).body = append(localctx.(*FalseblockContext).body, localctx.(*FalseblockContext)._term)

		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(212)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IIgnoreblockContext is an interface to support dynamic dispatch.
type IIgnoreblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsIgnoreblockContext differentiates from other interfaces.
	IsIgnoreblockContext()
}

type IgnoreblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyIgnoreblockContext() *IgnoreblockContext {
	var p = new(IgnoreblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ignoreblock
	return p
}

func (*IgnoreblockContext) IsIgnoreblockContext() {}

func NewIgnoreblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IgnoreblockContext {
	var p = new(IgnoreblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ignoreblock

	return p
}

func (s *IgnoreblockContext) GetParser() antlr.Parser { return s.parser }

func (s *IgnoreblockContext) Get_term() ITermContext { return s._term }

func (s *IgnoreblockContext) Set_term(v ITermContext) { s._term = v }

func (s *IgnoreblockContext) GetBody() []ITermContext { return s.body }

func (s *IgnoreblockContext) SetBody(v []ITermContext) { s.body = v }

func (s *IgnoreblockContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *IgnoreblockContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *IgnoreblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IgnoreblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IgnoreblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterIgnoreblock(s)
	}
}

func (s *IgnoreblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitIgnoreblock(s)
	}
}

func (p *BundParser) Ignoreblock() (localctx IIgnoreblockContext) {
	localctx = NewIgnoreblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BundParserRULE_ignoreblock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(BundParserT__11)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(215)

			var _x = p.Term()

			localctx.(*IgnoreblockContext)._term = _x
		}
		localctx.(*IgnoreblockContext).body = append(localctx.(*IgnoreblockContext).body, localctx.(*IgnoreblockContext)._term)

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) GetName() antlr.Token { return s.name }

func (s *LambdaContext) SetName(v antlr.Token) { s.name = v }

func (s *LambdaContext) Get_term() ITermContext { return s._term }

func (s *LambdaContext) Set_term(v ITermContext) { s._term = v }

func (s *LambdaContext) GetBody() []ITermContext { return s.body }

func (s *LambdaContext) SetBody(v []ITermContext) { s.body = v }

func (s *LambdaContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *LambdaContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *LambdaContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *BundParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BundParserRULE_lambda)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(BundParserT__0)
	}
	{
		p.SetState(224)

		var _m = p.Match(BundParserNAME)

		localctx.(*LambdaContext).name = _m
	}
	{
		p.SetState(225)
		p.Match(BundParserT__12)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(226)

			var _x = p.Term()

			localctx.(*LambdaContext)._term = _x
		}
		localctx.(*LambdaContext).body = append(localctx.(*LambdaContext).body, localctx.(*LambdaContext)._term)

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(232)
		p.Match(BundParserT__13)
	}

	return localctx
}

// IAlambdaContext is an interface to support dynamic dispatch.
type IAlambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsAlambdaContext differentiates from other interfaces.
	IsAlambdaContext()
}

type AlambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyAlambdaContext() *AlambdaContext {
	var p = new(AlambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_alambda
	return p
}

func (*AlambdaContext) IsAlambdaContext() {}

func NewAlambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlambdaContext {
	var p = new(AlambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_alambda

	return p
}

func (s *AlambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *AlambdaContext) Get_term() ITermContext { return s._term }

func (s *AlambdaContext) Set_term(v ITermContext) { s._term = v }

func (s *AlambdaContext) GetBody() []ITermContext { return s.body }

func (s *AlambdaContext) SetBody(v []ITermContext) { s.body = v }

func (s *AlambdaContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *AlambdaContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *AlambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterAlambda(s)
	}
}

func (s *AlambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitAlambda(s)
	}
}

func (p *BundParser) Alambda() (localctx IAlambdaContext) {
	localctx = NewAlambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BundParserRULE_alambda)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(BundParserT__14)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(235)

			var _x = p.Term()

			localctx.(*AlambdaContext)._term = _x
		}
		localctx.(*AlambdaContext).body = append(localctx.(*AlambdaContext).body, localctx.(*AlambdaContext)._term)

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(241)
		p.Match(BundParserT__13)
	}

	return localctx
}

// ILambda_cmdContext is an interface to support dynamic dispatch.
type ILambda_cmdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetBody returns the body rule context list.
	GetBody() []ITermContext

	// SetBody sets the body rule context list.
	SetBody([]ITermContext)

	// IsLambda_cmdContext differentiates from other interfaces.
	IsLambda_cmdContext()
}

type Lambda_cmdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	_term  ITermContext
	body   []ITermContext
}

func NewEmptyLambda_cmdContext() *Lambda_cmdContext {
	var p = new(Lambda_cmdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_lambda_cmd
	return p
}

func (*Lambda_cmdContext) IsLambda_cmdContext() {}

func NewLambda_cmdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lambda_cmdContext {
	var p = new(Lambda_cmdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_lambda_cmd

	return p
}

func (s *Lambda_cmdContext) GetParser() antlr.Parser { return s.parser }

func (s *Lambda_cmdContext) GetName() antlr.Token { return s.name }

func (s *Lambda_cmdContext) SetName(v antlr.Token) { s.name = v }

func (s *Lambda_cmdContext) Get_term() ITermContext { return s._term }

func (s *Lambda_cmdContext) Set_term(v ITermContext) { s._term = v }

func (s *Lambda_cmdContext) GetBody() []ITermContext { return s.body }

func (s *Lambda_cmdContext) SetBody(v []ITermContext) { s.body = v }

func (s *Lambda_cmdContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Lambda_cmdContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Lambda_cmdContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Lambda_cmdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lambda_cmdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lambda_cmdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterLambda_cmd(s)
	}
}

func (s *Lambda_cmdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitLambda_cmd(s)
	}
}

func (p *BundParser) Lambda_cmd() (localctx ILambda_cmdContext) {
	localctx = NewLambda_cmdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BundParserRULE_lambda_cmd)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(BundParserT__15)
	}
	{
		p.SetState(244)

		var _m = p.Match(BundParserCMD)

		localctx.(*Lambda_cmdContext).name = _m
	}
	{
		p.SetState(245)
		p.Match(BundParserT__16)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<BundParserT__0)|(1<<BundParserT__2)|(1<<BundParserT__4)|(1<<BundParserT__5)|(1<<BundParserT__6)|(1<<BundParserT__7)|(1<<BundParserT__8)|(1<<BundParserT__9)|(1<<BundParserT__10)|(1<<BundParserT__11)|(1<<BundParserT__14)|(1<<BundParserTRUE)|(1<<BundParserFALSE)|(1<<BundParserINTEGER)|(1<<BundParserUINTEGER)|(1<<BundParserFLOAT_NUMBER)|(1<<BundParserUFLOAT_NUMBER)|(1<<BundParserCOMPLEX_NUMBER)|(1<<BundParserSTRING)|(1<<BundParserGLOB)|(1<<BundParserTOBEGIN)|(1<<BundParserTOEND))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(BundParserDROP-32))|(1<<(BundParserDUPLICATE-32))|(1<<(BundParserNAME-32))|(1<<(BundParserCMD-32))|(1<<(BundParserEXECUTE-32))|(1<<(BundParserRETURN-32)))) != 0) {
		{
			p.SetState(246)

			var _x = p.Term()

			localctx.(*Lambda_cmdContext)._term = _x
		}
		localctx.(*Lambda_cmdContext).body = append(localctx.(*Lambda_cmdContext).body, localctx.(*Lambda_cmdContext)._term)

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(BundParserT__13)
	}

	return localctx
}

// ITrue_termContext is an interface to support dynamic dispatch.
type ITrue_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsTrue_termContext differentiates from other interfaces.
	IsTrue_termContext()
}

type True_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyTrue_termContext() *True_termContext {
	var p = new(True_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_true_term
	return p
}

func (*True_termContext) IsTrue_termContext() {}

func NewTrue_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *True_termContext {
	var p = new(True_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_true_term

	return p
}

func (s *True_termContext) GetParser() antlr.Parser { return s.parser }

func (s *True_termContext) GetValue() antlr.Token { return s.value }

func (s *True_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *True_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(BundParserTRUE, 0)
}

func (s *True_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *True_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *True_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterTrue_term(s)
	}
}

func (s *True_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitTrue_term(s)
	}
}

func (p *BundParser) True_term() (localctx ITrue_termContext) {
	localctx = NewTrue_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BundParserRULE_true_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)

		var _m = p.Match(BundParserTRUE)

		localctx.(*True_termContext).value = _m
	}

	return localctx
}

// IFalse_termContext is an interface to support dynamic dispatch.
type IFalse_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsFalse_termContext differentiates from other interfaces.
	IsFalse_termContext()
}

type False_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFalse_termContext() *False_termContext {
	var p = new(False_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_false_term
	return p
}

func (*False_termContext) IsFalse_termContext() {}

func NewFalse_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *False_termContext {
	var p = new(False_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_false_term

	return p
}

func (s *False_termContext) GetParser() antlr.Parser { return s.parser }

func (s *False_termContext) GetValue() antlr.Token { return s.value }

func (s *False_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *False_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(BundParserFALSE, 0)
}

func (s *False_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *False_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *False_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFalse_term(s)
	}
}

func (s *False_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFalse_term(s)
	}
}

func (p *BundParser) False_term() (localctx IFalse_termContext) {
	localctx = NewFalse_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BundParserRULE_false_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)

		var _m = p.Match(BundParserFALSE)

		localctx.(*False_termContext).value = _m
	}

	return localctx
}

// IString_termContext is an interface to support dynamic dispatch.
type IString_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsString_termContext differentiates from other interfaces.
	IsString_termContext()
}

type String_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyString_termContext() *String_termContext {
	var p = new(String_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_string_term
	return p
}

func (*String_termContext) IsString_termContext() {}

func NewString_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_termContext {
	var p = new(String_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_string_term

	return p
}

func (s *String_termContext) GetParser() antlr.Parser { return s.parser }

func (s *String_termContext) GetValue() antlr.Token { return s.value }

func (s *String_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *String_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(BundParserSTRING, 0)
}

func (s *String_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterString_term(s)
	}
}

func (s *String_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitString_term(s)
	}
}

func (p *BundParser) String_term() (localctx IString_termContext) {
	localctx = NewString_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BundParserRULE_string_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)

		var _m = p.Match(BundParserSTRING)

		localctx.(*String_termContext).value = _m
	}

	return localctx
}

// IGlob_termContext is an interface to support dynamic dispatch.
type IGlob_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsGlob_termContext differentiates from other interfaces.
	IsGlob_termContext()
}

type Glob_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyGlob_termContext() *Glob_termContext {
	var p = new(Glob_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_glob_term
	return p
}

func (*Glob_termContext) IsGlob_termContext() {}

func NewGlob_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Glob_termContext {
	var p = new(Glob_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_glob_term

	return p
}

func (s *Glob_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Glob_termContext) GetValue() antlr.Token { return s.value }

func (s *Glob_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Glob_termContext) GLOB() antlr.TerminalNode {
	return s.GetToken(BundParserGLOB, 0)
}

func (s *Glob_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Glob_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Glob_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterGlob_term(s)
	}
}

func (s *Glob_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitGlob_term(s)
	}
}

func (p *BundParser) Glob_term() (localctx IGlob_termContext) {
	localctx = NewGlob_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BundParserRULE_glob_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)

		var _m = p.Match(BundParserGLOB)

		localctx.(*Glob_termContext).value = _m
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) GetValue() antlr.Token { return s.value }

func (s *IntegerContext) SetValue(v antlr.Token) { s.value = v }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserINTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *BundParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BundParserRULE_integer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)

		var _m = p.Match(BundParserINTEGER)

		localctx.(*IntegerContext).value = _m
	}

	return localctx
}

// IUintegerContext is an interface to support dynamic dispatch.
type IUintegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsUintegerContext differentiates from other interfaces.
	IsUintegerContext()
}

type UintegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyUintegerContext() *UintegerContext {
	var p = new(UintegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_uinteger
	return p
}

func (*UintegerContext) IsUintegerContext() {}

func NewUintegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UintegerContext {
	var p = new(UintegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_uinteger

	return p
}

func (s *UintegerContext) GetParser() antlr.Parser { return s.parser }

func (s *UintegerContext) GetValue() antlr.Token { return s.value }

func (s *UintegerContext) SetValue(v antlr.Token) { s.value = v }

func (s *UintegerContext) UINTEGER() antlr.TerminalNode {
	return s.GetToken(BundParserUINTEGER, 0)
}

func (s *UintegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UintegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UintegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterUinteger(s)
	}
}

func (s *UintegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitUinteger(s)
	}
}

func (p *BundParser) Uinteger() (localctx IUintegerContext) {
	localctx = NewUintegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BundParserRULE_uinteger)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)

		var _m = p.Match(BundParserUINTEGER)

		localctx.(*UintegerContext).value = _m
	}

	return localctx
}

// IFloatContext is an interface to support dynamic dispatch.
type IFloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsFloatContext differentiates from other interfaces.
	IsFloatContext()
}

type FloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyFloatContext() *FloatContext {
	var p = new(FloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_float
	return p
}

func (*FloatContext) IsFloatContext() {}

func NewFloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatContext {
	var p = new(FloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_float

	return p
}

func (s *FloatContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatContext) GetValue() antlr.Token { return s.value }

func (s *FloatContext) SetValue(v antlr.Token) { s.value = v }

func (s *FloatContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserFLOAT_NUMBER, 0)
}

func (s *FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterFloat(s)
	}
}

func (s *FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitFloat(s)
	}
}

func (p *BundParser) Float() (localctx IFloatContext) {
	localctx = NewFloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BundParserRULE_float)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)

		var _m = p.Match(BundParserFLOAT_NUMBER)

		localctx.(*FloatContext).value = _m
	}

	return localctx
}

// IUfloatContext is an interface to support dynamic dispatch.
type IUfloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsUfloatContext differentiates from other interfaces.
	IsUfloatContext()
}

type UfloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyUfloatContext() *UfloatContext {
	var p = new(UfloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_ufloat
	return p
}

func (*UfloatContext) IsUfloatContext() {}

func NewUfloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UfloatContext {
	var p = new(UfloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_ufloat

	return p
}

func (s *UfloatContext) GetParser() antlr.Parser { return s.parser }

func (s *UfloatContext) GetValue() antlr.Token { return s.value }

func (s *UfloatContext) SetValue(v antlr.Token) { s.value = v }

func (s *UfloatContext) UFLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserUFLOAT_NUMBER, 0)
}

func (s *UfloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UfloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UfloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterUfloat(s)
	}
}

func (s *UfloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitUfloat(s)
	}
}

func (p *BundParser) Ufloat() (localctx IUfloatContext) {
	localctx = NewUfloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BundParserRULE_ufloat)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)

		var _m = p.Match(BundParserUFLOAT_NUMBER)

		localctx.(*UfloatContext).value = _m
	}

	return localctx
}

// IAllfloatContext is an interface to support dynamic dispatch.
type IAllfloatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsAllfloatContext differentiates from other interfaces.
	IsAllfloatContext()
}

type AllfloatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyAllfloatContext() *AllfloatContext {
	var p = new(AllfloatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_allfloat
	return p
}

func (*AllfloatContext) IsAllfloatContext() {}

func NewAllfloatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllfloatContext {
	var p = new(AllfloatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_allfloat

	return p
}

func (s *AllfloatContext) GetParser() antlr.Parser { return s.parser }

func (s *AllfloatContext) GetValue() antlr.Token { return s.value }

func (s *AllfloatContext) SetValue(v antlr.Token) { s.value = v }

func (s *AllfloatContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserFLOAT_NUMBER, 0)
}

func (s *AllfloatContext) UFLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserUFLOAT_NUMBER, 0)
}

func (s *AllfloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllfloatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AllfloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterAllfloat(s)
	}
}

func (s *AllfloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitAllfloat(s)
	}
}

func (p *BundParser) Allfloat() (localctx IAllfloatContext) {
	localctx = NewAllfloatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BundParserRULE_allfloat)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*AllfloatContext).value = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserFLOAT_NUMBER || _la == BundParserUFLOAT_NUMBER) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*AllfloatContext).value = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComplex_termContext is an interface to support dynamic dispatch.
type IComplex_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsComplex_termContext differentiates from other interfaces.
	IsComplex_termContext()
}

type Complex_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyComplex_termContext() *Complex_termContext {
	var p = new(Complex_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_complex_term
	return p
}

func (*Complex_termContext) IsComplex_termContext() {}

func NewComplex_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Complex_termContext {
	var p = new(Complex_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_complex_term

	return p
}

func (s *Complex_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Complex_termContext) GetValue() antlr.Token { return s.value }

func (s *Complex_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Complex_termContext) COMPLEX_NUMBER() antlr.TerminalNode {
	return s.GetToken(BundParserCOMPLEX_NUMBER, 0)
}

func (s *Complex_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Complex_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Complex_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterComplex_term(s)
	}
}

func (s *Complex_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitComplex_term(s)
	}
}

func (p *BundParser) Complex_term() (localctx IComplex_termContext) {
	localctx = NewComplex_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BundParserRULE_complex_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)

		var _m = p.Match(BundParserCOMPLEX_NUMBER)

		localctx.(*Complex_termContext).value = _m
	}

	return localctx
}

// ICall_termContext is an interface to support dynamic dispatch.
type ICall_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsCall_termContext differentiates from other interfaces.
	IsCall_termContext()
}

type Call_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyCall_termContext() *Call_termContext {
	var p = new(Call_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_call_term
	return p
}

func (*Call_termContext) IsCall_termContext() {}

func NewCall_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_termContext {
	var p = new(Call_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_call_term

	return p
}

func (s *Call_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_termContext) GetValue() antlr.Token { return s.value }

func (s *Call_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Call_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Call_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCall_term(s)
	}
}

func (s *Call_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCall_term(s)
	}
}

func (p *BundParser) Call_term() (localctx ICall_termContext) {
	localctx = NewCall_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BundParserRULE_call_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)

		var _m = p.Match(BundParserNAME)

		localctx.(*Call_termContext).value = _m
	}

	return localctx
}

// ICall_sysContext is an interface to support dynamic dispatch.
type ICall_sysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// GetSyscmd returns the syscmd token.
	GetSyscmd() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// SetSyscmd sets the syscmd token.
	SetSyscmd(antlr.Token)

	// IsCall_sysContext differentiates from other interfaces.
	IsCall_sysContext()
}

type Call_sysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
	syscmd antlr.Token
}

func NewEmptyCall_sysContext() *Call_sysContext {
	var p = new(Call_sysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_call_sys
	return p
}

func (*Call_sysContext) IsCall_sysContext() {}

func NewCall_sysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_sysContext {
	var p = new(Call_sysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_call_sys

	return p
}

func (s *Call_sysContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_sysContext) GetValue() antlr.Token { return s.value }

func (s *Call_sysContext) GetSyscmd() antlr.Token { return s.syscmd }

func (s *Call_sysContext) SetValue(v antlr.Token) { s.value = v }

func (s *Call_sysContext) SetSyscmd(v antlr.Token) { s.syscmd = v }

func (s *Call_sysContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(BundParserNAME)
}

func (s *Call_sysContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(BundParserNAME, i)
}

func (s *Call_sysContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Call_sysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_sysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_sysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCall_sys(s)
	}
}

func (s *Call_sysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCall_sys(s)
	}
}

func (p *BundParser) Call_sys() (localctx ICall_sysContext) {
	localctx = NewCall_sysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BundParserRULE_call_sys)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)

		var _m = p.Match(BundParserNAME)

		localctx.(*Call_sysContext).value = _m
	}
	{
		p.SetState(277)
		p.Match(BundParserT__17)
	}
	{
		p.SetState(278)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Call_sysContext).syscmd = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserNAME || _la == BundParserSYS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Call_sysContext).syscmd = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(279)
		p.Match(BundParserT__3)
	}

	return localctx
}

// ICmd_termContext is an interface to support dynamic dispatch.
type ICmd_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsCmd_termContext differentiates from other interfaces.
	IsCmd_termContext()
}

type Cmd_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyCmd_termContext() *Cmd_termContext {
	var p = new(Cmd_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_cmd_term
	return p
}

func (*Cmd_termContext) IsCmd_termContext() {}

func NewCmd_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_termContext {
	var p = new(Cmd_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_cmd_term

	return p
}

func (s *Cmd_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_termContext) GetValue() antlr.Token { return s.value }

func (s *Cmd_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Cmd_termContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Cmd_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCmd_term(s)
	}
}

func (s *Cmd_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCmd_term(s)
	}
}

func (p *BundParser) Cmd_term() (localctx ICmd_termContext) {
	localctx = NewCmd_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BundParserRULE_cmd_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)

		var _m = p.Match(BundParserCMD)

		localctx.(*Cmd_termContext).value = _m
	}

	return localctx
}

// ICmd_sysContext is an interface to support dynamic dispatch.
type ICmd_sysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// GetSyscmd returns the syscmd token.
	GetSyscmd() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// SetSyscmd sets the syscmd token.
	SetSyscmd(antlr.Token)

	// IsCmd_sysContext differentiates from other interfaces.
	IsCmd_sysContext()
}

type Cmd_sysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
	syscmd antlr.Token
}

func NewEmptyCmd_sysContext() *Cmd_sysContext {
	var p = new(Cmd_sysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_cmd_sys
	return p
}

func (*Cmd_sysContext) IsCmd_sysContext() {}

func NewCmd_sysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cmd_sysContext {
	var p = new(Cmd_sysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_cmd_sys

	return p
}

func (s *Cmd_sysContext) GetParser() antlr.Parser { return s.parser }

func (s *Cmd_sysContext) GetValue() antlr.Token { return s.value }

func (s *Cmd_sysContext) GetSyscmd() antlr.Token { return s.syscmd }

func (s *Cmd_sysContext) SetValue(v antlr.Token) { s.value = v }

func (s *Cmd_sysContext) SetSyscmd(v antlr.Token) { s.syscmd = v }

func (s *Cmd_sysContext) CMD() antlr.TerminalNode {
	return s.GetToken(BundParserCMD, 0)
}

func (s *Cmd_sysContext) SYS() antlr.TerminalNode {
	return s.GetToken(BundParserSYS, 0)
}

func (s *Cmd_sysContext) NAME() antlr.TerminalNode {
	return s.GetToken(BundParserNAME, 0)
}

func (s *Cmd_sysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cmd_sysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cmd_sysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterCmd_sys(s)
	}
}

func (s *Cmd_sysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitCmd_sys(s)
	}
}

func (p *BundParser) Cmd_sys() (localctx ICmd_sysContext) {
	localctx = NewCmd_sysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BundParserRULE_cmd_sys)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)

		var _m = p.Match(BundParserCMD)

		localctx.(*Cmd_sysContext).value = _m
	}
	{
		p.SetState(284)
		p.Match(BundParserT__17)
	}
	{
		p.SetState(285)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Cmd_sysContext).syscmd = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == BundParserNAME || _la == BundParserSYS) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Cmd_sysContext).syscmd = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(286)
		p.Match(BundParserT__3)
	}

	return localctx
}

// IBeginContext is an interface to support dynamic dispatch.
type IBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsBeginContext differentiates from other interfaces.
	IsBeginContext()
}

type BeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyBeginContext() *BeginContext {
	var p = new(BeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_begin
	return p
}

func (*BeginContext) IsBeginContext() {}

func NewBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BeginContext {
	var p = new(BeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_begin

	return p
}

func (s *BeginContext) GetParser() antlr.Parser { return s.parser }

func (s *BeginContext) GetValue() antlr.Token { return s.value }

func (s *BeginContext) SetValue(v antlr.Token) { s.value = v }

func (s *BeginContext) TOBEGIN() antlr.TerminalNode {
	return s.GetToken(BundParserTOBEGIN, 0)
}

func (s *BeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterBegin(s)
	}
}

func (s *BeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitBegin(s)
	}
}

func (p *BundParser) Begin() (localctx IBeginContext) {
	localctx = NewBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BundParserRULE_begin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)

		var _m = p.Match(BundParserTOBEGIN)

		localctx.(*BeginContext).value = _m
	}

	return localctx
}

// IEndContext is an interface to support dynamic dispatch.
type IEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsEndContext differentiates from other interfaces.
	IsEndContext()
}

type EndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyEndContext() *EndContext {
	var p = new(EndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_end
	return p
}

func (*EndContext) IsEndContext() {}

func NewEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndContext {
	var p = new(EndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_end

	return p
}

func (s *EndContext) GetParser() antlr.Parser { return s.parser }

func (s *EndContext) GetValue() antlr.Token { return s.value }

func (s *EndContext) SetValue(v antlr.Token) { s.value = v }

func (s *EndContext) TOEND() antlr.TerminalNode {
	return s.GetToken(BundParserTOEND, 0)
}

func (s *EndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterEnd(s)
	}
}

func (s *EndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitEnd(s)
	}
}

func (p *BundParser) End() (localctx IEndContext) {
	localctx = NewEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BundParserRULE_end)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)

		var _m = p.Match(BundParserTOEND)

		localctx.(*EndContext).value = _m
	}

	return localctx
}

// IDropContext is an interface to support dynamic dispatch.
type IDropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsDropContext differentiates from other interfaces.
	IsDropContext()
}

type DropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyDropContext() *DropContext {
	var p = new(DropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_drop
	return p
}

func (*DropContext) IsDropContext() {}

func NewDropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropContext {
	var p = new(DropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_drop

	return p
}

func (s *DropContext) GetParser() antlr.Parser { return s.parser }

func (s *DropContext) GetValue() antlr.Token { return s.value }

func (s *DropContext) SetValue(v antlr.Token) { s.value = v }

func (s *DropContext) DROP() antlr.TerminalNode {
	return s.GetToken(BundParserDROP, 0)
}

func (s *DropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDrop(s)
	}
}

func (s *DropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDrop(s)
	}
}

func (p *BundParser) Drop() (localctx IDropContext) {
	localctx = NewDropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BundParserRULE_drop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)

		var _m = p.Match(BundParserDROP)

		localctx.(*DropContext).value = _m
	}

	return localctx
}

// IDuplicateContext is an interface to support dynamic dispatch.
type IDuplicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsDuplicateContext differentiates from other interfaces.
	IsDuplicateContext()
}

type DuplicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyDuplicateContext() *DuplicateContext {
	var p = new(DuplicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_duplicate
	return p
}

func (*DuplicateContext) IsDuplicateContext() {}

func NewDuplicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DuplicateContext {
	var p = new(DuplicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_duplicate

	return p
}

func (s *DuplicateContext) GetParser() antlr.Parser { return s.parser }

func (s *DuplicateContext) GetValue() antlr.Token { return s.value }

func (s *DuplicateContext) SetValue(v antlr.Token) { s.value = v }

func (s *DuplicateContext) DUPLICATE() antlr.TerminalNode {
	return s.GetToken(BundParserDUPLICATE, 0)
}

func (s *DuplicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DuplicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DuplicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterDuplicate(s)
	}
}

func (s *DuplicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitDuplicate(s)
	}
}

func (p *BundParser) Duplicate() (localctx IDuplicateContext) {
	localctx = NewDuplicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BundParserRULE_duplicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(294)

		var _m = p.Match(BundParserDUPLICATE)

		localctx.(*DuplicateContext).value = _m
	}

	return localctx
}

// IExecute_termContext is an interface to support dynamic dispatch.
type IExecute_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsExecute_termContext differentiates from other interfaces.
	IsExecute_termContext()
}

type Execute_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyExecute_termContext() *Execute_termContext {
	var p = new(Execute_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_execute_term
	return p
}

func (*Execute_termContext) IsExecute_termContext() {}

func NewExecute_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Execute_termContext {
	var p = new(Execute_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_execute_term

	return p
}

func (s *Execute_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Execute_termContext) GetValue() antlr.Token { return s.value }

func (s *Execute_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Execute_termContext) EXECUTE() antlr.TerminalNode {
	return s.GetToken(BundParserEXECUTE, 0)
}

func (s *Execute_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Execute_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Execute_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterExecute_term(s)
	}
}

func (s *Execute_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitExecute_term(s)
	}
}

func (p *BundParser) Execute_term() (localctx IExecute_termContext) {
	localctx = NewExecute_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, BundParserRULE_execute_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)

		var _m = p.Match(BundParserEXECUTE)

		localctx.(*Execute_termContext).value = _m
	}

	return localctx
}

// IReturn_termContext is an interface to support dynamic dispatch.
type IReturn_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value token.
	GetValue() antlr.Token

	// SetValue sets the value token.
	SetValue(antlr.Token)

	// IsReturn_termContext differentiates from other interfaces.
	IsReturn_termContext()
}

type Return_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  antlr.Token
}

func NewEmptyReturn_termContext() *Return_termContext {
	var p = new(Return_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = BundParserRULE_return_term
	return p
}

func (*Return_termContext) IsReturn_termContext() {}

func NewReturn_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_termContext {
	var p = new(Return_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = BundParserRULE_return_term

	return p
}

func (s *Return_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_termContext) GetValue() antlr.Token { return s.value }

func (s *Return_termContext) SetValue(v antlr.Token) { s.value = v }

func (s *Return_termContext) RETURN() antlr.TerminalNode {
	return s.GetToken(BundParserRETURN, 0)
}

func (s *Return_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.EnterReturn_term(s)
	}
}

func (s *Return_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BundListener); ok {
		listenerT.ExitReturn_term(s)
	}
}

func (p *BundParser) Return_term() (localctx IReturn_termContext) {
	localctx = NewReturn_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, BundParserRULE_return_term)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)

		var _m = p.Match(BundParserRETURN)

		localctx.(*Return_termContext).value = _m
	}

	return localctx
}
