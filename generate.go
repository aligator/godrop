//go:generate gqlgen generate server
package godrop

import (
	_ "embed"
	"log"
	"net/http"
)

//go:embed schema/schema.graphqls
var schema []byte

// SchemaHandler just serves the schema file which is embedded.
type SchemaHandler struct {
	Logger log.Logger
}

func (h *SchemaHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write(schema)
	if err != nil {
		h.Logger.Println(err.Error())
	}
}
