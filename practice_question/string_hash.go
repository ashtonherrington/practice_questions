package practice_question

import "hash/fnv"

type StringHash struct {
	buckets       int
	actualBuckets map[uint32]map[string]int
}

func NewStringHash(buckets int) StringHash {
	return StringHash{
		buckets:       buckets,
		actualBuckets: make(map[uint32]map[string]int, buckets),
	}
}

func (s StringHash) Add(v string) {
	bucket := s.hash(v)

	if s.actualBuckets[bucket] == nil {
		s.actualBuckets[bucket] = make(map[string]int)
	}

	//Not optimal to call hash twice here, could abstract most of Get for reuse to "getInternal"
	existingCount := s.GetOccurances(v)

	s.actualBuckets[bucket][v] = existingCount + 1
}

func (s StringHash) GetOccurances(v string) int {
	bucket := s.hash(v)

	// if bucket has not been created, then it clearly cannot exist
	if s.actualBuckets[bucket] == nil {
		return 0
	}

	instances, ok := s.actualBuckets[bucket][v]
	if !ok {
		return 0
	}
	return int(instances)
}

func (s StringHash) hash(v string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(v))
	return h.Sum32() % uint32(s.buckets)
}
