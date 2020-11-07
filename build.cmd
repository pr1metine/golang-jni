@ECHO OFF

REM Adjust these lines to your environment
SET JAVA_HOME=C:\Java\jdk-11.0.7
SET MINGW_HOME=C:\MinGW\mingw32

REM Do not change the following lines
SET PATH=%MINGW_HOME%\bin;%JAVA_HOME%\bin;%PATH%
SET CFLAGS=-I. -I%JAVA_HOME%/include -I%JAVA_HOME%/include/win32
SET LDFLAGS=-L%JAVA_HOME%/lib -ljvm
SET CGO_CFLAGS=%CFLAGS%
SET CGO_LDFLAGS=%LDFLAGS%

rd  /q /s out 2>NUL

mkdir out
IF %ERRORLEVEL% NEQ 0 EXIT /b %ERRORLEVEL%

javac -d out VierBitAddierer.java
IF %ERRORLEVEL% NEQ 0 EXIT /b %ERRORLEVEL%

go build -buildmode=c-shared -o out\VierBitAddierer.dll .
IF %ERRORLEVEL% NEQ 0 EXIT /b %ERRORLEVEL%

java -Xcheck:jni -Djava.library.path=out -cp out VierBitAddierer
IF %ERRORLEVEL% NEQ 0 EXIT /b %ERRORLEVEL%
