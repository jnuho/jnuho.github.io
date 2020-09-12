package com.programmer.ct.hash.lv1;

import java.util.HashMap;
import java.util.Map;

public class P1 {

	public static void main(String[] args){
		P1 problem = new P1();
		String answer = problem.solution(new String[]{"mislav", "stanko", "mislav", "ana"},	new String[]{"stanko", "ana", "mislav"});
		System.out.println("ANSWER = " + answer);
	}

	/**
	 * 마라톤 전체참가자 - 완주 제외한 미완주 참가자 1명의 이름을 반환
	 *
	 * @param participant 참가자
	 * @param completion 완주자
	 * @return answer 미완주자
	 */
	public String solution(String[] participant, String[] completion) {
		Map<String, Integer> map = new HashMap<String, Integer>();

		for(String person : participant){
			map.put(person, map.getOrDefault(person, 0) + 1);
		}
		for(String person : completion) {
			if(map.containsKey(person) && map.get(person) > 1) {
				map.put(person, map.get(person) - 1);
			}
			else if(map.containsKey(person) && map.get(person) == 1) {
				map.remove(person);
			}
		}
		String answer = "";
		for(Map.Entry<String, Integer> entry : map.entrySet()) {
			answer = entry.getKey();
		}
		return answer;
	}
}

