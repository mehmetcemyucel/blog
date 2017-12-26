package com.cem;

import java.io.UnsupportedEncodingException;
import java.security.InvalidAlgorithmParameterException;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;

import javax.crypto.BadPaddingException;
import javax.crypto.Cipher;
import javax.crypto.IllegalBlockSizeException;
import javax.crypto.NoSuchPaddingException;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;

public class DESEncryptionDecryption {

	public static byte[] AES256Encryption2(String key, String initVector, String value)
			throws IllegalBlockSizeException, BadPaddingException, NoSuchAlgorithmException, NoSuchPaddingException,
			InvalidKeyException, InvalidAlgorithmParameterException, UnsupportedEncodingException {
		IvParameterSpec iv = new IvParameterSpec(initVector.getBytes("UTF-8"));
		SecretKeySpec skeySpec = new SecretKeySpec(key.getBytes("UTF-8"), "DES");

		Cipher cipher = Cipher.getInstance("DES/CBC/PKCS5PADDING");
		cipher.init(Cipher.ENCRYPT_MODE, skeySpec, iv);
		byte[] encrypted = cipher.doFinal(value.getBytes());
		return encrypted;
	}

	public static String AES256Decryption2(String key, String initVector, byte[] encrypted)
			throws UnsupportedEncodingException, InvalidKeyException, InvalidAlgorithmParameterException,
			NoSuchAlgorithmException, NoSuchPaddingException, IllegalBlockSizeException, BadPaddingException {
		IvParameterSpec iv = new IvParameterSpec(initVector.getBytes("UTF-8"));
		SecretKeySpec skeySpec = new SecretKeySpec(key.getBytes("UTF-8"), "DES");
		Cipher cipher = Cipher.getInstance("DES/CBC/PKCS5PADDING");
		cipher.init(Cipher.DECRYPT_MODE, skeySpec, iv);
		byte[] orginal = cipher.doFinal(encrypted);
		return new String(orginal);
	}

	public static void main(String[] args) {
		try {
			String key = "Bar12345";
			String initVector = "RandomIn";
			String data = "Merhaba Dünya ışçğüö";
			System.out.println(AES256Decryption2(key, initVector, AES256Encryption2(key, initVector, data)));
		} catch (Exception e) {
			e.printStackTrace();
		}
	}

}
