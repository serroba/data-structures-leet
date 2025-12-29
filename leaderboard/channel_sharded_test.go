package leaderboard

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestLBChanSharded_OverallAverageRatings(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	lb.RateAgent("Alice", 5, "2025-03-12")
	lb.RateAgent("Bob", 4, "2025-03-13")
	lb.RateAgent("Alice", 3, "2025-03-15")
	lb.RateAgent("Bob", 5, "2025-03-18")

	time.Sleep(50 * time.Millisecond)

	got := lb.GetAverageRatings()
	want := []string{"Bob,4.5", "Alice,4.0"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAverageRatings() = %v, want %v", got, want)
	}
}

func TestLBChanSharded_BestAgentsByMonth(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	lb.RateAgent("Alice", 5, "2025-02-02")
	lb.RateAgent("Bob", 3, "2025-02-05")
	lb.RateAgent("Charlie", 4, "2025-02-10")
	lb.RateAgent("Bob", 5, "2025-03-12")
	lb.RateAgent("Alice", 2, "2025-03-15")

	time.Sleep(50 * time.Millisecond)

	got := lb.GetBestAgentsByMonth("2025-02")
	want := []string{"Alice,5.0", "Charlie,4.0", "Bob,3.0"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetBestAgentsByMonth(\"2025-02\") = %v, want %v", got, want)
	}
}

func TestLBChanSharded_EmptyLeaderboard(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	got := lb.GetAverageRatings()
	if len(got) != 0 {
		t.Errorf("GetAverageRatings() on empty = %v, want empty", got)
	}

	gotMonth := lb.GetBestAgentsByMonth("2025-01")
	if len(gotMonth) != 0 {
		t.Errorf("GetBestAgentsByMonth on empty = %v, want empty", gotMonth)
	}
}

func TestLBChanSharded_SingleAgentSingleRating(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	lb.RateAgent("Solo", 5, "2025-01-15")

	time.Sleep(50 * time.Millisecond)

	got := lb.GetAverageRatings()
	want := []string{"Solo,5.0"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAverageRatings() = %v, want %v", got, want)
	}

	gotMonth := lb.GetBestAgentsByMonth("2025-01")
	if !reflect.DeepEqual(gotMonth, want) {
		t.Errorf("GetBestAgentsByMonth(\"2025-01\") = %v, want %v", gotMonth, want)
	}
}

func TestLBChanSharded_TieBreakingByName(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	lb.RateAgent("Zara", 4, "2025-01-01")
	lb.RateAgent("Alice", 4, "2025-01-02")
	lb.RateAgent("Mike", 4, "2025-01-03")

	time.Sleep(50 * time.Millisecond)

	got := lb.GetAverageRatings()
	want := []string{"Alice,4.0", "Mike,4.0", "Zara,4.0"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("GetAverageRatings() = %v, want %v", got, want)
	}
}

func TestLBChanSharded_QueryNonExistentMonth(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	lb.RateAgent("Alice", 5, "2025-01-01")

	time.Sleep(50 * time.Millisecond)

	got := lb.GetBestAgentsByMonth("2025-12")
	if len(got) != 0 {
		t.Errorf("GetBestAgentsByMonth for non-existent month = %v, want empty", got)
	}
}

func TestLBChanSharded_InvalidInputs(t *testing.T) {
	lb := NewLBChanSharded(4, 256)

	// Empty agent name
	lb.RateAgent("", 5, "2025-01-01")
	time.Sleep(20 * time.Millisecond)
	if got := lb.GetAverageRatings(); len(got) != 0 {
		t.Errorf("Empty agent should be ignored, got %v", got)
	}

	// Invalid rating (too low)
	lb.RateAgent("Test", 0, "2025-01-01")
	time.Sleep(20 * time.Millisecond)
	if got := lb.GetAverageRatings(); len(got) != 0 {
		t.Errorf("Rating 0 should be ignored, got %v", got)
	}

	// Invalid rating (too high)
	lb.RateAgent("Test", 6, "2025-01-01")
	time.Sleep(20 * time.Millisecond)
	if got := lb.GetAverageRatings(); len(got) != 0 {
		t.Errorf("Rating 6 should be ignored, got %v", got)
	}

	// Invalid date (too short)
	lb.RateAgent("Test", 5, "2025")
	time.Sleep(20 * time.Millisecond)
	if got := lb.GetAverageRatings(); len(got) != 0 {
		t.Errorf("Short date should be ignored, got %v", got)
	}

	// Valid input after invalids
	lb.RateAgent("Valid", 5, "2025-01-01")
	time.Sleep(20 * time.Millisecond)
	if got := lb.GetAverageRatings(); len(got) != 1 || got[0] != "Valid,5.0" {
		t.Errorf("Valid input after invalids should work, got %v", got)
	}
}

func TestLBChanSharded_ManyAgentsUniformAccess(t *testing.T) {
	lb := NewLBChanSharded(16, 4096)
	numAgents := 1000

	// Create agents with uniform rating distribution
	for i := 0; i < numAgents; i++ {
		agent := fmt.Sprintf("Agent%04d", i)
		// Each agent gets rating = (i % 5) + 1, so ratings 1-5
		lb.RateAgent(agent, (i%5)+1, "2025-01-15")
	}

	time.Sleep(100 * time.Millisecond)

	got := lb.GetAverageRatings()
	if len(got) != numAgents {
		t.Errorf("Expected %d agents, got %d", numAgents, len(got))
	}

	// Verify first few entries are sorted correctly (highest rating first)
	// Agents with rating 5: Agent0004, Agent0009, Agent0014, ...
	if len(got) > 0 && got[0] != "Agent0004,5.0" {
		t.Errorf("First agent should be Agent0004 with 5.0, got %s", got[0])
	}
}

func TestLBChanSharded_ManyAgentsMultipleRatings(t *testing.T) {
	lb := NewLBChanSharded(16, 4096)
	numAgents := 100

	// Each agent gets multiple ratings
	for i := 0; i < numAgents; i++ {
		agent := fmt.Sprintf("Agent%03d", i)
		lb.RateAgent(agent, 3, "2025-01-01")
		lb.RateAgent(agent, 4, "2025-01-15")
		lb.RateAgent(agent, 5, "2025-02-01")
	}

	time.Sleep(100 * time.Millisecond)

	got := lb.GetAverageRatings()
	if len(got) != numAgents {
		t.Errorf("Expected %d agents, got %d", numAgents, len(got))
	}

	// All agents should have average of 4.0 ((3+4+5)/3 = 4.0)
	for _, entry := range got {
		if entry[len(entry)-3:] != "4.0" {
			t.Errorf("Expected all agents to have 4.0 average, got %s", entry)
		}
	}

	// Check month-specific query
	janRatings := lb.GetBestAgentsByMonth("2025-01")
	if len(janRatings) != numAgents {
		t.Errorf("Expected %d agents in January, got %d", numAgents, len(janRatings))
	}

	// January average should be 3.5 ((3+4)/2 = 3.5)
	for _, entry := range janRatings {
		if entry[len(entry)-3:] != "3.5" {
			t.Errorf("Expected all agents to have 3.5 January average, got %s", entry)
		}
	}
}

func TestLBChanSharded_ConcurrentAccess(t *testing.T) {
	lb := NewLBChanSharded(16, 4096)
	numAgents := 100
	numGoroutines := 10
	ratingsPerGoroutine := 100

	var wg sync.WaitGroup
	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func(gid int) {
			defer wg.Done()
			for i := 0; i < ratingsPerGoroutine; i++ {
				agent := fmt.Sprintf("Agent%03d", (gid*ratingsPerGoroutine+i)%numAgents)
				lb.RateAgent(agent, (i%5)+1, "2025-01-15")
			}
		}(g)
	}
	wg.Wait()

	time.Sleep(100 * time.Millisecond)

	got := lb.GetAverageRatings()
	if len(got) != numAgents {
		t.Errorf("Expected %d agents after concurrent access, got %d", numAgents, len(got))
	}
}

// Benchmarks

func BenchmarkLBChanSharded_Write(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
	}
}

func BenchmarkLBChanSharded_Read(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)

	// Pre-populate with data
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}
	for i := 0; i < 10000; i++ {
		lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
	}
	time.Sleep(100 * time.Millisecond)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.GetAverageRatings()
	}
}

func BenchmarkLBChanSharded_ReadByMonth(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)

	// Pre-populate with data
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}
	for i := 0; i < 10000; i++ {
		lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
	}
	time.Sleep(100 * time.Millisecond)

	months := []string{"2025-01", "2025-02", "2025-03"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lb.GetBestAgentsByMonth(months[i%len(months)])
	}
}

func BenchmarkLBChanSharded_ConcurrentReads(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)

	// Pre-populate with data
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}
	for i := 0; i < 10000; i++ {
		lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
	}
	time.Sleep(100 * time.Millisecond)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lb.GetAverageRatings()
		}
	})
}

func BenchmarkLBChanSharded_ConcurrentWrites(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
			i++
		}
	})
}

func BenchmarkLBChanSharded_MixedReadHeavy(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}

	// Pre-populate
	for i := 0; i < 1000; i++ {
		lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
	}
	time.Sleep(50 * time.Millisecond)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 == 0 { // 10% writes, 90% reads
				lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
			} else {
				lb.GetAverageRatings()
			}
			i++
		}
	})
}

func BenchmarkLBChanSharded_MixedWriteHeavy(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)
	agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
	dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}

	// Pre-populate
	for i := 0; i < 1000; i++ {
		lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
	}
	time.Sleep(50 * time.Millisecond)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 < 9 { // 90% writes, 10% reads
				lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
			} else {
				lb.GetAverageRatings()
			}
			i++
		}
	})
}

func BenchmarkLBChanSharded_ScaleAgents(b *testing.B) {
	for _, numAgents := range []int{10, 100, 1000, 10000} {
		b.Run(fmt.Sprintf("agents_%d", numAgents), func(b *testing.B) {
			lb := NewLBChanSharded(16, 4096)

			// Pre-populate
			for i := 0; i < numAgents*10; i++ {
				agent := fmt.Sprintf("Agent%d", i%numAgents)
				lb.RateAgent(agent, (i%5)+1, "2025-01-01")
			}
			time.Sleep(100 * time.Millisecond)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lb.GetAverageRatings()
			}
		})
	}
}

func BenchmarkLBChanSharded_ContentionTest(b *testing.B) {
	for _, numGoroutines := range []int{1, 2, 4, 8, 16, 32} {
		b.Run(fmt.Sprintf("goroutines_%d", numGoroutines), func(b *testing.B) {
			lb := NewLBChanSharded(16, 4096)
			agents := []string{"Alice", "Bob", "Charlie", "Diana"}

			// Pre-populate
			for i := 0; i < 1000; i++ {
				lb.RateAgent(agents[i%len(agents)], (i%5)+1, "2025-01-01")
			}
			time.Sleep(50 * time.Millisecond)

			b.ResetTimer()
			var wg sync.WaitGroup
			opsPerGoroutine := b.N / numGoroutines

			for g := 0; g < numGoroutines; g++ {
				wg.Add(1)
				go func(id int) {
					defer wg.Done()
					for i := 0; i < opsPerGoroutine; i++ {
						if i%5 == 0 {
							lb.RateAgent(agents[i%len(agents)], (i%5)+1, "2025-01-01")
						} else {
							lb.GetAverageRatings()
						}
					}
				}(g)
			}
			wg.Wait()
		})
	}
}

func BenchmarkLBChanSharded_ScaleShards(b *testing.B) {
	for _, numShards := range []int{1, 2, 4, 8, 16, 32, 64} {
		b.Run(fmt.Sprintf("shards_%d", numShards), func(b *testing.B) {
			lb := NewLBChanSharded(numShards, 4096)
			agents := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Henry"}
			dates := []string{"2025-01-01", "2025-01-15", "2025-02-01", "2025-02-15", "2025-03-01"}

			// Pre-populate
			for i := 0; i < 1000; i++ {
				lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
			}
			time.Sleep(50 * time.Millisecond)

			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				i := 0
				for pb.Next() {
					if i%5 == 0 {
						lb.RateAgent(agents[i%len(agents)], (i%5)+1, dates[i%len(dates)])
					} else {
						lb.GetAverageRatings()
					}
					i++
				}
			})
		})
	}
}

func BenchmarkLBChanSharded_UniformAgentDistribution(b *testing.B) {
	lb := NewLBChanSharded(16, 4096)

	// Create 10000 agents for truly uniform distribution across shards
	agents := make([]string, 10000)
	for i := range agents {
		agents[i] = fmt.Sprintf("Agent%05d", i)
	}

	// Pre-populate with uniform access
	rng := rand.New(rand.NewSource(42))
	for i := 0; i < 100000; i++ {
		lb.RateAgent(agents[rng.Intn(len(agents))], rng.Intn(5)+1, "2025-01-01")
	}
	time.Sleep(200 * time.Millisecond)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		localRng := rand.New(rand.NewSource(rand.Int63()))
		for pb.Next() {
			if localRng.Intn(10) == 0 {
				lb.RateAgent(agents[localRng.Intn(len(agents))], localRng.Intn(5)+1, "2025-01-01")
			} else {
				lb.GetAverageRatings()
			}
		}
	})
}
