//go:build !devserver

package temporalcli

import (
	"fmt"
	"strings"

	"go.temporal.io/api/enums/v1"
)

func (t *TemporalServerStartDevCommand) run(_ *CommandContext, _ []string) error {
	return fmt.Errorf("server start-dev is unsupported in hardened build; rebuild CLI with -tags devserver to enable")
}

func toFriendlyIp(host string) string {
	return host
}

func persistentClusterID() string {
	return ""
}

func (t *TemporalServerStartDevCommand) prepareSearchAttributes() (map[string]enums.IndexedValueType, error) {
	opts, err := stringKeysValues(t.SearchAttribute)
	if err != nil {
		return nil, fmt.Errorf("invalid search attributes: %w", err)
	}
	attrs := make(map[string]enums.IndexedValueType, len(opts))
	for k, v := range opts {
		var valType enums.IndexedValueType
		for valTypeName, valTypeOrd := range enums.IndexedValueType_shorthandValue {
			if strings.EqualFold(v, valTypeName) {
				valType = enums.IndexedValueType(valTypeOrd)
				break
			}
		}
		if valType == 0 {
			return nil, fmt.Errorf("invalid search attribute value type %q", v)
		}
		attrs[k] = valType
	}
	return attrs, nil
}
