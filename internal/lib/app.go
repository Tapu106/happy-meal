package lib

import (
	"context"
	"encoding/json"
	"fmt"
)

func RequestHandler(ctx context.Context, event json.RawMessage) error {
	fmt.Println("Go!!!!!")
	return nil

}
