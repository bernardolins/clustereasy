package unitcont

import (
	"fmt"
	"github.com/bernardolins/clustereasy/setup/types"
)

func DefaultStaticIp(node types.Node) string {
	return fmt.Sprintf("[Match]\nName=%s\n[Network]\nAddress=%s/24", node.NodeIp(), node.NodeInterface())
}
