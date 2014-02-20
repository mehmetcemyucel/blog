package com.cem.aracYikama.strateji;

import com.cem.aracYikama.islem.IcTemizlik;
import com.cem.aracYikama.islem.Temizlik;

public class Ic {
	private static final Ic ic = new Ic();

	private Temizlik[] yapilacakTemizlikListesi = null;

	private Ic() {
		yapilacakTemizlikListesi = new Temizlik[1];
		IcTemizlik icTemizlik = new IcTemizlik();
		yapilacakTemizlikListesi[0] = icTemizlik;
	}

	public static final Temizlik ilkIslemiSoyle() {
		return ic.yapilacakTemizlikListesi[0];
	}
}