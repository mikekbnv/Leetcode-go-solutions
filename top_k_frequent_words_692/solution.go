package top_k_frequent_words_692

import (
	"sort"
	"strings"
)

type Alphabetic []string

var mp = make(map[string]int)

func (list Alphabetic) Len() int {
	return len(list)
}

func (list Alphabetic) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list Alphabetic) Less(i, j int) bool {
	if mp[list[i]] == mp[list[j]] {
		return strings.Compare(list[i], list[j]) == -1
	} else {
		return mp[list[i]] > mp[list[j]]
	}
}

func topKFrequent(words []string, k int) []string {
	mp = make(map[string]int)
	for _, v := range words {
		mp[v]++
	}
	slic := []string{}
	for k, _ := range mp {
		slic = append(slic, k)
	}
	sort.Sort(Alphabetic(slic))
	return slic[:k]
}
