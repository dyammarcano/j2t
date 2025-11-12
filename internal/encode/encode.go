package encode

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/alpkeskin/gotoon"
	"github.com/spf13/cobra"
)

func readInput(args []string) ([]byte, error) {
	if len(args) > 0 {
		path := args[0]
		return os.ReadFile(path)
	}

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("reading stdin: %w", err)
	}

	return data, nil
}

func encodeData(data []byte) (string, error) {
	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return "", fmt.Errorf("unmarshal JSON: %w", err)
	}

	return gotoon.Encode(value)
}

func Json2Toon(cmd *cobra.Command, args []string) error {
	data, err := readInput(args)
	if err != nil {
		return err
	}

	encoded, err := encodeData(data)
	if err != nil {
		return fmt.Errorf("encode failed: %w", err)
	}

	outPath, _ := cmd.Flags().GetString("output")
	if outPath == "" {
		_, _ = fmt.Fprintln(cmd.OutOrStdout(), encoded)
		return nil
	}

	if err := os.WriteFile(outPath, []byte(encoded), 0o644); err != nil {
		return fmt.Errorf("writing output file: %w", err)
	}

	cmd.Printf("Wrote encoded output to %s\n", outPath)

	return nil
}
