1
	パッケージ
	main関数
	スライス
	変数のゼロ値
	省略変数宣言 (:=)
	for
		for 相当: `for i := 0; i < n; i++ {}`
		while 相当: `for cond {}`
		for-each 相当: `for range iterable {}`
	Printf フォーマット
		だいたいCと一緒
		安直に `%v` (いい感じのフォーマットで出力) を使うことが多い
			特にエラー出力
		
	strings.Join (StringJoiner相当)
	mapのvalue側の初期値はそのゼロ値
	リファレンス(参照)
