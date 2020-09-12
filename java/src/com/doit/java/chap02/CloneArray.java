package com.doit.java.chap02;


public class CloneArray {
    public static void main(String[] args) {
        // 배열.clone()
        // 복제된 배열은 기존 배열을 참조하지 않음.
        int[] a = {1,2,3,4,5};
        int[] b = a.clone();
        b[2] = 0;
        for(int i : a) {
            System.out.print(i + " ");
        }
        System.out.println();
        for(int i : b) {
            System.out.print(i+ " ");
        }
        System.out.println();
    }
}
