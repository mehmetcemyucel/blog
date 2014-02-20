package com.cem.aracYikama.model;

import com.cem.aracYikama.islem.Temizlik;
import com.cem.aracYikama.strateji.Ic;

public class Otobus extends Arac {

	public Otobus(String marka, String model) {
		super(marka, model);
	}

	@Override
	public Temizlik ilkTemizlikIsleminiVer() {
		return Ic.ilkIslemiSoyle();
	}
}