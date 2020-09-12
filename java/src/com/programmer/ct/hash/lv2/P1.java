package com.programmer.ct.hash.lv2;

import java.util.HashSet;
import java.util.Set;

public class P1 {
	public static void main(String[] args) {
		P1 problem = new P1();
	}

	/**
	 * 한 번호가 다른 번호의 접두어인 경우가 있으면 true. 없으면 false. Examples:
	 * [119, 97674223, 1195524421]	false
	 * [123,456,789]	true
	 *
	 * @param phone_book 전화번호 부
	 * @return
	 */
	public boolean solution(String[] phone_book) {
		boolean answer = true;

		// brute force
//		Arrays.sort(phone_book);
//		for(int i =0; i< phone_book.length; i++) {
//			for(int j = i+1; j < phone_book.length; j++) {
//			    String large = phone_book[i].length() > phone_book[j].length() ? phone_book[i] : phone_book[j];
//				String small = phone_book[i].length() <= phone_book[j].length() ? phone_book[i] : phone_book[j];
//				String compare = large.substring(0,small.length());
//				if (compare.equals(small) ){
//					answer = false;
//				}
//
//			}
//		}

        Set<String> set = new HashSet<String>();
        for(int i =0 ;i<phone_book.length; i++){
            String var = phone_book[i].substring(0, phone_book[i].length());
            if( set.contains(var)) return false;
            set.add(var);


		}

		return answer;
	}
}
