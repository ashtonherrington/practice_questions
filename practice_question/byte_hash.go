package practice_question

type ByteHash struct {
	buckets       int
	actualBuckets map[int]map[byte]int
}

func NewByteHash(buckets int) ByteHash {
	return ByteHash{
		buckets:       buckets,
		actualBuckets: make(map[int]map[byte]int, buckets),
	}
}

func (s ByteHash) EqualContents(o ByteHash) bool {
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

func (s ByteHash) Add(v byte) {
	bucket := s.hash(v)

	if s.actualBuckets[bucket] == nil {
		s.actualBuckets[bucket] = make(map[byte]int)
	}

	//Not optimal to call hash twice here, could abstract most of Get for reusage to "getInternal"
	existingCount := s.GetOccurances(v)

	s.actualBuckets[bucket][v] = existingCount + 1
}

func (s ByteHash) GetOccurances(v byte) int {
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

func (s ByteHash) hash(v byte) int {
	return int(v) % s.buckets
}
