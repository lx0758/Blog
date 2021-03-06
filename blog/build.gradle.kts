import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    id("org.springframework.boot") version "2.4.3"
    id("io.spring.dependency-management") version "1.0.11.RELEASE"
    kotlin("jvm") version "1.4.30"
    kotlin("plugin.spring") version "1.4.30"
    id("org.jetbrains.kotlin.plugin.noarg") version "1.4.30"
}

group = "com.liux.blog"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_1_8

repositories {
    jcenter()
    mavenCentral()
}

dependencies {
    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")

    implementation("org.springframework.boot:spring-boot-starter-web")
    implementation("org.springframework.boot:spring-boot-starter-websocket")
    implementation("org.springframework.boot:spring-boot-starter-thymeleaf")
    implementation("org.springframework.boot:spring-boot-starter-security")
    implementation("org.springframework.boot:spring-boot-starter-data-redis")
    implementation("org.springframework.boot:spring-boot-starter-mail")

    implementation("org.springframework.session:spring-session-data-redis")

    implementation("com.github.pagehelper:pagehelper-spring-boot-starter:1.3.0")
    implementation("org.mybatis.spring.boot:mybatis-spring-boot-starter:2.1.4")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
    implementation("com.github.ua-parser:uap-java:1.5.2")
    implementation("com.github.penggle:kaptcha:2.3.2")

    implementation("org.commonmark:commonmark:0.17.1")
    implementation("org.commonmark:commonmark-ext-autolink:0.17.1")
    implementation("org.commonmark:commonmark-ext-gfm-tables:0.17.1")
    implementation("org.commonmark:commonmark-ext-gfm-strikethrough:0.17.1")

    developmentOnly("org.springframework.boot:spring-boot-devtools")
    runtimeOnly("mysql:mysql-connector-java")

    testImplementation("org.springframework.boot:spring-boot-starter-test")
    testImplementation("org.springframework.security:spring-security-test")
}

tasks.withType<KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict")
        jvmTarget = "1.8"
    }
}

tasks.withType<Test> {
    useJUnitPlatform()
}

// http://www.kotlincn.net/docs/reference/compiler-plugins.html
noArg {
    annotation("com.liux.blog.bean.NoArgsConstructor")
}


