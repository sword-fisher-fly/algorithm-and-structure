package binarysearch

func mergeTwoSortedArray(nums1, nums2 []int) []int {
	result := make([]int, len(nums1)+len(nums2))
	idx1, idx2 := 0, 0
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			result[idx1+idx2] = nums1[idx1]
			idx1++
		} else {
			result[idx1+idx2] = nums2[idx2]
			idx2++
		}
	}

	for idx1 < len(nums1) {
		result[idx1+idx2] = nums1[idx1]
		idx1++
	}

	for idx2 < len(nums2) {
		result[idx1+idx2] = nums2[idx2]
		idx2++
	}

	return result
}

func FindMedianInTwoSortedArrayByForce(nums1, nums2 []int) int {
	mergedArray := mergeTwoSortedArray(nums1, nums2)
	if len(mergedArray) == 0 {
		return -1
	}

	if len(mergedArray)%2 == 1 {
		return mergedArray[len(mergedArray)/2]
	}

	return (mergedArray[len(mergedArray)/2] + mergedArray[len(mergedArray)/2-1]) / 2
}

// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
func FindMediaInTwoSortedArrayII(nums1, nums2 []int) int {

	return -1
}

/*
class Solution {
  public double findMedianSortedArrays(int[] nums1, int[] nums2) {
        int m = nums1.length;
        int n = nums2.length;
        int left = (m + n + 1) / 2;
        int right = (m + n + 2) / 2;
        return (findKth(nums1, 0, nums2, 0, left) + findKth(nums1, 0, nums2, 0, right)) / 2.0;
    }
    //i: nums1的起始位置 j: nums2的起始位置
    public int findKth(int[] nums1, int i, int[] nums2, int j, int k){
        if( i >= nums1.length) return nums2[j + k - 1];//nums1为空数组
        if( j >= nums2.length) return nums1[i + k - 1];//nums2为空数组
        if(k == 1){
            return Math.min(nums1[i], nums2[j]);
        }
        int midVal1 = (i + k / 2 - 1 < nums1.length) ? nums1[i + k / 2 - 1] : Integer.MAX_VALUE;
        int midVal2 = (j + k / 2 - 1 < nums2.length) ? nums2[j + k / 2 - 1] : Integer.MAX_VALUE;
        if(midVal1 < midVal2){
            return findKth(nums1, i + k / 2, nums2, j , k - k / 2);
        }else{
            return findKth(nums1, i, nums2, j + k / 2 , k - k / 2);
        }
    }
}
*/
