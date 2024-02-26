package command

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func Exec(timeout int, command string, args []string) (result string, err error) {
	ctx := context.Background()
	if timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
	}

	combined := append([]string{command}, args...)
	joined := strings.Join(combined, " ")
	cmd := exec.CommandContext(ctx, "bash", "-c", joined)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("timout exceeded, your command is taking too long to execute (timeout: %d seconds)", timeout)
		}

		return "", fmt.Errorf("%s: %s", err, stderr.String())
	} else {
		return out.String(), nil
	}
}
