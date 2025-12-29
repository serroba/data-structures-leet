package leaderboard

import (
	"fmt"
	"math"
	"sort"
)

type agg struct {
	sum   int64
	count int64
}

type entry struct {
	name string
	avg  float64
}

func (a agg) avg1() float64 {
	if a.count == 0 {
		return 0
	}

	return math.Round((float64(a.sum)/float64(a.count))*10) / 10
}

func formatAndSort(m map[string]agg) []string {
	entries := make([]entry, 0, len(m))
	for name, a := range m {
		if a.count == 0 {
			continue
		}

		entries = append(entries, entry{name: name, avg: a.avg1()})
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].avg != entries[j].avg {
			return entries[i].avg > entries[j].avg
		}

		return entries[i].name < entries[j].name
	})

	out := make([]string, len(entries))
	for i, e := range entries {
		out[i] = fmt.Sprintf("%s,%.1f", e.name, e.avg)
	}

	return out
}
