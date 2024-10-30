package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spoitler/glyph-browser/internal/engine"
	"os"
)

var url string

var rootCmd = &cobra.Command{
	Use:   "glyph",
	Short: "Glyph is a browser in cli",
	Long:  "Glyph is a browser in cli. To navigate it only require the keyboard.",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			fmt.Println("URL is required, pass it by the flag --url.")
			return
		}
		glyphEngine := engine.New()
		glyphEngine.Run()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL to process")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
