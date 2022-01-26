# grep-archive
Grep archive search in any files on the filesystem, in archive and even inner archive.

Supported archive format are : 
- Tar Formats
  - tar.gz 
  - tar 
  - tgz 
  - tar.bz2 
- Zip Format
  - jar
  - ear 
  - war 


# Usage

## List all files recursively in archives:

In the exmple below find all files in an archive
```
─$ /mnt/hgfs/VM_Shared/grep-archive -scan .    
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/LICENSE.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/NOTICE.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/RELEASE-NOTES.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/INSTALL.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/Makedefs.in Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/Makefile.in Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/configure.in Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/man/README.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/man/jsvc.1.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/.indent.pro Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/Makefile.in Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/arguments.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/arguments.h Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/debug.c Found    
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/debug.h Found    
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/dso-dlfcn.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/unix/native/dso-dyld.c Found 
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/rprocess.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/mclib.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/mclib.h Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/private.h Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/registry.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/rprocess.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/service.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/src/utils.c Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/xdocs/index.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/apps/prunmgr/prunmgr.rc Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/apps/prunsrv/prunsrv.rc Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/resources/commons.bmp Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/resources/procrunr.ico Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/resources/procruns.ico Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/resources/procrunw.ico Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon-native.tar.gz:commons-daemon-1.2.4-native-src/windows/resources/susers.bmp Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/MANIFEST.MF Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/maven/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/maven/commons-daemon/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/maven/commons-daemon/commons-daemon/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/LICENSE.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/NOTICE.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/Daemon.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/DaemonContext.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/DaemonController.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/DaemonInitException.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/DaemonListener.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/DaemonPermission.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/DaemonUserSignal.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonConfiguration.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonLoader$1.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonLoader$Context.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonLoader$Controller.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonLoader.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonWrapper$Invoker.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:org/apache/commons/daemon/support/DaemonWrapper.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/maven/commons-daemon/commons-daemon/pom.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/commons-daemon.jar:META-INF/maven/commons-daemon/commons-daemon/pom.properties Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/configtest.bat Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/digest.bat Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/makebase.bat Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/setclasspath.bat Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/shutdown.bat Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/startup.bat Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:META-INF/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:META-INF/MANIFEST.MF Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:META-INF/LICENSE Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:META-INF/NOTICE Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:module-info.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/AsyncFileHandler$LogEntry.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/AsyncFileHandler$LoggerThread.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/AsyncFileHandler.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ClassLoaderLogManager$1.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ClassLoaderLogManager$ClassLoaderLogInfo.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ClassLoaderLogManager$Cleaner.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ClassLoaderLogManager$LogNode.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ClassLoaderLogManager$RootLogger.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/ClassLoaderLogManager.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/DateFormatCache$1.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/DateFormatCache$Cache.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/DateFormatCache.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/FileHandler$1.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/FileHandler.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/JdkLoggerFormatter.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/OneLineFormatter$IndentingPrintWriter.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/OneLineFormatter$MillisHandling.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/OneLineFormatter$ThreadNameCache.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/OneLineFormatter.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/VerbatimFormatter.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/WebappProperties.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/logging/ Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/logging/DirectJDKLog.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/logging/Log.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/logging/LogConfigurationException.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-juli.jar:org/apache/juli/logging/LogFactory.class Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/CMakeLists.txt Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/mkcerts Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/org/apache/tomcat/jni/Echo.java Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/org/apache/tomcat/jni/SSLServer.java Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/org/apache/tomcat/jni/SSL.properties Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/org/apache/tomcat/jni/LocalServer.java Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/org/apache/tomcat/jni/Echo.properties Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/examples/org/apache/tomcat/jni/Local.properties Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/.gitignore Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/build.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/index.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/miscellaneous/changelog.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/miscellaneous/project.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/style.xsl Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2016.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2009.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2018.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2014.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2011.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2021.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2012.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2020.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2008.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2017.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2015.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/project.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2013.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2019.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/news/2010.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/build.xml Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/images/update.gif Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/images/void.gif Found
File C:\Users\user\Downloads\apache-tomcat-9.0.58.tar.gz:apache-tomcat-9.0.58/bin/tomcat-native.tar.gz:tomcat-native-1.2.31-src/xdocs/images/code.gif Found
.....
INFO: 1 scanned file, 408 scanned in archive, 409 files match file pattern, 0 line match line pattern
```

## Search for a string :

In the sample below we search for the string __Appender__ in any file log4j.* located in directory __/home/kali__
```
└─$ ./grep-archive -scan /home/kali -filter 'log4j.*' -grep-string 'Appender'
MATCH FOUND IN /home/kali/.BurpSuite/bapps/xxx/build/libs/hackvertor-all.jar:log4j.properties
     log4j.appender.CONSOLE=org.apache.log4j.ConsoleAppender
MATCH FOUND IN /home/kali/Audit/SAMPLE.war:WEB-INF/lib/mylib-1.14.jar:log4j.properties
     log4j.appender.A1=org.apache.log4j.RollingFileAppender
     log4j.appender.A2=org.apache.log4j.ConsoleAppender
```

log4j.property has been found in archive __/home/kali/.BurpSuite/bapps/xxx/build/libs/hackvertor-all.jar__ but also in archive __WEB-INF/lib/renault-1.14.jar__ contained in __/home/kali/Audit/FLUME.war__
