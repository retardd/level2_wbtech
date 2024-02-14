package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	fields    string
	delimiter string
	separated bool
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "cut",
		Short: "A utility similar to the cut command",
		Run:   cut,
	}

	rootCmd.Flags().StringVarP(&fields, "fields", "f", "", "Select fields (columns)")
	rootCmd.Flags().StringVarP(&delimiter, "delimiter", "d", "\t", "Use a different delimiter")
	rootCmd.Flags().BoolVarP(&separated, "separated", "s", false, "Only output lines with the delimiter")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cut(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if separated && !strings.Contains(line, delimiter) {
			continue
		}
		fieldsList := strings.Split(line, delimiter)
		if fields == "" {
			fmt.Println(line)
		} else {
			selectedFields := parseFields(fieldsList)
			fmt.Println(strings.Join(selectedFields, delimiter))
		}
	}
}

func parseFields(fieldsList []string) []string {
	var selected []string
	fieldsIndexes := strings.Split(fields, ",")
	for _, index := range fieldsIndexes {
		if i, ok := parseIndex(index); ok && i >= 1 && i <= len(fieldsList) {
			selected = append(selected, fieldsList[i-1])
		}
	}
	return selected
}

func parseIndex(field string) (int, bool) {
	if field == "" {
		return 0, false
	}
	if i := atoi(field); i != 0 {
		return i, true
	}
	return 0, false
}

func atoi(s string) int {
	n := 0
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		n = n*10 + int(ch-'0')
	}
	return n
}
