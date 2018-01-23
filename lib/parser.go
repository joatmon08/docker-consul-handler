package lib

import (
	"encoding/json"
	"sort"
	"strings"
)

type entry struct {
	Key         string `json:"Key"`
	CreateIndex int    `json:"CreateIndex"`
}

func convertDataToConsulEntry(data []byte) ([]entry, error) {
	var entries []entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

func sortByCreateIndex(entryList *[]entry) {
	var entries []entry
	entries = *entryList
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].CreateIndex > entries[j].CreateIndex
	})
}

func extractNetworkID(e entry) string {
	slicedKey := strings.Split(e.Key, "/")
	if len(slicedKey) > 5 {
		return slicedKey[len(slicedKey)-2]
	}
	return ""
}

// Network is a structure of ID and its CreateIndex.
type Network struct {
	ID          string
	CreateIndex int
}

// HasNewNetwork sets the newest network. It ignore modifications
// or blank networks.
func (n *Network) HasNewNetwork(data []byte) (bool, error) {
	isNewNetwork := false
	entries, err := convertDataToConsulEntry(data)
	if err != nil {
		return isNewNetwork, err
	}
	if len(entries) == 0 {
		return isNewNetwork, nil
	}
	sortByCreateIndex(&entries)
	entry := extractNetworkID(entries[0])
	if (entries[0].CreateIndex != n.CreateIndex) && (entry != "") {
		n.ID = entry
		n.CreateIndex = entries[0].CreateIndex
		return true, nil
	}
	return isNewNetwork, nil
}
