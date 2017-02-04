package ZFic

import "github.com/russross/blackfriday"

func GetMD(input []byte) []byte {
	return blackfriday.MarkdownCommon(input)
}
