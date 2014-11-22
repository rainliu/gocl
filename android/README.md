GOCL on Android
======

# Install Ant.
0. install  http://archive.apache.org/dist/ant/binaries/apache-ant-1.9.2-bin.tar.gz to /usr/local
2. "export ANT_HOME=/usr/local/apache-ant"

# Install Android SDK.
0. install http://dl.google.com/android/android-sdk_r23.0.2-darwin.tgz to /usr/local
1. "export ANDROID_HOME=/usr/local/android-sdk"
2. "$ANDROID_HOME/tools/android update sdk --no-ui --all --filter build-tools-19.1.0" 
3. "$ANDROID_HOME/tools/android update sdk --no-ui --all --filter platform-tools" 
4. "$ANDROID_HOME/tools/android update sdk --no-ui --all --filter android-19"

# Install Android NDK.
0. install http://dl.google.com/android/ndk/android-ndk-r9d-darwin-x86_64.tar.bz2 to /usr/local
1. "export NDK_ROOT=/usr/local/android-ndk"
2. "copy gocl/android/include/CL/*.* into $NDK_ROOT/platform/android-19/arch-arm/usr/include"
3. "copy gocl/android/lib/*.* into $NDK_ROOT/platform/android-19/arch-arm/usr/lib"
4. "$NDK_ROOT/build/tools/make-standalone-toolchain.sh --platform=android-19 --install-dir=$NDK_ROOT --system=darwin-x86_64"

# Update PATH for the above.
1. "export PATH=$PATH:$ANDROID_HOME/tools"
2. "export PATH=$PATH:$ANDROID_HOME/platform-tools"
3. "export PATH=$PATH:$NDK_ROOT"
4. "export PATH=$PATH:$ANT_HOME/bin"

# Install Go from source code.
1. "cd $GOROOT/src"
2. "CC=clang ./make.bash"
3. "CC_FOR_TARGET=$NDK_ROOT/bin/arm-linux-androideabi-gcc CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 ./make.bash"
4. "go env | grep CC="


# Install Go.Mobile
1. go get -d -t golang.org/x/mobile/...
2. "cd $GOPATH/src/golang.org/x/mobile/example/basic"
3. build android: GO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 go build -tags="cl11" -ldflags="-shared" -o jni/armeabi/libbasic.so
4. build darwin:  CC=clang go build -tags 'cl11' gocl/cl