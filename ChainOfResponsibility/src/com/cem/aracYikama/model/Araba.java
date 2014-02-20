package com.cem.aracYikama.model;

import com.cem.aracYikama.islem.Temizlik;
import com.cem.aracYikama.strateji.IcDis;

public class Araba extends Arac {

	public Araba(String marka, String model) {
		super(marka, model);
	}

	@Override
	public Temizlik ilkTemizlikIsleminiVer() {
		return IcDis.ilkIslemiSoyle();
	}
}