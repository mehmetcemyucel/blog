package com.cem;

import java.util.ArrayList;
import java.util.List;

public class FreeMemory {

	private static final int MEGABYTE_CONSTANT = 1024 * 1024;

	public static void main(String[] args) {
		long freeMemory = Runtime.getRuntime().freeMemory() / MEGABYTE_CONSTANT;
		long totalMemory = Runtime.getRuntime().totalMemory()
				/ MEGABYTE_CONSTANT;
		long maxMemory = Runtime.getRuntime().maxMemory() / MEGABYTE_CONSTANT;

		System.out.println("freeMemory=\t" + freeMemory);
		System.out.println("totalMemory=\t" + totalMemory
				+ " (initial heap miktari)");
		System.out.println("maxMemory=\t" + maxMemory);

		List<String> fillHeapList = new ArrayList<String>();

		for (int i = 0; i < 1000000; i++) {
			fillHeapList.add(String.valueOf(9999));
		}

		System.out.println("\nKOD CALISTI\n");

		freeMemory = Runtime.getRuntime().freeMemory() / MEGABYTE_CONSTANT;
		totalMemory = Runtime.getRuntime().totalMemory() / MEGABYTE_CONSTANT;
		maxMemory = Runtime.getRuntime().maxMemory() / MEGABYTE_CONSTANT;

		System.out.println("freeMemory=\t" + freeMemory);
		System.out
				.println("totalMemory=\t"
						+ totalMemory
						+ " (initial degerleri astiktan sonraki allocate edilen heap miktari)");
		System.out.println("maxMemory=\t" + maxMemory);

		System.out.println("\nVM'in kullamakta oldugu miktar=\t"
				+ (totalMemory - freeMemory));

	}
}
