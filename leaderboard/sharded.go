package leaderboard

import (
	"hash/fnv"
	"sync"
)

type LBSharded struct {
	shards []lbShard
	mask   uint64
}

type lbShard struct {
	mu      sync.Mutex
	overall map[string]agg
	byMonth map[string]map[string]agg
}

func NewLBSharded(numShards int) *LBSharded {
	if numShards < 1 {
		numShards = 16
	}
	// force power of two for cheap masking
	n := 1
	for n < numShards {
		n <<= 1
	}

	s := make([]lbShard, n)
	for i := range s {
		s[i].overall = make(map[string]agg, 1024/n+1)
		s[i].byMonth = make(map[string]map[string]agg, 128/n+1)
	}

	return &LBSharded{
		shards: s,
		mask:   uint64(n - 1),
	}
}

func (lb *LBSharded) shardFor(agent string) *lbShard {
	h := fnv.New64a()
	_, _ = h.Write([]byte(agent))

	return &lb.shards[h.Sum64()&lb.mask]
}

func (lb *LBSharded) RateAgent(agent string, rating int, date string) {
	if agent == "" || rating < 1 || rating > 5 || len(date) < 7 {
		return
	}

	month := date[:7]

	sh := lb.shardFor(agent)

	sh.mu.Lock()
	defer sh.mu.Unlock()

	a := sh.overall[agent]
	a.sum += int64(rating)
	a.count++
	sh.overall[agent] = a

	m := sh.byMonth[month]
	if m == nil {
		m = make(map[string]agg, 256)
		sh.byMonth[month] = m
	}

	ma := m[agent]
	ma.sum += int64(rating)
	ma.count++
	m[agent] = ma
}

func (lb *LBSharded) GetAverageRatings() []string {
	merged := make(map[string]agg, 4096)

	for i := range lb.shards {
		sh := &lb.shards[i]
		sh.mu.Lock()

		for k, v := range sh.overall {
			acc := merged[k]
			acc.sum += v.sum
			acc.count += v.count
			merged[k] = acc
		}

		sh.mu.Unlock()
	}

	return formatAndSort(merged)
}

func (lb *LBSharded) GetBestAgentsByMonth(month string) []string {
	merged := make(map[string]agg, 2048)

	for i := range lb.shards {
		sh := &lb.shards[i]
		sh.mu.Lock()

		m := sh.byMonth[month]
		if m != nil {
			for k, v := range m {
				acc := merged[k]
				acc.sum += v.sum
				acc.count += v.count
				merged[k] = acc
			}
		}

		sh.mu.Unlock()
	}

	return formatAndSort(merged)
}
