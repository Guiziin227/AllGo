package xml

import "github.com/clbanning/mxj/v2"

type ConversorXML struct{}

func (x ConversorXML) Converter(dados string) (map[string]any, error) {
	xml, err := mxj.NewMapXml([]byte(dados))
	if err != nil {
		return nil, err
	}

	return xml.Old(), nil
}

func (x ConversorXML) Formato() string {
	return "XML"
}
