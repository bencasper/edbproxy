package mock

import (
	"testing"
	"fmt"
)

func TestXmlFmt(*testing.T) {
	xml := DeliveryXml("106752228588756992")
	fmt.Print(xml)
}
