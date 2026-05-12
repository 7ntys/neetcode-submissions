func findMin(nums []int) int {
	l, r := 0, len(nums) - 1

	for l < r {
		mid := l + (r -l) / 2

		if nums[r] > nums[mid] {
			// Left side is sorted : 
			r = mid
		} else {
			l = mid + 1
			// Ok now i get it, its because mid itself COULD be the minimum of the array, thats why
		}
	}

	return nums[l]
}

/*
BRUTE FORCE : 
Linear search for min in an array : O(n)

BINARY SEARCH :
My idea here is that if we can find the pivot of the rotated array, we can deduce the min element : 

first, we know the condition would be l < r BECAUSE we are not looking for a specific element, and a border 
what should be the rule ? Lets look at some examples:  

[3,4,5,6,1,2]

Here : l = 0, r = 5 => mid = 2

we have nums[l] = 3, nums[mid] = 5, nums[r] = 2
Which side is sorted ? that can be checked with : 
if nums[l] < nums[mid] {
		left side sorted
	}
if nums[mid] < nums[r] {
	right side sorted
}

I still remember we need to have nums[l] <= nums[mid], but i don't remember why
Lets try to implement and see if it makes sense

Ok after first run, i realized i was missing some condition : left side sorted doesn't mean the minimum is necessarily there

after trying some things, i still don't get it :( , curr timer = 10 min

*/