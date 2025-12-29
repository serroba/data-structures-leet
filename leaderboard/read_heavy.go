package leaderboard

import "sync"

type LBReadHeavy struct {
	mu      sync.RWMutex
	overall map[string]agg
	byMonth map[string]map[string]agg
}

func NewLBReadHeavy() *LBReadHeavy {
	return &LBReadHeavy{
		overall: make(map[string]agg, 1024),
		byMonth: make(map[string]map[string]agg, 128),
	}
}

func (lb *LBReadHeavy) RateAgent(agent string, rating int, date string) {
	if agent == "" || rating < 1 || rating > 5 || len(date) < 7 {
		return
	}

	month := date[:7]

	lb.mu.Lock()
	defer lb.mu.Unlock()

	a := lb.overall[agent]
	a.sum += int64(rating)
	a.count++
	lb.overall[agent] = a

	m := lb.byMonth[month]
	if m == nil {
		m = make(map[string]agg, 256)
		lb.byMonth[month] = m
	}

	ma := m[agent]
	ma.sum += int64(rating)
	ma.count++
	m[agent] = ma
}

func (lb *LBReadHeavy) GetAverageRatings() []string {
	lb.mu.RLock()

	snap := make(map[string]agg, len(lb.overall))
	for k, v := range lb.overall {
		snap[k] = v
	}

	lb.mu.RUnlock()

	return formatAndSort(snap)
}

func (lb *LBReadHeavy) GetBestAgentsByMonth(month string) []string {
	lb.mu.RLock()

	m := lb.byMonth[month]
	if m == nil {
		lb.mu.RUnlock()

		return nil
	}

	snap := make(map[string]agg, len(m))
	for k, v := range m {
		snap[k] = v
	}

	lb.mu.RUnlock()

	return formatAndSort(snap)
}
