// Code generated from Bund.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 38, 441,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5,
	17, 176, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 5, 18, 200, 10, 18, 3, 19, 5, 19, 203, 10, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 7, 21, 212, 10, 21, 12,
	21, 14, 21, 215, 11, 21, 3, 21, 6, 21, 218, 10, 21, 13, 21, 14, 21, 219,
	5, 21, 222, 10, 21, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 228, 10, 23, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29,
	3, 29, 7, 29, 242, 10, 29, 12, 29, 14, 29, 245, 11, 29, 3, 30, 6, 30, 248,
	10, 30, 13, 30, 14, 30, 249, 3, 31, 6, 31, 253, 10, 31, 13, 31, 14, 31,
	254, 3, 32, 6, 32, 258, 10, 32, 13, 32, 14, 32, 259, 3, 33, 6, 33, 263,
	10, 33, 13, 33, 14, 33, 264, 3, 34, 3, 34, 3, 34, 3, 34, 7, 34, 271, 10,
	34, 12, 34, 14, 34, 274, 11, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 7, 35, 282, 10, 35, 12, 35, 14, 35, 285, 11, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 36, 6, 36, 293, 10, 36, 13, 36, 14, 36, 294, 3, 36, 3,
	36, 3, 37, 3, 37, 3, 37, 7, 37, 302, 10, 37, 12, 37, 14, 37, 305, 11, 37,
	3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 5, 40, 314, 10, 40, 3,
	40, 6, 40, 317, 10, 40, 13, 40, 14, 40, 318, 3, 40, 5, 40, 322, 10, 40,
	3, 40, 3, 40, 5, 40, 326, 10, 40, 3, 40, 6, 40, 329, 10, 40, 13, 40, 14,
	40, 330, 3, 40, 5, 40, 334, 10, 40, 3, 40, 5, 40, 337, 10, 40, 3, 41, 7,
	41, 340, 10, 41, 12, 41, 14, 41, 343, 11, 41, 3, 41, 3, 41, 6, 41, 347,
	10, 41, 13, 41, 14, 41, 348, 3, 41, 6, 41, 352, 10, 41, 13, 41, 14, 41,
	353, 3, 41, 5, 41, 357, 10, 41, 3, 42, 3, 42, 3, 43, 5, 43, 362, 10, 43,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 370, 10, 44, 3, 44, 7,
	44, 373, 10, 44, 12, 44, 14, 44, 376, 11, 44, 3, 44, 3, 44, 3, 44, 3, 44,
	3, 44, 5, 44, 383, 10, 44, 3, 44, 7, 44, 386, 10, 44, 12, 44, 14, 44, 389,
	11, 44, 3, 44, 5, 44, 392, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 7,
	45, 399, 10, 45, 12, 45, 14, 45, 402, 11, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 7, 45, 412, 10, 45, 12, 45, 14, 45, 415, 11,
	45, 3, 45, 3, 45, 3, 45, 5, 45, 420, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46,
	5, 46, 426, 10, 46, 5, 46, 428, 10, 46, 3, 47, 3, 47, 5, 47, 432, 10, 47,
	3, 47, 5, 47, 435, 10, 47, 3, 48, 3, 48, 3, 48, 5, 48, 440, 10, 48, 5,
	283, 400, 413, 2, 49, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71,
	37, 73, 38, 75, 2, 77, 2, 79, 2, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91,
	2, 93, 2, 95, 2, 3, 2, 16, 4, 2, 87, 87, 119, 119, 6, 2, 40, 40, 45, 45,
	47, 47, 62, 64, 5, 2, 44, 44, 65, 66, 98, 98, 4, 2, 12, 12, 15, 15, 5,
	2, 11, 12, 15, 15, 34, 34, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2, 71, 71, 103,
	103, 4, 2, 45, 45, 47, 47, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 6, 2,
	12, 12, 15, 15, 36, 36, 94, 94, 3, 2, 94, 94, 4, 2, 67, 92, 99, 124, 3,
	2, 99, 124, 2, 483, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2,
	2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2,
	2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3,
	2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31,
	3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2,
	39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2,
	2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2,
	2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2,
	2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3,
	2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 3, 97, 3, 2, 2, 2, 5, 99,
	3, 2, 2, 2, 7, 102, 3, 2, 2, 2, 9, 104, 3, 2, 2, 2, 11, 106, 3, 2, 2, 2,
	13, 109, 3, 2, 2, 2, 15, 116, 3, 2, 2, 2, 17, 121, 3, 2, 2, 2, 19, 127,
	3, 2, 2, 2, 21, 133, 3, 2, 2, 2, 23, 140, 3, 2, 2, 2, 25, 142, 3, 2, 2,
	2, 27, 144, 3, 2, 2, 2, 29, 147, 3, 2, 2, 2, 31, 150, 3, 2, 2, 2, 33, 175,
	3, 2, 2, 2, 35, 199, 3, 2, 2, 2, 37, 202, 3, 2, 2, 2, 39, 206, 3, 2, 2,
	2, 41, 221, 3, 2, 2, 2, 43, 223, 3, 2, 2, 2, 45, 227, 3, 2, 2, 2, 47, 229,
	3, 2, 2, 2, 49, 231, 3, 2, 2, 2, 51, 233, 3, 2, 2, 2, 53, 235, 3, 2, 2,
	2, 55, 237, 3, 2, 2, 2, 57, 239, 3, 2, 2, 2, 59, 247, 3, 2, 2, 2, 61, 252,
	3, 2, 2, 2, 63, 257, 3, 2, 2, 2, 65, 262, 3, 2, 2, 2, 67, 266, 3, 2, 2,
	2, 69, 277, 3, 2, 2, 2, 71, 292, 3, 2, 2, 2, 73, 298, 3, 2, 2, 2, 75, 308,
	3, 2, 2, 2, 77, 310, 3, 2, 2, 2, 79, 336, 3, 2, 2, 2, 81, 356, 3, 2, 2,
	2, 83, 358, 3, 2, 2, 2, 85, 361, 3, 2, 2, 2, 87, 391, 3, 2, 2, 2, 89, 419,
	3, 2, 2, 2, 91, 427, 3, 2, 2, 2, 93, 434, 3, 2, 2, 2, 95, 439, 3, 2, 2,
	2, 97, 98, 7, 93, 2, 2, 98, 4, 3, 2, 2, 2, 99, 100, 7, 61, 2, 2, 100, 101,
	7, 61, 2, 2, 101, 6, 3, 2, 2, 2, 102, 103, 7, 42, 2, 2, 103, 8, 3, 2, 2,
	2, 104, 105, 7, 43, 2, 2, 105, 10, 3, 2, 2, 2, 106, 107, 7, 42, 2, 2, 107,
	108, 7, 44, 2, 2, 108, 12, 3, 2, 2, 2, 109, 110, 7, 42, 2, 2, 110, 111,
	7, 104, 2, 2, 111, 112, 7, 110, 2, 2, 112, 113, 7, 113, 2, 2, 113, 114,
	7, 99, 2, 2, 114, 115, 7, 118, 2, 2, 115, 14, 3, 2, 2, 2, 116, 117, 7,
	42, 2, 2, 117, 118, 7, 107, 2, 2, 118, 119, 7, 112, 2, 2, 119, 120, 7,
	118, 2, 2, 120, 16, 3, 2, 2, 2, 121, 122, 7, 42, 2, 2, 122, 123, 7, 119,
	2, 2, 123, 124, 7, 107, 2, 2, 124, 125, 7, 112, 2, 2, 125, 126, 7, 118,
	2, 2, 126, 18, 3, 2, 2, 2, 127, 128, 7, 42, 2, 2, 128, 129, 7, 118, 2,
	2, 129, 130, 7, 116, 2, 2, 130, 131, 7, 119, 2, 2, 131, 132, 7, 103, 2,
	2, 132, 20, 3, 2, 2, 2, 133, 134, 7, 42, 2, 2, 134, 135, 7, 104, 2, 2,
	135, 136, 7, 99, 2, 2, 136, 137, 7, 110, 2, 2, 137, 138, 7, 117, 2, 2,
	138, 139, 7, 103, 2, 2, 139, 22, 3, 2, 2, 2, 140, 141, 7, 95, 2, 2, 141,
	24, 3, 2, 2, 2, 142, 143, 7, 48, 2, 2, 143, 26, 3, 2, 2, 2, 144, 145, 7,
	93, 2, 2, 145, 146, 7, 93, 2, 2, 146, 28, 3, 2, 2, 2, 147, 148, 7, 95,
	2, 2, 148, 149, 7, 95, 2, 2, 149, 30, 3, 2, 2, 2, 150, 151, 7, 48, 2, 2,
	151, 152, 7, 42, 2, 2, 152, 32, 3, 2, 2, 2, 153, 154, 7, 118, 2, 2, 154,
	155, 7, 116, 2, 2, 155, 156, 7, 119, 2, 2, 156, 176, 7, 103, 2, 2, 157,
	158, 7, 86, 2, 2, 158, 159, 7, 116, 2, 2, 159, 160, 7, 119, 2, 2, 160,
	176, 7, 103, 2, 2, 161, 162, 7, 86, 2, 2, 162, 163, 7, 84, 2, 2, 163, 164,
	7, 87, 2, 2, 164, 176, 7, 71, 2, 2, 165, 176, 7, 86, 2, 2, 166, 167, 7,
	123, 2, 2, 167, 168, 7, 103, 2, 2, 168, 176, 7, 117, 2, 2, 169, 170, 7,
	91, 2, 2, 170, 171, 7, 103, 2, 2, 171, 176, 7, 117, 2, 2, 172, 173, 7,
	91, 2, 2, 173, 174, 7, 71, 2, 2, 174, 176, 7, 85, 2, 2, 175, 153, 3, 2,
	2, 2, 175, 157, 3, 2, 2, 2, 175, 161, 3, 2, 2, 2, 175, 165, 3, 2, 2, 2,
	175, 166, 3, 2, 2, 2, 175, 169, 3, 2, 2, 2, 175, 172, 3, 2, 2, 2, 176,
	34, 3, 2, 2, 2, 177, 178, 7, 104, 2, 2, 178, 179, 7, 99, 2, 2, 179, 180,
	7, 110, 2, 2, 180, 181, 7, 117, 2, 2, 181, 200, 7, 103, 2, 2, 182, 183,
	7, 72, 2, 2, 183, 184, 7, 99, 2, 2, 184, 185, 7, 110, 2, 2, 185, 186, 7,
	117, 2, 2, 186, 200, 7, 103, 2, 2, 187, 188, 7, 72, 2, 2, 188, 189, 7,
	67, 2, 2, 189, 190, 7, 78, 2, 2, 190, 191, 7, 85, 2, 2, 191, 200, 7, 71,
	2, 2, 192, 200, 7, 72, 2, 2, 193, 194, 7, 112, 2, 2, 194, 200, 7, 113,
	2, 2, 195, 196, 7, 80, 2, 2, 196, 200, 7, 113, 2, 2, 197, 198, 7, 80, 2,
	2, 198, 200, 7, 81, 2, 2, 199, 177, 3, 2, 2, 2, 199, 182, 3, 2, 2, 2, 199,
	187, 3, 2, 2, 2, 199, 192, 3, 2, 2, 2, 199, 193, 3, 2, 2, 2, 199, 195,
	3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 36, 3, 2, 2, 2, 201, 203, 5, 83,
	42, 2, 202, 201, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2,
	204, 205, 5, 41, 21, 2, 205, 38, 3, 2, 2, 2, 206, 207, 9, 2, 2, 2, 207,
	208, 5, 41, 21, 2, 208, 40, 3, 2, 2, 2, 209, 213, 5, 75, 38, 2, 210, 212,
	5, 77, 39, 2, 211, 210, 3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211, 3,
	2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 222, 3, 2, 2, 2, 215, 213, 3, 2, 2,
	2, 216, 218, 7, 50, 2, 2, 217, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219,
	217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 222, 3, 2, 2, 2, 221, 209,
	3, 2, 2, 2, 221, 217, 3, 2, 2, 2, 222, 42, 3, 2, 2, 2, 223, 224, 5, 79,
	40, 2, 224, 44, 3, 2, 2, 2, 225, 228, 5, 87, 44, 2, 226, 228, 5, 89, 45,
	2, 227, 225, 3, 2, 2, 2, 227, 226, 3, 2, 2, 2, 228, 46, 3, 2, 2, 2, 229,
	230, 7, 60, 2, 2, 230, 48, 3, 2, 2, 2, 231, 232, 7, 61, 2, 2, 232, 50,
	3, 2, 2, 2, 233, 234, 7, 49, 2, 2, 234, 52, 3, 2, 2, 2, 235, 236, 7, 46,
	2, 2, 236, 54, 3, 2, 2, 2, 237, 238, 7, 96, 2, 2, 238, 56, 3, 2, 2, 2,
	239, 243, 5, 93, 47, 2, 240, 242, 5, 95, 48, 2, 241, 240, 3, 2, 2, 2, 242,
	245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 58, 3,
	2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 248, 9, 3, 2, 2, 247, 246, 3, 2, 2,
	2, 248, 249, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250,
	60, 3, 2, 2, 2, 251, 253, 9, 4, 2, 2, 252, 251, 3, 2, 2, 2, 253, 254, 3,
	2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 62, 3, 2, 2,
	2, 256, 258, 7, 35, 2, 2, 257, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259,
	257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 64, 3, 2, 2, 2, 261, 263, 7,
	38, 2, 2, 262, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 262, 3, 2, 2,
	2, 264, 265, 3, 2, 2, 2, 265, 66, 3, 2, 2, 2, 266, 267, 7, 37, 2, 2, 267,
	268, 7, 37, 2, 2, 268, 272, 3, 2, 2, 2, 269, 271, 10, 5, 2, 2, 270, 269,
	3, 2, 2, 2, 271, 274, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2,
	2, 2, 273, 275, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 275, 276, 8, 34, 2, 2,
	276, 68, 3, 2, 2, 2, 277, 278, 7, 49, 2, 2, 278, 279, 7, 44, 2, 2, 279,
	283, 3, 2, 2, 2, 280, 282, 11, 2, 2, 2, 281, 280, 3, 2, 2, 2, 282, 285,
	3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 284, 286, 3, 2,
	2, 2, 285, 283, 3, 2, 2, 2, 286, 287, 7, 44, 2, 2, 287, 288, 7, 49, 2,
	2, 288, 289, 3, 2, 2, 2, 289, 290, 8, 35, 2, 2, 290, 70, 3, 2, 2, 2, 291,
	293, 9, 6, 2, 2, 292, 291, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 292,
	3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 297, 8, 36,
	2, 2, 297, 72, 3, 2, 2, 2, 298, 299, 7, 37, 2, 2, 299, 303, 7, 35, 2, 2,
	300, 302, 10, 5, 2, 2, 301, 300, 3, 2, 2, 2, 302, 305, 3, 2, 2, 2, 303,
	301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 306, 3, 2, 2, 2, 305, 303,
	3, 2, 2, 2, 306, 307, 8, 37, 3, 2, 307, 74, 3, 2, 2, 2, 308, 309, 9, 7,
	2, 2, 309, 76, 3, 2, 2, 2, 310, 311, 9, 8, 2, 2, 311, 78, 3, 2, 2, 2, 312,
	314, 5, 83, 42, 2, 313, 312, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 321,
	3, 2, 2, 2, 315, 317, 9, 8, 2, 2, 316, 315, 3, 2, 2, 2, 317, 318, 3, 2,
	2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 322, 3, 2, 2, 2,
	320, 322, 5, 81, 41, 2, 321, 316, 3, 2, 2, 2, 321, 320, 3, 2, 2, 2, 322,
	323, 3, 2, 2, 2, 323, 325, 9, 9, 2, 2, 324, 326, 9, 10, 2, 2, 325, 324,
	3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 328, 3, 2, 2, 2, 327, 329, 9, 8,
	2, 2, 328, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2,
	330, 331, 3, 2, 2, 2, 331, 337, 3, 2, 2, 2, 332, 334, 5, 83, 42, 2, 333,
	332, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 337,
	5, 81, 41, 2, 336, 313, 3, 2, 2, 2, 336, 333, 3, 2, 2, 2, 337, 80, 3, 2,
	2, 2, 338, 340, 9, 8, 2, 2, 339, 338, 3, 2, 2, 2, 340, 343, 3, 2, 2, 2,
	341, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 344, 3, 2, 2, 2, 343,
	341, 3, 2, 2, 2, 344, 346, 7, 48, 2, 2, 345, 347, 9, 8, 2, 2, 346, 345,
	3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 348, 349, 3, 2,
	2, 2, 349, 357, 3, 2, 2, 2, 350, 352, 9, 8, 2, 2, 351, 350, 3, 2, 2, 2,
	352, 353, 3, 2, 2, 2, 353, 351, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354,
	355, 3, 2, 2, 2, 355, 357, 7, 48, 2, 2, 356, 341, 3, 2, 2, 2, 356, 351,
	3, 2, 2, 2, 357, 82, 3, 2, 2, 2, 358, 359, 9, 10, 2, 2, 359, 84, 3, 2,
	2, 2, 360, 362, 7, 15, 2, 2, 361, 360, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2,
	362, 363, 3, 2, 2, 2, 363, 364, 7, 12, 2, 2, 364, 86, 3, 2, 2, 2, 365,
	374, 7, 41, 2, 2, 366, 369, 7, 94, 2, 2, 367, 370, 5, 85, 43, 2, 368, 370,
	11, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369, 368, 3, 2, 2, 2, 370, 373, 3, 2,
	2, 2, 371, 373, 10, 11, 2, 2, 372, 366, 3, 2, 2, 2, 372, 371, 3, 2, 2,
	2, 373, 376, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375,
	377, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 377, 392, 7, 41, 2, 2, 378, 387,
	7, 36, 2, 2, 379, 382, 7, 94, 2, 2, 380, 383, 5, 85, 43, 2, 381, 383, 11,
	2, 2, 2, 382, 380, 3, 2, 2, 2, 382, 381, 3, 2, 2, 2, 383, 386, 3, 2, 2,
	2, 384, 386, 10, 12, 2, 2, 385, 379, 3, 2, 2, 2, 385, 384, 3, 2, 2, 2,
	386, 389, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388,
	390, 3, 2, 2, 2, 389, 387, 3, 2, 2, 2, 390, 392, 7, 36, 2, 2, 391, 365,
	3, 2, 2, 2, 391, 378, 3, 2, 2, 2, 392, 88, 3, 2, 2, 2, 393, 394, 7, 41,
	2, 2, 394, 395, 7, 41, 2, 2, 395, 396, 7, 41, 2, 2, 396, 400, 3, 2, 2,
	2, 397, 399, 5, 91, 46, 2, 398, 397, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2,
	400, 401, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 401, 403, 3, 2, 2, 2, 402,
	400, 3, 2, 2, 2, 403, 404, 7, 41, 2, 2, 404, 405, 7, 41, 2, 2, 405, 420,
	7, 41, 2, 2, 406, 407, 7, 36, 2, 2, 407, 408, 7, 36, 2, 2, 408, 409, 7,
	36, 2, 2, 409, 413, 3, 2, 2, 2, 410, 412, 5, 91, 46, 2, 411, 410, 3, 2,
	2, 2, 412, 415, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2,
	414, 416, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 416, 417, 7, 36, 2, 2, 417,
	418, 7, 36, 2, 2, 418, 420, 7, 36, 2, 2, 419, 393, 3, 2, 2, 2, 419, 406,
	3, 2, 2, 2, 420, 90, 3, 2, 2, 2, 421, 428, 10, 13, 2, 2, 422, 425, 7, 94,
	2, 2, 423, 426, 5, 85, 43, 2, 424, 426, 11, 2, 2, 2, 425, 423, 3, 2, 2,
	2, 425, 424, 3, 2, 2, 2, 426, 428, 3, 2, 2, 2, 427, 421, 3, 2, 2, 2, 427,
	422, 3, 2, 2, 2, 428, 92, 3, 2, 2, 2, 429, 432, 9, 14, 2, 2, 430, 432,
	5, 51, 26, 2, 431, 429, 3, 2, 2, 2, 431, 430, 3, 2, 2, 2, 432, 435, 3,
	2, 2, 2, 433, 435, 9, 15, 2, 2, 434, 431, 3, 2, 2, 2, 434, 433, 3, 2, 2,
	2, 435, 94, 3, 2, 2, 2, 436, 440, 5, 93, 47, 2, 437, 440, 9, 8, 2, 2, 438,
	440, 5, 51, 26, 2, 439, 436, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 439, 438,
	3, 2, 2, 2, 440, 96, 3, 2, 2, 2, 46, 2, 175, 199, 202, 213, 219, 221, 227,
	243, 249, 254, 259, 264, 272, 283, 294, 303, 313, 318, 321, 325, 330, 333,
	336, 341, 348, 353, 356, 361, 369, 372, 374, 382, 385, 387, 391, 400, 413,
	419, 425, 427, 431, 434, 439, 4, 8, 2, 2, 2, 3, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "';;'", "'('", "')'", "'(*'", "'(float'", "'(int'", "'(uint'",
	"'(true'", "'(false'", "']'", "'.'", "'[['", "']]'", "'.('", "", "", "",
	"", "", "", "", "':'", "';'", "'/'", "','", "'^'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "TRUE",
	"FALSE", "INTEGER", "UINTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING",
	"TOBEGIN", "TOEND", "SLASH", "DROP", "DUPLICATE", "NAME", "CMD", "SYS",
	"EXECUTE", "RETURN", "COMMENT", "BLOCK_COMMENT", "WS", "SHEBANG",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "TRUE", "FALSE", "INTEGER",
	"UINTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "TOBEGIN", "TOEND",
	"SLASH", "DROP", "DUPLICATE", "NAME", "CMD", "SYS", "EXECUTE", "RETURN",
	"COMMENT", "BLOCK_COMMENT", "WS", "SHEBANG", "NON_ZERO_DIGIT", "DIGIT",
	"EXPONENT_OR_POINT_FLOAT", "POINT_FLOAT", "SIGN", "RN", "SHORT_STRING",
	"LONG_STRING", "LONG_STRING_ITEM", "ID_START", "ID_CONTINUE",
}

type BundLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewBundLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *BundLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewBundLexer(input antlr.CharStream) *BundLexer {
	l := new(BundLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Bund.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BundLexer tokens.
const (
	BundLexerT__0            = 1
	BundLexerT__1            = 2
	BundLexerT__2            = 3
	BundLexerT__3            = 4
	BundLexerT__4            = 5
	BundLexerT__5            = 6
	BundLexerT__6            = 7
	BundLexerT__7            = 8
	BundLexerT__8            = 9
	BundLexerT__9            = 10
	BundLexerT__10           = 11
	BundLexerT__11           = 12
	BundLexerT__12           = 13
	BundLexerT__13           = 14
	BundLexerT__14           = 15
	BundLexerTRUE            = 16
	BundLexerFALSE           = 17
	BundLexerINTEGER         = 18
	BundLexerUINTEGER        = 19
	BundLexerDECIMAL_INTEGER = 20
	BundLexerFLOAT_NUMBER    = 21
	BundLexerSTRING          = 22
	BundLexerTOBEGIN         = 23
	BundLexerTOEND           = 24
	BundLexerSLASH           = 25
	BundLexerDROP            = 26
	BundLexerDUPLICATE       = 27
	BundLexerNAME            = 28
	BundLexerCMD             = 29
	BundLexerSYS             = 30
	BundLexerEXECUTE         = 31
	BundLexerRETURN          = 32
	BundLexerCOMMENT         = 33
	BundLexerBLOCK_COMMENT   = 34
	BundLexerWS              = 35
	BundLexerSHEBANG         = 36
)
