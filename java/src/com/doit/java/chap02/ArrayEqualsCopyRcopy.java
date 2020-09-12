package com.doit.java.chap02;

public class ArrayEqualsCopyRcopy {
    public static boolean arrEquals(int[] a, int[] b) {
        if(a.length != b.length)
            return false;
        for(int i =0; i<a.length; i++) {
            if (a[i] != b[i])
                return false;
        }
        return true;
    }

    /**
     * Copy elements
     * @param a, b
     */
    public static void copy(int[] a, int[] b){
    	if( a== null || b== null || a.length != b.length)
    	    throw new IllegalArgumentException();

    }
    /**
     * Copy in reverse order
     * @param a, b
     */
    public static void rcopy(int[] a, int[] b) {
        if( a== null || b== null || a.length != b.length)
            throw new IllegalArgumentException();
    }

}
