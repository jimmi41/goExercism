package electionday
import "strconv"
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var initVotes int
    initVotes = initialVotes
    return &(initVotes)
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
        return 0
    }

    return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if counter == nil {
        return
    }

    *counter += increment
}


func NewElectionResult(name string, votes int) *ElectionResult {

    return &ElectionResult{
        Name:  name,
        Votes: votes,
    }
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return result.Name+" ("+ strconv.Itoa(result.Votes)+")"
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] = results[candidate] - 1
}
