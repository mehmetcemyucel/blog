package com.cem;

public class Timer {

	private long startTime = 0, endTime = 0;

	public void start() {
		this.startTime = System.currentTimeMillis();
	}

	public void end() {
		this.endTime = System.currentTimeMillis();
	}

	public long getStartTime() {
		return this.startTime;
	}

	public long getEndTime() {
		return this.endTime;
	}

	public long getTotalTimeInMillis() {
		return this.endTime - this.startTime;
	}

	public double getTotalTimeInSec() {
		return ((double) (this.endTime - this.startTime)) / 1000;
	}
}