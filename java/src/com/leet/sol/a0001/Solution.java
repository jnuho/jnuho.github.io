package com.leet.sol.a0001;


import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

public class Solution {
	/*  Given an array of integers, return indices of the two numbers such that
	 * they add up to a specific target. You may assume that each input would have
	 * exactly one solution, and you may not use the same element twice.
	 * Example: Given nums = [2, 7, 11, 15], target = 9
	 * Because nums[0] + nums[1] = 2 + 7 = 9, return [0, 1].
	 */
	public static void main(String[] args) {

		Solution sol = new Solution();
		int[] indices = null;

		//test1
		int[] nums = {2, 7, 11, 15};
		int target = 9;
		indices = sol.twoSum(nums, target);

		if( indices!=null && indices.length ==2)
			System.out.println("[" + indices[0] + ", " + indices[1] + "]");

		//test2
		indices = sol.twoSum2(nums, target);

		if( indices!=null && indices.length ==2)
			System.out.println("[" + indices[0] + ", " + indices[1] + "]");
	}

	public int[] twoSum2(int[] nums, int target) {

		int[] indices = this.getIndices(nums, target); //must return [0, 1]
		Arrays.sort(indices);

		return indices;
	}

	public int[] twoSum(int[] nums, int target) {

		int[] indices = this.getIndices(nums, target); //must return [0, 1]
		Arrays.sort(indices);

		return indices;
	}

	public int[] getIndices(int[] nums, int target) {

		List<Integer> list = Arrays.stream(nums).boxed().collect(Collectors.toList());

		if (list !=null && list.size() > 0) {
			for(int i =0; i<list.size(); i++) {
				int temp = target - list.get(i);

				if(list.contains(temp) && i != list.indexOf(temp))
					return new int[] {list.indexOf(temp), i};
			}
		}

		return new int[] {};
	}
}