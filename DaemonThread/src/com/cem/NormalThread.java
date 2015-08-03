package com.cem;

public class NormalThread extends Thread {
	@Override
	public void run() {
		for (int i = 0; i < 2; i++) {
			System.out.println("Normal thread is running");
			try {
				Thread.sleep(1000);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}
}