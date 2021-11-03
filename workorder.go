package ftk_tools

import (
	"bufio"
	"os"
	"strings"
)

type WorkOrderEntry struct {
	ResourceID          string
	RefID               string
	URI                 string
	ContainerIndicator1 string
	ContainerIndicator2 string
	ContainerIndicator3 string
	Title               string
	ComponentID         string
}

type WorkOrder []WorkOrderEntry

func ParseWorkOrder(path string) (WorkOrder, error) {
	wo, err := os.Open(path)
	if err != nil {
		return WorkOrder{}, err
	}
	workOrder := WorkOrder{}
	scanner := bufio.NewScanner(wo)
	scanner.Scan() // skip the header
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, "\t")
		entry := WorkOrderEntry{
			ResourceID:          cols[0],
			RefID:               cols[1],
			URI:                 cols[2],
			ContainerIndicator1: cols[3],
			ContainerIndicator2: cols[4],
			ContainerIndicator3: cols[5],
			Title:               cols[6],
			ComponentID:         cols[7],
		}
		workOrder = append(workOrder, entry)
	}
	return workOrder, nil
}

func (wo WorkOrder) GetCUIDs() []string {
	cuids := []string{}
	for _, entry := range wo {
		cuids = append(cuids, entry.ComponentID)
	}
	return cuids
}
