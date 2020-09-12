package com.doit.java.chap02;

public class ReverseArray {
    public static void swap(int[] arr, int s, int e) {
        int temp = arr[s];
        arr[s] = arr[e];
        arr[e] = temp;
    }

    public static void reverse(int[] arr) {
        for (int i=0; i< arr.length/2; i++) {
            swap(arr, i, arr.length -i-1);
        }

    }

    public static void main(String[] args) {
        int[] arr = {1,2,3,4,5};
        reverse(arr);
        for(int i =0; i< arr.length; i++) {
            System.out.print(arr[i] + " " );
        }
        System.out.println();

    }
}
