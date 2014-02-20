package com.cem.aracYikama.model;

import com.cem.aracYikama.islem.Temizlik;
import com.cem.aracYikama.strateji.Dis;

public class Motosiklet extends Arac {

	public Motosiklet(String marka, String model) {
		super(marka, model);
	}

	@Override
	public Temizlik ilkTemizlikIsleminiVer() {
		return Dis.ilkIslemiSoyle();
	}
}
