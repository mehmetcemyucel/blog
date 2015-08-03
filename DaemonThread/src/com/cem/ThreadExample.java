package com.cem;

public class ThreadExample {
	public static void main(String[] args) {
		Thread daemon = new DaemonThread();
		daemon.setDaemon(true);
		Thread normal = new NormalThread();

		daemon.start();
		normal.start();
	}
}