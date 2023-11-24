package doc

import (
	"fmt"
	"regexp"
)

// Inject insert a text into the target document at the place indicated by the marker with the given name.
// The beginning of the insertion point is indicated by the line containing only the text
// "<!-- begin:NAME -->", and the end is indicated by the line containing only the text
// "<!-- end:NAME -->", where NAME is the name of the marker.
func Inject(target []byte, name string, source []byte) ([]byte, error) {
	begin := fmt.Sprintf("<!-- begin:%s -->", name)
	end := fmt.Sprintf("<!-- end:%s -->", name)

	re, err := regexp.Compile("(?sm)^\\s*" + begin + "\\s*$.*" + end)
	if err != nil {
		return nil, err
	}

	tmp := make([]byte, 0, len(source)+len(begin)+1+len(end))

	tmp = append(tmp, []byte(begin)...)
	tmp = append(tmp, '\n')
	tmp = append(tmp, source...)
	tmp = append(tmp, []byte(end)...)

	return re.ReplaceAll(target, tmp), nil
}
