# expl0rer

## 設定

`expl0rer.exe` と同じフォルダに `expl0rer.toml` を置くと、起動するツールのパスを変更できます。

キーには、既定の `exe` パスの最後のファイル名から拡張子を除いた値を指定します。

例: `C:\Program Files\vim\gvim.exe` のキーは `gvim` です。

```toml
[tools]
gvim = 'C:\Program Files\vim\gvim.exe'
```
