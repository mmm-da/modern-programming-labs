package output

import "lab4/pkg/document"

type Output interface {
	SaveFile(document.Document, string)
}
