package com.doit.java.chap02;

import java.util.Random;

public class MaxArrayRand {

    public static int maxOf(int[] arr){
        int max = arr[0];
        for(int i =0; i< arr.length; i++) {
            if (arr[i] > max) {
                max = arr[i];
            }
        }
        return max;
    }

    public static void main(String[] args) {
        int sizeOfArr = 10;
        int[] arr = new int[sizeOfArr];

        int n = 100;
//        int seed = 999;
//        Random rand = new Random(seed); //pseudorandom; same seed same rand number
        Random rand = new Random(); // different seed each time (based on current time)
        for (int i = 0; i < arr.length; i++) {
            // 0 ~ n-1 사이의 난수 (random number)
            arr[i] = rand.nextInt(n);
        }

        System.out.print("Array of size "  + arr.length + " has a Max Val = ");
        System.out.println(maxOf(arr));
    }
}
