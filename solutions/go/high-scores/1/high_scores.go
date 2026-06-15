package highscores

import (
	"sort"
)

type HighScores struct {
	scores []int
}

func NewHighScores(scores []int) *HighScores {
	copiedScores := make([]int, len(scores))
	copy(copiedScores, scores)

	return &HighScores{
		scores: copiedScores,
	}
}

func (s *HighScores) Scores() []int {
	return s.scores
}

func (s *HighScores) Latest() int {
	if len(s.scores) == 0 {
		return 0
	}
	return s.scores[len(s.scores)-1]
}

func (s *HighScores) PersonalBest() int {
	if len(s.scores) == 0 {
		return 0
	}

	best := s.scores[0]
	for _, score := range s.scores {
		if score > best {
			best = score
		}
	}
	return best
}

func (s *HighScores) TopThree() []int {
	
	sortedScores := make([]int, len(s.scores))
	copy(sortedScores, s.scores)

	sort.Slice(sortedScores, func(i, j int) bool {
		return sortedScores[i] > sortedScores[j]
	})
    
	limit := 3
	if len(sortedScores) < 3 {
		limit = len(sortedScores)
	}

	return sortedScores[:limit]
}
