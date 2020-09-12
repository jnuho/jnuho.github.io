package com.java.thread.image;

public class MainThread {
	public static  void main(String[] args) {
		int imageRows = 768;
		int range = imageRows;
		ImageProcessingThread[] workers = new ImageProcessingThread[range];

		for(int i =0; i < workers.length; i++) {
			workers[i] = new ImageProcessingThread(i);
			workers[i].start();
		}

		for(int i =0; i < workers.length; i++) {
			try{
				workers[i].join();
			} catch(InterruptedException e) {
				e.printStackTrace();
			}
		}
		System.out.println("\n...MAIN THREAD...\n");
	}
}
