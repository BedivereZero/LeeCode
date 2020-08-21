package algorithms

import "log"

type bucket struct {
	min int
	max int
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	if min == max {
		return 0
	}

	log.Printf("Min: %d, Max: %d", min, max)

	bucketSize := (max - min) / (len(nums) - 1)
	if bucketSize == 0 {
		bucketSize = 1
	}
	log.Printf("Bucket size: %d", bucketSize)

	buckets := make([]*bucket, (max-min)/bucketSize+1)
	log.Printf("Buckets: %v", buckets)

	for _, n := range nums {
		i := (n - min) / bucketSize
		if buckets[i] == nil {
			buckets[i] = &bucket{min: n, max: n}
		} else {
			if n < buckets[i].min {
				buckets[i].min = n
			}
			if n > buckets[i].max {
				buckets[i].max = n
			}
		}
	}

	log.Printf("Buckets: %v", buckets)

	gap := 0
	b := buckets[0]
	for i := 1; i < len(buckets); i++ {
		if buckets[i] == nil {
			continue
		}

		if gap < buckets[i].min-b.max {
			gap = buckets[i].min - b.max
		}

		b = buckets[i]
	}

	return gap
}
