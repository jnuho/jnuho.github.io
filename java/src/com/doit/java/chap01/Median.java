package com.doit.java.chap01;

import java.util.Scanner;

public class Median {
    public static int getMedian(Integer[] nums) {
        int a = nums[0];
        int b = nums[1];
        int c = nums[2];
        if(a >= b) {
            if (b >= c)
                return b;
            else if (a >= c)
                return c;
            else
                return a;
        }
        else if(b <= c) { // a<b
            return b;
        }
        else if (a >= c) {//b> c & b >a
            return a;
        }
        else
            return c;
    }


    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        final int SIZE = 3;
        Integer[] arr = new Integer[SIZE];

        for (int i =0; i< arr.length; i++) {
            System.out.print("정수 " + (i+1) + " : ");
            arr[i] = sc.nextInt();
        }

        int med = getMedian(arr);
        System.out.println("\nMEDIAN = " + med);

    }
}
