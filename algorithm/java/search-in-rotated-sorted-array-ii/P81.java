func search(nums []int, target int) bool {
    if len(nums) == 0 {
        return false
    }
    l, r := 0, len(nums)-1
    
    for l <= r {
        mid := (l+r)/2
        if nums[mid] == target {
            return true
        } else if nums[l] < nums[r] {
            if nums[mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else if nums[mid] > nums[l] || (nums[mid] == nums[l] && nums[l] > nums[r]){
            // mid at left part
            if nums[mid] > target  && target >= nums[l] {
                r = mid - 1
            } else {
                l = mid + 1
            }
        } else if nums[mid] < nums[r]  || (nums[mid] == nums[r] && nums[l] > nums[r]) {
            // mid at right part
            if (nums[mid] < target && target <= nums[r]) {
                l = mid + 1
            } else {
                r = mid - 1
            }
        } else if(nums[mid] == nums[r] && nums[mid] == nums[l]) {
            l++
            r--
        }
    }
    return false
}
