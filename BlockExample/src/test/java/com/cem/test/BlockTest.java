package com.cem.test;

import static org.junit.Assert.*;

import java.util.ArrayList;

import org.apache.commons.lang.StringUtils;
import org.junit.Test;

import com.cem.Block;

public class BlockTest {

	@Test
	public void reliabilityNumberTest() {
		ArrayList<String> transactionList = new ArrayList<String>();
		transactionList.add("Transaction1");
		transactionList.add("Transaction2");
		Block block = new Block(1, "00000000000000000000", "01012018223000",transactionList );
		assertTrue(block.getHash().startsWith(StringUtils.leftPad("", 6, "0")));
	}

}
