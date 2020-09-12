package com.java.thread.pool;

public class CountPrimes implements Runnable {

	int count = 0;
	int min, max;

	public CountPrimes(int min, int max){
		this.min = min;
		this.max = max;
	}

	private static boolean isPrime(int val) {
		int sqr = (int)Math.sqrt(val);
		for(int i=2; i <= sqr; i++) {
			if (val % i == 0)
				return false;
		}
		return true;
	}

	@Override
	public void run() {
		for (int i = min; i<=max;i++){
			if(isPrime(i))
				count++;
		}
	}

	public int getCount() {
		return count;
	}

	public void setCount(int count) {
		this.count = count;
	}

	// Total elapsed time: 3.195 seconds
//	public static void main(String[] args){
//		long startTime = System.currentTimeMillis();
//		CountPrimes n=new CountPrimes(1,10000000);
//		n.run();
//		long elapsedTime = System.currentTimeMillis() - startTime;
//		System.out.println(n.getCount());
//		System.out.println("Total elapsed time: " + (elapsedTime/1000.0) + " seconds");
//	}
}
