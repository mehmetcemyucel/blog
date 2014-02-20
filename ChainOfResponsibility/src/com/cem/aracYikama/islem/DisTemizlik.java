package com.cem.aracYikama.islem;

import com.cem.aracYikama.model.Arac;

public class DisTemizlik extends Temizlik {

	@Override
	public void temizle(Arac arac) {
		System.out.println(arac.getMarka() + " marka " + arac.getModel()
				+ " model aracin dis temizligi tamamlandi.");
	}

}