package com.cem;

public class DaemonThread extends Thread {
	@Override
	public void run() {
		while (true) {
			System.out.println("Daemon thread is running");
			try {
				Thread.sleep(3000);
			} catch (InterruptedException e) {
				e.printStackTrace();
			}
		}
	}
}