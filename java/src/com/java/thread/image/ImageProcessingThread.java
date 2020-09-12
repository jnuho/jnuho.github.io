package com.java.thread.image;

public class ImageProcessingThread extends Thread{

	private int rowNo;

	public ImageProcessingThread(int rowNo){
		this.rowNo = rowNo;
	}

	@Override
	public void run() {
		System.out.println("PROCESSING...ROW..." + rowNo);
	}

	public int getRowNo() {
		return rowNo;
	}

	public void setRowNo(int rowNo) {
		this.rowNo = rowNo;
	}

}
