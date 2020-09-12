package com.doit.java.chap01;

import java.util.Scanner;

public class Max {
    public static int getMax(Integer[] nums) {
        int max = nums[0];
        for(int i =0; i < nums.length; i++) {
            if (nums[i] > max)
                max = nums[i];
        }
        return max;
    }


    public static void main(String[] args) {
        final Scanner sc = new Scanner(System.in);
        final int SIZE = 3;
        Integer[] arr = new Integer[SIZE];

        for (int i =0; i< arr.length; i++) {
            System.out.print("정수 " + (i+1) + " : ");
            arr[i] = sc.nextInt();
        }

        int max = getMax(arr);
        System.out.println("\nMAX = " + max);

    }
}
