# Ahenk Go

Ahenk Agent for Liderahenk Central Management System - Go Implementation

## !!! Yardım Gerekli !!!
deb paketi yapmak, kurmak, güncellemek işlemleri arasında kayboldum.

ulaşmaya çalıştığım structure: ahenk-go dosyası /usr/bin/ahenk-go dizininde daemon olarak çalışacak, .service file doğru yere koyulacak, ayrıca -daha tam belirlemedim- farklı bir dizin içinde pluginler bulunacak.

SORU-1 deb paketi paketlerken illa remote bir git reposundan kodu çekmek gerekir mi? (dh-make-golang böyle çalıştığı için soruyorum)

SORU-2 makefile, dpkg-buildpackage, dpkg-deb, gbp buildpackage, dh-make, debhelper komutları arasındaki fark nedir? Neden bunca farklı komut var ve bu kadar karışık...

SORU-3 makefile dpkg-buildpackage kullanmadan önce dosyaları doğru yere yerleştirmek ve gerekli konfigrasyonları yapmak için mi? Eğer öyle ise doğru yer neresi, DESTDIR neresi olmalı ki dpkg-buildpackage .deb dosyasını oluştururken kuracağı paketleri doğru anlasın?

SORU-4 systemd dosyalarını kurmak için makefile veya dpkg-buildpackage kullanarak yapabileceğim bir şey var mı? Yoksa git.pardus.org.tr > eta-pulse-config deposundaki örnek programdaki gibi `rules` dosyası içinden değişiklik yapmalıyım? Eğer evet ise yeni sürümlerde kaldırıldığı yazıyor, eta-pulse-config'deki gibi compat 9 sürümünü kullanmaya çalışınca dpkg-buildpacke şöyle error veriyor: `dh: warning: Compatibility levels before 10 are deprecated (level 9 in use)` ne yapmamı önerirsin(iz)?

YAPMAK İSTEDİĞİM FARKLI AŞAMALAR:
1- debian paketleme için gerekli dosyaları (control,rules...) oluşturma
  - manuel
  - dh-make?
2- .deb paketleme. öyle ki oluşan bu deb paketi çalıştırıldığında cmd/ahenk-go içinde buildlenmiş dosyaları /usr/bin/ahenk-go'ya, ahenk-go.service dosyasını /lib/systemd/system altına .... koysun.
  - dpkg-buildpackage?
  - gbp buildpackage?
2.1- Ayrıca source code paketi oluşturma ve bu paketten tekrar .deb oluşturabilme

3- .deb paketi kurma ve paketi sistemden kaldırma işlemleri. 
  - deb zaten kendi mi yapıyor bunu?

4- koduma eklemeler yaptım ve tekrardan .deb paketi paketlemek istiyorum, güncellemeleri hangi komut ile veya hangi şekilde yapmalıyım?

diğer branchların son durumu: 
- main: en son 1 hafta önce denediğim dosyalar ve scripts/docker-build.sh son hâli (docker üzerinde denediğim build işleme, dh-make-golang ile)
- debian/sid: main'deki dosyaları buildlediğim, dh-make-golang [guide](https://people.debian.org/~stapelberg/2015/07/27/dh-make-golang.html)'i sonucu oluşturduğum branch

özellikle bu guide üzerinden ve scripts/docker-build.sh üzerinden çalıştığım için bir hatam varsa buralarda olma ihtimali büyük
---


### Useful Links
| Explanation | Link |
| ----------- | ---- |
| Ahenk-go documentation | [Liderahenk/ahenk-docs](https://git.aliberksandikci.com.tr/Liderahenk/ahenk-docs/) |
| Current Python Implementation of Ahenk | [Pardus-LiderAhenk/ahenk](https://github.com/Pardus-LiderAhenk/ahenk/) |


