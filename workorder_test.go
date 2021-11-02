package ftk_tools

import (
	"path/filepath"
	"testing"
)

var testWO = filepath.Join("test", "test_wo.tsv")

func TestWorkorder(t *testing.T) {
	var workOrder WorkOrder
	var err error

	t.Run("Parse A Work Order", func(t *testing.T) {
		workOrder, err = ParseWorkOrder(testWO)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Ensure Entry Count", func(t *testing.T) {
		want := 9
		got := len(workOrder)
		if want != got {
			t.Errorf("Wanted %d got %d", want, got)
		}
	})

	t.Run("Ensure Value", func(t *testing.T) {
		want := "Gallery Files Backup"
		got := workOrder[8].Title
		if want != got {
			t.Errorf("Wanted %s got %s", want, got)
		}
	})
}
