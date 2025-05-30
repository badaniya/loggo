/*
Copyright © 2022 Aurelio Calegari, et al.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"github.com/badaniya/loggo/internal/loggo"
	"github.com/badaniya/loggo/internal/reader"
	"github.com/spf13/cobra"
)

// streamCmd represents the stream command
var streamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Continuously stream log input source",
	Long: `Continuously stream log entries from an input stream such
as the standard input (through pipe) or a input file. Note that
if it's reading from a file, it automatically detects file 
rotation and continue to stream. For example:

	loggo stream --file <file-path>
	<some arbitrary input> | loggo stream`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := cmd.Flag("file").Value.String()
		templateFile := cmd.Flag("template").Value.String()
		reader := reader.MakeReader(fileName, nil)
		app := loggo.NewLoggoApp(reader, templateFile)
		app.Run()
	},
}

func init() {
	rootCmd.AddCommand(streamCmd)
	streamCmd.Flags().
		StringP("file", "f", "", "Input Log File")
	streamCmd.Flags().
		StringP("template", "t", "", "Rendering Template")
}
