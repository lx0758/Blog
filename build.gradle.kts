group = "com.liux.blog"
version = "0.0.1-SNAPSHOT"

allprojects {
    buildscript {
        repositories {
            maven {
                setUrl("https://maven.aliyun.com/repository/public/")
            }
            mavenCentral()
            mavenLocal()
        }
    }
    repositories {
        maven {
            setUrl("https://maven.aliyun.com/repository/public/")
        }
        mavenCentral()
        mavenLocal()
    }
}
