GOCL on Android
======

# Install Ant.
install  http://archive.apache.org/dist/ant/binaries/apache-ant-1.9.2-bin.tar.gz to /usr/local
1. "export ANT_HOME=/usr/local/apache-ant"

# Install Android SDK.
install http://dl.google.com/android/android-sdk_r23.0.2-darwin.tgz to /usr/local
1. "export ANDROID_HOME=/usr/local/android-sdk"
2. "$ANDROID_HOME/tools/android update sdk --no-ui --all --filter build-tools-19.1.0" 
3. "$ANDROID_HOME/tools/android update sdk --no-ui --all --filter platform-tools" 
4. "$ANDROID_HOME/tools/android update sdk --no-ui --all --filter android-19"

# Install Android NDK.
install http://dl.google.com/android/ndk/android-ndk-r9d-darwin-x86_64.tar.bz2 to /usr/local
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
5. "build android: CGO_ENABLED=1 GOOS=android GOARCH=arm GOARM=7 go build -ldflags="-shared" -o jni/armeabi/libbasic.so"
6. "build darwin:  CC=clang go build -tags 'cl12' gocl/cl"

# Copy the local version of go.mobile to GOPATH.
ADD . /gopath/src/golang.org/x/mobile

# Install dependencies. This will not overwrite the local copy.
RUN go get -d -t golang.org/x/mobile/...

WORKDIR /gopath/src/golang.org/x/mobile