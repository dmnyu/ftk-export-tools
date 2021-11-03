package cmd

import (
	"errors"
	"fmt"
	ftk_tools "github.com/dmnyu/ftk-tools"
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

/* given a directory input ensure there is a metadata folder that contains a work order
// and that there is a cuid entry in the payload for each entry
// report any DS_Store files
// report any empty directories
// report any directories not on workorder
*/

func init() {
	verifyCmd.Flags().StringVar(&inputDirectory, "input-dir", "", "the location of the transfer")
	rootCmd.AddCommand(verifyCmd)
}

var verifyCmd = &cobra.Command{
	Use: "verify",
	Run: func(cmd *cobra.Command, args []string) {
		if fi, err := os.Stat(inputDirectory); err == nil {
			if fi.IsDir() {
				if err := scan(); err != nil {
					log.Fatal(err)
				}
			} else {
				log.Printf("[FATAL] %s is not a directory", inputDirectory)
			}
		} else if errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does *not* exist
			log.Printf("[FATAL] %s does not exist", inputDirectory)
		} else {
			log.Printf("[FATAL] %s", err.Error())
		}
	},
}

var (
	workOrderPtn = regexp.MustCompile("aspace_wo.tsv$")
)

func scan() error {
	log.Printf("[INFO] Running scanner on %s", inputDirectory)
	payloadFiles, err := ioutil.ReadDir(inputDirectory)
	if err != nil {
		return err
	}
	//determine there is a metadata folder
	if containsFile("metadata", payloadFiles) != true {
		return fmt.Errorf("[FATAL] %s does not contain a metadata directory", inputDirectory)
	}

	log.Printf("[INFO] %s contains a metadata directory", inputDirectory)

	metadataDir := filepath.Join(inputDirectory, "metadata")
	metadataFiles, err := ioutil.ReadDir(metadataDir)
	if err != nil {
		return err
	}

	wo, err := containsWorkOrder(metadataFiles)
	if err != nil {
		return err
	}

	workOrderFile := *wo
	workOrderPath := filepath.Join(metadataDir, workOrderFile.Name())

	log.Printf("[INFO] %s contains a work order file: %s", inputDirectory, workOrderFile.Name())

	workOrder, err := ftk_tools.ParseWorkOrder(workOrderPath)
	if err != nil {
		return fmt.Errorf("[FATAL] %s", err.Error())
	}

	for _, cuid := range workOrder.GetCUIDs() {
		if containsFile(cuid, payloadFiles) != true {
			return fmt.Errorf("[FATAL] component %s does not exist in %s", cuid, inputDirectory)
		} else {
			payloadDirectory := filepath.Join(inputDirectory, cuid)
			log.Printf("[INFO] scanning %s", payloadDirectory)
		}
	}

	log.Printf("[INFO] scan complete on %s", inputDirectory)
	return nil
}

func containsWorkOrder(dir []fs.FileInfo) (*fs.FileInfo, error) {
	for _, file := range dir {
		if workOrderPtn.MatchString(file.Name()) == true {
			return &file, nil
		}
	}
	return nil, fmt.Errorf("[FATAL] %s did not contain a work order file")
}

func containsFile(name string, dir []fs.FileInfo) bool {
	for _, file := range dir {
		if file.Name() == name {
			return true
		}
	}
	return false
}

