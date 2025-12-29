package leaderboard

import (
	"hash/maphash"
	"time"
)

type LBChanSharded struct {
	shards []chanShard
	mask   uint64
	seed   maphash.Seed
}

type chanShard struct {
	in   chan rateMsg
	req  chan snapReq
	stop chan struct{}
}

type rateMsg struct {
	agent  string
	month  string
	rating int
}

type snapReq struct {
	month string
	resp  chan map[string]agg
}

func NewLBChanSharded(numShards int, queueSize int) *LBChanSharded {
	if numShards < 1 {
		numShards = 64
	}

	if queueSize < 1 {
		queueSize = 4096
	}

	n := 1
	for n < numShards {
		n <<= 1
	}

	lb := &LBChanSharded{
		shards: make([]chanShard, n),
		mask:   uint64(n - 1),
		seed:   maphash.MakeSeed(),
	}

	for i := range lb.shards {
		sh := &lb.shards[i]
		sh.in = make(chan rateMsg, queueSize)
		sh.req = make(chan snapReq, 8)

		sh.stop = make(chan struct{})
		go runShard(sh)
	}

	return lb
}

func (lb *LBChanSharded) shardIndex(agent string) int {
	var h maphash.Hash
	h.SetSeed(lb.seed)
	h.WriteString(agent)

	return int(h.Sum64() & lb.mask)
}

func (lb *LBChanSharded) RateAgent(agent string, rating int, date string) {
	if agent == "" || rating < 1 || rating > 5 || len(date) < 7 {
		return
	}

	msg := rateMsg{agent: agent, month: date[:7], rating: rating}

	sh := &lb.shards[lb.shardIndex(agent)]

	// Non-blocking enqueue: under overload, you must choose a policy.
	// Here we drop. Alternatives: block, or return bool, or count drops.
	select {
	case sh.in <- msg:
	default:
		// drop under pressure
	}
}

func (lb *LBChanSharded) GetAverageRatings() []string {
	merged := make(map[string]agg, 4096)

	for i := range lb.shards {
		snap := lb.getShardSnap(i, "")
		for k, v := range snap {
			a := merged[k]
			a.sum += v.sum
			a.count += v.count
			merged[k] = a
		}
	}

	return formatAndSort(merged)
}

func (lb *LBChanSharded) GetBestAgentsByMonth(month string) []string {
	merged := make(map[string]agg, 2048)

	for i := range lb.shards {
		snap := lb.getShardSnap(i, month)
		for k, v := range snap {
			a := merged[k]
			a.sum += v.sum
			a.count += v.count
			merged[k] = a
		}
	}

	return formatAndSort(merged)
}

func (lb *LBChanSharded) getShardSnap(i int, month string) map[string]agg {
	resp := make(chan map[string]agg, 1)
	req := snapReq{month: month, resp: resp}

	sh := &lb.shards[i]

	// bounded wait so a stuck shard doesn't hang reads forever
	select {
	case sh.req <- req:
	case <-time.After(200 * time.Millisecond):
		return nil
	}

	select {
	case snap := <-resp:
		return snap
	case <-time.After(200 * time.Millisecond):
		return nil
	}
}

func runShard(sh *chanShard) {
	overall := make(map[string]agg, 1024)
	byMonth := make(map[string]map[string]agg, 128)

	for {
		select {
		case msg := <-sh.in:
			a := overall[msg.agent]
			a.sum += int64(msg.rating)
			a.count++
			overall[msg.agent] = a

			m := byMonth[msg.month]
			if m == nil {
				m = make(map[string]agg, 256)
				byMonth[msg.month] = m
			}

			ma := m[msg.agent]
			ma.sum += int64(msg.rating)
			ma.count++
			m[msg.agent] = ma

		case req := <-sh.req:
			if req.month == "" {
				snap := make(map[string]agg, len(overall))
				for k, v := range overall {
					snap[k] = v
				}

				req.resp <- snap

				continue
			}

			m := byMonth[req.month]
			if m == nil {
				req.resp <- nil

				continue
			}

			snap := make(map[string]agg, len(m))
			for k, v := range m {
				snap[k] = v
			}

			req.resp <- snap

		case <-sh.stop:
			return
		}
	}
}
