package leaderboard

import "sync"

type LBWriteHeavy struct {
	mu      sync.Mutex
	overall map[string]agg
	byMonth map[string]map[string]agg // month -> agent -> agg
}

func NewLBWriteHeavy() *LBWriteHeavy {
	return &LBWriteHeavy{
		overall: make(map[string]agg, 1024),
		byMonth: make(map[string]map[string]agg, 128),
	}
}

func (lb *LBWriteHeavy) RateAgent(agent string, rating int, date string) {
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

func (lb *LBWriteHeavy) GetAverageRatings() []string {
	// Copy under lock; sort outside.
	lb.mu.Lock()

	snap := make(map[string]agg, len(lb.overall))
	for k, v := range lb.overall {
		snap[k] = v
	}

	lb.mu.Unlock()

	return formatAndSort(snap)
}

func (lb *LBWriteHeavy) GetBestAgentsByMonth(month string) []string {
	lb.mu.Lock()

	m := lb.byMonth[month]
	if m == nil {
		lb.mu.Unlock()

		return nil
	}

	snap := make(map[string]agg, len(m))
	for k, v := range m {
		snap[k] = v
	}

	lb.mu.Unlock()

	return formatAndSort(snap)
}
