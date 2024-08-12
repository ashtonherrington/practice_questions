package practice_question

type IntHash struct {
	buckets       int
	actualBuckets map[int]map[int]int
}

func NewIntHash(buckets int) IntHash {
	return IntHash{
		buckets:       buckets,
		actualBuckets: make(map[int]map[int]int, buckets),
	}
}

func (s IntHash) EqualContents(o IntHash) bool {
	if len(s.actualBuckets) != len(o.actualBuckets) {
		return false
	}

	for key, bucket := range s.actualBuckets {
		if len(bucket) != len(o.actualBuckets[key]) {
			return false
		}

		for bucketKey, value := range bucket {
			if value != o.actualBuckets[key][bucketKey] {
				return false
			}
		}
	}
	return true
}

func (s IntHash) Add(v int) {
	bucket := s.hash(v)

	if s.actualBuckets[bucket] == nil {
		s.actualBuckets[bucket] = make(map[int]int)
	}

	//Not optimal to call hash twice here, could abstract most of Get for reusage to "getInternal"
	existingCount := s.GetOccurances(v)

	s.actualBuckets[bucket][v] = existingCount + 1
}

func (s IntHash) GetOccurances(v int) int {
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

func (s IntHash) hash(v int) int {
	return v % s.buckets
}
