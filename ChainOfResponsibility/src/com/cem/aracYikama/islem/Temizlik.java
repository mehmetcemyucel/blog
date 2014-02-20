package com.cem.aracYikama.islem;

import com.cem.aracYikama.model.Arac;

public abstract class Temizlik implements ITemizlik {

	protected Temizlik siradakiIslem;

	public void siradakiIslemAta(Temizlik islem) {
		this.siradakiIslem = islem;
	}

	public void basla(Arac arac) {
		this.temizle(arac);
		if (siradakiIslem != null) {
			siradakiIslem.basla(arac);
		}
	}

}