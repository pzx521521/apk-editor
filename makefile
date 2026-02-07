apkEditor:
	rm -f release/signed4.apk && rm -f release/notsigned.apk&& rm -f release/notsigned_my.apk && rm -f release/not4.apk && rm -f release/signed.apk
	go run main.go release/app-release.apk release/signed.apk
#	go run zipalign.go  release/not4.apk  release/notsigned.apk
#	/Users/parapeng/Library/Android/sdk/build-tools/30.0.3/zipalign  -v 4 release/not4.apk release/notsigned.apk
	#/Users/parapeng/Library/Android/sdk/build-tools/30.0.3/zipalign -c -v 4 release/notsigned.apk
	#/Users/parapeng/Library/Android/sdk/build-tools/30.0.3/apksigner sign --ks test.keystore --ks-key-alias "key0" --ks-pass pass:123456 --key-pass pass:123456 --out release/signed.apk  release/notsigned.apk
	adb install ./release/signed.apk
install:
	adb install /Users/parapeng/Downloads/app-new.apk
build:
	go build -ldflags "-s -w" -o apkEditor ./cmd
build-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o apkEditor.exe ./cmd
upload:
	aws s3 sync ./release s3://app/html2apk