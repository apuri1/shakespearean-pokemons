package shakespearean

import (
	"net/http"
	"testing"
)

func TestProcessIncoming(t *testing.T) {

	name := "NotAPokemon"

	strVal, intVal := processIncoming(name)

	if strVal != "" || intVal != 1 {
		t.Errorf("Wrong return vals, got %v and %v", strVal, intVal)
	}

}

func TestCallTranslatorEmptyArg(t *testing.T) {

	emptyStr := ""

	strVal, intVal := callTranslator(emptyStr)

	if strVal != "" || intVal != http.StatusOK {
		if intVal != http.StatusTooManyRequests {
			t.Errorf("Wrong return vals, got %v and %v", strVal, intVal)
		}
	}
}
