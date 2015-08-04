package com.cem;

public class TimeMeasure {

	// public static void main(String[] args) {
	//
	// long beginTime = 0, endTime = 0;
	//
	// String s = "";
	//
	// beginTime = System.currentTimeMillis();
	//
	// for (int i = 0; i < 100000; i++) {
	// s = s.concat(".");
	// }
	//
	// endTime = System.currentTimeMillis();
	//
	// System.out.println(((double) (endTime - beginTime)) / 1000);
	// }

	public static void main(String[] args) {

		Timer t = new Timer();
		String s = "";

		t.start();
		for (int i = 0; i < 100000; i++) {
			s = s.concat(".");
		}

		t.end();
		System.out.println(t.getTotalTimeInSec());
	}
}
