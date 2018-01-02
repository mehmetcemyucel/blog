package com.cem;

import java.util.Collections;
import java.util.List;

import org.apache.commons.lang.StringUtils;

public class Block {

	private final long index;
	private final String previousHash;
	private final String hash;
	private final String timestamp;
	private final java.util.List<String> transactionList;
	private long nonce;

	private static final int RELIABILITY_NUMBER = 6;

	private String generateHash(long index, String previousHash, String timestamp, List<String> transactionList,
			long nonce) {
		return org.apache.commons.codec.digest.DigestUtils
				.sha256Hex(index + previousHash + timestamp + transactionList + nonce);
	}

	public Block(long index, String previousHash, String timestamp, List<String> transactionList) {
		this.index = index;
		this.previousHash = previousHash;
		this.hash = generateValidHash();
		this.timestamp = timestamp;
		this.transactionList = Collections.unmodifiableList(transactionList);
	}

	private String generateValidHash() {
		long nonce = 0;
		String generatedHash = "";
		do {
			generatedHash = generateHash(index, previousHash, timestamp, transactionList, ++nonce);
		} while (!generatedHash.startsWith(StringUtils.leftPad("", RELIABILITY_NUMBER, "0")));
		setNonce(nonce);
		return generatedHash;
	}

	public long getNonce() {
		return nonce;
	}

	private void setNonce(long nonce) {
		this.nonce = nonce;
	}

	public long getIndex() {
		return index;
	}

	public String getPreviousHash() {
		return previousHash;
	}

	public String getHash() {
		return hash;
	}

	public String getTimestamp() {
		return timestamp;
	}

	public java.util.List<String> getTransactionList() {
		return transactionList;
	}

}
