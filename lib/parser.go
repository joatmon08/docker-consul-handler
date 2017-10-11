package lib

import (
	"encoding/json"
	"sort"
	"strings"
)

type entry struct {
	Key         string `json:"Key"`
	ModifyIndex int    `json:"ModifyIndex"`
}

func convertDataToConsulEntry(data []byte) ([]entry, error) {
	var entries []entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}

func sortByModifyIndex(entry_list *[]entry) {
	var entries []entry
	entries = *entry_list
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ModifyIndex > entries[j].ModifyIndex
	})
}

func extractNetworkID(e entry) string {
	slicedKey := strings.Split(e.Key, "/")
	if len(slicedKey) > 5 {
		return slicedKey[len(slicedKey)-2]
	}
	return ""
}

func GetNewestNetwork(data []byte) (string, error) {
	entries, err := convertDataToConsulEntry(data)
	if err != nil {
		return "", err
	}
	if len(entries) == 0 {
		return "", nil
	}
	sortByModifyIndex(&entries)
	entry := extractNetworkID(entries[0])
	return entry, nil
}
