package com.java.thread.pool;

import java.util.concurrent.ConcurrentLinkedQueue;

public class ThreadInThePool extends Thread{
	ConcurrentLinkedQueue<Runnable> taskQueue;

	private boolean isRunning = true;
	private int ct = 0;

	private static int thrdCt=0;
	private static int total = 0;
	private static long startTime,elapsedTime;

	public ThreadInThePool(ConcurrentLinkedQueue<Runnable> queue){
		taskQueue=queue;
	}

	//this will stops the loop in the run() method before finishing
	//all sub-tasks in the taskQueue.
	public void abort(){
		isRunning = false;
	}

	public void run(){
		try{
			while(isRunning){
				Runnable subTask=taskQueue.poll();

				if(subTask==null){
					break; //break and the thread terminates when all the sub-tasks
					//in the taskQueue are over!
				}
				subTask.run();
				CountPrimes tsk = (CountPrimes) subTask;
				ct+=tsk.getCount();
			}
			addToTotal(ct);

		}
		finally{
			increseThrdCt();
			if (thrdCt==8){
				System.out.println(total);
				elapsedTime= System.currentTimeMillis() - startTime;
				System.out.println("Total elapsed time: " + (elapsedTime/1000.0) + " seconds");

			}
		}
	}

	synchronized private static void increseThrdCt(){
		thrdCt++;
	}

	synchronized private static void addToTotal(int x) {
		total = total + x;
	}

	// Total elapsed time: 0.971 seconds
	// much faster than CountPrimes > main 메소드
	public static void main(String[] args){
		startTime = System.currentTimeMillis();
		ConcurrentLinkedQueue<Runnable> queue = new ConcurrentLinkedQueue<Runnable>();
		int lo = 1;
		int hi = 10000000/20;

		for (int i=0;i<20;i++){
			CountPrimes obj = new CountPrimes(lo,hi);
			queue.add(obj);
			lo = hi;
			hi = hi+10000000/20;

		}

		ThreadInThePool[] thrdPool = new ThreadInThePool[8];
		for (int i=0;i<8;i++){
			thrdPool[i] = new ThreadInThePool(queue);
			thrdPool[i].start();
		}
	}
}
