package main

func init() {
	colors = []int{
		0x0000ff, 0x0001ff, 0x0002ff, 0x0003ff, 0x0004ff, 0x0005ff, 0x0006ff, 0x0007ff,
		0x0008ff, 0x0009ff, 0x000aff, 0x000bff, 0x000cff, 0x000dff, 0x000eff, 0x000fff,
		0x0010ff, 0x0011ff, 0x0012ff, 0x0013ff, 0x0014ff, 0x0015ff, 0x0016ff, 0x0017ff,
		0x0018ff, 0x0019ff, 0x001aff, 0x001bff, 0x001cff, 0x001dff, 0x001eff, 0x001fff,
		0x0020ff, 0x0021ff, 0x0022ff, 0x0023ff, 0x0024ff, 0x0025ff, 0x0026ff, 0x0027ff,
		0x0028ff, 0x0029ff, 0x002aff, 0x002bff, 0x002cff, 0x002dff, 0x002eff, 0x002fff,
		0x0030ff, 0x0031ff, 0x0032ff, 0x0033ff, 0x0034ff, 0x0035ff, 0x0036ff, 0x0037ff,
		0x0038ff, 0x0039ff, 0x003aff, 0x003bff, 0x003cff, 0x003dff, 0x003eff, 0x003fff,
		0x0040ff, 0x0041ff, 0x0042ff, 0x0043ff, 0x0044ff, 0x0045ff, 0x0046ff, 0x0047ff,
		0x0048ff, 0x0049ff, 0x004aff, 0x004bff, 0x004cff, 0x004dff, 0x004eff, 0x004fff,
		0x0050ff, 0x0051ff, 0x0052ff, 0x0053ff, 0x0054ff, 0x0055ff, 0x0056ff, 0x0057ff,
		0x0058ff, 0x0059ff, 0x005aff, 0x005bff, 0x005cff, 0x005dff, 0x005eff, 0x005fff,
		0x0060ff, 0x0061ff, 0x0062ff, 0x0063ff, 0x0064ff, 0x0065ff, 0x0066ff, 0x0067ff,
		0x0068ff, 0x0069ff, 0x006aff, 0x006bff, 0x006cff, 0x006dff, 0x006eff, 0x006fff,
		0x0070ff, 0x0071ff, 0x0072ff, 0x0073ff, 0x0074ff, 0x0075ff, 0x0076ff, 0x0077ff,
		0x0078ff, 0x0079ff, 0x007aff, 0x007bff, 0x007cff, 0x007dff, 0x007eff, 0x007fff,
		0x0080ff, 0x0081ff, 0x0082ff, 0x0083ff, 0x0084ff, 0x0085ff, 0x0086ff, 0x0087ff,
		0x0088ff, 0x0089ff, 0x008aff, 0x008bff, 0x008cff, 0x008dff, 0x008eff, 0x008fff,
		0x0090ff, 0x0091ff, 0x0092ff, 0x0093ff, 0x0094ff, 0x0095ff, 0x0096ff, 0x0097ff,
		0x0098ff, 0x0099ff, 0x009aff, 0x009bff, 0x009cff, 0x009dff, 0x009eff, 0x009fff,
		0x00a0ff, 0x00a1ff, 0x00a2ff, 0x00a3ff, 0x00a4ff, 0x00a5ff, 0x00a6ff, 0x00a7ff,
		0x00a8ff, 0x00a9ff, 0x00aaff, 0x00aaff, 0x00abff, 0x00acff, 0x00adff, 0x00aeff,
		0x00afff, 0x00b0ff, 0x00b1ff, 0x00b2ff, 0x00b3ff, 0x00b4ff, 0x00b5ff, 0x00b6ff,
		0x00b7ff, 0x00b8ff, 0x00b9ff, 0x00baff, 0x00bbff, 0x00bcff, 0x00bdff, 0x00beff,
		0x00bfff, 0x00c0ff, 0x00c1ff, 0x00c2ff, 0x00c3ff, 0x00c4ff, 0x00c5ff, 0x00c6ff,
		0x00c7ff, 0x00c8ff, 0x00c9ff, 0x00caff, 0x00cbff, 0x00ccff, 0x00cdff, 0x00ceff,
		0x00cfff, 0x00d0ff, 0x00d1ff, 0x00d2ff, 0x00d3ff, 0x00d4ff, 0x00d5ff, 0x00d6ff,
		0x00d7ff, 0x00d8ff, 0x00d9ff, 0x00daff, 0x00dbff, 0x00dcff, 0x00ddff, 0x00deff,
		0x00dfff, 0x00e0ff, 0x00e1ff, 0x00e2ff, 0x00e3ff, 0x00e4ff, 0x00e5ff, 0x00e6ff,
		0x00e7ff, 0x00e8ff, 0x00e9ff, 0x00eaff, 0x00ebff, 0x00ecff, 0x00edff, 0x00eeff,
		0x00efff, 0x00f0ff, 0x00f1ff, 0x00f2ff, 0x00f3ff, 0x00f4ff, 0x00f5ff, 0x00f6ff,
		0x00f7ff, 0x00f8ff, 0x00f9ff, 0x00faff, 0x00fbff, 0x00fcff, 0x00fdff, 0x00feff,
		0x00ffff, 0x00fffe, 0x00fffd, 0x00fffc, 0x00fffb, 0x00fffa, 0x00fff9, 0x00fff8,
		0x00fff7, 0x00fff6, 0x00fff5, 0x00fff4, 0x00fff3, 0x00fff2, 0x00fff1, 0x00fff0,
		0x00ffef, 0x00ffee, 0x00ffed, 0x00ffec, 0x00ffeb, 0x00ffea, 0x00ffe9, 0x00ffe8,
		0x00ffe7, 0x00ffe6, 0x00ffe5, 0x00ffe4, 0x00ffe3, 0x00ffe2, 0x00ffe1, 0x00ffe0,
		0x00ffdf, 0x00ffde, 0x00ffdd, 0x00ffdc, 0x00ffdb, 0x00ffda, 0x00ffd9, 0x00ffd8,
		0x00ffd7, 0x00ffd6, 0x00ffd5, 0x00ffd4, 0x00ffd3, 0x00ffd2, 0x00ffd1, 0x00ffd0,
		0x00ffcf, 0x00ffce, 0x00ffcd, 0x00ffcc, 0x00ffcb, 0x00ffca, 0x00ffc9, 0x00ffc8,
		0x00ffc7, 0x00ffc6, 0x00ffc5, 0x00ffc4, 0x00ffc3, 0x00ffc2, 0x00ffc1, 0x00ffc0,
		0x00ffbf, 0x00ffbe, 0x00ffbd, 0x00ffbc, 0x00ffbb, 0x00ffba, 0x00ffb9, 0x00ffb8,
		0x00ffb7, 0x00ffb6, 0x00ffb5, 0x00ffb4, 0x00ffb3, 0x00ffb2, 0x00ffb1, 0x00ffb0,
		0x00ffaf, 0x00ffae, 0x00ffad, 0x00ffac, 0x00ffab, 0x00ffaa, 0x00ffa9, 0x00ffa8,
		0x00ffa7, 0x00ffa6, 0x00ffa5, 0x00ffa4, 0x00ffa3, 0x00ffa2, 0x00ffa1, 0x00ffa0,
		0x00ff9f, 0x00ff9e, 0x00ff9d, 0x00ff9c, 0x00ff9b, 0x00ff9a, 0x00ff99, 0x00ff98,
		0x00ff97, 0x00ff96, 0x00ff95, 0x00ff94, 0x00ff93, 0x00ff92, 0x00ff91, 0x00ff90,
		0x00ff8f, 0x00ff8e, 0x00ff8d, 0x00ff8c, 0x00ff8b, 0x00ff8a, 0x00ff89, 0x00ff88,
		0x00ff87, 0x00ff86, 0x00ff85, 0x00ff84, 0x00ff83, 0x00ff82, 0x00ff81, 0x00ff80,
		0x00ff7f, 0x00ff7e, 0x00ff7d, 0x00ff7c, 0x00ff7b, 0x00ff7a, 0x00ff79, 0x00ff78,
		0x00ff77, 0x00ff76, 0x00ff75, 0x00ff74, 0x00ff73, 0x00ff72, 0x00ff71, 0x00ff70,
		0x00ff6f, 0x00ff6e, 0x00ff6d, 0x00ff6c, 0x00ff6b, 0x00ff6a, 0x00ff69, 0x00ff68,
		0x00ff67, 0x00ff66, 0x00ff65, 0x00ff64, 0x00ff63, 0x00ff62, 0x00ff61, 0x00ff60,
		0x00ff5f, 0x00ff5e, 0x00ff5d, 0x00ff5c, 0x00ff5b, 0x00ff5a, 0x00ff59, 0x00ff58,
		0x00ff57, 0x00ff56, 0x00ff55, 0x00ff54, 0x00ff53, 0x00ff52, 0x00ff51, 0x00ff50,
		0x00ff4f, 0x00ff4e, 0x00ff4d, 0x00ff4c, 0x00ff4b, 0x00ff4a, 0x00ff49, 0x00ff48,
		0x00ff47, 0x00ff46, 0x00ff45, 0x00ff44, 0x00ff43, 0x00ff42, 0x00ff41, 0x00ff40,
		0x00ff3f, 0x00ff3e, 0x00ff3d, 0x00ff3c, 0x00ff3b, 0x00ff3a, 0x00ff39, 0x00ff38,
		0x00ff37, 0x00ff36, 0x00ff35, 0x00ff34, 0x00ff33, 0x00ff32, 0x00ff31, 0x00ff30,
		0x00ff2f, 0x00ff2e, 0x00ff2d, 0x00ff2c, 0x00ff2b, 0x00ff2a, 0x00ff29, 0x00ff28,
		0x00ff27, 0x00ff26, 0x00ff25, 0x00ff24, 0x00ff23, 0x00ff22, 0x00ff21, 0x00ff20,
		0x00ff1f, 0x00ff1e, 0x00ff1d, 0x00ff1c, 0x00ff1b, 0x00ff1a, 0x00ff19, 0x00ff18,
		0x00ff17, 0x00ff16, 0x00ff15, 0x00ff14, 0x00ff13, 0x00ff12, 0x00ff11, 0x00ff10,
		0x00ff0f, 0x00ff0e, 0x00ff0d, 0x00ff0c, 0x00ff0b, 0x00ff0a, 0x00ff09, 0x00ff08,
		0x00ff07, 0x00ff06, 0x00ff05, 0x00ff04, 0x00ff03, 0x00ff02, 0x00ff01, 0x00ff00,
		0x00ff00, 0x01ff00, 0x02ff00, 0x03ff00, 0x04ff00, 0x05ff00, 0x06ff00, 0x07ff00,
		0x08ff00, 0x09ff00, 0x0aff00, 0x0bff00, 0x0cff00, 0x0dff00, 0x0eff00, 0x0fff00,
		0x10ff00, 0x11ff00, 0x12ff00, 0x13ff00, 0x14ff00, 0x15ff00, 0x16ff00, 0x17ff00,
		0x18ff00, 0x19ff00, 0x1aff00, 0x1bff00, 0x1cff00, 0x1dff00, 0x1eff00, 0x1fff00,
		0x20ff00, 0x21ff00, 0x22ff00, 0x23ff00, 0x24ff00, 0x25ff00, 0x26ff00, 0x27ff00,
		0x28ff00, 0x29ff00, 0x2aff00, 0x2bff00, 0x2cff00, 0x2dff00, 0x2eff00, 0x2fff00,
		0x30ff00, 0x31ff00, 0x32ff00, 0x33ff00, 0x34ff00, 0x35ff00, 0x36ff00, 0x37ff00,
		0x38ff00, 0x39ff00, 0x3aff00, 0x3bff00, 0x3cff00, 0x3dff00, 0x3eff00, 0x3fff00,
		0x40ff00, 0x41ff00, 0x42ff00, 0x43ff00, 0x44ff00, 0x45ff00, 0x46ff00, 0x47ff00,
		0x48ff00, 0x49ff00, 0x4aff00, 0x4bff00, 0x4cff00, 0x4dff00, 0x4eff00, 0x4fff00,
		0x50ff00, 0x51ff00, 0x52ff00, 0x53ff00, 0x54ff00, 0x55ff00, 0x56ff00, 0x57ff00,
		0x58ff00, 0x59ff00, 0x5aff00, 0x5bff00, 0x5cff00, 0x5dff00, 0x5eff00, 0x5fff00,
		0x60ff00, 0x61ff00, 0x62ff00, 0x63ff00, 0x64ff00, 0x65ff00, 0x66ff00, 0x67ff00,
		0x68ff00, 0x69ff00, 0x6aff00, 0x6bff00, 0x6cff00, 0x6dff00, 0x6eff00, 0x6fff00,
		0x70ff00, 0x71ff00, 0x72ff00, 0x73ff00, 0x74ff00, 0x75ff00, 0x76ff00, 0x77ff00,
		0x78ff00, 0x79ff00, 0x7aff00, 0x7bff00, 0x7cff00, 0x7dff00, 0x7eff00, 0x7fff00,
		0x80ff00, 0x81ff00, 0x82ff00, 0x83ff00, 0x84ff00, 0x85ff00, 0x86ff00, 0x87ff00,
		0x88ff00, 0x89ff00, 0x8aff00, 0x8bff00, 0x8cff00, 0x8dff00, 0x8eff00, 0x8fff00,
		0x90ff00, 0x91ff00, 0x92ff00, 0x93ff00, 0x94ff00, 0x95ff00, 0x96ff00, 0x97ff00,
		0x98ff00, 0x99ff00, 0x9aff00, 0x9bff00, 0x9cff00, 0x9dff00, 0x9eff00, 0x9fff00,
		0xa0ff00, 0xa1ff00, 0xa2ff00, 0xa3ff00, 0xa4ff00, 0xa5ff00, 0xa6ff00, 0xa7ff00,
		0xa8ff00, 0xa9ff00, 0xaaff00, 0xabff00, 0xacff00, 0xadff00, 0xaeff00, 0xafff00,
		0xb0ff00, 0xb1ff00, 0xb2ff00, 0xb3ff00, 0xb4ff00, 0xb5ff00, 0xb6ff00, 0xb7ff00,
		0xb8ff00, 0xb9ff00, 0xbaff00, 0xbbff00, 0xbcff00, 0xbdff00, 0xbeff00, 0xbfff00,
		0xc0ff00, 0xc1ff00, 0xc2ff00, 0xc3ff00, 0xc4ff00, 0xc5ff00, 0xc6ff00, 0xc7ff00,
		0xc8ff00, 0xc9ff00, 0xcaff00, 0xcbff00, 0xccff00, 0xcdff00, 0xceff00, 0xcfff00,
		0xd0ff00, 0xd1ff00, 0xd2ff00, 0xd3ff00, 0xd4ff00, 0xd5ff00, 0xd6ff00, 0xd7ff00,
		0xd8ff00, 0xd9ff00, 0xdaff00, 0xdbff00, 0xdcff00, 0xddff00, 0xdeff00, 0xdfff00,
		0xe0ff00, 0xe1ff00, 0xe2ff00, 0xe3ff00, 0xe4ff00, 0xe5ff00, 0xe6ff00, 0xe7ff00,
		0xe8ff00, 0xe9ff00, 0xeaff00, 0xebff00, 0xecff00, 0xedff00, 0xeeff00, 0xefff00,
		0xf0ff00, 0xf1ff00, 0xf2ff00, 0xf3ff00, 0xf4ff00, 0xf5ff00, 0xf6ff00, 0xf7ff00,
		0xf8ff00, 0xf9ff00, 0xfaff00, 0xfbff00, 0xfcff00, 0xfdff00, 0xfeff00, 0xffff00,
		0xfffe00, 0xfffd00, 0xfffc00, 0xfffb00, 0xfffa00, 0xfff900, 0xfff800, 0xfff700,
		0xfff600, 0xfff500, 0xfff400, 0xfff300, 0xfff200, 0xfff100, 0xfff000, 0xffef00,
		0xffee00, 0xffed00, 0xffec00, 0xffeb00, 0xffea00, 0xffe900, 0xffe800, 0xffe700,
		0xffe600, 0xffe500, 0xffe400, 0xffe300, 0xffe200, 0xffe100, 0xffe000, 0xffdf00,
		0xffde00, 0xffdd00, 0xffdc00, 0xffdb00, 0xffda00, 0xffd900, 0xffd800, 0xffd700,
		0xffd600, 0xffd500, 0xffd400, 0xffd300, 0xffd200, 0xffd100, 0xffd000, 0xffcf00,
		0xffce00, 0xffcd00, 0xffcc00, 0xffcb00, 0xffca00, 0xffc900, 0xffc800, 0xffc700,
		0xffc600, 0xffc500, 0xffc400, 0xffc300, 0xffc200, 0xffc100, 0xffc000, 0xffbf00,
		0xffbe00, 0xffbd00, 0xffbc00, 0xffbb00, 0xffba00, 0xffb900, 0xffb800, 0xffb700,
		0xffb600, 0xffb500, 0xffb400, 0xffb300, 0xffb200, 0xffb100, 0xffb000, 0xffaf00,
		0xffae00, 0xffad00, 0xffac00, 0xffab00, 0xffaa00, 0xffaa00, 0xffa900, 0xffa800,
		0xffa700, 0xffa600, 0xffa500, 0xffa400, 0xffa300, 0xffa200, 0xffa100, 0xffa000,
		0xff9f00, 0xff9e00, 0xff9d00, 0xff9c00, 0xff9b00, 0xff9a00, 0xff9900, 0xff9800,
		0xff9700, 0xff9600, 0xff9500, 0xff9400, 0xff9300, 0xff9200, 0xff9100, 0xff9000,
		0xff8f00, 0xff8e00, 0xff8d00, 0xff8c00, 0xff8b00, 0xff8a00, 0xff8900, 0xff8800,
		0xff8700, 0xff8600, 0xff8500, 0xff8400, 0xff8300, 0xff8200, 0xff8100, 0xff8000,
		0xff7f00, 0xff7e00, 0xff7d00, 0xff7c00, 0xff7b00, 0xff7a00, 0xff7900, 0xff7800,
		0xff7700, 0xff7600, 0xff7500, 0xff7400, 0xff7300, 0xff7200, 0xff7100, 0xff7000,
		0xff6f00, 0xff6e00, 0xff6d00, 0xff6c00, 0xff6b00, 0xff6a00, 0xff6900, 0xff6800,
		0xff6700, 0xff6600, 0xff6500, 0xff6400, 0xff6300, 0xff6200, 0xff6100, 0xff6000,
		0xff5f00, 0xff5e00, 0xff5d00, 0xff5c00, 0xff5b00, 0xff5a00, 0xff5900, 0xff5800,
		0xff5700, 0xff5600, 0xff5500, 0xff5400, 0xff5300, 0xff5200, 0xff5100, 0xff5000,
		0xff4f00, 0xff4e00, 0xff4d00, 0xff4c00, 0xff4b00, 0xff4a00, 0xff4900, 0xff4800,
		0xff4700, 0xff4600, 0xff4500, 0xff4400, 0xff4300, 0xff4200, 0xff4100, 0xff4000,
		0xff3f00, 0xff3e00, 0xff3d00, 0xff3c00, 0xff3b00, 0xff3a00, 0xff3900, 0xff3800,
		0xff3700, 0xff3600, 0xff3500, 0xff3400, 0xff3300, 0xff3200, 0xff3100, 0xff3000,
		0xff2f00, 0xff2e00, 0xff2d00, 0xff2c00, 0xff2b00, 0xff2a00, 0xff2900, 0xff2800,
		0xff2700, 0xff2600, 0xff2500, 0xff2400, 0xff2300, 0xff2200, 0xff2100, 0xff2000,
		0xff1f00, 0xff1e00, 0xff1d00, 0xff1c00, 0xff1b00, 0xff1a00, 0xff1900, 0xff1800,
		0xff1700, 0xff1600, 0xff1500, 0xff1400, 0xff1300, 0xff1200, 0xff1100, 0xff1000,
		0xff0f00, 0xff0e00, 0xff0d00, 0xff0c00, 0xff0b00, 0xff0a00, 0xff0900, 0xff0800,
		0xff0700, 0xff0600, 0xff0500, 0xff0400, 0xff0300, 0xff0200, 0xff0100, 0xff0000,
	}
}